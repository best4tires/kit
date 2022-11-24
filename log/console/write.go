package console

import (
	"io"
	"os"

	"github.com/best4tires/kit/log/entry"
)

// Formatter defines the interface for formatting entries for console output
type Formatter interface {
	Format(entry.Entry) string
}

// Option is the console-writer option type
type Option func(w *Writer)

// WithStream configures the console writer to use the provided io.Writer
func WithStream(s io.WriteCloser) Option {
	return func(w *Writer) {
		w.writer = s
	}
}

// WithFormatter configures the console writer to use the provided Formatter
func WithFormatter(f Formatter) Option {
	return func(w *Writer) {
		w.formatter = f
	}
}

// Writer is the standard console writer and implements the log.Writer interface
type Writer struct {
	writer    io.WriteCloser
	formatter Formatter
}

// NewWriter creates a new console-writer with the provided options
func NewWriter(opts ...Option) *Writer {
	// choose formatter dependant on terminal type
	tty := false
	if soStat, err := os.Stdout.Stat(); err == nil {
		tty = (soStat.Mode()&os.ModeCharDevice == os.ModeCharDevice)
	}
	var fmtr Formatter = TextFormatter{}
	if tty {
		fmtr = ColorFormatter{}
	}

	w := &Writer{
		writer:    os.Stdout,
		formatter: fmtr,
	}
	for _, o := range opts {
		o(w)
	}
	return w
}

// Close closes the Writer
func (w *Writer) Close() {
	w.writer.Close()
}

// Write writes a log.Entry to the writer stream
func (w *Writer) Write(e entry.Entry) {
	w.writer.Write([]byte(w.formatter.Format(e) + "\n"))
}
