package main

import (
	"fmt"

	"github.com/best4tires/kit/log"
)

func main() {
	log.Debugf("a simple debug message %d", 42)
	log.Infof("a simple info message %q", "dude")
	log.Warnf("a simple warn message %T", true)
	log.Errorf("a simple error message %f", 1.23)
	defer log.Fatalf("finally ... %v", fmt.Errorf("an expected error occurred"))
}
