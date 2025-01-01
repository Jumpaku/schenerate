package postgres_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/postgres"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTableLister_List(t *testing.T) {
	testcases := []struct {
		name string
		ddls []string
		want []postgres.Table
	}{
		{
			name: "empty",
			want: []postgres.Table{},
		},
		{
			name: "tables",
			ddls: []string{
				list_ddl00AllTypes,
				list_ddl02ForeignKeys,
				list_ddl03ForeignLoop1,
				list_ddl04ForeignLoop2,
				list_ddl05ForeignLoop3,
				list_ddl07UniqueKeysConstraint,
				list_ddl08UniqueKeysColumn,
			},
			want: []postgres.Table{
				{Schema: "public", Name: "A"},
				{Schema: "public", Name: "C_1"},
				{Schema: "public", Name: "C_2"},
				{Schema: "public", Name: "C_3"},
				{Schema: "public", Name: "C_4"},
				{Schema: "public", Name: "C_5"},
				{Schema: "public", Name: "D_1"},
				{Schema: "public", Name: "E_1"},
				{Schema: "public", Name: "E_2"},
				{Schema: "public", Name: "F_1"},
				{Schema: "public", Name: "F_2"},
				{Schema: "public", Name: "F_3"},
				{Schema: "public", Name: "H"},
				{Schema: "public", Name: "I"},
			},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			database := fmt.Sprintf(`tablelister_%d_%d`, number, time.Now().Unix())
			q, teardown := postgres.Setup(t, database, testcase.ddls)
			defer teardown()

			got, err := postgres.ListTables(context.Background(), q)

			require.Nil(t, err)
			for i, w := range testcase.want {
				require.Equalf(t, w.Schema, got[i].Schema, `index=%d`, i)
				require.Equalf(t, w.Name, got[i].Name, `index=%d`, i)
			}
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

//go:embed testdata/list/ddl_07_unique_keys_constraint.sql
var list_ddl07UniqueKeysConstraint string

//go:embed testdata/list/ddl_08_unique_keys_column.sql
var list_ddl08UniqueKeysColumn string
