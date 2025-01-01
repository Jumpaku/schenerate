package spanner_test

import (
	"context"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/spanner"
)

func Example_generateWithSchema() {
	ctx := context.Background()
	q, err := spanner.Open(ctx, "<project>", "<instance>", "<database>")
	if err != nil {
		panic(err)
	}
	defer q.Close()

	err = spanner.GenerateWithSchema(context.Background(), q,
		[]string{"Table"},
		func(w *files.Writer, schemas spanner.Schemas) error {
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
