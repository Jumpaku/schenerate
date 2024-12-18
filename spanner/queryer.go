package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
)

func Open(project, instance, database string) (queryer, error) {
	c, err := spanner.NewClient(context.Background(), fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, database))
	if err != nil {
		return queryer{}, fmt.Errorf("failed to create spanner client: %w", err)
	}

	return NewClient(c), nil
}

func NewClient(client *spanner.Client) queryer {
	return queryer{client: client}
}

type queryer struct {
	client *spanner.Client
}

func (q queryer) Close() {
	q.client.Close()
}

func query[Record any](ctx context.Context, tx *spanner.ReadOnlyTransaction, stmt spanner.Statement) (records []Record, err error) {
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
