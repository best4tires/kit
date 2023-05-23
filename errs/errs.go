package errs

import "fmt"

var (
	errNotFound         error = fmt.Errorf("not found")
	errBadArgs          error = fmt.Errorf("bad arguments")
	errNotAuthenticated error = fmt.Errorf("not authenticated")
	errForbidden        error = fmt.Errorf("forbidden")
)

func NotFound() error {
	return errNotFound
}

func BadArgs() error {
	return errBadArgs
}

func NotAuthenticated() error {
	return errNotAuthenticated
}

func Forbidden() error {
	return errForbidden
}
