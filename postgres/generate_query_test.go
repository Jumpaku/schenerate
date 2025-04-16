package postgres_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
	"github.com/Jumpaku/schenerate/postgres"
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
			dbPath := fmt.Sprintf(`gwq_%d_%d`, number, time.Now().Unix())
			q, teardown := postgres.Setup(t, dbPath, nil)
			defer teardown()

			var got []T
			err := postgres.GenerateWithQuery(
				context.Background(),
				q,
				testcase.inStmt,
				nil,
				func(_ *files.Writer, rows []T) error {
					got = rows
					return nil
				})
			require.Nil(t, err)
			require.ElementsMatch(t, got, testcase.want)
		})
	}
}
