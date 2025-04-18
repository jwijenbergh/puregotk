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

type FileInputStreamClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *FileInputStreamClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type FileInputStreamPrivate struct {
	_ structs.HostLayout
}

func (x *FileInputStreamPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// GFileInputStream provides input streams that take their
// content from a file.
//
// GFileInputStream implements #GSeekable, which allows the input
// stream to jump to arbitrary positions in the file, provided the
// filesystem of the file allows it. To find the position of a file
// input stream, use g_seekable_tell(). To find out if a file input
// stream supports seeking, use g_seekable_can_seek().
// To position a file input stream, use g_seekable_seek().
type FileInputStream struct {
	InputStream
}

var xFileInputStreamGLibType func() types.GType

func FileInputStreamGLibType() types.GType {
	return xFileInputStreamGLibType()
}

func FileInputStreamNewFromInternalPtr(ptr uintptr) *FileInputStream {
	cls := &FileInputStream{}
	cls.Ptr = ptr
	return cls
}

var xFileInputStreamQueryInfo func(uintptr, string, uintptr, **glib.Error) uintptr

// Queries a file input stream the given @attributes. This function blocks
// while querying the stream. For the asynchronous (non-blocking) version
// of this function, see g_file_input_stream_query_info_async(). While the
// stream is blocked, the stream will set the pending flag internally, and
// any other operations on the stream will fail with %G_IO_ERROR_PENDING.
func (x *FileInputStream) QueryInfo(AttributesVar string, CancellableVar *Cancellable) (*FileInfo, error) {
	var cls *FileInfo
	var cerr *glib.Error

	cret := xFileInputStreamQueryInfo(x.GoPointer(), AttributesVar, CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &FileInfo{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xFileInputStreamQueryInfoAsync func(uintptr, string, int, uintptr, uintptr, uintptr)

// Queries the stream information asynchronously.
// When the operation is finished @callback will be called.
// You can then call g_file_input_stream_query_info_finish()
// to get the result of the operation.
//
// For the synchronous version of this function,
// see g_file_input_stream_query_info().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be set
func (x *FileInputStream) QueryInfoAsync(AttributesVar string, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	xFileInputStreamQueryInfoAsync(x.GoPointer(), AttributesVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xFileInputStreamQueryInfoFinish func(uintptr, uintptr, **glib.Error) uintptr

// Finishes an asynchronous info query operation.
func (x *FileInputStream) QueryInfoFinish(ResultVar AsyncResult) (*FileInfo, error) {
	var cls *FileInfo
	var cerr *glib.Error

	cret := xFileInputStreamQueryInfoFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &FileInfo{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

func (c *FileInputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *FileInputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Tests if the stream supports the #GSeekableIface.
func (x *FileInputStream) CanSeek() bool {

	cret := XGSeekableCanSeek(x.GoPointer())
	return cret
}

// Tests if the length of the stream can be adjusted with
// g_seekable_truncate().
func (x *FileInputStream) CanTruncate() bool {

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
func (x *FileInputStream) Seek(OffsetVar int64, TypeVar glib.SeekType, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGSeekableSeek(x.GoPointer(), OffsetVar, TypeVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Tells the current position within the stream.
func (x *FileInputStream) Tell() int64 {

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
func (x *FileInputStream) Truncate(OffsetVar int64, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGSeekableTruncate(x.GoPointer(), OffsetVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFileInputStreamGLibType, lib, "g_file_input_stream_get_type")

	core.PuregoSafeRegister(&xFileInputStreamQueryInfo, lib, "g_file_input_stream_query_info")
	core.PuregoSafeRegister(&xFileInputStreamQueryInfoAsync, lib, "g_file_input_stream_query_info_async")
	core.PuregoSafeRegister(&xFileInputStreamQueryInfoFinish, lib, "g_file_input_stream_query_info_finish")

}
