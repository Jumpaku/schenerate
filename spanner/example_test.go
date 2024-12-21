package spanner_test

import (
	spanner2 "cloud.google.com/go/spanner"
	"context"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/spanner"
)

func Example_processSchema() {
	ctx := context.Background()
	client, err := spanner2.NewClient(ctx, "projects/<project>/instances/<instance>/databases/<database>")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	tx := client.ReadOnlyTransaction()
	defer tx.Close()

	err = spanner.ProcessSchema(
		context.Background(),
		tx,
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
