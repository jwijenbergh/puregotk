// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Provides an interface for implementing seekable functionality on I/O Streams.
type SeekableIface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *SeekableIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GSeekable is implemented by streams (implementations of
// #GInputStream or #GOutputStream) that support seeking.
//
// Seekable streams largely fall into two categories: resizable and
// fixed-size.
//
// #GSeekable on fixed-sized streams is approximately the same as POSIX
// lseek() on a block device (for example: attempting to seek past the
// end of the device is an error).  Fixed streams typically cannot be
// truncated.
//
// #GSeekable on resizable streams is approximately the same as POSIX
// lseek() on a normal file.  Seeking past the end and writing data will
// usually cause the stream to resize by introducing zero bytes.
type Seekable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	CanSeek() bool
	CanTruncate() bool
	Seek(OffsetVar int64, TypeVar glib.SeekType, CancellableVar *Cancellable) bool
	Tell() int64
	Truncate(OffsetVar int64, CancellableVar *Cancellable) bool
}

var xSeekableGLibType func() types.GType

func SeekableGLibType() types.GType {
	return xSeekableGLibType()
}

type SeekableBase struct {
	Ptr uintptr
}

func (x *SeekableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *SeekableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Tests if the stream supports the #GSeekableIface.
func (x *SeekableBase) CanSeek() bool {

	cret := XGSeekableCanSeek(x.GoPointer())
	return cret
}

// Tests if the length of the stream can be adjusted with
// g_seekable_truncate().
func (x *SeekableBase) CanTruncate() bool {

	cret := XGSeekableCanTruncate(x.GoPointer())
	return cret
}

// Seeks in the stream by the given @offset, modified by @type.
//
// Attempting to seek past the end of the stream will have different
// results depending on if the stream is fixed-sized or resizable.  If
// the stream is resizable then seeking past the end and then writing
// will result in zeros filling the empty space.  Seeking past the end
// of a resizable stream and reading will result in EOF.  Seeking past
// the end of a fixed-sized stream will fail.
//
// Any operation that would result in a negative offset will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *SeekableBase) Seek(OffsetVar int64, TypeVar glib.SeekType, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGSeekableSeek(x.GoPointer(), OffsetVar, TypeVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Tells the current position within the stream.
func (x *SeekableBase) Tell() int64 {

	cret := XGSeekableTell(x.GoPointer())
	return cret
}

// Sets the length of the stream to @offset. If the stream was previously
// larger than @offset, the extra data is discarded. If the stream was
// previously shorter than @offset, it is extended with NUL ('\0') bytes.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
func (x *SeekableBase) Truncate(OffsetVar int64, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGSeekableTruncate(x.GoPointer(), OffsetVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var XGSeekableCanSeek func(uintptr) bool
var XGSeekableCanTruncate func(uintptr) bool
var XGSeekableSeek func(uintptr, int64, glib.SeekType, uintptr, **glib.Error) bool
var XGSeekableTell func(uintptr) int64
var XGSeekableTruncate func(uintptr, int64, uintptr, **glib.Error) bool

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSeekableGLibType, lib, "g_seekable_get_type")

	core.PuregoSafeRegister(&XGSeekableCanSeek, lib, "g_seekable_can_seek")
	core.PuregoSafeRegister(&XGSeekableCanTruncate, lib, "g_seekable_can_truncate")
	core.PuregoSafeRegister(&XGSeekableSeek, lib, "g_seekable_seek")
	core.PuregoSafeRegister(&XGSeekableTell, lib, "g_seekable_tell")
	core.PuregoSafeRegister(&XGSeekableTruncate, lib, "g_seekable_truncate")

}
