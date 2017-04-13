// +build release nodebug

package errutil

func Bug(_ string, _ ...interface{})           {}
func BugOn(_ bool, _ string, _ ...interface{}) {}
