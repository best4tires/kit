package main

import (
	"github.com/best4tires/kit/log"
	"github.com/best4tires/kit/log/rotate"
)

func createLogger() *log.DefaultLogger {
	w, err := rotate.NewWriter(
		"logs/example.log",
		rotate.WithFileSize(2*rotate.MB),
		rotate.WithFileCount(2),
	)
	if err != nil {
		panic(err)
	}
	return log.NewDefaultLogger(
		"rotate-example",
		w,
	)
}

func main() {
	l := createLogger()
	log.Install(l)
	defer l.Close()

	log.Infof("some interesting stuff")

	for i := 0; i < 100043; i++ {
		log.Infof("this is log-entry number (%d)", i)
	}
}
