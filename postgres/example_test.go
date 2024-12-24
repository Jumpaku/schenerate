package postgres_test

import (
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/postgres"
)

func Example_processSchema() {
	ctx := context.Background()
	q, err := postgres.Open("postgres://<user>:<password>@<host>:<port>/<dbname>")
	if err != nil {
		panic(err)
	}
	defer q.Close()

	err = postgres.ProcessSchema(ctx, q,
		[]string{"Table"},
		func(_ *sqlgogen.Writer, schemas postgres.Schemas) error {
			for _, schema := range schemas {
				// do something with schemas
				fmt.Printf("%+v\n", schema.Name)
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
}
