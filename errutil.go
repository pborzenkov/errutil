package errutil

import (
	"fmt"
)

// First returns first non-nil error out of errs, or nil.
func First(errs ...error) error {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}

// FatalIf panics if err is not nil.
func FatalIf(err error) {
	if err == nil {
		return
	}
	panic(fmt.Sprintf("FATAL: %v", err))
}
