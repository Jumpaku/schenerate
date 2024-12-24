package files

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Writer struct {
	current  string
	contents map[string]*bytes.Buffer
}

var _ io.Writer = (*Writer)(nil)

func (w *Writer) Add(name string) {
	w.current = name
	w.contents[name] = bytes.NewBuffer(nil)
}

func (w *Writer) Write(b []byte) (int, error) {
	return w.contents[w.current].Write(b)
}

func (w *Writer) SaveAll() error {
	for name, content := range w.contents {
		if err := saveContent(name, content); err != nil {
			return err
		}
	}
	return nil
}

func saveContent(name string, content *bytes.Buffer) error {
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf(`fail to create file %s: %w`, name, err)
	}
	defer f.Close()
	if _, err := io.Copy(f, content); err != nil {
		return fmt.Errorf(`fail to write file %s: %w`, name, err)
	}
	return nil
}
