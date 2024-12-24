package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type queryer struct {
	conn *pgx.Conn
}

func Open(connStr string) (queryer, error) {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return queryer{}, fmt.Errorf("failed to connect: %w", err)
	}
	return queryer{conn: conn}, nil
}

func (q queryer) Close() error {
	return q.conn.Close(context.Background())
}

func query[Record any](ctx context.Context, q queryer, stmt string, args ...any) (records []Record, err error) {
	rows, err := q.conn.Query(ctx, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	for rows.Next() {
		record, err := pgx.RowToStructByNameLax[Record](rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}
		records = append(records, record)
	}

	return records, nil
}
