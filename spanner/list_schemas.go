package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"github.com/samber/lo"
	"slices"
)

func ListSchemas(ctx context.Context, q queryer, tables []string) (schemas Schemas, err error) {
	tx := q.client.ReadOnlyTransaction()
	defer tx.Close()

	schemaMap, err := queryTables(ctx, tx, tables)
	if err != nil {
		return nil, fmt.Errorf(`fail to get schema: %w`, err)
	}

	columnsMap, err := queryTableColumns(ctx, tx, tables)
	if err != nil {
		return nil, fmt.Errorf(`fail to get columns: %w`, err)
	}
	for t, c := range columnsMap {
		s, ok := schemaMap[t]
		if !ok {
			return nil, fmt.Errorf(`fail to get columns of %s: table not found`, t)
		}
		s.Columns = c
	}

	primaryKeysMap, err := queryTablePrimaryKeys(ctx, tx, tables)
	if err != nil {
		return nil, fmt.Errorf(`fail to get primary keys: %w`, err)
	}
	for t, pk := range primaryKeysMap {
		s, ok := schemaMap[t]
		if !ok {
			return nil, fmt.Errorf(`fail to get primary key of %s: table not found`, t)
		}
		s.PrimaryKey = pk
	}

	foreignKeysMap, err := queryTableForeignKeys(ctx, tx, tables)
	if err != nil {
		return nil, fmt.Errorf(`fail to get foreign keys: %w`, err)
	}
	for t, fks := range foreignKeysMap {
		s, ok := schemaMap[t]
		if !ok {
			return nil, fmt.Errorf(`fail to get foreign keys of %s: table not found`, t)
		}
		s.ForeignKeys = fks
	}

	indexesMap, err := queryTableIndexes(ctx, tx, tables)
	if err != nil {
		return nil, fmt.Errorf(`fail to get indexes: %w`, err)
	}
	for t, idx := range indexesMap {
		s, ok := schemaMap[t]
		if !ok {
			return nil, fmt.Errorf(`fail to get indexes of %s: table not found`, t)
		}
		s.Indexes = idx
	}

	ks := lo.Keys(schemaMap)
	slices.Sort(ks)

	return lo.Map(ks, func(t string, _ int) Schema { return *schemaMap[t] }), nil
}

func queryTables(ctx context.Context, tx *spanner.ReadOnlyTransaction, tables []string) (map[string]*Schema, error) {
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
WHERE TABLE_NAME IN UNNEST(@Tables)
ORDER BY TABLE_NAME`,
		Params: map[string]interface{}{"Tables": tables},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get tables: %w`, err)
	}

	return lo.SliceToMap(rows, func(t recordTable) (string, *Schema) {
		return t.Name, &Schema{
			Name:        t.Name,
			Type:        t.Type,
			Parent:      t.Parent,
			ForeignKeys: nil,
			Indexes:     nil,
		}
	}), nil
}

func queryTableColumns(ctx context.Context, tx *spanner.ReadOnlyTransaction, tables []string) (map[string][]Column, error) {
	type column struct {
		Name     string `db:"Name"`
		Type     string `db:"Type"`
		Nullable bool   `db:"Nullable"`
	}
	type tableColumns struct {
		TableName    string    `db:"TableName"`
		TableColumns []*column `db:"TableColumns"`
	}
	rows, err := query[tableColumns](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query column information
SELECT
    TableName,
	ARRAY(
	    SELECT AS STRUCT
			COLUMN_NAME AS Name,
			SPANNER_TYPE AS Type,
			(IS_NULLABLE = 'YES') AS Nullable,
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_NAME = TableName
		ORDER BY ORDINAL_POSITION
	) AS TableColumns
FROM UNNEST(@Tables) AS TableName`,
		Params: map[string]interface{}{"Tables": tables},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get columns: %w`, err)
	}

	return lo.SliceToMap(rows, func(tc tableColumns) (string, []Column) {
		return tc.TableName, lo.Map(tc.TableColumns, func(c *column, _ int) Column {
			return Column{Name: c.Name, Type: c.Type, Nullable: c.Nullable}
		})
	}), nil
}

func queryTablePrimaryKeys(ctx context.Context, tx *spanner.ReadOnlyTransaction, tables []string) (map[string][]string, error) {
	type tablePk struct {
		TableName       string   `db:"TableName"`
		TablePrimaryKey []string `db:"TablePrimaryKey"`
	}
	rows, err := query[tablePk](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query primary key information
SELECT
    TableName,
    ARRAY(
		SELECT AS VALUE kcu.COLUMN_NAME
		FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS AS tc
			JOIN INFORMATION_SCHEMA.KEY_COLUMN_USAGE AS kcu
			ON kcu.CONSTRAINT_NAME = tc.CONSTRAINT_NAME 
				AND kcu.TABLE_NAME = tc.TABLE_NAME
		WHERE kcu.TABLE_NAME = TableName AND tc.CONSTRAINT_TYPE = 'PRIMARY KEY'
		ORDER BY kcu.ORDINAL_POSITION
	) AS TablePrimaryKey
FROM UNNEST(@Tables) AS TableName`,
		Params: map[string]interface{}{"Tables": tables},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get primary key: %w`, err)
	}
	return lo.SliceToMap(rows, func(item tablePk) (string, []string) {
		return item.TableName, item.TablePrimaryKey
	}), nil
}

func queryTableForeignKeys(ctx context.Context, tx *spanner.ReadOnlyTransaction, tables []string) (map[string][]ForeignKey, error) {
	type fkRow struct {
		Name           string   `db:"Name"`
		Key            []string `db:"Key"`
		ReferenceTable string   `db:"ReferenceTable"`
		ReferenceKey   []string `db:"ReferenceKey"`
	}
	type tableFk struct {
		TableName        string   `db:"TableName"`
		TableForeignKeys []*fkRow `db:"TableForeignKeys"`
	}
	rows, err := query[tableFk](ctx, tx, spanner.Statement{
		//language=SQL
		SQL: `--sql query foreign key information
SELECT
    TableName,
    ARRAY(
		SELECT AS STRUCT
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
		WHERE tc.CONSTRAINT_TYPE = 'FOREIGN KEY' AND tc.TABLE_NAME = TableName
		ORDER BY Name
	) AS TableForeignKeys
FROM UNNEST(@Tables) AS TableName`,
		Params: map[string]interface{}{"Tables": tables},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get foreign keys: %w`, err)
	}
	foreignKeysMap := map[string][]ForeignKey{}
	for _, row := range rows {
		foreignKeys := lo.Map(row.TableForeignKeys, func(fkRow *fkRow, _ int) ForeignKey {
			return ForeignKey{
				Name: fkRow.Name,
				Key:  fkRow.Key,
				Reference: ForeignKeyReference{
					Table: fkRow.ReferenceTable,
					Key:   fkRow.ReferenceKey,
				},
			}
		})
		foreignKeysMap[row.TableName] = foreignKeys
	}
	return foreignKeysMap, nil
}

func queryTableIndexes(ctx context.Context, tx *spanner.ReadOnlyTransaction, tables []string) (map[string][]Index, error) {
	type idxKey struct {
		Name   string `db:"Name"`
		IsDesc bool   `db:"IsDesc"`
	}
	type idxRow struct {
		Name     string `db:"Name"`
		IsUnique bool   `db:"IsUnique"`
		Key      []*idxKey
	}
	type tableIdx struct {
		TableName    string    `db:"TableName"`
		TableIndexes []*idxRow `db:"TableIndexes"`
	}
	rows, err := query[tableIdx](ctx, tx, spanner.Statement{
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
    TableName,
    ARRAY(
		SELECT AS STRUCT
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
			idx.TABLE_NAME = TableName
			AND INDEX_TYPE = "INDEX"
			AND idx.INDEX_NAME NOT IN (SELECT Name FROM EXCLUDE_FK_BACKING)
		ORDER BY Name
	) AS TableIndexes
FROM UNNEST(@Tables) AS TableName`,
		Params: map[string]interface{}{"Tables": tables},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get unique keys: %w`, err)
	}
	indexMap := map[string][]Index{}
	for _, row := range rows {
		var index []Index
		for _, row := range row.TableIndexes {
			index = append(index, Index{
				Name:   row.Name,
				Unique: row.IsUnique,
				Key: lo.Map(row.Key, func(idxKey *idxKey, _ int) IndexKeyElem {
					return IndexKeyElem{Name: idxKey.Name, Desc: idxKey.IsDesc}
				}),
			})
		}
		indexMap[row.TableName] = index
	}
	return indexMap, nil
}
