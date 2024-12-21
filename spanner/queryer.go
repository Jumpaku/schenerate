package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
)

type Queryer interface {
	Query(ctx context.Context, statement spanner.Statement) *spanner.RowIterator
}

func query[Record any](ctx context.Context, tx Queryer, stmt spanner.Statement) (records []Record, err error) {
	rows := tx.Query(ctx, stmt)
	for {
		row, err := rows.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			}
			return nil, fmt.Errorf("failed to query: %w", err)
		}
		var record Record
		if err := row.ToStructLenient(&record); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}
		records = append(records, record)
	}

	return records, nil
}
