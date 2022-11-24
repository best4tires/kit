package entry

import (
	"fmt"
	"time"
)

// Level defines the level of a log.Entry
type Level string

const (
	// LevelDebug denotes typical debug log entries
	LevelDebug Level = "DEBUG"
	// LevelInfo denotes just informational log entries
	LevelInfo Level = "INFO "
	// LevelWarn denotes warnings, which are not as critical as errors.
	LevelWarn Level = "WARN "
	// LevelError denotes errors on which the system keeps on running but action should be taken.
	LevelError Level = "ERROR"
	// LevelFatal denotes fatal errors in which the system should panic immediately
	LevelFatal Level = "FATAL"
	// LevelAccess denotes http-access messages
	LevelAccess Level = "ACCSS"
	// LevelImportant denotes important messages, which may be handled separately from usual INFO messages
	LevelImportant Level = "IMPNT"
)

// Entry defines a general log.Entry
type Entry struct {
	Time      time.Time
	Level     Level
	Program   string
	Component string
	Message   string
}

// Make creates a new log entry with the passed parameters
func Make(level Level, s string, args ...interface{}) Entry {
	return Entry{
		Time:    time.Now(),
		Level:   level,
		Message: fmt.Sprintf(s, args...),
	}
}
