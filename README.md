# schenerate

schenerate is a Go library that helps to implement code generators based on database schemas.

Currently, the following databases are supported:

* Spanner: https://pkg.go.dev/github.com/Jumpaku/schenerate/spanner
* SQLite3: https://pkg.go.dev/github.com/Jumpaku/schenerate/sqlite3
* PostgreSQL: https://pkg.go.dev/github.com/Jumpaku/schenerate/postgres

## Installation

```shell
go get github.com/Jumpaku/schenerate@latest
```

## Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
	"github.com/Jumpaku/schenerate/sqlite3"
)

func main() {
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
```