package sqlite3

import (
	"cmp"
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib"
	"github.com/samber/lo"
	"slices"
)

func NewSchemaProcessor(queryer queryer) sqlgogen.SchemaProcessor[Schemas] {
	return schemaProcessor{queryer: queryer}
}

type schemaProcessor struct {
	queryer queryer
}

func (p schemaProcessor) Process(ctx context.Context, tables []sqlgogen.Table, handler sqlgogen.SchemaProcessHandler[Schemas]) error {
	var schemas []Schema
	for _, t := range tables {
		schema, err := queryTable(ctx, p.queryer, t.Name)
		if err != nil {
			return fmt.Errorf(`fail to get schema of %s: %w`, t.Name, err)
		}
		schema.Columns, err = queryColumns(ctx, p.queryer, t.Name)
		if err != nil {
			return fmt.Errorf(`fail to get columns of %s: %w`, t.Name, err)
		}
		schema.PrimaryKey, err = queryPrimaryKey(ctx, p.queryer, t.Name)
		if err != nil {
			return fmt.Errorf(`fail to get primary key of %s: %w`, t.Name, err)
		}
		schema.ForeignKeys, err = queryForeignKeys(ctx, p.queryer, t.Name)
		if err != nil {
			return fmt.Errorf(`fail to get foreign keys of %s: %w`, t.Name, err)
		}
		schema.Indexes, err = queryIndexes(ctx, p.queryer, t.Name)
		if err != nil {
			return fmt.Errorf(`fail to get indexes of %s: %w`, t.Name, err)
		}

		schemas = append(schemas, schema)
	}

	return handler(schemas)
}

func queryTable(ctx context.Context, q queryer, table string) (Schema, error) {
	type recordTable struct {
		Schema string `db:"Schema"`
		Name   string `db:"Name"`
		Type   string `db:"Type"`
	}
	rows, err := Query[recordTable](ctx, q, sqlgogen.Statement{
		//language=SQL
		Stmt: `--sql query table information
SELECT 
	"schema" AS Schema,
	"name" AS Name,
	"type" AS Type
FROM pragma_table_list(?)`,
		Args: []interface{}{table},
	})
	if err != nil {
		return Schema{}, fmt.Errorf(`fail to get table %s: %w`, table, err)
	}
	if len(rows) == 0 {
		return Schema{}, fmt.Errorf(`fail to get table %s: not found`, table)
	}

	record := rows[0]

	return Schema{Name: record.Name, Type: record.Type}, nil
}

func queryColumns(ctx context.Context, q queryer, table string) ([]Column, error) {
	type column struct {
		Name     string `db:"Name"`
		Type     string `db:"Type"`
		Nullable bool   `db:"Nullable"`
	}
	rows, err := Query[column](ctx, q, sqlgogen.Statement{
		//language=SQL
		Stmt: `--sql query column information
SELECT 
	"name" AS Name,
	"type" AS Type,
	"notnull" = 0 AS Nullable
FROM pragma_table_info(?)
ORDER BY "cid"`,
		Args: []interface{}{table},
	})
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
	rows, err := Query[name](ctx, q, sqlgogen.Statement{
		//language=SQL
		Stmt: `--sql query primary key information
SELECT "name" AS Name
FROM pragma_table_info(?)
WHERE "pk" > 0
ORDER BY "pk"`,
		Args: []interface{}{table},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get primary key of %s: %w`, table, err)
	}
	return lo.Map(rows, func(item name, _ int) string { return item.Name }), nil
}

func queryForeignKeys(ctx context.Context, q queryer, table string) ([]ForeignKey, error) {
	type fkRow struct {
		Id           int64  `db:"Id"`
		Seq          int64  `db:"Seq"`
		ForeignTable string `db:"ForeignTable"`
		FromColumn   string `db:"FromColumn"`
		ToColumn     string `db:"ToColumn"`
	}
	rows, err := Query[fkRow](ctx, q, sqlgogen.Statement{
		//language=SQL
		Stmt: `--sql query foreign key information
SELECT
	"id" AS Id,
	"seq" AS Seq,
	"table" AS ForeignTable,
	"from" AS FromColumn,
	"to" AS ToColumn
FROM pragma_foreign_key_list(?)
ORDER BY "id", "seq"`,
		Args: []interface{}{table},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get foreign keys of %s: %w`, table, err)
	}

	group := lo.GroupBy(rows, func(fkRow fkRow) int64 { return fkRow.Id })
	groupIDs := lo.MapToSlice(group, func(id int64, _ []fkRow) int64 { return id })
	slices.Sort(groupIDs)

	var foreignKeys []ForeignKey
	for _, id := range groupIDs {
		g := group[id]
		foreignKeys = append(foreignKeys, ForeignKey{
			Key: lo.Map(g, func(fkRow fkRow, _ int) string { return fkRow.FromColumn }),
			Reference: ForeignKeyReference{
				Table: g[0].ForeignTable,
				Key:   lo.Map(g, func(fkRow fkRow, _ int) string { return fkRow.ToColumn }),
			},
		})
	}

	return foreignKeys, nil
}

func queryIndexes(ctx context.Context, q queryer, table string) ([]Index, error) {
	type idxRow struct {
		Seq      int64  `db:"Seq"`
		Name     string `db:"Name"`
		Origin   string `db:"Origin"`
		IsUnique bool   `db:"IsUnique"`
		ColName  string `db:"ColName"`
		IsDesc   bool   `db:"IsDesc"`
	}
	rows, err := Query[idxRow](ctx, q, sqlgogen.Statement{
		//language=SQL
		Stmt: `--sql query index information
SELECT
    pil."seq" AS Seq,
    pil."name" AS Name,
    pil."origin" AS Origin,
    pil."unique" AS IsUnique,
    pii."name" AS ColName,
    pii."desc" AS IsDesc
FROM pragma_index_list(?) AS pil
    JOIN pragma_index_xinfo(pil.name) AS pii
WHERE pii."name" <> ''
ORDER BY pil."origin", pil."name", pil."seq", pii."seqno"`,
		Args: []interface{}{table},
	})
	if err != nil {
		return nil, fmt.Errorf(`fail to get unique keys of %s: %w`, table, err)
	}
	type groupKey struct {
		Origin string
		Name   string
		Seq    int64
	}
	group := lo.GroupBy(rows, func(idxRow idxRow) groupKey {
		return groupKey{Origin: idxRow.Origin, Name: idxRow.Name, Seq: idxRow.Seq}
	})
	groupIDs := lo.MapToSlice(group, func(k groupKey, _ []idxRow) groupKey { return k })
	slices.SortFunc(groupIDs, func(a, b groupKey) int {
		if a.Origin != b.Origin {
			return cmp.Compare(a.Origin, b.Origin)
		}
		if a.Name != b.Name {
			return cmp.Compare(a.Name, b.Name)
		}
		return cmp.Compare(a.Seq, b.Seq)
	})

	var index []Index
	for _, id := range groupIDs {
		g := group[id]
		index = append(index, Index{
			Name:   g[0].Name,
			Origin: IndexOrigin(g[0].Origin),
			Unique: g[0].IsUnique,
			Key: lo.Map(g, func(idxRow idxRow, _ int) IndexKeyElem {
				return IndexKeyElem{Name: idxRow.ColName, Desc: idxRow.IsDesc}
			}),
		})
	}

	return index, nil
}
