package sqlite3_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/sql-gogen-lib/files"
	"github.com/Jumpaku/sql-gogen-lib/sqlite3"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSchemaProcessor_Process(t *testing.T) {
	testcases := []struct {
		name string
		ddls []string
		in   []string
		want sqlite3.Schemas
	}{
		{
			name: "all types",
			ddls: []string{schema_processor_ddl00AllTypes},
			in:   []string{"A"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "A",
					Type: "table",
					Columns: []sqlite3.Column{
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
					PrimaryKey: []string{"PK"},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_A_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK", Desc: false}}},
					},
				},
			},
		},
		{
			name: "foreign keys",
			ddls: []string{schema_processor_ddl02ForeignKeys},
			in:   []string{"C_1", "C_2", "C_3", "C_4", "C_5"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "C_1",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_C_1_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_11", Desc: false}, {Name: "PK_12", Desc: false}}},
					},
				},
				sqlite3.Schema{
					Name: "C_2",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_22", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_21", "PK_22"}, Reference: sqlite3.ForeignKeyReference{Table: "C_1", Key: []string{"PK_11", "PK_12"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_C_2_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_21", Desc: false}, {Name: "PK_22", Desc: false}}},
					},
				},
				sqlite3.Schema{
					Name: "C_3",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_31", Type: "INT64", Nullable: false},
						{Name: "PK_32", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_31", "PK_32"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_31", "PK_32"}, Reference: sqlite3.ForeignKeyReference{Table: "C_2", Key: []string{"PK_21", "PK_22"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_C_3_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_31", Desc: false}, {Name: "PK_32", Desc: false}}},
					},
				},
				sqlite3.Schema{
					Name: "C_4",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_41", Type: "INT64", Nullable: false},
						{Name: "PK_42", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_41", "PK_42"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_41", "PK_42"}, Reference: sqlite3.ForeignKeyReference{Table: "C_2", Key: []string{"PK_21", "PK_22"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_C_4_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_41", Desc: false}, {Name: "PK_42", Desc: false}}},
					},
				},
				sqlite3.Schema{
					Name: "C_5",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_51", Type: "INT64", Nullable: false},
						{Name: "PK_52", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_51", "PK_52"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_51", "PK_52"}, Reference: sqlite3.ForeignKeyReference{Table: "C_4", Key: []string{"PK_41", "PK_42"}}},
						{Key: []string{"PK_51", "PK_52"}, Reference: sqlite3.ForeignKeyReference{Table: "C_3", Key: []string{"PK_31", "PK_32"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_C_5_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_51", Desc: false}, {Name: "PK_52", Desc: false}}},
					},
				},
			},
		},
		{
			name: "foreign loop 1",
			ddls: []string{schema_processor_ddl03ForeignLoop1},
			in:   []string{"D_1"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "D_1",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_11"}, Reference: sqlite3.ForeignKeyReference{Table: "D_1", Key: []string{"PK_12"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_D_1_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_11", Desc: false}, {Name: "PK_12", Desc: false}}},
					},
				},
			},
		},
		{
			name: "foreign loop 2",
			ddls: []string{schema_processor_ddl04ForeignLoop2},
			in:   []string{"E_1", "E_2"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "E_1",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_11", "PK_12"}, Reference: sqlite3.ForeignKeyReference{Table: "E_2", Key: []string{"PK_21", "PK_22"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_E_1_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_11", Desc: false}, {Name: "PK_12", Desc: false}}},
					},
				},
				sqlite3.Schema{
					Name: "E_2",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_22", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_21", "PK_22"}, Reference: sqlite3.ForeignKeyReference{Table: "E_1", Key: []string{"PK_11", "PK_12"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_E_2_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_21", Desc: false}, {Name: "PK_22", Desc: false}}},
					},
				},
			},
		},
		{
			name: "foreign loop 3",
			ddls: []string{schema_processor_ddl05ForeignLoop3},
			in:   []string{"F_1", "F_2", "F_3"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "F_1",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_11", Type: "INT64", Nullable: false},
						{Name: "PK_12", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_11", "PK_12"}, Reference: sqlite3.ForeignKeyReference{Table: "F_3", Key: []string{"PK_31", "PK_32"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_F_1_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_11", Desc: false}, {Name: "PK_12", Desc: false}}},
					},
				},
				sqlite3.Schema{
					Name: "F_2",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_21", Type: "INT64", Nullable: false},
						{Name: "PK_22", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_21", "PK_22"}, Reference: sqlite3.ForeignKeyReference{Table: "F_1", Key: []string{"PK_11", "PK_12"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_F_2_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_21", Desc: false}, {Name: "PK_22", Desc: false}}},
					},
				},
				sqlite3.Schema{
					Name: "F_3",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK_31", Type: "INT64", Nullable: false},
						{Name: "PK_32", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK_31", "PK_32"},
					ForeignKeys: []sqlite3.ForeignKey{
						{Key: []string{"PK_31", "PK_32"}, Reference: sqlite3.ForeignKeyReference{Table: "F_2", Key: []string{"PK_21", "PK_22"}}},
					},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_F_3_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK_31", Desc: false}, {Name: "PK_32", Desc: false}}},
					},
				},
			},
		},
		{
			name: "unique keys index",
			ddls: []string{schema_processor_ddl06UniqueKeysIndex},
			in:   []string{"G"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "G",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK", Type: "INT64", Nullable: false},
						{Name: "C1", Type: "INT64", Nullable: false},
						{Name: "C2", Type: "INT64", Nullable: false},
						{Name: "C3", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK"},
					Indexes: []sqlite3.Index{
						{Name: "UQ_G_C1", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}}},
						{Name: "UQ_G_C1_C2", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C1_C2_C3", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C2", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C1_C3", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C1_C3_C2", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C3", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C2", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}}},
						{Name: "UQ_G_C2_C1", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "UQ_G_C2_C1_C3", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C1", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C2_C3", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "UQ_G_C2_C3_C1", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C3", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "UQ_G_C3", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}}},
						{Name: "UQ_G_C3_C1", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "UQ_G_C3_C1_C2", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C1", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C3_C2", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "UQ_G_C3_C2_C1", Origin: "c", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C2", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "sqlite_autoindex_G_1", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK", Desc: false}}},
					},
				},
			},
		},
		{
			name: "unique keys constraint",
			ddls: []string{schema_processor_ddl07UniqueKeysConstraint},
			in:   []string{"H"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "H",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK", Type: "INT64", Nullable: false},
						{Name: "C1", Type: "INT64", Nullable: false},
						{Name: "C2", Type: "INT64", Nullable: false},
						{Name: "C3", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK"},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_H_16", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK", Desc: false}}},
						{Name: "sqlite_autoindex_H_1", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}}},
						{Name: "sqlite_autoindex_H_10", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C2", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "sqlite_autoindex_H_11", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C3", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "sqlite_autoindex_H_12", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C3", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "sqlite_autoindex_H_13", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C1", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "sqlite_autoindex_H_14", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C1", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "sqlite_autoindex_H_15", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C2", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "sqlite_autoindex_H_2", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}}},
						{Name: "sqlite_autoindex_H_3", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}}},
						{Name: "sqlite_autoindex_H_4", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "sqlite_autoindex_H_5", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "sqlite_autoindex_H_6", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}, {Name: "C3", Desc: false}}},
						{Name: "sqlite_autoindex_H_7", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C2", Desc: false}}},
						{Name: "sqlite_autoindex_H_8", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}, {Name: "C1", Desc: false}}},
						{Name: "sqlite_autoindex_H_9", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}, {Name: "C3", Desc: false}}},
					},
				},
			},
		},
		{
			name: "unique keys column",
			ddls: []string{schema_processor_ddl08UniqueKeysColumn},
			in:   []string{"I"},
			want: sqlite3.Schemas{
				sqlite3.Schema{
					Name: "I",
					Type: "table",
					Columns: []sqlite3.Column{
						{Name: "PK", Type: "INT64", Nullable: false},
						{Name: "C1", Type: "INT64", Nullable: false},
						{Name: "C2", Type: "INT64", Nullable: false},
						{Name: "C3", Type: "INT64", Nullable: false},
					},
					PrimaryKey: []string{"PK"},
					Indexes: []sqlite3.Index{
						{Name: "sqlite_autoindex_I_4", Origin: "pk", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "PK", Desc: false}}},
						{Name: "sqlite_autoindex_I_1", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C1", Desc: false}}},
						{Name: "sqlite_autoindex_I_2", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C2", Desc: false}}},
						{Name: "sqlite_autoindex_I_3", Origin: "u", Unique: true, Key: []sqlite3.IndexKeyElem{{Name: "C3", Desc: false}}},
					},
				},
			},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			dbPath := fmt.Sprintf(`test_%d_%d.sqlite`, number, time.Now().Unix())
			q, teardown := sqlite3.Setup(t, dbPath, testcase.ddls)
			defer teardown()

			var got sqlite3.Schemas
			err := sqlite3.ProcessSchema(context.Background(), q, testcase.in, func(w *files.Writer, schemas sqlite3.Schemas) error {
				got = schemas
				return nil
			})

			require.Nil(t, err)
			equalSchema(t, testcase.want, got)
		})
	}
}

func equalSchema(t *testing.T, want, got sqlite3.Schemas) {
	t.Helper()

	require.Equal(t, want, got)
}

//go:embed testdata/schema_processor/ddl_00_all_types.sql
var schema_processor_ddl00AllTypes string

//go:embed testdata/schema_processor/ddl_02_foreign_keys.sql
var schema_processor_ddl02ForeignKeys string

//go:embed testdata/schema_processor/ddl_03_foreign_loop_1.sql
var schema_processor_ddl03ForeignLoop1 string

//go:embed testdata/schema_processor/ddl_04_foreign_loop_2.sql
var schema_processor_ddl04ForeignLoop2 string

//go:embed testdata/schema_processor/ddl_05_foreign_loop_3.sql
var schema_processor_ddl05ForeignLoop3 string

//go:embed testdata/schema_processor/ddl_06_unique_keys_index.sql
var schema_processor_ddl06UniqueKeysIndex string

//go:embed testdata/schema_processor/ddl_07_unique_keys_constraint.sql
var schema_processor_ddl07UniqueKeysConstraint string

//go:embed testdata/schema_processor/ddl_08_unique_keys_column.sql
var schema_processor_ddl08UniqueKeysColumn string
