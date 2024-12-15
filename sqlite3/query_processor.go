package sqlite3

import (
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
	records, err := Query[Record](ctx, p.queryer, stmt)
	if err != nil {
		fmt.Errorf("failed to query: %w", err)
	}
	return handler(records)
}
