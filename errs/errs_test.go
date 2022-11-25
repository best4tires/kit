package errs

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrs(t *testing.T) {
	err1 := NotFound()
	err2 := NotFound()
	if !errors.Is(err1, err2) {
		t.Fatalf("errors not equal")
	}

	err3 := fmt.Errorf("wrap no found: %w", err1)
	if !errors.Is(err3, err1) {
		t.Fatalf("error not wrapped")
	}
}
