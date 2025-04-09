package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"github.com/samber/lo"
)

func ListTables(ctx context.Context, q queryer) ([]string, error) {
	tx := q.client.ReadOnlyTransaction()
	defer tx.Close()

	type table struct {
		Name string `db:"Name"`
	}
	records, err := query[table](ctx, tx, spanner.Statement{
		SQL: `SELECT TABLE_NAME AS Name
FROM INFORMATION_SCHEMA.TABLES
WHERE TABLE_SCHEMA = ''
ORDER BY TABLE_NAME`,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	return lo.Map(records, func(r table, _ int) string { return r.Name }), nil
}
