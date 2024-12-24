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
				table_lister_ddl00AllTypes,
				table_lister_ddl02ForeignKeys,
				table_lister_ddl03ForeignLoop1,
				table_lister_ddl04ForeignLoop2,
				table_lister_ddl05ForeignLoop3,
				table_lister_ddl07UniqueKeysConstraint,
				table_lister_ddl08UniqueKeysColumn,
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

//go:embed testdata/table_lister/ddl_07_unique_keys_constraint.sql
var table_lister_ddl07UniqueKeysConstraint string

//go:embed testdata/table_lister/ddl_08_unique_keys_column.sql
var table_lister_ddl08UniqueKeysColumn string
