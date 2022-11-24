package log

import "github.com/best4tires/kit/log/entry"

// Filter is a writer midlleware for a log-chain
type Filter struct {
	accept func(e entry.Entry) bool
	next   Writer
}

// NewFilter creates a new Filter
func NewFilter(accept func(e entry.Entry) bool, next Writer) *Filter {
	return &Filter{
		accept: accept,
		next:   next,
	}
}

// Close closes the Filter and all related resources
func (f *Filter) Close() {
	f.next.Close()
}

// Write writes to the filter
func (f *Filter) Write(e entry.Entry) {
	if !f.accept(e) {
		return
	}
	f.next.Write(e)
}
