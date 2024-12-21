package sqlite3

import (
	"context"
	"fmt"
	"github.com/samber/lo"
)

type Table struct {
	Schema string
	Name   string
}

func ListTables(ctx context.Context, q queryer) ([]Table, error) {
	type table struct {
		Schema string `db:"Schema"`
		Name   string `db:"Name"`
		Type   string `db:"Type"`
	}
	records, err := query[table](ctx, q, `SELECT
	"schema" AS Schema,
	"name" AS Name,
	"type" AS Type
FROM pragma_table_list()
ORDER BY "schema", "name"`)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	return lo.Map(records, func(r table, _ int) Table {
		return Table{Schema: r.Schema, Name: r.Name}
	}), nil
}
