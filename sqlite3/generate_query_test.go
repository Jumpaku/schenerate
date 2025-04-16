package sqlite3_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
	"github.com/Jumpaku/schenerate/sqlite3"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGenerateWithQuery(t *testing.T) {
	type T struct {
		Name string `db:"name"`
		Age  int64  `db:"age"`
	}
	testcases := []struct {
		name   string
		inStmt string
		want   []T
	}{
		{
			name:   "none",
			inStmt: `SELECT * FROM (SELECT 'A' AS "name", 1 AS "age") WHERE FALSE`,
			want:   []T{},
		},
		{
			name:   "one",
			inStmt: `SELECT 'A' AS "name", 1 AS "age"`,
			want:   []T{{Name: "A", Age: 1}},
		},
		{
			name:   "two",
			inStmt: `SELECT 'A' AS "name", 1 AS "age" UNION SELECT 'B' AS "name", 2 AS "age"`,
			want:   []T{{Name: "A", Age: 1}, {Name: "B", Age: 2}},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			dbPath := fmt.Sprintf(`test_%d_%d.sqlite`, number, time.Now().Unix())
			q, teardown := sqlite3.Setup(t, dbPath, nil)
			defer teardown()

			var got []T
			err := sqlite3.GenerateWithQuery[T](context.Background(), q, testcase.inStmt, nil, func(w *files.Writer, rows []T) error {
				got = rows
				return nil
			})

			require.Nil(t, err)
			require.ElementsMatch(t, got, testcase.want)
		})
	}
}
