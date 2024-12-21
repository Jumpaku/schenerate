package sqlite3_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/sqlite3"
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
				table_lister_ddl00AllTypes,
				table_lister_ddl02ForeignKeys,
				table_lister_ddl03ForeignLoop1,
				table_lister_ddl04ForeignLoop2,
				table_lister_ddl05ForeignLoop3,
				table_lister_ddl06UniqueKeysIndex,
				table_lister_ddl07UniqueKeysConstraint,
				table_lister_ddl08UniqueKeysColumn,
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

//go:embed testdata/table_lister/ddl_00_all_types.sql
var table_lister_ddl00AllTypes string

//go:embed testdata/table_lister/ddl_02_foreign_keys.sql
var table_lister_ddl02ForeignKeys string

//go:embed testdata/table_lister/ddl_03_foreign_loop_1.sql
var table_lister_ddl03ForeignLoop1 string

//go:embed testdata/table_lister/ddl_04_foreign_loop_2.sql
var table_lister_ddl04ForeignLoop2 string

//go:embed testdata/table_lister/ddl_05_foreign_loop_3.sql
var table_lister_ddl05ForeignLoop3 string

//go:embed testdata/table_lister/ddl_06_unique_keys_index.sql
var table_lister_ddl06UniqueKeysIndex string

//go:embed testdata/table_lister/ddl_07_unique_keys_constraint.sql
var table_lister_ddl07UniqueKeysConstraint string

//go:embed testdata/table_lister/ddl_08_unique_keys_column.sql
var table_lister_ddl08UniqueKeysColumn string
