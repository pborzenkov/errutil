package errutil

// First returns first non-nil error out of errs, or nil.
func First(errs ...error) error {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}
