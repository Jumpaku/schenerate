package postgres

import (
	"context"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/samber/lo"
	"slices"
)

type SchemaProcessHandler func(out *files.Writer, schemas Schemas) error

func ProcessSchema(ctx context.Context, q queryer, tables []string, handler SchemaProcessHandler) error {
	var schemas []Schema
	for _, t := range tables {
		schema, err := queryTable(ctx, q, t)
		if err != nil {
			return fmt.Errorf(`fail to get schema of %s: %w`, t, err)
		}
		schema.Columns, err = queryColumns(ctx, q, t)
		if err != nil {
			return fmt.Errorf(`fail to get columns of %s: %w`, t, err)
		}
		schema.PrimaryKey, err = queryPrimaryKey(ctx, q, t)
		if err != nil {
			return fmt.Errorf(`fail to get primary key of %s: %w`, t, err)
		}
		schema.ForeignKeys, err = queryForeignKeys(ctx, q, t)
		if err != nil {
			return fmt.Errorf(`fail to get foreign keys of %s: %w`, t, err)
		}
		schema.UniqueKeys, err = queryUniqueKeys(ctx, q, t)
		if err != nil {
			return fmt.Errorf(`fail to get unique keys of %s: %w`, t, err)
		}
		schema.Indexes, err = queryIndexes(ctx, q, t)
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

func queryTable(ctx context.Context, q queryer, table string) (Schema, error) {
	type recordTable struct {
		Schema string `db:"Schema"`
		Name   string `db:"Name"`
		Type   string `db:"Type"`
	}
	rows, err := query[recordTable](ctx, q,
		//language=SQL
		`--sql query table information
SELECT
	table_schema AS "Schema",
	table_name AS "Name",
	table_type AS "Type"
FROM information_schema.tables
WHERE table_name = $1
ORDER BY table_name`, table)
	if err != nil {
		return Schema{}, fmt.Errorf(`fail to get table %s: %w`, table, err)
	}
	if len(rows) == 0 {
		return Schema{}, fmt.Errorf(`fail to get table %s: not found`, table)
	}

	record := rows[0]

	return Schema{Schema: record.Schema, Name: record.Name, Type: record.Type}, nil
}

func queryColumns(ctx context.Context, q queryer, table string) ([]Column, error) {
	type column struct {
		Name     string `db:"Name"`
		Type     string `db:"Type"`
		Nullable bool   `db:"Nullable"`
	}
	rows, err := query[column](ctx, q,
		//language=SQL
		`--sql query column information
SELECT 
	column_name AS "Name",
	data_type AS "Type",
	is_nullable = 'YES' AS "Nullable"
FROM information_schema.columns
WHERE table_name = $1
ORDER BY ordinal_position`, table)
	if err != nil {
		return nil, fmt.Errorf(`fail to get columns of %s: %w`, table, err)
	}

	return lo.Map(rows, func(item column, _ int) Column {
		return Column{Name: item.Name, Type: item.Type, Nullable: item.Nullable}
	}), nil
}

func queryPrimaryKey(ctx context.Context, q queryer, table string) ([]string, error) {
	type name struct {
		Name string `db:"Name"`
	}
	rows, err := query[name](ctx, q,
		//language=SQL
		`--sql query primary key information
SELECT
    kcu.column_name AS "Name"
FROM information_schema.table_constraints AS tc
    JOIN information_schema.key_column_usage AS kcu
    	ON kcu.constraint_name = tc.constraint_name
WHERE kcu.table_name = $1 AND tc.constraint_type = 'PRIMARY KEY'
ORDER BY kcu.ordinal_position`, table)
	if err != nil {
		return nil, fmt.Errorf(`fail to get primary key of %s: %w`, table, err)
	}
	return lo.Map(rows, func(item name, _ int) string { return item.Name }), nil
}

func queryForeignKeys(ctx context.Context, q queryer, table string) ([]ForeignKey, error) {
	type fkRow struct {
		Name             string `db:"Name"`
		ReferencedSchema string `db:"ReferencedSchema"`
		ReferencedTable  string `db:"ReferencedTable"`
		ReferencingKey   string `db:"ReferencingKey"`
		ReferencedKey    string `db:"ReferencedKey"`
	}
	rows, err := query[fkRow](ctx, q,
		//language=SQL
		`--sql query foreign key information
SELECT
    tc.constraint_name AS "Name",
    ctu.table_schema AS "ReferencedSchema",
    ctu.table_name AS "ReferencedTable",
    kcu1.column_name AS "ReferencingKey",
    kcu2.column_name AS "ReferencedKey"
FROM
    information_schema.table_constraints tc
        JOIN information_schema.referential_constraints rc
            ON rc.constraint_name = tc.constraint_name
        JOIN information_schema.constraint_table_usage ctu
            ON ctu.constraint_name = rc.unique_constraint_name
        JOIN information_schema.key_column_usage kcu1
            ON kcu1.constraint_name = rc.constraint_name
        JOIN information_schema.key_column_usage kcu2
            ON kcu2.constraint_name = rc.unique_constraint_name
                AND kcu2.ordinal_position = kcu1.ordinal_position
WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name = $1
ORDER BY "Name", kcu1.ordinal_position;`, table)
	if err != nil {
		return nil, fmt.Errorf(`fail to get foreign keys of %s: %w`, table, err)
	}

	group := lo.GroupBy(rows, func(fkRow fkRow) string { return fkRow.Name })
	groupNames := lo.Keys(group)
	slices.Sort(groupNames)

	var foreignKeys []ForeignKey
	for _, id := range groupNames {
		g := group[id]
		foreignKeys = append(foreignKeys, ForeignKey{
			Name: g[0].Name,
			Key:  lo.Map(g, func(fkRow fkRow, _ int) string { return fkRow.ReferencingKey }),
			Reference: ForeignKeyReference{
				Schema: g[0].ReferencedSchema,
				Table:  g[0].ReferencedTable,
				Key:    lo.Map(g, func(fkRow fkRow, _ int) string { return fkRow.ReferencedKey }),
			},
		})
	}

	return foreignKeys, nil
}

func queryUniqueKeys(ctx context.Context, q queryer, table string) ([]UniqueKey, error) {
	type ukRow struct {
		Name       string `db:"Name"`
		ColumnName string `db:"ColumnName"`
	}

	rows, err := query[ukRow](ctx, q,
		//language=SQL
		`--sql query primary key information
SELECT
    tc.constraint_name AS "Name",
    kcu.column_name AS "ColumnName"
FROM information_schema.table_constraints AS tc
	 JOIN information_schema.key_column_usage AS kcu
		  ON kcu.constraint_name = tc.constraint_name
WHERE kcu.table_name = $1 AND tc.constraint_type = 'UNIQUE'
ORDER BY tc.constraint_name, kcu.ordinal_position;`, table)
	if err != nil {
		return nil, fmt.Errorf(`fail to get unique keys of %s: %w`, table, err)
	}
	group := lo.GroupBy(rows, func(ukRow ukRow) string { return ukRow.Name })
	groupNames := lo.Keys(group)
	slices.Sort(groupNames)

	var uniqueKeys []UniqueKey
	for _, name := range groupNames {
		g := group[name]

		uniqueKeys = append(uniqueKeys, UniqueKey{
			Name: name,
			Key:  lo.Map(g, func(ukRow ukRow, _ int) string { return ukRow.ColumnName }),
		})
	}

	return uniqueKeys, nil
}

func queryIndexes(ctx context.Context, q queryer, table string) ([]Index, error) {
	type idxRow struct {
		Name      string `db:"Name"`
		Unique    bool   `db:"Unique"`
		KeyColumn string `db:"KeyColumn"`
	}
	rows, err := query[idxRow](ctx, q,
		//language=SQL
		`-- sql query index information
SELECT cls.relname  AS "Name",
       idx.indisunique AS "Unique",
       attr.attname AS "KeyColumn"
FROM pg_catalog.pg_index AS idx
         JOIN pg_catalog.pg_class AS cls ON idx.indexrelid = cls.oid
         JOIN pg_catalog.pg_class AS tbl ON idx.indrelid = tbl.oid
         JOIN pg_catalog.pg_attribute AS attr ON idx.indexrelid = attr.attrelid
WHERE tbl.relname = $1 AND NOT idx.indisprimary
ORDER BY cls.relname, attr.attnum;`, table)
	if err != nil {
		return nil, fmt.Errorf(`fail to get unique keys of %s: %w`, table, err)
	}
	group := lo.GroupBy(rows, func(idxRow idxRow) string { return idxRow.Name })
	groupNames := lo.Keys(group)
	slices.Sort(groupNames)
	var index []Index
	for _, name := range groupNames {
		g := group[name]
		index = append(index, Index{
			Name:   g[0].Name,
			Unique: g[0].Unique,
			Key:    lo.Map(g, func(idxKey idxRow, _ int) string { return idxKey.KeyColumn }),
		})
	}

	return index, nil
}
