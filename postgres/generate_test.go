package postgres_test

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
	"github.com/Jumpaku/schenerate/postgres"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGenerateWithSchema(t *testing.T) {
	testcases := []struct {
		name string
		ddls []string
		in   []string
		want postgres.Schemas
	}{
		{
			name: "all types",
			ddls: []string{generate_ddl00AllTypes},
			in:   []string{"A"},
			want: postgres.Schemas{
				postgres.Schema{
					Schema: "public",
					Name:   "A",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK", Type: "integer", Nullable: false},
						{Name: "Col_01", Type: "bigint", Nullable: true},
						{Name: "Col_02", Type: "bigint", Nullable: false},
						{Name: "Col_04", Type: "bigint", Nullable: false},
						{Name: "Col_05", Type: "bit", Nullable: true},
						{Name: "Col_06", Type: "bit", Nullable: false},
						{Name: "Col_07", Type: "bit varying", Nullable: true},
						{Name: "Col_08", Type: "bit varying", Nullable: false},
						{Name: "Col_09", Type: "boolean", Nullable: true},
						{Name: "Col_10", Type: "boolean", Nullable: false},
						{Name: "Col_11", Type: "bytea", Nullable: true},
						{Name: "Col_12", Type: "bytea", Nullable: false},
						{Name: "Col_13", Type: "character", Nullable: true},
						{Name: "Col_14", Type: "character", Nullable: false},
						{Name: "Col_15", Type: "character varying", Nullable: true},
						{Name: "Col_16", Type: "character varying", Nullable: false},
						{Name: "Col_17", Type: "date", Nullable: true},
						{Name: "Col_18", Type: "date", Nullable: false},
						{Name: "Col_19", Type: "double precision", Nullable: true},
						{Name: "Col_20", Type: "double precision", Nullable: false},
						{Name: "Col_21", Type: "integer", Nullable: true},
						{Name: "Col_22", Type: "integer", Nullable: false},
						{Name: "Col_23", Type: "json", Nullable: true},
						{Name: "Col_24", Type: "json", Nullable: false},
						{Name: "Col_25", Type: "money", Nullable: true},
						{Name: "Col_26", Type: "money", Nullable: false},
						{Name: "Col_27", Type: "numeric", Nullable: true},
						{Name: "Col_28", Type: "numeric", Nullable: false},
						{Name: "Col_29", Type: "real", Nullable: true},
						{Name: "Col_30", Type: "real", Nullable: false},
						{Name: "Col_31", Type: "smallint", Nullable: true},
						{Name: "Col_32", Type: "smallint", Nullable: false},
						{Name: "Col_34", Type: "smallint", Nullable: false},
						{Name: "Col_36", Type: "integer", Nullable: false},
						{Name: "Col_37", Type: "text", Nullable: true},
						{Name: "Col_38", Type: "text", Nullable: false},
						{Name: "Col_39", Type: "time without time zone", Nullable: true},
						{Name: "Col_40", Type: "time without time zone", Nullable: false},
						{Name: "Col_41", Type: "time with time zone", Nullable: true},
						{Name: "Col_42", Type: "time with time zone", Nullable: false},
						{Name: "Col_43", Type: "timestamp without time zone", Nullable: true},
						{Name: "Col_44", Type: "timestamp without time zone", Nullable: false},
						{Name: "Col_45", Type: "timestamp with time zone", Nullable: true},
						{Name: "Col_46", Type: "timestamp with time zone", Nullable: false},
						{Name: "Col_47", Type: "uuid", Nullable: true},
						{Name: "Col_48", Type: "uuid", Nullable: false},
						{Name: "Col_49", Type: "xml", Nullable: true},
						{Name: "Col_50", Type: "xml", Nullable: false},
					},
					PrimaryKey: []string{"PK"},
				},
			},
		},
		{
			name: "foreign keys",
			ddls: []string{generate_ddl02ForeignKeys},
			in:   []string{"C_1", "C_2", "C_3", "C_4", "C_5"},
			want: postgres.Schemas{
				postgres.Schema{
					Schema: "public",
					Name:   "C_1",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_11", Type: "integer", Nullable: false},
						{Name: "PK_12", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
				},
				postgres.Schema{
					Schema: "public",
					Name:   "C_2",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_21", Type: "integer", Nullable: false},
						{Name: "PK_22", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_C_2_1", Key: []string{"PK_21", "PK_22"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "C_1", Key: []string{"PK_11", "PK_12"}}},
					},
				},
				postgres.Schema{
					Schema: "public",
					Name:   "C_3",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_31", Type: "integer", Nullable: false},
						{Name: "PK_32", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_31", "PK_32"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_C_3_2", Key: []string{"PK_31", "PK_32"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "C_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
				postgres.Schema{
					Schema: "public",
					Name:   "C_4",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_41", Type: "integer", Nullable: false},
						{Name: "PK_42", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_41", "PK_42"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_C_4_2", Key: []string{"PK_41", "PK_42"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "C_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
				postgres.Schema{
					Schema: "public",
					Name:   "C_5",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_51", Type: "integer", Nullable: false},
						{Name: "PK_52", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_51", "PK_52"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_C_5_3", Key: []string{"PK_51", "PK_52"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "C_3", Key: []string{"PK_31", "PK_32"}}},
						{Name: "FK_C_5_4", Key: []string{"PK_51", "PK_52"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "C_4", Key: []string{"PK_41", "PK_42"}}},
					},
				},
			},
		},
		{
			name: "foreign loop 1",
			ddls: []string{generate_ddl03ForeignLoop1},
			in:   []string{"D_1"},
			want: postgres.Schemas{
				postgres.Schema{
					Schema: "public",
					Name:   "D_1",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_11", Type: "integer", Nullable: false},
						{Name: "PK_12", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_D_1_1", Key: []string{"PK_11"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "D_1", Key: []string{"PK_12"}}},
					},
					UniqueKeys: []postgres.UniqueKey{
						{Name: "D_1_PK_12_key", Key: []string{"PK_12"}},
					},
					Indexes: []postgres.Index{
						{Name: "D_1_PK_12_key", Unique: true, Key: []string{"PK_12"}},
					},
				},
			},
		},
		{
			name: "foreign loop 2",
			ddls: []string{generate_ddl04ForeignLoop2},
			in:   []string{"E_1", "E_2"},
			want: postgres.Schemas{
				postgres.Schema{
					Schema: "public",
					Name:   "E_1",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_11", Type: "integer", Nullable: false},
						{Name: "PK_12", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_E_1_2", Key: []string{"PK_11", "PK_12"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "E_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
				postgres.Schema{
					Schema: "public",
					Name:   "E_2",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_21", Type: "integer", Nullable: false},
						{Name: "PK_22", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_E_2_1", Key: []string{"PK_21", "PK_22"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "E_1", Key: []string{"PK_11", "PK_12"}}},
					},
				},
			},
		},
		{
			name: "foreign loop 3",
			ddls: []string{generate_ddl05ForeignLoop3},
			in:   []string{"F_1", "F_2", "F_3"},
			want: postgres.Schemas{
				postgres.Schema{
					Schema: "public",
					Name:   "F_1",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_11", Type: "integer", Nullable: false},
						{Name: "PK_12", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_11", "PK_12"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_F_1_3", Key: []string{"PK_11", "PK_12"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "F_3", Key: []string{"PK_31", "PK_32"}}},
					},
				},
				postgres.Schema{
					Schema: "public",
					Name:   "F_2",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_21", Type: "integer", Nullable: false},
						{Name: "PK_22", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_21", "PK_22"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_F_2_1", Key: []string{"PK_21", "PK_22"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "F_1", Key: []string{"PK_11", "PK_12"}}},
					},
				},
				postgres.Schema{
					Schema: "public",
					Name:   "F_3",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK_31", Type: "integer", Nullable: false},
						{Name: "PK_32", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK_31", "PK_32"},
					ForeignKeys: []postgres.ForeignKey{
						{Name: "FK_F_3_2", Key: []string{"PK_31", "PK_32"}, Reference: postgres.ForeignKeyReference{Schema: "public", Table: "F_2", Key: []string{"PK_21", "PK_22"}}},
					},
				},
			},
		},
		{
			name: "unique keys constraint",
			ddls: []string{generate_ddl07UniqueKeysConstraint},
			in:   []string{"H"},
			want: postgres.Schemas{
				postgres.Schema{
					Schema: "public",
					Name:   "H",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK", Type: "integer", Nullable: false},
						{Name: "C1", Type: "integer", Nullable: false},
						{Name: "C2", Type: "integer", Nullable: false},
						{Name: "C3", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK"},
					UniqueKeys: []postgres.UniqueKey{
						{Name: "UQ_H_C1", Key: []string{"C1"}},
						{Name: "UQ_H_C1_C2", Key: []string{"C1", "C2"}},
						{Name: "UQ_H_C1_C2_C3", Key: []string{"C1", "C2", "C3"}},
						{Name: "UQ_H_C1_C3", Key: []string{"C1", "C3"}},
						{Name: "UQ_H_C1_C3_C2", Key: []string{"C1", "C3", "C2"}},
						{Name: "UQ_H_C2", Key: []string{"C2"}},
						{Name: "UQ_H_C2_C1", Key: []string{"C2", "C1"}},
						{Name: "UQ_H_C2_C1_C3", Key: []string{"C2", "C1", "C3"}},
						{Name: "UQ_H_C2_C3", Key: []string{"C2", "C3"}},
						{Name: "UQ_H_C2_C3_C1", Key: []string{"C2", "C3", "C1"}},
						{Name: "UQ_H_C3", Key: []string{"C3"}},
						{Name: "UQ_H_C3_C1", Key: []string{"C3", "C1"}},
						{Name: "UQ_H_C3_C1_C2", Key: []string{"C3", "C1", "C2"}},
						{Name: "UQ_H_C3_C2", Key: []string{"C3", "C2"}},
						{Name: "UQ_H_C3_C2_C1", Key: []string{"C3", "C2", "C1"}},
					},
					Indexes: []postgres.Index{
						{Name: "UQ_H_C1", Unique: true, Key: []string{"C1"}},
						{Name: "UQ_H_C1_C2", Unique: true, Key: []string{"C1", "C2"}},
						{Name: "UQ_H_C1_C2_C3", Unique: true, Key: []string{"C1", "C2", "C3"}},
						{Name: "UQ_H_C1_C3", Unique: true, Key: []string{"C1", "C3"}},
						{Name: "UQ_H_C1_C3_C2", Unique: true, Key: []string{"C1", "C3", "C2"}},
						{Name: "UQ_H_C2", Unique: true, Key: []string{"C2"}},
						{Name: "UQ_H_C2_C1", Unique: true, Key: []string{"C2", "C1"}},
						{Name: "UQ_H_C2_C1_C3", Unique: true, Key: []string{"C2", "C1", "C3"}},
						{Name: "UQ_H_C2_C3", Unique: true, Key: []string{"C2", "C3"}},
						{Name: "UQ_H_C2_C3_C1", Unique: true, Key: []string{"C2", "C3", "C1"}},
						{Name: "UQ_H_C3", Unique: true, Key: []string{"C3"}},
						{Name: "UQ_H_C3_C1", Unique: true, Key: []string{"C3", "C1"}},
						{Name: "UQ_H_C3_C1_C2", Unique: true, Key: []string{"C3", "C1", "C2"}},
						{Name: "UQ_H_C3_C2", Unique: true, Key: []string{"C3", "C2"}},
						{Name: "UQ_H_C3_C2_C1", Unique: true, Key: []string{"C3", "C2", "C1"}},
					},
				},
			},
		},
		{
			name: "unique keys column",
			ddls: []string{generate_ddl08UniqueKeysColumn},
			in:   []string{"I"},
			want: postgres.Schemas{
				postgres.Schema{
					Schema: "public",
					Name:   "I",
					Type:   "BASE TABLE",
					Columns: []postgres.Column{
						{Name: "PK", Type: "integer", Nullable: false},
						{Name: "C1", Type: "integer", Nullable: false},
						{Name: "C2", Type: "integer", Nullable: false},
						{Name: "C3", Type: "integer", Nullable: false},
					},
					PrimaryKey: []string{"PK"},
					UniqueKeys: []postgres.UniqueKey{
						{Name: "I_C1_key", Key: []string{"C1"}},
						{Name: "I_C2_key", Key: []string{"C2"}},
						{Name: "I_C3_key", Key: []string{"C3"}},
					},
					Indexes: []postgres.Index{
						{Name: "I_C1_key", Unique: true, Key: []string{"C1"}},
						{Name: "I_C2_key", Unique: true, Key: []string{"C2"}},
						{Name: "I_C3_key", Unique: true, Key: []string{"C3"}},
					},
				},
			},
		},
	}
	for number, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			dbPath := fmt.Sprintf(`schemaprocessor_%d_%d`, number, time.Now().Unix())
			q, teardown := postgres.Setup(t, dbPath, testcase.ddls)
			defer teardown()

			var got postgres.Schemas
			err := postgres.GenerateWithSchema(
				context.Background(),
				q,
				testcase.in,
				func(_ *files.Writer, schemas postgres.Schemas) error {
					got = schemas
					return nil
				})
			require.Nil(t, err)
			equalSchema(t, testcase.want, got)
		})
	}
}

func equalSchema(t *testing.T, want, got postgres.Schemas) {
	t.Helper()

	require.Equal(t, want, got)
}

//go:embed testdata/generate/ddl_00_all_types.sql
var generate_ddl00AllTypes string

//go:embed testdata/generate/ddl_02_foreign_keys.sql
var generate_ddl02ForeignKeys string

//go:embed testdata/generate/ddl_03_foreign_loop_1.sql
var generate_ddl03ForeignLoop1 string

//go:embed testdata/generate/ddl_04_foreign_loop_2.sql
var generate_ddl04ForeignLoop2 string

//go:embed testdata/generate/ddl_05_foreign_loop_3.sql
var generate_ddl05ForeignLoop3 string

//go:embed testdata/generate/ddl_07_unique_keys_constraint.sql
var generate_ddl07UniqueKeysConstraint string

//go:embed testdata/generate/ddl_08_unique_keys_column.sql
var generate_ddl08UniqueKeysColumn string
