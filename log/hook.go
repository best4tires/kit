package log

import (
	"fmt"

	"github.com/best4tires/kit/log/entry"
)

// Hook specifies a func, which is called before a log-entry is passed to the next writer - mostly used for amending logs entries with additional information
type Hook struct {
	hookFnc func(e entry.Entry) entry.Entry
}

// NewHook creates a new hook with the passed hook function
func NewHook(hookFnc func(e entry.Entry) entry.Entry) *Hook {
	return &Hook{
		hookFnc: hookFnc,
	}
}

// Debugf logs a debug entry
func (lh *Hook) Debugf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelDebug, s, args...)
	logger.Log(lh.hookFnc(e))
}

// Infof logs an info entry
func (lh *Hook) Infof(s string, args ...interface{}) {
	e := entry.Make(entry.LevelInfo, s, args...)
	logger.Log(lh.hookFnc(e))
}

// Warnf logs a warn entry
func (lh *Hook) Warnf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelWarn, s, args...)
	logger.Log(lh.hookFnc(e))
}

// Errorf logs an error entry
func (lh *Hook) Errorf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelError, s, args...)
	logger.Log(lh.hookFnc(e))
}

// Fatalf logs a fatla entry and panics
func (lh *Hook) Fatalf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelFatal, s, args...)
	logger.Log(lh.hookFnc(e))
	panic(fmt.Sprintf(s, args...))
}

// Importantf logs an important entry
func (lh *Hook) Importantf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelImportant, s, args...)
	logger.Log(lh.hookFnc(e))
}

// Accessf logs an access entry
func (lh *Hook) Accessf(s string, args ...interface{}) {
	e := entry.Make(entry.LevelAccess, s, args...)
	logger.Log(lh.hookFnc(e))
}

// ComponentHook creates a new hook, which modifies the component field of a log entry
func ComponentHook(comp string) *Hook {
	return NewHook(func(e entry.Entry) entry.Entry {
		e.Component = comp
		return e
	})
}
