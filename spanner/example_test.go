package spanner_test

import (
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib"
	"github.com/Jumpaku/sql-gogen-lib/spanner"
)

func ExampleNewSchemaProcessor() {
	db, err := spanner.Open("<project>", "instance", "database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := spanner.NewSchemaProcessor(db)
	err = p.Process(
		context.Background(),
		[]sqlgogen.Table{{Name: "Table"}},
		func(schemas spanner.Schemas) error {
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
