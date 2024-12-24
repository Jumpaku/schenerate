package postgres

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	"testing"
)

var dataSource = flag.String("data-source", "", "-data-source=<data source>")

func Setup(t *testing.T, database string, ddls []string) (q queryer, teardown func()) {
	t.Helper()

	if *dataSource == "" {
		t.Skip(`postgres dataSource is required`)
		return queryer{}, nil
	}

	dataSource := *dataSource

	ctx := context.Background()
	{
		{
			db, err := pgx.Connect(ctx, dataSource)
			if err != nil {
				t.Fatalf(`fail to create postgres admin client: %v`, err)
			}
			defer db.Close(ctx)

			_, err = db.Exec(ctx, fmt.Sprintf(`CREATE DATABASE "%s"`, database))
			if err != nil {
				t.Fatalf(`fail to create postgres database in %q: %+v`, dataSource, err)
			}
		}

		{
			c, err := pgx.ParseConfig(dataSource)
			if err != nil {
				t.Fatalf(`fail to parse postgres dataSource: %+v`, err)
			}
			c.Database = database
			db, err := pgx.ConnectConfig(ctx, c)
			if err != nil {
				t.Fatalf(`fail to create postgres client: %+v`, err)
			}
			defer db.Close(ctx)
			for i, ddl := range ddls {
				_, err := db.Exec(ctx, ddl)
				if err != nil {
					t.Fatalf(`fail to execute DDL[%d]: %q: %+v`, i, ddl, err)
				}
			}
		}
	}

	c, err := pgx.ParseConfig(dataSource)
	if err != nil {
		t.Fatalf(`fail to parse postgres dataSource: %+v`, err)
	}
	c.Database = database
	db, err := pgx.ConnectConfig(ctx, c)
	if err != nil {
		t.Fatalf(`fail to create postgres client with %q : %+v`, dataSource, err)
	}
	teardown = func() {
		db.Close(ctx)
	}
	return queryer{conn: db}, teardown
}
