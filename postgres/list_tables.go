package postgres

import (
	"context"
	"fmt"
	"github.com/samber/lo"
)

type Table struct {
	Catalog string
	Schema  string
	Name    string
}

func ListTables(ctx context.Context, q queryer) ([]Table, error) {
	type table struct {
		Catalog string `db:"Catalog"`
		Schema  string `db:"Schema"`
		Name    string `db:"Name"`
	}
	records, err := query[table](ctx, q,
		`SELECT
    table_catalog AS "Catalog",
    table_schema AS "Schema",
	table_name AS "Name"
FROM information_schema.tables
WHERE table_schema NOT IN ('information_schema', 'pg_catalog')
ORDER BY table_catalog, table_schema, table_name`)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	return lo.Map(records, func(r table, _ int) Table {
		return Table{
			Catalog: r.Catalog,
			Schema:  r.Schema,
			Name:    r.Name,
		}
	}), nil
}
