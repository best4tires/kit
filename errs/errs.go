package errs

import "fmt"

var (
	errNotFound error = fmt.Errorf("not found")
)

func NotFound() error {
	return errNotFound
}
