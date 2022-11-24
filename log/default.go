package log

import (
	"github.com/best4tires/kit/log/entry"
)

// Writer defines the general log-writer interface
type Writer interface {
	Write(e entry.Entry)
	Close()
}

// DefaultLogger is the default Logger implementation based on one Writer
type DefaultLogger struct {
	name   string
	writer Writer
}

// NewDefaultLogger creates a new StdLogger
func NewDefaultLogger(name string, w Writer) *DefaultLogger {
	l := &DefaultLogger{
		name:   name,
		writer: w,
	}
	return l
}

// Close closes the writer associated with the logger
func (l *DefaultLogger) Close() {
	l.writer.Close()
}

// Log logs writes an entry to the associated logger
func (l *DefaultLogger) Log(e entry.Entry) {
	if e.Program == "" {
		e.Program = l.name
	}
	l.writer.Write(e)
}
