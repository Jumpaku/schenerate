package sqlite3_test

import (
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib"
	"github.com/Jumpaku/sql-gogen-lib/sqlite3"
)

func ExampleNewSchemaProcessor() {
	db, err := sqlite3.Open("db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := sqlite3.NewSchemaProcessor(db)
	err = p.Process(
		context.Background(),
		[]sqlgogen.Table{{Name: "Table"}},
		func(schemas sqlite3.Schemas) error {
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
