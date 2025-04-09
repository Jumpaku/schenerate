package spanner_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/schenerate/spanner"
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
				list_ddl00AllTypes,
				list_ddl01Interleave,
				list_ddl02ForeignKeys,
				list_ddl03ForeignLoop1,
				list_ddl04ForeignLoop2,
				list_ddl05ForeignLoop3,
				list_ddl06UniqueKeys,
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

//go:embed testdata/list/ddl_00_all_types.sql
var list_ddl00AllTypes string

//go:embed testdata/list/ddl_01_interleave.sql
var list_ddl01Interleave string

//go:embed testdata/list/ddl_02_foreign_keys.sql
var list_ddl02ForeignKeys string

//go:embed testdata/list/ddl_03_foreign_loop_1.sql
var list_ddl03ForeignLoop1 string

//go:embed testdata/list/ddl_04_foreign_loop_2.sql
var list_ddl04ForeignLoop2 string

//go:embed testdata/list/ddl_05_foreign_loop_3.sql
var list_ddl05ForeignLoop3 string

//go:embed testdata/list/ddl_06_unique_keys.sql
var list_ddl06UniqueKeys string
