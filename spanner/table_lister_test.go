package spanner_test

import (
	"context"
	_ "embed"
	"fmt"
	sqlgogen "github.com/Jumpaku/sql-gogen-lib"
	"github.com/Jumpaku/sql-gogen-lib/spanner"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTableLister_List(t *testing.T) {
	testcases := []struct {
		name string
		ddls []string
		want []sqlgogen.Table
	}{
		{
			name: "empty",
			want: []sqlgogen.Table{},
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
			want: []sqlgogen.Table{
				{Catalog: "", Schema: "", Name: "A"},
				{Catalog: "", Schema: "", Name: "B_1"},
				{Catalog: "", Schema: "", Name: "B_2"},
				{Catalog: "", Schema: "", Name: "B_3"},
				{Catalog: "", Schema: "", Name: "B_4"},
				{Catalog: "", Schema: "", Name: "C_1"},
				{Catalog: "", Schema: "", Name: "C_2"},
				{Catalog: "", Schema: "", Name: "C_3"},
				{Catalog: "", Schema: "", Name: "C_4"},
				{Catalog: "", Schema: "", Name: "C_5"},
				{Catalog: "", Schema: "", Name: "D_1"},
				{Catalog: "", Schema: "", Name: "E_1"},
				{Catalog: "", Schema: "", Name: "E_2"},
				{Catalog: "", Schema: "", Name: "F_1"},
				{Catalog: "", Schema: "", Name: "F_2"},
				{Catalog: "", Schema: "", Name: "F_3"},
				{Catalog: "", Schema: "", Name: "G"},
			},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			database := fmt.Sprintf(`tablelister_%d_%d`, number, time.Now().Unix())
			q, teardown := spanner.Setup(t, database, testcase.ddls)
			defer teardown()

			sut := spanner.NewTableLister(q)

			got, err := sut.List(context.Background())

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
