package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func query[Record any](ctx context.Context, sqlite3 *sql.DB, stmt string, args ...any) (records []Record, err error) {
	dbx := sqlx.NewDb(sqlite3, "sqlite3")

	rows, err := dbx.QueryxContext(ctx, stmt, args...)
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
