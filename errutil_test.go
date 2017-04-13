package errutil

import (
	"errors"
	"testing"
)

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
