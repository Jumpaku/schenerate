package sqlite3_test

import (
	"context"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/sqlite3"
)

func Example_generateWithSchema() {
	q, err := sqlite3.Open("db.sqlite")
	if err != nil {
		panic(err)
	}
	defer q.Close()

	err = sqlite3.GenerateWithSchema(context.Background(), q,
		[]string{"Table"},
		func(w *files.Writer, schemas sqlite3.Schemas) error {
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
