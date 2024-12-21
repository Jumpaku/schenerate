package spanner_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/spanner"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTableLister_List(t *testing.T) {
	testcases := []struct {
		name string
		ddls []string
		want []string
	}{
		{
			name: "empty",
			want: []string{},
		},
		{
			name: "tables",
			ddls: []string{
				table_lister_ddl00AllTypes,
				table_lister_ddl01Interleave,
				table_lister_ddl02ForeignKeys,
				table_lister_ddl03ForeignLoop1,
				table_lister_ddl04ForeignLoop2,
				table_lister_ddl05ForeignLoop3,
				table_lister_ddl06UniqueKeys,
			},
			want: []string{
				"A",
				"B_1",
				"B_2",
				"B_3",
				"B_4",
				"C_1",
				"C_2",
				"C_3",
				"C_4",
				"C_5",
				"D_1",
				"E_1",
				"E_2",
				"F_1",
				"F_2",
				"F_3",
				"G",
			},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			database := fmt.Sprintf(`tablelister_%d_%d`, number, time.Now().Unix())
			q, teardown := spanner.Setup(t, database, testcase.ddls)
			defer teardown()

			got, err := spanner.ListTables(context.Background(), q)

			require.Nil(t, err)
			require.Equal(t, testcase.want, got)
		})
	}
}

//go:embed testdata/table_lister/ddl_00_all_types.sql
var table_lister_ddl00AllTypes string

//go:embed testdata/table_lister/ddl_01_interleave.sql
var table_lister_ddl01Interleave string

//go:embed testdata/table_lister/ddl_02_foreign_keys.sql
var table_lister_ddl02ForeignKeys string

//go:embed testdata/table_lister/ddl_03_foreign_loop_1.sql
var table_lister_ddl03ForeignLoop1 string

//go:embed testdata/table_lister/ddl_04_foreign_loop_2.sql
var table_lister_ddl04ForeignLoop2 string

//go:embed testdata/table_lister/ddl_05_foreign_loop_3.sql
var table_lister_ddl05ForeignLoop3 string

//go:embed testdata/table_lister/ddl_06_unique_keys.sql
var table_lister_ddl06UniqueKeys string
