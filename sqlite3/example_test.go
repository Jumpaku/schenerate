package sqlite3_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/sqlite3"
)

func Example_processSchema() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = sqlite3.ProcessSchema(
		context.Background(),
		db,
		[]string{"Table"},
		func(out *files.Writer, schemas sqlite3.Schemas) error {
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
