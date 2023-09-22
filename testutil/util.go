package testutil

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, want, have any) {
	t.Helper()

	if !reflect.DeepEqual(want, have) {
		t.Errorf("want %v, have %v", want, have)
	}
}

func AssertNoErr(t *testing.T, err error, msg string, args ...any) {
	t.Helper()

	if err != nil {
		t.Errorf("assert-no-err: %s: %v", fmt.Sprintf(msg, args...), err)
	}
}

func AssertErr(t *testing.T, err error, msg string, args ...any) {
	t.Helper()

	if err == nil {
		t.Errorf("assert-err: %s: %v", fmt.Sprintf(msg, args...), err)
	}
}
