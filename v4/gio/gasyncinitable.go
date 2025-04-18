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

// Provides an interface for asynchronous initializing object such that
// initialization may fail.
type AsyncInitableIface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *AsyncInitableIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// This is the asynchronous version of #GInitable; it behaves the same
// in all ways except that initialization is asynchronous. For more details
// see the descriptions on #GInitable.
//
// A class may implement both the #GInitable and #GAsyncInitable interfaces.
//
// Users of objects implementing this are not intended to use the interface
// method directly; instead it will be used automatically in various ways.
// For C applications you generally just call g_async_initable_new_async()
// directly, or indirectly via a foo_thing_new_async() wrapper. This will call
// g_async_initable_init_async() under the cover, calling back with %NULL and
// a set %GError on failure.
//
// A typical implementation might look something like this:
//
// |[&lt;!-- language="C" --&gt;
//
//	enum {
//	   NOT_INITIALIZED,
//	   INITIALIZING,
//	   INITIALIZED
//	};
//
// static void
// _foo_ready_cb (Foo *self)
//
//	{
//	  GList *l;
//
//	  self-&gt;priv-&gt;state = INITIALIZED;
//
//	  for (l = self-&gt;priv-&gt;init_results; l != NULL; l = l-&gt;next)
//	    {
//	      GTask *task = l-&gt;data;
//
//	      if (self-&gt;priv-&gt;success)
//	        g_task_return_boolean (task, TRUE);
//	      else
//	        g_task_return_new_error (task, ...);
//	      g_object_unref (task);
//	    }
//
//	  g_list_free (self-&gt;priv-&gt;init_results);
//	  self-&gt;priv-&gt;init_results = NULL;
//	}
//
// static void
// foo_init_async (GAsyncInitable       *initable,
//
//	int                   io_priority,
//	GCancellable         *cancellable,
//	GAsyncReadyCallback   callback,
//	gpointer              user_data)
//
//	{
//	  Foo *self = FOO (initable);
//	  GTask *task;
//
//	  task = g_task_new (initable, cancellable, callback, user_data);
//	  g_task_set_name (task, G_STRFUNC);
//
//	  switch (self-&gt;priv-&gt;state)
//	    {
//	      case NOT_INITIALIZED:
//	        _foo_get_ready (self);
//	        self-&gt;priv-&gt;init_results = g_list_append (self-&gt;priv-&gt;init_results,
//	                                                  task);
//	        self-&gt;priv-&gt;state = INITIALIZING;
//	        break;
//	      case INITIALIZING:
//	        self-&gt;priv-&gt;init_results = g_list_append (self-&gt;priv-&gt;init_results,
//	                                                  task);
//	        break;
//	      case INITIALIZED:
//	        if (!self-&gt;priv-&gt;success)
//	          g_task_return_new_error (task, ...);
//	        else
//	          g_task_return_boolean (task, TRUE);
//	        g_object_unref (task);
//	        break;
//	    }
//	}
//
// static gboolean
// foo_init_finish (GAsyncInitable       *initable,
//
//	GAsyncResult         *result,
//	GError              **error)
//
//	{
//	  g_return_val_if_fail (g_task_is_valid (result, initable), FALSE);
//
//	  return g_task_propagate_boolean (G_TASK (result), error);
//	}
//
// static void
// foo_async_initable_iface_init (gpointer g_iface,
//
//	gpointer data)
//
//	{
//	  GAsyncInitableIface *iface = g_iface;
//
//	  iface-&gt;init_async = foo_init_async;
//	  iface-&gt;init_finish = foo_init_finish;
//	}
//
// ]|
type AsyncInitable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	InitAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr)
	InitFinish(ResVar AsyncResult) bool
	NewFinish(ResVar AsyncResult) *gobject.Object
}

var xAsyncInitableGLibType func() types.GType

func AsyncInitableGLibType() types.GType {
	return xAsyncInitableGLibType()
}

type AsyncInitableBase struct {
	Ptr uintptr
}

func (x *AsyncInitableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *AsyncInitableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Starts asynchronous initialization of the object implementing the
// interface. This must be done before any real use of the object after
// initial construction. If the object also implements #GInitable you can
// optionally call g_initable_init() instead.
//
// This method is intended for language bindings. If writing in C,
// g_async_initable_new_async() should typically be used instead.
//
// When the initialization is finished, @callback will be called. You can
// then call g_async_initable_init_finish() to get the result of the
// initialization.
//
// Implementations may also support cancellation. If @cancellable is not
// %NULL, then initialization can be cancelled by triggering the cancellable
// object from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. If @cancellable is not %NULL, and
// the object doesn't support cancellable initialization, the error
// %G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// As with #GInitable, if the object is not initialized, or initialization
// returns with an error, then all operations on the object except
// g_object_ref() and g_object_unref() are considered to be invalid, and
// have undefined behaviour. They will often fail with g_critical() or
// g_warning(), but this must not be relied on.
//
// Callers should not assume that a class which implements #GAsyncInitable can
// be initialized multiple times; for more information, see g_initable_init().
// If a class explicitly supports being initialized multiple times,
// implementation requires yielding all subsequent calls to init_async() on the
// results of the first call.
//
// For classes that also support the #GInitable interface, the default
// implementation of this method will run the g_initable_init() function
// in a thread, so if you want to support asynchronous initialization via
// threads, just implement the #GAsyncInitable interface without overriding
// any interface methods.
func (x *AsyncInitableBase) InitAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGAsyncInitableInitAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Finishes asynchronous initialization and returns the result.
// See g_async_initable_init_async().
func (x *AsyncInitableBase) InitFinish(ResVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGAsyncInitableInitFinish(x.GoPointer(), ResVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Finishes the async construction for the various g_async_initable_new
// calls, returning the created object or %NULL on error.
func (x *AsyncInitableBase) NewFinish(ResVar AsyncResult) (*gobject.Object, error) {
	var cls *gobject.Object
	var cerr *glib.Error

	cret := XGAsyncInitableNewFinish(x.GoPointer(), ResVar.GoPointer(), &cerr)

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

var XGAsyncInitableInitAsync func(uintptr, int, uintptr, uintptr, uintptr)
var XGAsyncInitableInitFinish func(uintptr, uintptr, **glib.Error) bool
var XGAsyncInitableNewFinish func(uintptr, uintptr, **glib.Error) uintptr

var xAsyncInitableNewvAsync func(types.GType, uint, *gobject.Parameter, int, uintptr, uintptr, uintptr)

// Helper function for constructing #GAsyncInitable object. This is
// similar to g_object_newv() but also initializes the object asynchronously.
//
// When the initialization is finished, @callback will be called. You can
// then call g_async_initable_new_finish() to get the new object and check
// for any errors.
func AsyncInitableNewvAsync(ObjectTypeVar types.GType, NParametersVar uint, ParametersVar *gobject.Parameter, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	xAsyncInitableNewvAsync(ObjectTypeVar, NParametersVar, ParametersVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAsyncInitableNewvAsync, lib, "g_async_initable_newv_async")

	core.PuregoSafeRegister(&xAsyncInitableGLibType, lib, "g_async_initable_get_type")

	core.PuregoSafeRegister(&XGAsyncInitableInitAsync, lib, "g_async_initable_init_async")
	core.PuregoSafeRegister(&XGAsyncInitableInitFinish, lib, "g_async_initable_init_finish")
	core.PuregoSafeRegister(&XGAsyncInitableNewFinish, lib, "g_async_initable_new_finish")

}
