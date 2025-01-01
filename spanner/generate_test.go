package spanner_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/spanner"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSchemaProcessor_Process(t *testing.T) {
	testcases := []struct {
		name string
		ddls []string
		in   []string
		want spanner.Schemas
	}{
		{
			name: "all types",
			ddls: []string{generate_ddl00AllTypes},
			in:   []string{"A"},
			want: spanner.Schemas{
				spanner.Schema{
					Name:   "A",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK", Type: "INT64", Nullable: false},
						{Name: "Col_01", Type: "BOOL", Nullable: true},
						{Name: "Col_02", Type: "BOOL", Nullable: false},
						{Name: "Col_03", Type: "BYTES(50)", Nullable: true},
						{Name: "Col_04", Type: "BYTES(50)", Nullable: false},
						{Name: "Col_05", Type: "DATE", Nullable: true},
						{Name: "Col_06", Type: "DATE", Nullable: false},
						{Name: "Col_07", Type: "FLOAT64", Nullable: true},
						{Name: "Col_08", Type: "FLOAT64", Nullable: false},
						{Name: "Col_09", Type: "INT64", Nullable: true},
						{Name: "Col_10", Type: "INT64", Nullable: false},
						{Name: "Col_11", Type: "JSON", Nullable: true},
						{Name: "Col_12", Type: "JSON", Nullable: false},
						{Name: "Col_13", Type: "NUMERIC", Nullable: true},
						{Name: "Col_14", Type: "NUMERIC", Nullable: false},
						{Name: "Col_15", Type: "STRING(50)", Nullable: true},
						{Name: "Col_16", Type: "STRING(50)", Nullable: false},
						{Name: "Col_17", Type: "TIMESTAMP", Nullable: true},
						{Name: "Col_18", Type: "TIMESTAMP", Nullable: false},
					},
					PrimaryKey: []string{"PK"}}},
		},
		{
			name: "interleave",
			ddls: []string{generate_ddl01Interleave},
			in:   []string{"B_1", "B_2", "B_3", "B_4"},
			want: spanner.Schemas{
				spanner.Schema{
					Name:   "B_1",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11"},
				},
				spanner.Schema{
					Name:   "B_2",
					Type:   "BASE TABLE",
					Parent: "B_1",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_21", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_21"},
				}, spanner.Schema{
					Name:   "B_3",
					Type:   "BASE TABLE",
					Parent: "B_2",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_31", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_21", "PK_31"},
				},
				spanner.Schema{
					Name:   "B_4",
					Type:   "BASE TABLE",
					Parent: "B_2",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_41", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_21", "PK_41"},
				},
			},
		},
		{
			name: "foreign keys",
			ddls: []string{generate_ddl02ForeignKeys},
			in:   []string{"C_1", "C_2", "C_3", "C_4", "C_5"},
			want: spanner.Schemas{
				spanner.Schema{
					Name:   "C_1",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
				},
				spanner.Schema{
					Name:   "C_2",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_22", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []spanner.ForeignKey{
						{
							Name:      "FK_C_2_1",
							Key:       []string{"PK_21", "PK_22"},
							Reference: spanner.ForeignKeyReference{Table: "C_1", Key: []string{"PK_11", "PK_12"}},
						},
					},
				},
				spanner.Schema{
					Name:   "C_3",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_31", Type: "INT64", Nullable: false},
						{Name: "PK_32", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_31", "PK_32"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_C_3_2", Key: []string{"PK_31", "PK_32"}, Reference: spanner.ForeignKeyReference{Table: "C_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
				spanner.Schema{
					Name:   "C_4",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_41", Type: "INT64", Nullable: false},
						{Name: "PK_42", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_41", "PK_42"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_C_4_2", Key: []string{"PK_41", "PK_42"}, Reference: spanner.ForeignKeyReference{Table: "C_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
				spanner.Schema{
					Name:   "C_5",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_51", Type: "INT64", Nullable: false},
						{Name: "PK_52", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_51", "PK_52"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_C_5_3", Key: []string{"PK_51", "PK_52"}, Reference: spanner.ForeignKeyReference{Table: "C_3", Key: []string{"PK_31", "PK_32"}}},
						{Name: "FK_C_5_4", Key: []string{"PK_51", "PK_52"}, Reference: spanner.ForeignKeyReference{Table: "C_4", Key: []string{"PK_41", "PK_42"}}},
					},
				},
			},
		},
		{
			name: "foreign loop 1",
			ddls: []string{generate_ddl03ForeignLoop1},
			in:   []string{"D_1"},
			want: spanner.Schemas{
				spanner.Schema{
					Name:   "D_1",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_D_1_1", Key: []string{"PK_11"}, Reference: spanner.ForeignKeyReference{Table: "D_1", Key: []string{"PK_12"}}},
					},
				},
			},
		},
		{
			name: "foreign loop 2",
			ddls: []string{generate_ddl04ForeignLoop2},
			in:   []string{"E_1", "E_2"},
			want: spanner.Schemas{
				spanner.Schema{
					Name:   "E_1",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_E_1_2", Key: []string{"PK_11", "PK_12"}, Reference: spanner.ForeignKeyReference{Table: "E_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
				spanner.Schema{
					Name:   "E_2",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_22", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_E_2_1", Key: []string{"PK_21", "PK_22"}, Reference: spanner.ForeignKeyReference{Table: "E_1", Key: []string{"PK_11", "PK_12"}}},
					},
				},
			},
		},
		{
			name: "foreign loop 3",
			ddls: []string{generate_ddl05ForeignLoop3},
			in:   []string{"F_1", "F_2", "F_3"},
			want: spanner.Schemas{
				spanner.Schema{
					Name:   "F_1",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_F_1_3", Key: []string{"PK_11", "PK_12"}, Reference: spanner.ForeignKeyReference{Table: "F_3", Key: []string{"PK_31", "PK_32"}}},
					},
				},
				spanner.Schema{
					Name:   "F_2",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_22", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_F_2_1", Key: []string{"PK_21", "PK_22"}, Reference: spanner.ForeignKeyReference{Table: "F_1", Key: []string{"PK_11", "PK_12"}}},
					},
				},
				spanner.Schema{
					Name:   "F_3",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK_31", Type: "INT64", Nullable: false},
						{Name: "PK_32", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_31", "PK_32"},
					ForeignKeys: []spanner.ForeignKey{
						{Name: "FK_F_3_2", Key: []string{"PK_31", "PK_32"}, Reference: spanner.ForeignKeyReference{Table: "F_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
			},
		},
		{
			name: "unique keys index",
			ddls: []string{generate_ddl06UniqueKeys},
			in:   []string{"G"},
			want: spanner.Schemas{
				spanner.Schema{
					Name:   "G",
					Type:   "BASE TABLE",
					Parent: "",
					Columns: []spanner.Column{
						{Name: "PK", Type: "INT64", Nullable: false},
						{Name: "C1", Type: "INT64", Nullable: false},
						{Name: "C2", Type: "INT64", Nullable: false},
						{Name: "C3", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK"},
					Indexes: []spanner.Index{
						{Name: "UQ_G_C1", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C1", Desc: false}}},
						{Name: "UQ_G_C1_C2", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C1_C2_C3", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C2", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C1_C3", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C1_C3_C2", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C3", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C2", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C2", Desc: false}}},
						{Name: "UQ_G_C2_C1", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "UQ_G_C2_C1_C3", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C1", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C2_C3", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C2_C3_C1", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C3", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "UQ_G_C3", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C3", Desc: false}}},
						{Name: "UQ_G_C3_C1", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "UQ_G_C3_C1_C2", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C1", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C3_C2", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C3_C2_C1", Unique: true, Key: []spanner.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C2", Desc: false}, {Name: "C1", Desc: false}}},
					},
				},
			},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			dbPath := fmt.Sprintf(`schemaprocessor_%d_%d`, number, time.Now().Unix())
			q, teardown := spanner.Setup(t, dbPath, testcase.ddls)
			defer teardown()

			var got spanner.Schemas
			err := spanner.GenerateWithSchema(
				context.Background(),
				q,
				testcase.in,
				func(_ *files.Writer, schemas spanner.Schemas) error {
					got = schemas
					return nil
				})
			require.Nil(t, err)
			equalSchema(t, testcase.want, got)
		})
	}
}

func equalSchema(t *testing.T, want, got spanner.Schemas) {
	t.Helper()

	require.Equal(t, want, got)
}

//go:embed testdata/generate/ddl_00_all_types.sql
var generate_ddl00AllTypes string

//go:embed testdata/generate/ddl_01_interleave.sql
var generate_ddl01Interleave string

//go:embed testdata/generate/ddl_02_foreign_keys.sql
var generate_ddl02ForeignKeys string

//go:embed testdata/generate/ddl_03_foreign_loop_1.sql
var generate_ddl03ForeignLoop1 string

//go:embed testdata/generate/ddl_04_foreign_loop_2.sql
var generate_ddl04ForeignLoop2 string

//go:embed testdata/generate/ddl_05_foreign_loop_3.sql
var generate_ddl05ForeignLoop3 string

//go:embed testdata/generate/ddl_06_unique_keys.sql
var generate_ddl06UniqueKeys string
