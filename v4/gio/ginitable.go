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

// Provides an interface for initializing object such that initialization
// may fail.
type InitableIface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *InitableIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GInitable is implemented by objects that can fail during
// initialization. If an object implements this interface then
// it must be initialized as the first thing after construction,
// either via g_initable_init() or g_async_initable_init_async()
// (the latter is only available if it also implements #GAsyncInitable).
//
// If the object is not initialized, or initialization returns with an
// error, then all operations on the object except g_object_ref() and
// g_object_unref() are considered to be invalid, and have undefined
// behaviour. They will often fail with g_critical() or g_warning(), but
// this must not be relied on.
//
// Users of objects implementing this are not intended to use
// the interface method directly, instead it will be used automatically
// in various ways. For C applications you generally just call
// g_initable_new() directly, or indirectly via a foo_thing_new() wrapper.
// This will call g_initable_init() under the cover, returning %NULL and
// setting a #GError on failure (at which point the instance is
// unreferenced).
//
// For bindings in languages where the native constructor supports
// exceptions the binding could check for objects implementing %GInitable
// during normal construction and automatically initialize them, throwing
// an exception on failure.
type Initable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	Init(CancellableVar *Cancellable) bool
}

var xInitableGLibType func() types.GType

func InitableGLibType() types.GType {
	return xInitableGLibType()
}

type InitableBase struct {
	Ptr uintptr
}

func (x *InitableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *InitableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Initializes the object implementing the interface.
//
// This method is intended for language bindings. If writing in C,
// g_initable_new() should typically be used instead.
//
// The object must be initialized before any real use after initial
// construction, either with this function or g_async_initable_init_async().
//
// Implementations may also support cancellation. If @cancellable is not %NULL,
// then initialization can be cancelled by triggering the cancellable object
// from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. If @cancellable is not %NULL and
// the object doesn't support cancellable initialization the error
// %G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// If the object is not initialized, or initialization returns with an
// error, then all operations on the object except g_object_ref() and
// g_object_unref() are considered to be invalid, and have undefined
// behaviour. See the [introduction][ginitable] for more details.
//
// Callers should not assume that a class which implements #GInitable can be
// initialized multiple times, unless the class explicitly documents itself as
// supporting this. Generally, a class’ implementation of init() can assume
// (and assert) that it will only be called once. Previously, this documentation
// recommended all #GInitable implementations should be idempotent; that
// recommendation was relaxed in GLib 2.54.
//
// If a class explicitly supports being initialized multiple times, it is
// recommended that the method is idempotent: multiple calls with the same
// arguments should return the same results. Only the first call initializes
// the object; further calls return the result of the first call.
//
// One reason why a class might need to support idempotent initialization is if
// it is designed to be used via the singleton pattern, with a
// #GObjectClass.constructor that sometimes returns an existing instance.
// In this pattern, a caller would expect to be able to call g_initable_init()
// on the result of g_object_new(), regardless of whether it is in fact a new
// instance.
func (x *InitableBase) Init(CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGInitableInit(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var XGInitableInit func(uintptr, uintptr, **glib.Error) bool

var xInitableNewv func(types.GType, uint, []gobject.Parameter, uintptr, **glib.Error) uintptr

// Helper function for constructing #GInitable object. This is
// similar to g_object_newv() but also initializes the object
// and returns %NULL, setting an error on failure.
func InitableNewv(ObjectTypeVar types.GType, NParametersVar uint, ParametersVar []gobject.Parameter, CancellableVar *Cancellable) (*gobject.Object, error) {
	var cls *gobject.Object
	var cerr *glib.Error

	cret := xInitableNewv(ObjectTypeVar, NParametersVar, ParametersVar, CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &gobject.Object{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xInitableNewv, lib, "g_initable_newv")

	core.PuregoSafeRegister(&xInitableGLibType, lib, "g_initable_get_type")

	core.PuregoSafeRegister(&XGInitableInit, lib, "g_initable_init")

}
