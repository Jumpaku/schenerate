package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib"
)

func NewQueryProcessor[Record any](queryer queryer) sqlgogen.QueryProcessor[Record] {
	return queryProcessor[Record]{queryer: queryer}
}

type queryProcessor[Record any] struct {
	queryer queryer
}

func (p queryProcessor[Record]) Process(ctx context.Context, stmt sqlgogen.Statement, handler sqlgogen.QueryProcessHandler[Record]) error {
	tx := p.queryer.client.ReadOnlyTransaction()
	defer tx.Close()

	records, err := query[Record](ctx, tx, spanner.Statement{
		SQL:    stmt.Stmt,
		Params: stmt.ArgsMap(),
	})
	if err != nil {
		return fmt.Errorf("failed to query: %w", err)
	}
	return handler(records)
}
