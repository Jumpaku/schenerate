package postgres

import (
	"context"
	"fmt"
	"github.com/Jumpaku/schenerate/files"
)

type Generator func(out *files.Writer, schemas Schemas) error

func GenerateWithSchema(ctx context.Context, q queryer, tables []string, generator Generator) error {
	schemas, err := ListSchemas(ctx, q, tables)
	if err != nil {
		return fmt.Errorf(`fail to list schemas: %w`, err)
	}

	w := &files.Writer{}
	if err := generator(w, schemas); err != nil {
		return err
	}

	if err := w.SaveAll(); err != nil {
		return fmt.Errorf(`fail to save files writer: %w`, err)
	}
	return nil
}
