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

// The interface for pollable output streams.
//
// The default implementation of @can_poll always returns %TRUE.
//
// The default implementation of @write_nonblocking calls
// g_pollable_output_stream_is_writable(), and then calls
// g_output_stream_write() if it returns %TRUE. This means you only
// need to override it if it is possible that your @is_writable
// implementation may return %TRUE when the stream is not actually
// writable.
//
// The default implementation of @writev_nonblocking calls
// g_pollable_output_stream_write_nonblocking() for each vector, and converts
// its return value and error (if set) to a #GPollableReturn. You should
// override this where possible to avoid having to allocate a #GError to return
// %G_IO_ERROR_WOULD_BLOCK.
type PollableOutputStreamInterface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *PollableOutputStreamInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GPollableOutputStream is implemented by #GOutputStreams that
// can be polled for readiness to write. This can be used when
// interfacing with a non-GIO API that expects
// UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
type PollableOutputStream interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	CanPoll() bool
	CreateSource(CancellableVar *Cancellable) *glib.Source
	IsWritable() bool
	WriteNonblocking(BufferVar []byte, CountVar uint, CancellableVar *Cancellable) int
	WritevNonblocking(VectorsVar []OutputVector, NVectorsVar uint, BytesWrittenVar uint, CancellableVar *Cancellable) PollableReturn
}

var xPollableOutputStreamGLibType func() types.GType

func PollableOutputStreamGLibType() types.GType {
	return xPollableOutputStreamGLibType()
}

type PollableOutputStreamBase struct {
	Ptr uintptr
}

func (x *PollableOutputStreamBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *PollableOutputStreamBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Checks if @stream is actually pollable. Some classes may implement
// #GPollableOutputStream but have only certain instances of that
// class be pollable. If this method returns %FALSE, then the behavior
// of other #GPollableOutputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant;
// a stream cannot switch from pollable to non-pollable or vice versa.
func (x *PollableOutputStreamBase) CanPoll() bool {

	cret := XGPollableOutputStreamCanPoll(x.GoPointer())
	return cret
}

// Creates a #GSource that triggers when @stream can be written, or
// @cancellable is triggered or an error occurs. The callback on the
// source is of the #GPollableSourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that
// the stream may not actually be writable even after the source
// triggers, so you should use g_pollable_output_stream_write_nonblocking()
// rather than g_output_stream_write() from the callback.
func (x *PollableOutputStreamBase) CreateSource(CancellableVar *Cancellable) *glib.Source {

	cret := XGPollableOutputStreamCreateSource(x.GoPointer(), CancellableVar.GoPointer())
	return cret
}

// Checks if @stream can be written.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_output_stream_write()
// after this returns %TRUE would still block. To guarantee
// non-blocking behavior, you should always use
// g_pollable_output_stream_write_nonblocking(), which will return a
// %G_IO_ERROR_WOULD_BLOCK error rather than blocking.
func (x *PollableOutputStreamBase) IsWritable() bool {

	cret := XGPollableOutputStreamIsWritable(x.GoPointer())
	return cret
}

// Attempts to write up to @count bytes from @buffer to @stream, as
// with g_output_stream_write(). If @stream is not currently writable,
// this will immediately return %G_IO_ERROR_WOULD_BLOCK, and you can
// use g_pollable_output_stream_create_source() to create a #GSource
// that will be triggered when @stream is writable.
//
// Note that since this method never blocks, you cannot actually
// use @cancellable to cancel it. However, it will return an error
// if @cancellable has already been cancelled when you call, which
// may happen if you call this method after a source triggers due
// to having been cancelled.
//
// Also note that if %G_IO_ERROR_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same @buffer and
// @count in the next write call.
func (x *PollableOutputStreamBase) WriteNonblocking(BufferVar []byte, CountVar uint, CancellableVar *Cancellable) (int, error) {
	var cerr *glib.Error

	cret := XGPollableOutputStreamWriteNonblocking(x.GoPointer(), BufferVar, CountVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Attempts to write the bytes contained in the @n_vectors @vectors to @stream,
// as with g_output_stream_writev(). If @stream is not currently writable,
// this will immediately return %@G_POLLABLE_RETURN_WOULD_BLOCK, and you can
// use g_pollable_output_stream_create_source() to create a #GSource
// that will be triggered when @stream is writable. @error will *not* be
// set in that case.
//
// Note that since this method never blocks, you cannot actually
// use @cancellable to cancel it. However, it will return an error
// if @cancellable has already been cancelled when you call, which
// may happen if you call this method after a source triggers due
// to having been cancelled.
//
// Also note that if %G_POLLABLE_RETURN_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same @vectors and
// @n_vectors in the next write call.
func (x *PollableOutputStreamBase) WritevNonblocking(VectorsVar []OutputVector, NVectorsVar uint, BytesWrittenVar uint, CancellableVar *Cancellable) (PollableReturn, error) {
	var cerr *glib.Error

	cret := XGPollableOutputStreamWritevNonblocking(x.GoPointer(), VectorsVar, NVectorsVar, BytesWrittenVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var XGPollableOutputStreamCanPoll func(uintptr) bool
var XGPollableOutputStreamCreateSource func(uintptr, uintptr) *glib.Source
var XGPollableOutputStreamIsWritable func(uintptr) bool
var XGPollableOutputStreamWriteNonblocking func(uintptr, []byte, uint, uintptr, **glib.Error) int
var XGPollableOutputStreamWritevNonblocking func(uintptr, []OutputVector, uint, uint, uintptr, **glib.Error) PollableReturn

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xPollableOutputStreamGLibType, lib, "g_pollable_output_stream_get_type")

	core.PuregoSafeRegister(&XGPollableOutputStreamCanPoll, lib, "g_pollable_output_stream_can_poll")
	core.PuregoSafeRegister(&XGPollableOutputStreamCreateSource, lib, "g_pollable_output_stream_create_source")
	core.PuregoSafeRegister(&XGPollableOutputStreamIsWritable, lib, "g_pollable_output_stream_is_writable")
	core.PuregoSafeRegister(&XGPollableOutputStreamWriteNonblocking, lib, "g_pollable_output_stream_write_nonblocking")
	core.PuregoSafeRegister(&XGPollableOutputStreamWritevNonblocking, lib, "g_pollable_output_stream_writev_nonblocking")

}
