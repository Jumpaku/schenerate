package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type queryer struct {
	dbx *sqlx.DB
}

func Open(dsn string) (queryer, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return queryer{}, fmt.Errorf("failed to open: %w", err)
	}
	return queryer{dbx: sqlx.NewDb(db, "sqlite3")}, nil
}

func (q queryer) Close() error {
	return q.dbx.Close()
}

func query[Record any](ctx context.Context, q queryer, stmt string, args ...any) (records []Record, err error) {
	rows, err := q.dbx.QueryxContext(ctx, stmt, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	for rows.Next() {
		var record Record
		err := rows.StructScan(&record)
		if err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}
		records = append(records, record)
	}

	return records, nil
}
