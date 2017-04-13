// +build !release,!nodebug

package errutil

import (
	"fmt"
	"testing"
)

func testFunc(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	f()

	return nil
}

func TestBug(t *testing.T) {
	format := "horrendous error %d %d"
	data := []interface{}{5, 7}

	err := testFunc(func() { Bug(format, data...) })
	if err == nil {
		t.Fatal("expected testFunc to fail, but it's passed")
	}

	if want, got := fmt.Sprintf("BUG: "+format, data...), err.Error(); want != got {
		t.Errorf("unexpected error value, want = %q, got = %q", want, got)
	}
}

func TestBugOn(t *testing.T) {
	var tests = []struct {
		cond   bool
		format string
		data   []interface{}
	}{
		{false, "this one should pass", nil},
		{true, "dreadful situation %v", []interface{}{"boo"}},
		{true, "w/o arguments", nil},
	}

	for _, test := range tests {
		err := testFunc(func() { BugOn(test.cond, test.format, test.data...) })
		if !test.cond {
			if err != nil {
				t.Fatal("expected testFunc(%v, %v, %v) to pass, but it's failed",
					test.cond, test.format, test.data)
			}
			continue
		}

		if want, got := fmt.Sprintf("BUG: "+test.format, test.data...), err.Error(); want != got {
			t.Errorf("unexpected error value, want = %q, got = %q", want, got)
		}
	}
}
