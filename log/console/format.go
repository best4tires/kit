package console

import (
	"fmt"

	"github.com/best4tires/kit/log/entry"
)

type TextFormatter struct{}

func (f TextFormatter) Format(e entry.Entry) string {
	return fmt.Sprintf("%s [%s] [%s] [%s] %s", e.Time.Format("2006-01-02T15:04:05.000"), e.Program, e.Component, e.Level, e.Message)
}

const (
	fgBlack int = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

// ColorFormatter implements the "Formatter" interface nad generated colored console output
type ColorFormatter struct {
}

// Format formats an entry
func (f ColorFormatter) Format(e entry.Entry) string {
	var cr int
	switch e.Level {
	case entry.LevelDebug:
		cr = fgWhite
	case entry.LevelInfo:
		cr = fgCyan
	case entry.LevelWarn:
		cr = fgYellow
	case entry.LevelError:
		cr = fgRed
	case entry.LevelFatal:
		cr = fgMagenta
	case entry.LevelImportant:
		cr = fgBlue
	case entry.LevelAccess:
		cr = fgCyan
	default:
		cr = fgWhite
	}
	s := fmt.Sprintf("%s [%s] [%s] [%s] %s", e.Time.Format("2006-01-02T15:04:05.000"), e.Program, e.Component, e.Level, e.Message)
	return f.colored(cr, s)
}

func (f ColorFormatter) colored(cr int, s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", cr, s)
}
