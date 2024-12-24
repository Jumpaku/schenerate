package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/samber/lo"
)

type SchemaProcessHandler func(out *files.Writer, schemas Schemas) error

func ProcessSchema(ctx context.Context, q queryer, tables []string, handler SchemaProcessHandler) error {
	tx := q.client.ReadOnlyTransaction()
	defer tx.Close()

	var schemas []Schema
	for _, t := range tables {
		schema, err := queryTable(ctx, tx, t)
		if err != nil {
			return fmt.Errorf(`fail to get schema of %s: %w`, t, err)
		}
		schema.Columns, err = queryColumns(ctx, tx, t)
		if err != nil {
			return fmt.Errorf(`fail to get columns of %s: %w`, t, err)
		}
		schema.PrimaryKey, err = queryPrimaryKey(ctx, tx, t)
		if err != nil {
			return fmt.Errorf(`fail to get primary key of %s: %w`, t, err)
		}
		schema.ForeignKeys, err = queryForeignKeys(ctx, tx, t)
		if err != nil {
			return fmt.Errorf(`fail to get foreign keys of %s: %w`, t, err)
		}
		schema.Indexes, err = queryIndexes(ctx, tx, t)
		if err != nil {
			return fmt.Errorf(`fail to get indexes of %s: %w`, t, err)
		}

		schemas = append(schemas, schema)
	}

	w := &files.Writer{}
	if err := handler(w, schemas); err != nil {
		return err
	}

	if err := w.SaveAll(); err != nil {
		return fmt.Errorf(`fail to save files writer: %w`, err)
	}
	return nil
}

func queryTable(ctx context.Context, tx *spanner.ReadOnlyTransaction, table string) (Schema, error) {
	type recordTable struct {
		Name   string `db:"Name"`
		Type   string `db:"Type"`
		Parent string `db:"Parent"`
	}
	rows, err := query[recordTable](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query table information
SELECT
	TABLE_NAME AS Name,
	TABLE_TYPE AS Type,
	IFNULL(PARENT_TABLE_NAME, '') AS Parent
FROM INFORMATION_SCHEMA.TABLES
WHERE TABLE_NAME = @Table
ORDER BY TABLE_NAME`,
		Params: map[string]interface{}{"Table": table},
	})
	if err != nil {
		return Schema{}, fmt.Errorf(`fail to get table %s: %w`, table, err)
	}
	if len(rows) == 0 {
		return Schema{}, fmt.Errorf(`fail to get table %s: not found`, table)
	}

	record := rows[0]

	return Schema{Name: record.Name, Type: record.Type, Parent: record.Parent}, nil
}

func queryColumns(ctx context.Context, tx *spanner.ReadOnlyTransaction, table string) ([]Column, error) {
	type column struct {
		Name     string `db:"Name"`
		Type     string `db:"Type"`
		Nullable bool   `db:"Nullable"`
	}
	rows, err := query[column](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query column information
SELECT
	COLUMN_NAME AS Name,
	SPANNER_TYPE AS Type,
	(IS_NULLABLE = 'YES') AS Nullable,
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_NAME = @Table
ORDER BY ORDINAL_POSITION`,
		Params: map[string]interface{}{"Table": table},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get columns of %s: %w`, table, err)
	}

	return lo.Map(rows, func(item column, _ int) Column {
		return Column{Name: item.Name, Type: item.Type, Nullable: item.Nullable}
	}), nil
}

func queryPrimaryKey(ctx context.Context, tx *spanner.ReadOnlyTransaction, table string) ([]string, error) {
	type name struct {
		Name string `db:"Name"`
	}
	rows, err := query[name](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query primary key information
SELECT kcu.COLUMN_NAME AS Name
FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS AS tc
	JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE AS kcu
	ON kcu.CONSTRAINT_NAME = tc.CONSTRAINT_NAME 
        AND kcu.TABLE_NAME = tc.TABLE_NAME
WHERE kcu.TABLE_NAME = @Table AND tc.CONSTRAINT_TYPE = 'PRIMARY KEY'
ORDER BY kcu.ORDINAL_POSITION`,
		Params: map[string]interface{}{"Table": table},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get primary key of %s: %w`, table, err)
	}
	return lo.Map(rows, func(item name, _ int) string { return item.Name }), nil
}

func queryForeignKeys(ctx context.Context, tx *spanner.ReadOnlyTransaction, table string) ([]ForeignKey, error) {
	type fkRow struct {
		Name           string   `db:"Name"`
		Key            []string `db:"Key"`
		ReferenceTable string   `db:"ReferenceTable"`
		ReferenceKey   []string `db:"ReferenceKey"`
	}
	rows, err := query[fkRow](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query foreign key information
SELECT
	tc.CONSTRAINT_NAME AS Name,
	ARRAY(
		SELECT kcu.COLUMN_NAME
		FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu 
		WHERE kcu.CONSTRAINT_NAME = tc.CONSTRAINT_NAME
		ORDER BY kcu.ORDINAL_POSITION
	) AS Key,
	ctu.TABLE_NAME AS ReferenceTable,
	ARRAY(
		SELECT kcu.COLUMN_NAME
		FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu 
		WHERE kcu.CONSTRAINT_NAME = rc.UNIQUE_CONSTRAINT_NAME
		ORDER BY kcu.ORDINAL_POSITION
	) AS ReferenceKey
FROM
	INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
	JOIN INFORMATION_SCHEMA.REFERENTIAL_CONSTRAINTS rc ON rc.CONSTRAINT_NAME = tc.CONSTRAINT_NAME
	JOIN INFORMATION_SCHEMA.CONSTRAINT_TABLE_USAGE ctu ON ctu.CONSTRAINT_NAME = rc.UNIQUE_CONSTRAINT_NAME
WHERE tc.CONSTRAINT_TYPE = 'FOREIGN KEY' AND tc.TABLE_NAME = @Table
ORDER BY Name`,
		Params: map[string]interface{}{"Table": table},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get foreign keys of %s: %w`, table, err)
	}

	var foreignKeys []ForeignKey
	for _, row := range rows {
		foreignKeys = append(foreignKeys, ForeignKey{
			Name: row.Name,
			Key:  row.Key,
			Reference: ForeignKeyReference{
				Table: row.ReferenceTable,
				Key:   row.ReferenceKey,
			},
		})
	}

	return foreignKeys, nil
}

func queryIndexes(ctx context.Context, tx *spanner.ReadOnlyTransaction, table string) ([]Index, error) {
	type idxKey struct {
		Name   string `db:"Name"`
		IsDesc bool   `db:"IsDesc"`
	}
	type idxRow struct {
		Name     string `db:"Name"`
		IsUnique bool   `db:"IsUnique"`
		Key      []*idxKey
	}
	rows, err := query[idxRow](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query unique key information
WITH
	EXCLUDE_FK_BACKING AS (
		SELECT rc.UNIQUE_CONSTRAINT_NAME AS Name
		FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
			JOIN INFORMATION_SCHEMA.REFERENTIAL_CONSTRAINTS rc ON rc.CONSTRAINT_NAME = tc.CONSTRAINT_NAME
			JOIN INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc2 ON tc2.CONSTRAINT_NAME = rc.UNIQUE_CONSTRAINT_NAME
		WHERE tc.CONSTRAINT_TYPE = 'FOREIGN KEY'
	)
SELECT
	idx.INDEX_NAME AS Name,
	idx.IS_UNIQUE AS IsUnique,
	ARRAY(
		SELECT AS STRUCT
	    	idxc.COLUMN_NAME AS Name,
	    	idxc.COLUMN_ORDERING = 'DESC' AS IsDesc
		FROM INFORMATION_SCHEMA.INDEX_COLUMNS idxc
		WHERE idx.INDEX_NAME = idxc.INDEX_NAME AND idx.TABLE_NAME = idxc.TABLE_NAME
		ORDER BY idxc.ORDINAL_POSITION
	) AS Key,
FROM INFORMATION_SCHEMA.INDEXES idx
WHERE
	idx.TABLE_NAME = @Table 
	AND INDEX_TYPE = "INDEX"
	AND idx.INDEX_NAME NOT IN (SELECT Name FROM EXCLUDE_FK_BACKING)
ORDER BY Name`,
		Params: map[string]interface{}{"Table": table},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get unique keys of %s: %w`, table, err)
	}
	var index []Index
	for _, row := range rows {
		index = append(index, Index{
			Name:   row.Name,
			Unique: row.IsUnique,
			Key: lo.Map(row.Key, func(idxKey *idxKey, _ int) IndexKeyElem {
				return IndexKeyElem{Name: idxKey.Name, Desc: idxKey.IsDesc}
			}),
		})
	}

	return index, nil
}
