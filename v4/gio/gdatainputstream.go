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

type DataInputStreamClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *DataInputStreamClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type DataInputStreamPrivate struct {
	_ structs.HostLayout
}

func (x *DataInputStreamPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Data input stream implements #GInputStream and includes functions for
// reading structured data directly from a binary input stream.
type DataInputStream struct {
	BufferedInputStream
}

var xDataInputStreamGLibType func() types.GType

func DataInputStreamGLibType() types.GType {
	return xDataInputStreamGLibType()
}

func DataInputStreamNewFromInternalPtr(ptr uintptr) *DataInputStream {
	cls := &DataInputStream{}
	cls.Ptr = ptr
	return cls
}

var xNewDataInputStream func(uintptr) uintptr

// Creates a new data input stream for the @base_stream.
func NewDataInputStream(BaseStreamVar *InputStream) *DataInputStream {
	var cls *DataInputStream

	cret := xNewDataInputStream(BaseStreamVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &DataInputStream{}
	cls.Ptr = cret
	return cls
}

var xDataInputStreamGetByteOrder func(uintptr) DataStreamByteOrder

// Gets the byte order for the data input stream.
func (x *DataInputStream) GetByteOrder() DataStreamByteOrder {

	cret := xDataInputStreamGetByteOrder(x.GoPointer())
	return cret
}

var xDataInputStreamGetNewlineType func(uintptr) DataStreamNewlineType

// Gets the current newline type for the @stream.
func (x *DataInputStream) GetNewlineType() DataStreamNewlineType {

	cret := xDataInputStreamGetNewlineType(x.GoPointer())
	return cret
}

var xDataInputStreamReadByte func(uintptr, uintptr, **glib.Error) byte

// Reads an unsigned 8-bit/1-byte value from @stream.
func (x *DataInputStream) ReadByte(CancellableVar *Cancellable) (byte, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadByte(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadInt16 func(uintptr, uintptr, **glib.Error) int16

// Reads a 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
func (x *DataInputStream) ReadInt16(CancellableVar *Cancellable) (int16, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadInt16(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadInt32 func(uintptr, uintptr, **glib.Error) int32

// Reads a signed 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *DataInputStream) ReadInt32(CancellableVar *Cancellable) (int32, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadInt32(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadInt64 func(uintptr, uintptr, **glib.Error) int64

// Reads a 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *DataInputStream) ReadInt64(CancellableVar *Cancellable) (int64, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadInt64(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadLine func(uintptr, uint, uintptr, **glib.Error) []byte

// Reads a line from the data input stream.  Note that no encoding
// checks or conversion is performed; the input is not guaranteed to
// be UTF-8, and may in fact have embedded NUL characters.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *DataInputStream) ReadLine(LengthVar uint, CancellableVar *Cancellable) ([]byte, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadLine(x.GoPointer(), LengthVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadLineAsync func(uintptr, int, uintptr, uintptr, uintptr)

// The asynchronous version of g_data_input_stream_read_line().  It is
// an error to have two outstanding calls to this function.
//
// When the operation is finished, @callback will be called. You
// can then call g_data_input_stream_read_line_finish() to get
// the result of the operation.
func (x *DataInputStream) ReadLineAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	xDataInputStreamReadLineAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xDataInputStreamReadLineFinish func(uintptr, uintptr, uint, **glib.Error) []byte

// Finish an asynchronous call started by
// g_data_input_stream_read_line_async().  Note the warning about
// string encoding in g_data_input_stream_read_line() applies here as
// well.
func (x *DataInputStream) ReadLineFinish(ResultVar AsyncResult, LengthVar uint) ([]byte, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadLineFinish(x.GoPointer(), ResultVar.GoPointer(), LengthVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadLineFinishUtf8 func(uintptr, uintptr, uint, **glib.Error) string

// Finish an asynchronous call started by
// g_data_input_stream_read_line_async().
func (x *DataInputStream) ReadLineFinishUtf8(ResultVar AsyncResult, LengthVar uint) (string, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadLineFinishUtf8(x.GoPointer(), ResultVar.GoPointer(), LengthVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadLineUtf8 func(uintptr, uint, uintptr, **glib.Error) string

// Reads a UTF-8 encoded line from the data input stream.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *DataInputStream) ReadLineUtf8(LengthVar uint, CancellableVar *Cancellable) (string, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadLineUtf8(x.GoPointer(), LengthVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadUint16 func(uintptr, uintptr, **glib.Error) uint16

// Reads an unsigned 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
func (x *DataInputStream) ReadUint16(CancellableVar *Cancellable) (uint16, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadUint16(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadUint32 func(uintptr, uintptr, **glib.Error) uint32

// Reads an unsigned 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order() and g_data_input_stream_set_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *DataInputStream) ReadUint32(CancellableVar *Cancellable) (uint32, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadUint32(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadUint64 func(uintptr, uintptr, **glib.Error) uint64

// Reads an unsigned 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation,
// see g_data_input_stream_get_byte_order().
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *DataInputStream) ReadUint64(CancellableVar *Cancellable) (uint64, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadUint64(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadUntil func(uintptr, string, uint, uintptr, **glib.Error) string

// Reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// Note that, in contrast to g_data_input_stream_read_until_async(),
// this function consumes the stop character that it finds.
//
// Don't use this function in new code.  Its functionality is
// inconsistent with g_data_input_stream_read_until_async().  Both
// functions will be marked as deprecated in a future release.  Use
// g_data_input_stream_read_upto() instead, but note that that function
// does not consume the stop character.
func (x *DataInputStream) ReadUntil(StopCharsVar string, LengthVar uint, CancellableVar *Cancellable) (string, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadUntil(x.GoPointer(), StopCharsVar, LengthVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadUntilAsync func(uintptr, string, int, uintptr, uintptr, uintptr)

// The asynchronous version of g_data_input_stream_read_until().
// It is an error to have two outstanding calls to this function.
//
// Note that, in contrast to g_data_input_stream_read_until(),
// this function does not consume the stop character that it finds.  You
// must read it for yourself.
//
// When the operation is finished, @callback will be called. You
// can then call g_data_input_stream_read_until_finish() to get
// the result of the operation.
//
// Don't use this function in new code.  Its functionality is
// inconsistent with g_data_input_stream_read_until().  Both functions
// will be marked as deprecated in a future release.  Use
// g_data_input_stream_read_upto_async() instead.
func (x *DataInputStream) ReadUntilAsync(StopCharsVar string, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	xDataInputStreamReadUntilAsync(x.GoPointer(), StopCharsVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xDataInputStreamReadUntilFinish func(uintptr, uintptr, uint, **glib.Error) string

// Finish an asynchronous call started by
// g_data_input_stream_read_until_async().
func (x *DataInputStream) ReadUntilFinish(ResultVar AsyncResult, LengthVar uint) (string, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadUntilFinish(x.GoPointer(), ResultVar.GoPointer(), LengthVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadUpto func(uintptr, string, int, uint, uintptr, **glib.Error) string

// Reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// In contrast to g_data_input_stream_read_until(), this function
// does not consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is
// specified.
//
// The returned string will always be nul-terminated on success.
func (x *DataInputStream) ReadUpto(StopCharsVar string, StopCharsLenVar int, LengthVar uint, CancellableVar *Cancellable) (string, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadUpto(x.GoPointer(), StopCharsVar, StopCharsLenVar, LengthVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamReadUptoAsync func(uintptr, string, int, int, uintptr, uintptr, uintptr)

// The asynchronous version of g_data_input_stream_read_upto().
// It is an error to have two outstanding calls to this function.
//
// In contrast to g_data_input_stream_read_until(), this function
// does not consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is
// specified.
//
// When the operation is finished, @callback will be called. You
// can then call g_data_input_stream_read_upto_finish() to get
// the result of the operation.
func (x *DataInputStream) ReadUptoAsync(StopCharsVar string, StopCharsLenVar int, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	xDataInputStreamReadUptoAsync(x.GoPointer(), StopCharsVar, StopCharsLenVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xDataInputStreamReadUptoFinish func(uintptr, uintptr, uint, **glib.Error) string

// Finish an asynchronous call started by
// g_data_input_stream_read_upto_async().
//
// Note that this function does not consume the stop character. You
// have to use g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto_async() again.
//
// The returned string will always be nul-terminated on success.
func (x *DataInputStream) ReadUptoFinish(ResultVar AsyncResult, LengthVar uint) (string, error) {
	var cerr *glib.Error

	cret := xDataInputStreamReadUptoFinish(x.GoPointer(), ResultVar.GoPointer(), LengthVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDataInputStreamSetByteOrder func(uintptr, DataStreamByteOrder)

// This function sets the byte order for the given @stream. All subsequent
// reads from the @stream will be read in the given @order.
func (x *DataInputStream) SetByteOrder(OrderVar DataStreamByteOrder) {

	xDataInputStreamSetByteOrder(x.GoPointer(), OrderVar)

}

var xDataInputStreamSetNewlineType func(uintptr, DataStreamNewlineType)

// Sets the newline type for the @stream.
//
// Note that using G_DATA_STREAM_NEWLINE_TYPE_ANY is slightly unsafe. If a read
// chunk ends in "CR" we must read an additional byte to know if this is "CR" or
// "CR LF", and this might block if there is no more data available.
func (x *DataInputStream) SetNewlineType(TypeVar DataStreamNewlineType) {

	xDataInputStreamSetNewlineType(x.GoPointer(), TypeVar)

}

func (c *DataInputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *DataInputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Tests if the stream supports the #GSeekableIface.
func (x *DataInputStream) CanSeek() bool {

	cret := XGSeekableCanSeek(x.GoPointer())
	return cret
}

// Tests if the length of the stream can be adjusted with
// g_seekable_truncate().
func (x *DataInputStream) CanTruncate() bool {

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
func (x *DataInputStream) Seek(OffsetVar int64, TypeVar glib.SeekType, CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGSeekableSeek(x.GoPointer(), OffsetVar, TypeVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Tells the current position within the stream.
func (x *DataInputStream) Tell() int64 {

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
func (x *DataInputStream) Truncate(OffsetVar int64, CancellableVar *Cancellable) (bool, error) {
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

	core.PuregoSafeRegister(&xDataInputStreamGLibType, lib, "g_data_input_stream_get_type")

	core.PuregoSafeRegister(&xNewDataInputStream, lib, "g_data_input_stream_new")

	core.PuregoSafeRegister(&xDataInputStreamGetByteOrder, lib, "g_data_input_stream_get_byte_order")
	core.PuregoSafeRegister(&xDataInputStreamGetNewlineType, lib, "g_data_input_stream_get_newline_type")
	core.PuregoSafeRegister(&xDataInputStreamReadByte, lib, "g_data_input_stream_read_byte")
	core.PuregoSafeRegister(&xDataInputStreamReadInt16, lib, "g_data_input_stream_read_int16")
	core.PuregoSafeRegister(&xDataInputStreamReadInt32, lib, "g_data_input_stream_read_int32")
	core.PuregoSafeRegister(&xDataInputStreamReadInt64, lib, "g_data_input_stream_read_int64")
	core.PuregoSafeRegister(&xDataInputStreamReadLine, lib, "g_data_input_stream_read_line")
	core.PuregoSafeRegister(&xDataInputStreamReadLineAsync, lib, "g_data_input_stream_read_line_async")
	core.PuregoSafeRegister(&xDataInputStreamReadLineFinish, lib, "g_data_input_stream_read_line_finish")
	core.PuregoSafeRegister(&xDataInputStreamReadLineFinishUtf8, lib, "g_data_input_stream_read_line_finish_utf8")
	core.PuregoSafeRegister(&xDataInputStreamReadLineUtf8, lib, "g_data_input_stream_read_line_utf8")
	core.PuregoSafeRegister(&xDataInputStreamReadUint16, lib, "g_data_input_stream_read_uint16")
	core.PuregoSafeRegister(&xDataInputStreamReadUint32, lib, "g_data_input_stream_read_uint32")
	core.PuregoSafeRegister(&xDataInputStreamReadUint64, lib, "g_data_input_stream_read_uint64")
	core.PuregoSafeRegister(&xDataInputStreamReadUntil, lib, "g_data_input_stream_read_until")
	core.PuregoSafeRegister(&xDataInputStreamReadUntilAsync, lib, "g_data_input_stream_read_until_async")
	core.PuregoSafeRegister(&xDataInputStreamReadUntilFinish, lib, "g_data_input_stream_read_until_finish")
	core.PuregoSafeRegister(&xDataInputStreamReadUpto, lib, "g_data_input_stream_read_upto")
	core.PuregoSafeRegister(&xDataInputStreamReadUptoAsync, lib, "g_data_input_stream_read_upto_async")
	core.PuregoSafeRegister(&xDataInputStreamReadUptoFinish, lib, "g_data_input_stream_read_upto_finish")
	core.PuregoSafeRegister(&xDataInputStreamSetByteOrder, lib, "g_data_input_stream_set_byte_order")
	core.PuregoSafeRegister(&xDataInputStreamSetNewlineType, lib, "g_data_input_stream_set_newline_type")

}
