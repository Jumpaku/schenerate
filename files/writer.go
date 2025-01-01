package files

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type Writer struct {
	contents []struct {
		path    string
		content *bytes.Buffer
	}
}

var _ io.Writer = (*Writer)(nil)

func (w *Writer) Add(path string) {
	w.contents = append(w.contents, struct {
		path    string
		content *bytes.Buffer
	}{
		path:    path,
		content: bytes.NewBuffer(nil),
	})
}

func (w *Writer) Write(b []byte) (int, error) {
	if len(w.contents) == 0 {
		panic("file path not added")
	}
	return w.contents[len(w.contents)-1].content.Write(b)
}

func (w *Writer) SaveAll() error {
	for _, content := range w.contents {
		if err := saveContent(content.path, content.content); err != nil {
			return err
		}
	}
	return nil
}

func saveContent(name string, content *bytes.Buffer) (err error) {
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf(`fail to create file %s: %w`, name, err)
	}
	defer func(f *os.File) {
		err = errors.Join(err, f.Close())
	}(f)
	if _, e := io.Copy(f, content); e != nil {
		err = fmt.Errorf(`fail to write file %s: %w`, name, e)
		return err
	}
	return nil
}
