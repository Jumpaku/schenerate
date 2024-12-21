package spanner_test

import (
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/spanner"
)

func Example_processSchema() {
	ctx := context.Background()
	q, err := spanner.Open(ctx, "<project>", "<instance>", "<database>")
	if err != nil {
		panic(err)
	}
	defer q.Close()

	err = spanner.ProcessSchema(context.Background(), q,
		[]string{"Table"},
		func(_ *sqlgogen.Writer, schemas spanner.Schemas) error {
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
