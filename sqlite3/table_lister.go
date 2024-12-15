package sqlite3

import (
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib"
	"github.com/samber/lo"
)

func NewTableLister(queryer queryer) sqlgogen.TableLister {
	return tableLister{queryer: queryer}
}

type tableLister struct {
	queryer queryer
}

func (l tableLister) List(ctx context.Context) ([]sqlgogen.Table, error) {
	type table struct {
		Schema string `db:"Schema"`
		Name   string `db:"Name"`
		Type   string `db:"Type"`
	}
	records, err := Query[table](ctx, l.queryer, sqlgogen.Statement{
		Stmt: `SELECT
	"schema" AS Schema,
	"name" AS Name,
	"type" AS Type
FROM pragma_table_list()
ORDER BY "schema", "name"`,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	return lo.Map(records, func(r table, _ int) sqlgogen.Table {
		return sqlgogen.Table{Schema: r.Schema, Name: r.Name}
	}), nil
}
