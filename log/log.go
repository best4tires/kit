package log

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/best4tires/kit/log/console"
	"github.com/best4tires/kit/log/entry"
)

// Debugf logs a debug entry to the installed logger
func Debugf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelDebug, s, args...))
}

// Infof logs an info entry to the installed logger
func Infof(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelInfo, s, args...))
}

// Warnf logs a warn entry to the installed logger
func Warnf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelWarn, s, args...))
}

// Errorf logs an error entry to the installed logger
func Errorf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelError, s, args...))
}

// Fatalf logs a fatal entry to the installed logger and panics
func Fatalf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelFatal, s, args...))
	panic(fmt.Sprintf(s, args...))
}

// Warnf logs a warn entry to the installed logger
func Importantf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelImportant, s, args...))
}

// Accessf logs a access entry to the installed logger
func Accessf(s string, args ...interface{}) {
	logger.Log(entry.Make(entry.LevelAccess, s, args...))
}

// Logger defines the general logger interface
type Logger interface {
	Log(e entry.Entry)
	Close()
}

var logger Logger = NewDefaultLogger("default", console.NewWriter())

// Install installs the provided Logger as the global logger
func Install(l Logger) {
	logger = l
}

func Close() {
	logger.Close()
}

func DebugStack() {
	dbs := debug.Stack()
	for i, l := range strings.Split(string(dbs), "\n")[7:] {
		Debugf("stack %03d: %s", i, l)
	}
}
