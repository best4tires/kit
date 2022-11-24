package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/best4tires/kit/log/entry"
)

// Writer implements the log.Writer interface.
// It performs logging into a single log.file
type Writer struct {
	file io.WriteCloser
}

// NewWriter creates a new file.Writer
func NewWriter(path string) (*Writer, error) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Writer{
		file: f,
	}, nil
}

// Close closes the associated file
func (w *Writer) Close() {
	w.file.Close()
}

// Write writes an entry to the file
func (w *Writer) Write(e entry.Entry) {
	fmt.Fprintf(w.file, "%s [%s] [%s] [%s] %s\n", e.Time.Format("2006-01-02T15:04:05.000"), e.Program, e.Component, e.Level, e.Message)
}
