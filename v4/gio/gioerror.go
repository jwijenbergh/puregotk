// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xIoErrorFromErrno func(int) IOErrorEnum

// Converts errno.h error codes into GIO error codes. The fallback
// value %G_IO_ERROR_FAILED is returned for error codes not currently
// handled (but note that future GLib releases may return a more
// specific value instead).
//
// As %errno is global and may be modified by intermediate function
// calls, you should save its value as soon as the call which sets it
func IoErrorFromErrno(ErrNoVar int) IOErrorEnum {

	cret := xIoErrorFromErrno(ErrNoVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xIoErrorFromErrno, lib, "g_io_error_from_errno")

}
