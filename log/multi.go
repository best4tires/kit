package log

import "github.com/best4tires/kit/log/entry"

// MultiWriter implements the Writer interface
// it provides a logger middleware which writes to multiple next writers
type MultiWriter struct {
	targets []Writer
}

// NewMultiWriter creates a new MultiWriter
func NewMultiWriter(targets ...Writer) *MultiWriter {
	return &MultiWriter{
		targets: targets,
	}
}

// Close closes the writer and all related resources
func (mw *MultiWriter) Close() {
	for _, t := range mw.targets {
		t.Close()
	}
}

// Write writes an entry to all installed targets
func (mw *MultiWriter) Write(e entry.Entry) {
	for _, t := range mw.targets {
		t.Write(e)
	}
}
