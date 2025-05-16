package sqlite3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemas_BuildGraph(t *testing.T) {
	testcases := []struct {
		name string
		sut  Schemas
		want [][]int
	}{
		{
			name: "empty",
			sut:  Schemas{},
			want: [][]int{},
		},
		{
			name: "ddl_02_foreign_keys",
			sut: Schemas{
				{
					Name: "C_1",
				},
				{
					Name: "C_2",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "C_1"}},
					},
				},
				{
					Name: "C_3",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "C_2"}},
					},
				},
				{
					Name: "C_4",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "C_2"}},
					},
				},
				{
					Name: "C_5",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "C_3"}},
						{Reference: ForeignKeyReference{Table: "C_4"}},
					},
				},
			},
			want: [][]int{{}, {0}, {1}, {1}, {2, 3}},
		},
		{
			name: "ddl_03_foreign_loop_1",
			sut: Schemas{
				{
					Name: "D_1",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "D_1"}},
					},
				},
			},
			want: [][]int{{0}},
		},
		{
			name: "ddl_04_foreign_loop_2",
			sut: Schemas{
				{
					Name: "E_1",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "E_2"}},
					},
				},
				{
					Name: "E_2",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "E_1"}},
					},
				},
			},
			want: [][]int{{1}, {0}},
		},
		{
			name: "ddl_05_foreign_loop_3",
			sut: Schemas{
				{
					Name: "F_1",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "F_3"}},
					},
				},
				{
					Name: "F_2",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "F_1"}},
					},
				},
				{
					Name: "F_3",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "F_2"}},
					},
				},
			},
			want: [][]int{{2}, {0}, {1}},
		},
		{
			name: "ddl_07_foreign_loop_3",
			sut: Schemas{
				{
					Name: "F_1",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "F_3"}},
					},
				},
				{
					Name: "F_2",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "F_1"}},
					},
				},
				{
					Name: "F_3",
					ForeignKeys: []ForeignKey{
						{Reference: ForeignKeyReference{Table: "F_2"}},
					},
				},
			},
			want: [][]int{{2}, {0}, {1}},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sut.BuildGraph()
			assert.Equal(t, got.Len(), len(tt.want))
			for i := 0; i < got.Len(); i++ {
				assert.ElementsMatch(t, got.References(i), tt.want[i])
			}
		})
	}
}
