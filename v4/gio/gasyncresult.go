// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Interface definition for #GAsyncResult.
type AsyncResultIface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *AsyncResultIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Provides a base class for implementing asynchronous function results.
//
// Asynchronous operations are broken up into two separate operations
// which are chained together by a #GAsyncReadyCallback. To begin
// an asynchronous operation, provide a #GAsyncReadyCallback to the
// asynchronous function. This callback will be triggered when the
// operation has completed, and must be run in a later iteration of
// the [thread-default main context][g-main-context-push-thread-default]
// from where the operation was initiated. It will be passed a
// #GAsyncResult instance filled with the details of the operation's
// success or failure, the object the asynchronous function was
// started for and any error codes returned. The asynchronous callback
// function is then expected to call the corresponding "_finish()"
// function, passing the object the function was called for, the
// #GAsyncResult instance, and (optionally) an @error to grab any
// error conditions that may have occurred.
//
// The "_finish()" function for an operation takes the generic result
// (of type #GAsyncResult) and returns the specific result that the
// operation in question yields (e.g. a #GFileEnumerator for a
// "enumerate children" operation). If the result or error status of the
// operation is not needed, there is no need to call the "_finish()"
// function; GIO will take care of cleaning up the result and error
// information after the #GAsyncReadyCallback returns. You can pass
// %NULL for the #GAsyncReadyCallback if you don't need to take any
// action at all after the operation completes. Applications may also
// take a reference to the #GAsyncResult and call "_finish()" later;
// however, the "_finish()" function may be called at most once.
//
// Example of a typical asynchronous operation flow:
// |[&lt;!-- language="C" --&gt;
// void _theoretical_frobnitz_async (Theoretical         *t,
//
//	GCancellable        *c,
//	GAsyncReadyCallback  cb,
//	gpointer             u);
//
// gboolean _theoretical_frobnitz_finish (Theoretical   *t,
//
//	GAsyncResult  *res,
//	GError       **e);
//
// static void
// frobnitz_result_func (GObject      *source_object,
//
//	GAsyncResult *res,
//	gpointer      user_data)
//
//	{
//	  gboolean success = FALSE;
//
//	  success = _theoretical_frobnitz_finish (source_object, res, NULL);
//
//	  if (success)
//	    g_printf ("Hurray!\n");
//	  else
//	    g_printf ("Uh oh!\n");
//
//	  ...
//
// }
//
// int main (int argc, void *argv[])
//
//	{
//	   ...
//
//	   _theoretical_frobnitz_async (theoretical_data,
//	                                NULL,
//	                                frobnitz_result_func,
//	                                NULL);
//
//	   ...
//	}
//
// ]|
//
// The callback for an asynchronous operation is called only once, and is
// always called, even in the case of a cancelled operation. On cancellation
// the result is a %G_IO_ERROR_CANCELLED error.
//
// ## I/O Priority # {#io-priority}
//
// Many I/O-related asynchronous operations have a priority parameter,
// which is used in certain cases to determine the order in which
// operations are executed. They are not used to determine system-wide
// I/O scheduling. Priorities are integers, with lower numbers indicating
// higher priority. It is recommended to choose priorities between
// %G_PRIORITY_LOW and %G_PRIORITY_HIGH, with %G_PRIORITY_DEFAULT
// as a default.
type AsyncResult interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetSourceObject() *gobject.Object
	GetUserData() uintptr
	IsTagged(SourceTagVar uintptr) bool
	LegacyPropagateError() bool
}

var xAsyncResultGLibType func() types.GType

func AsyncResultGLibType() types.GType {
	return xAsyncResultGLibType()
}

type AsyncResultBase struct {
	Ptr uintptr
}

func (x *AsyncResultBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *AsyncResultBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the source object from a #GAsyncResult.
func (x *AsyncResultBase) GetSourceObject() *gobject.Object {
	var cls *gobject.Object

	cret := XGAsyncResultGetSourceObject(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

// Gets the user data from a #GAsyncResult.
func (x *AsyncResultBase) GetUserData() uintptr {

	cret := XGAsyncResultGetUserData(x.GoPointer())
	return cret
}

// Checks if @res has the given @source_tag (generally a function
// pointer indicating the function @res was created by).
func (x *AsyncResultBase) IsTagged(SourceTagVar uintptr) bool {

	cret := XGAsyncResultIsTagged(x.GoPointer(), SourceTagVar)
	return cret
}

// If @res is a #GSimpleAsyncResult, this is equivalent to
// g_simple_async_result_propagate_error(). Otherwise it returns
// %FALSE.
//
// This can be used for legacy error handling in async *_finish()
// wrapper functions that traditionally handled #GSimpleAsyncResult
// error returns themselves rather than calling into the virtual method.
// This should not be used in new code; #GAsyncResult errors that are
// set by virtual methods should also be extracted by virtual methods,
// to enable subclasses to chain up correctly.
func (x *AsyncResultBase) LegacyPropagateError() (bool, error) {
	var cerr *glib.Error

	cret := XGAsyncResultLegacyPropagateError(x.GoPointer())
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var XGAsyncResultGetSourceObject func(uintptr) uintptr
var XGAsyncResultGetUserData func(uintptr) uintptr
var XGAsyncResultIsTagged func(uintptr, uintptr) bool
var XGAsyncResultLegacyPropagateError func(uintptr) bool

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAsyncResultGLibType, lib, "g_async_result_get_type")

	core.PuregoSafeRegister(&XGAsyncResultGetSourceObject, lib, "g_async_result_get_source_object")
	core.PuregoSafeRegister(&XGAsyncResultGetUserData, lib, "g_async_result_get_user_data")
	core.PuregoSafeRegister(&XGAsyncResultIsTagged, lib, "g_async_result_is_tagged")
	core.PuregoSafeRegister(&XGAsyncResultLegacyPropagateError, lib, "g_async_result_legacy_propagate_error")

}
