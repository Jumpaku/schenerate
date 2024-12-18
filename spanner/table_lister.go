package spanner

import (
	"cloud.google.com/go/spanner"
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
	tx := l.queryer.client.ReadOnlyTransaction()
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
	return lo.Map(records, func(r table, _ int) sqlgogen.Table {
		return sqlgogen.Table{Name: r.Name}
	}), nil
}
