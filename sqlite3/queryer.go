package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Open(sqlite3ConnectionString string) (queryer, error) {
	dbx, err := sqlx.Open("sqlite3", sqlite3ConnectionString)
	if err != nil {
		return queryer{}, fmt.Errorf("failed to connect to sqlite3: %w", err)
	}
	return queryer{db: dbx}, nil
}

func New(db *sql.DB) queryer {
	dbx := sqlx.NewDb(db, "sqlite3")
	return queryer{db: dbx}
}

type queryer struct {
	db *sqlx.DB
}

func (q queryer) Close() error {
	return q.db.Close()
}

func Query[Record any](ctx context.Context, q queryer, stmt sqlgogen.Statement) (records []Record, err error) {
	rows, err := q.db.QueryxContext(ctx, stmt.Stmt, stmt.Args...)
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
