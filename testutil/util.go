package testutil

import (
	"fmt"
	"reflect"
	"testing"
)

const Verbose = "verbose"

func AssertEqual(t *testing.T, want, have any, opts ...any) {
	t.Helper()

	format := "want %v, have %v"

	for _, opt := range opts {
		if opt == Verbose {
			format = "want %#v, have %#v"
		}
	}

	if !reflect.DeepEqual(want, have) {
		t.Fatalf(format, want, have)
	}
}

func AssertNoErr(t *testing.T, err error, msg string, args ...any) {
	t.Helper()

	if err != nil {
		t.Fatalf("assert-no-err: %s: %v", fmt.Sprintf(msg, args...), err)
	}
}

func AssertErr(t *testing.T, err error, msg string, args ...any) {
	t.Helper()

	if err == nil {
		t.Fatalf("assert-err: %s: %v", fmt.Sprintf(msg, args...), err)
	}
}
