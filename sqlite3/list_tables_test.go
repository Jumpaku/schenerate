package sqlite3_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/schenerate/sqlite3"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListTables(t *testing.T) {
	testcases := []struct {
		name string
		ddls []string
		want []sqlite3.Table
	}{
		{
			name: "empty",
			want: []sqlite3.Table{
				{
					Schema: "main",
					Name:   "sqlite_schema",
				},
				{
					Schema: "temp",
					Name:   "sqlite_temp_schema",
				},
			},
		},
		{
			name: "tables",
			ddls: []string{
				list_ddl00AllTypes,
				list_ddl02ForeignKeys,
				list_ddl03ForeignLoop1,
				list_ddl04ForeignLoop2,
				list_ddl05ForeignLoop3,
				list_ddl06UniqueKeysIndex,
				list_ddl07UniqueKeysConstraint,
				list_ddl08UniqueKeysColumn,
			},
			want: []sqlite3.Table{
				{Schema: "main", Name: "A"},
				{Schema: "main", Name: "C_1"},
				{Schema: "main", Name: "C_2"},
				{Schema: "main", Name: "C_3"},
				{Schema: "main", Name: "C_4"},
				{Schema: "main", Name: "C_5"},
				{Schema: "main", Name: "D_1"},
				{Schema: "main", Name: "E_1"},
				{Schema: "main", Name: "E_2"},
				{Schema: "main", Name: "F_1"},
				{Schema: "main", Name: "F_2"},
				{Schema: "main", Name: "F_3"},
				{Schema: "main", Name: "G"},
				{Schema: "main", Name: "H"},
				{Schema: "main", Name: "I"},
				{Schema: "main", Name: "sqlite_schema"},
				{Schema: "temp", Name: "sqlite_temp_schema"},
			},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			dbPath := fmt.Sprintf(`test_%d_%d.sqlite`, number, time.Now().Unix())
			q, teardown := sqlite3.Setup(t, dbPath, testcase.ddls)
			defer teardown()

			got, err := sqlite3.ListTables(context.Background(), q)

			require.Nil(t, err)
			require.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testdata/list/ddl_00_all_types.sql
var list_ddl00AllTypes string

//go:embed testdata/list/ddl_02_foreign_keys.sql
var list_ddl02ForeignKeys string

//go:embed testdata/list/ddl_03_foreign_loop_1.sql
var list_ddl03ForeignLoop1 string

//go:embed testdata/list/ddl_04_foreign_loop_2.sql
var list_ddl04ForeignLoop2 string

//go:embed testdata/list/ddl_05_foreign_loop_3.sql
var list_ddl05ForeignLoop3 string

//go:embed testdata/list/ddl_06_unique_keys_index.sql
var list_ddl06UniqueKeysIndex string

//go:embed testdata/list/ddl_07_unique_keys_constraint.sql
var list_ddl07UniqueKeysConstraint string

//go:embed testdata/list/ddl_08_unique_keys_column.sql
var list_ddl08UniqueKeysColumn string
