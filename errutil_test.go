package errutil

import (
	"errors"
	"fmt"
	"testing"
)

func recovered(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	f()
	return
}

func TestFirst(t *testing.T) {
	err1 := errors.New("first error")
	err2 := errors.New("second error")

	var tests = []struct {
		in  []error
		out error
	}{
		{[]error{nil, nil, nil, nil}, nil},
		{[]error{err1, nil, nil, nil}, err1},
		{[]error{nil, nil, err1, err2}, err1},
		{[]error{nil, err2, nil, err1}, err2},
		{[]error{}, nil},
	}

	for _, test := range tests {
		out := First(test.in...)
		if out != test.out {
			t.Errorf("First(%v) = '%v', expected '%v'", test.in, out, test.out)
		}
	}
}

func TestFatalIf(t *testing.T) {
	var tests = []struct {
		in   error
		fail bool
	}{
		{errors.New("test error"), true},
		{nil, false},
	}

	for _, test := range tests {
		err := recovered(func() { FatalIf(test.in) })
		if !test.fail {
			if err != nil {
				t.Errorf("FatalIf(%v): expected to pass, but it hasn't", test.in)
			}
			continue
		}
		if err == nil {
			t.Errorf("FatalIf(%v): expected to fail, but it hasn't", test.in)
		}
		if want, got := fmt.Sprintf("FATAL: %v", test.in), err.Error(); want != got {
			t.Errorf("FatalIf(%v): unexpected error, want = %q, got = %q", test.in, want, got)
		}
	}
}
