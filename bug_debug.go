// +build !release,!nodebug

package errutil

import (
	"fmt"
)

// Bug causes panic with formatted message msg.  Does nothing when built with
// 'release' or 'nodebug' tag.
func Bug(format string, msg ...interface{}) {
	panic(fmt.Sprintf("BUG: "+format, msg...))
}

// BugOn causes panic with formatted message msg if cond is true.  Does nothing
// when built with 'release' or 'nodebug' tag.
func BugOn(cond bool, format string, msg ...interface{}) {
	if !cond {
		return
	}
	Bug(format, msg...)
}
