package main

import (
	"fmt"
	"os"
	"time"

	"github.com/best4tires/kit/log"
	"github.com/best4tires/kit/log/console"
	"github.com/best4tires/kit/log/entry"
	"github.com/best4tires/kit/log/file"
)

func createLogger() *log.DefaultLogger {
	w := console.NewWriter(
		console.WithFormatter(console.ColorFormatter{}),
		console.WithStream(os.Stderr),
	)

	ew, err := file.NewWriter("error.log")
	if err != nil {
		panic(err)
	}
	ef := log.NewFilter(
		func(e entry.Entry) bool {
			return e.Level == entry.LevelError ||
				e.Level == entry.LevelFatal
		},
		ew,
	)

	return log.NewDefaultLogger(
		"simple-example",
		log.NewMultiWriter(w, ef),
	)
}

func main() {
	l := createLogger()
	log.Install(l)
	defer l.Close()

	log.Debugf("a simple debug message %d", 42)
	log.Infof("a simple info message %q", "dude")
	log.Warnf("a simple warn message %T", true)
	log.Errorf("a simple error message %f", 1.23)
	defer log.Fatalf("finally ... %v", fmt.Errorf("an expected error occurred"))

	c := newComponent("comp-x")
	c.do()
}

type component struct {
	name string
	*log.Hook
}

func newComponent(name string) *component {
	c := &component{
		name: name,
		Hook: log.ComponentHook(name),
	}
	return c
}

func (c *component) do() {
	c.Infof("did stuff")
	c.Warnf("go to work in %s", 5*time.Minute)
	c.Errorf("go to work failed ...")
}
