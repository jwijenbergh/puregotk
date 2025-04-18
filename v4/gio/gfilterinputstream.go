// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

type FilterInputStreamClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *FilterInputStreamClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Base class for input stream implementations that perform some
// kind of filtering operation on a base stream. Typical examples
// of filtering operations are character set conversion, compression
// and byte order flipping.
type FilterInputStream struct {
	InputStream
}

var xFilterInputStreamGLibType func() types.GType

func FilterInputStreamGLibType() types.GType {
	return xFilterInputStreamGLibType()
}

func FilterInputStreamNewFromInternalPtr(ptr uintptr) *FilterInputStream {
	cls := &FilterInputStream{}
	cls.Ptr = ptr
	return cls
}

var xFilterInputStreamGetBaseStream func(uintptr) uintptr

// Gets the base stream for the filter stream.
func (x *FilterInputStream) GetBaseStream() *InputStream {
	var cls *InputStream

	cret := xFilterInputStreamGetBaseStream(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &InputStream{}
	cls.Ptr = cret
	return cls
}

var xFilterInputStreamGetCloseBaseStream func(uintptr) bool

// Returns whether the base stream will be closed when @stream is
// closed.
func (x *FilterInputStream) GetCloseBaseStream() bool {

	cret := xFilterInputStreamGetCloseBaseStream(x.GoPointer())
	return cret
}

var xFilterInputStreamSetCloseBaseStream func(uintptr, bool)

// Sets whether the base stream will be closed when @stream is closed.
func (x *FilterInputStream) SetCloseBaseStream(CloseBaseVar bool) {

	xFilterInputStreamSetCloseBaseStream(x.GoPointer(), CloseBaseVar)

}

func (c *FilterInputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *FilterInputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFilterInputStreamGLibType, lib, "g_filter_input_stream_get_type")

	core.PuregoSafeRegister(&xFilterInputStreamGetBaseStream, lib, "g_filter_input_stream_get_base_stream")
	core.PuregoSafeRegister(&xFilterInputStreamGetCloseBaseStream, lib, "g_filter_input_stream_get_close_base_stream")
	core.PuregoSafeRegister(&xFilterInputStreamSetCloseBaseStream, lib, "g_filter_input_stream_set_close_base_stream")

}
