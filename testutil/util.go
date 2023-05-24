package testutil

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, want, have any) {
	if reflect.DeepEqual(want, have) {
		return
	}
	t.Fatalf("want %v, have %v", want, have)
}

func AssertNoErr(t *testing.T, err error, msg string, args ...any) {
	if err == nil {
		return
	}
	t.Fatalf("%s: %v", fmt.Sprintf(msg, args...), err)
}