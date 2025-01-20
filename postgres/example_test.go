package postgres_test

import (
	"context"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
	"github.com/Jumpaku/schenerate/postgres"
)

func Example_generateWithSchema() {
	ctx := context.Background()
	q, err := postgres.Open("postgres://<user>:<password>@<host>:<port>/<dbname>")
	if err != nil {
		panic(err)
	}
	defer q.Close()

	err = postgres.GenerateWithSchema(ctx, q,
		[]string{"Table"},
		func(w *files.Writer, schemas postgres.Schemas) error {
			for _, schema := range schemas {
				// do something with schemas
				w.Add(schema.Name)
				fmt.Fprintf(w, "%+v\n", schema.Name)
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
}
