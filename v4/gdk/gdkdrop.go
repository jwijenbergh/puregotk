// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The `GdkDrop` object represents the target of an ongoing DND operation.
//
// Possible drop sites get informed about the status of the ongoing drag
// operation with events of type %GDK_DRAG_ENTER, %GDK_DRAG_LEAVE,
// %GDK_DRAG_MOTION and %GDK_DROP_START. The `GdkDrop` object can be obtained
// from these [class@Gdk.Event] types using [method@Gdk.DNDEvent.get_drop].
//
// The actual data transfer is initiated from the target side via an async
// read, using one of the `GdkDrop` methods for this purpose:
// [method@Gdk.Drop.read_async] or [method@Gdk.Drop.read_value_async].
//
// GTK provides a higher level abstraction based on top of these functions,
// and so they are not normally needed in GTK applications. See the
// "Drag and Drop" section of the GTK documentation for more information.
type Drop struct {
	gobject.Object
}

var xDropGLibType func() types.GType

func DropGLibType() types.GType {
	return xDropGLibType()
}

func DropNewFromInternalPtr(ptr uintptr) *Drop {
	cls := &Drop{}
	cls.Ptr = ptr
	return cls
}

var xDropFinish func(uintptr, DragAction)

// Ends the drag operation after a drop.
//
// The @action must be a single action selected from the actions
// available via [method@Gdk.Drop.get_actions].
func (x *Drop) Finish(ActionVar DragAction) {

	xDropFinish(x.GoPointer(), ActionVar)

}

var xDropGetActions func(uintptr) DragAction

// Returns the possible actions for this `GdkDrop`.
//
// If this value contains multiple actions - i.e.
// [func@Gdk.DragAction.is_unique] returns %FALSE for the result -
// [method@Gdk.Drop.finish] must choose the action to use when
// accepting the drop. This will only happen if you passed
// %GDK_ACTION_ASK as one of the possible actions in
// [method@Gdk.Drop.status]. %GDK_ACTION_ASK itself will not
// be included in the actions returned by this function.
//
// This value may change over the lifetime of the [class@Gdk.Drop]
// both as a response to source side actions as well as to calls to
// [method@Gdk.Drop.status] or [method@Gdk.Drop.finish]. The source
// side will not change this value anymore once a drop has started.
func (x *Drop) GetActions() DragAction {

	cret := xDropGetActions(x.GoPointer())
	return cret
}

var xDropGetDevice func(uintptr) uintptr

// Returns the `GdkDevice` performing the drop.
func (x *Drop) GetDevice() *Device {
	var cls *Device

	cret := xDropGetDevice(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Device{}
	cls.Ptr = cret
	return cls
}

var xDropGetDisplay func(uintptr) uintptr

// Gets the `GdkDisplay` that @self was created for.
func (x *Drop) GetDisplay() *Display {
	var cls *Display

	cret := xDropGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Display{}
	cls.Ptr = cret
	return cls
}

var xDropGetDrag func(uintptr) uintptr

// If this is an in-app drag-and-drop operation, returns the `GdkDrag`
// that corresponds to this drop.
//
// If it is not, %NULL is returned.
func (x *Drop) GetDrag() *Drag {
	var cls *Drag

	cret := xDropGetDrag(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Drag{}
	cls.Ptr = cret
	return cls
}

var xDropGetFormats func(uintptr) *ContentFormats

// Returns the `GdkContentFormats` that the drop offers the data
// to be read in.
func (x *Drop) GetFormats() *ContentFormats {

	cret := xDropGetFormats(x.GoPointer())
	return cret
}

var xDropGetSurface func(uintptr) uintptr

// Returns the `GdkSurface` performing the drop.
func (x *Drop) GetSurface() *Surface {
	var cls *Surface

	cret := xDropGetSurface(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Surface{}
	cls.Ptr = cret
	return cls
}

var xDropReadAsync func(uintptr, []string, int, uintptr, uintptr, uintptr)

// Asynchronously read the dropped data from a `GdkDrop`
// in a format that complies with one of the mime types.
func (x *Drop) ReadAsync(MimeTypesVar []string, IoPriorityVar int, CancellableVar *gio.Cancellable, CallbackVar *gio.AsyncReadyCallback, UserDataVar uintptr) {

	xDropReadAsync(x.GoPointer(), MimeTypesVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xDropReadFinish func(uintptr, uintptr, string, **glib.Error) uintptr

// Finishes an async drop read operation.
//
// Note that you must not use blocking read calls on the returned stream
// in the GTK thread, since some platforms might require communication with
// GTK to complete the data transfer. You can use async APIs such as
// g_input_stream_read_bytes_async().
//
// See [method@Gdk.Drop.read_async].
func (x *Drop) ReadFinish(ResultVar gio.AsyncResult, OutMimeTypeVar string) (*gio.InputStream, error) {
	var cls *gio.InputStream
	var cerr *glib.Error

	cret := xDropReadFinish(x.GoPointer(), ResultVar.GoPointer(), OutMimeTypeVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &gio.InputStream{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xDropReadValueAsync func(uintptr, types.GType, int, uintptr, uintptr, uintptr)

// Asynchronously request the drag operation's contents converted
// to the given @type.
//
// When the operation is finished @callback will be called. You must
// then call [method@Gdk.Drop.read_value_finish] to get the resulting
// `GValue`.
//
// For local drag-and-drop operations that are available in the given
// `GType`, the value will be copied directly. Otherwise, GDK will
// try to use [func@Gdk.content_deserialize_async] to convert the data.
func (x *Drop) ReadValueAsync(TypeVar types.GType, IoPriorityVar int, CancellableVar *gio.Cancellable, CallbackVar *gio.AsyncReadyCallback, UserDataVar uintptr) {

	xDropReadValueAsync(x.GoPointer(), TypeVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xDropReadValueFinish func(uintptr, uintptr, **glib.Error) *gobject.Value

// Finishes an async drop read.
//
// See [method@Gdk.Drop.read_value_async].
func (x *Drop) ReadValueFinish(ResultVar gio.AsyncResult) (*gobject.Value, error) {
	var cerr *glib.Error

	cret := xDropReadValueFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDropStatus func(uintptr, DragAction, DragAction)

// Selects all actions that are potentially supported by the destination.
//
// When calling this function, do not restrict the passed in actions to
// the ones provided by [method@Gdk.Drop.get_actions]. Those actions may
// change in the future, even depending on the actions you provide here.
//
// The @preferred action is a hint to the drag-and-drop mechanism about which
// action to use when multiple actions are possible.
//
// This function should be called by drag destinations in response to
// %GDK_DRAG_ENTER or %GDK_DRAG_MOTION events. If the destination does
// not yet know the exact actions it supports, it should set any possible
// actions first and then later call this function again.
func (x *Drop) Status(ActionsVar DragAction, PreferredVar DragAction) {

	xDropStatus(x.GoPointer(), ActionsVar, PreferredVar)

}

func (c *Drop) GoPointer() uintptr {
	return c.Ptr
}

func (c *Drop) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xDropGLibType, lib, "gdk_drop_get_type")

	core.PuregoSafeRegister(&xDropFinish, lib, "gdk_drop_finish")
	core.PuregoSafeRegister(&xDropGetActions, lib, "gdk_drop_get_actions")
	core.PuregoSafeRegister(&xDropGetDevice, lib, "gdk_drop_get_device")
	core.PuregoSafeRegister(&xDropGetDisplay, lib, "gdk_drop_get_display")
	core.PuregoSafeRegister(&xDropGetDrag, lib, "gdk_drop_get_drag")
	core.PuregoSafeRegister(&xDropGetFormats, lib, "gdk_drop_get_formats")
	core.PuregoSafeRegister(&xDropGetSurface, lib, "gdk_drop_get_surface")
	core.PuregoSafeRegister(&xDropReadAsync, lib, "gdk_drop_read_async")
	core.PuregoSafeRegister(&xDropReadFinish, lib, "gdk_drop_read_finish")
	core.PuregoSafeRegister(&xDropReadValueAsync, lib, "gdk_drop_read_value_async")
	core.PuregoSafeRegister(&xDropReadValueFinish, lib, "gdk_drop_read_value_finish")
	core.PuregoSafeRegister(&xDropStatus, lib, "gdk_drop_status")

}
