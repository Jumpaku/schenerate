package sqlite3_test

import (
	"context"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
	"github.com/Jumpaku/schenerate/sqlite3"
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

func Example_generateWithQuery() {
	type T struct {
		Name string `db:"name"`
		Age  int64  `db:"age"`
	}
	q, err := sqlite3.Open("db.sqlite")
	if err != nil {
		panic(err)
	}
	defer q.Close()

	err = sqlite3.GenerateWithQuery(context.Background(), q,
		`SELECT 'A' AS "name", 1 AS "age" UNION SELECT 'B' AS "name", 2 AS "age"`,
		nil,
		func(w *files.Writer, rows []T) error {
			w.Add("records.txt")
			for _, row := range rows {
				// do something with a row
				fmt.Fprintf(w, "%+v\n", row)
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
}
