package postgres

import (
	"context"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
)

type GeneratorWithQuery[RecordStruct any] func(out *files.Writer, rows []RecordStruct) error

func GenerateWithQuery[RecordStruct any](ctx context.Context, q queryer, stmt string, params []any, generator GeneratorWithQuery[RecordStruct]) error {
	rows, err := QueryRows[RecordStruct](ctx, q, stmt, params)
	if err != nil {
		return fmt.Errorf(`fail to query rows: %w`, err)
	}

	w := &files.Writer{}
	if err := generator(w, rows); err != nil {
		return fmt.Errorf(`fail to process row: %w`, err)
	}

	if err := w.SaveAll(); err != nil {
		return fmt.Errorf(`fail to save files writer: %w`, err)
	}
	return nil
}
