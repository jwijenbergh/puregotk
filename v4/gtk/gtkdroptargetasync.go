// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

type DropTargetAsyncClass struct {
	_ structs.HostLayout
}

func (x *DropTargetAsyncClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkDropTargetAsync` is an event controller to receive Drag-and-Drop
// operations, asynchronously.
//
// It is the more complete but also more complex method of handling drop
// operations compared to [class@Gtk.DropTarget], and you should only use
// it if `GtkDropTarget` doesn't provide all the features you need.
//
// To use a `GtkDropTargetAsync` to receive drops on a widget, you create
// a `GtkDropTargetAsync` object, configure which data formats and actions
// you support, connect to its signals, and then attach it to the widget
// with [method@Gtk.Widget.add_controller].
//
// During a drag operation, the first signal that a `GtkDropTargetAsync`
// emits is [signal@Gtk.DropTargetAsync::accept], which is meant to determine
// whether the target is a possible drop site for the ongoing drop. The
// default handler for the ::accept signal accepts the drop if it finds
// a compatible data format and an action that is supported on both sides.
//
// If it is, and the widget becomes a target, you will receive a
// [signal@Gtk.DropTargetAsync::drag-enter] signal, followed by
// [signal@Gtk.DropTargetAsync::drag-motion] signals as the pointer moves,
// optionally a [signal@Gtk.DropTargetAsync::drop] signal when a drop happens,
// and finally a [signal@Gtk.DropTargetAsync::drag-leave] signal when the
// pointer moves off the widget.
//
// The ::drag-enter and ::drag-motion handler return a `GdkDragAction`
// to update the status of the ongoing operation. The ::drop handler
// should decide if it ultimately accepts the drop and if it does, it
// should initiate the data transfer and finish the operation by calling
// [method@Gdk.Drop.finish].
//
// Between the ::drag-enter and ::drag-leave signals the widget is a
// current drop target, and will receive the %GTK_STATE_FLAG_DROP_ACTIVE
// state, which can be used by themes to style the widget as a drop target.
type DropTargetAsync struct {
	EventController
}

var xDropTargetAsyncGLibType func() types.GType

func DropTargetAsyncGLibType() types.GType {
	return xDropTargetAsyncGLibType()
}

func DropTargetAsyncNewFromInternalPtr(ptr uintptr) *DropTargetAsync {
	cls := &DropTargetAsync{}
	cls.Ptr = ptr
	return cls
}

var xNewDropTargetAsync func(*gdk.ContentFormats, gdk.DragAction) uintptr

// Creates a new `GtkDropTargetAsync` object.
func NewDropTargetAsync(FormatsVar *gdk.ContentFormats, ActionsVar gdk.DragAction) *DropTargetAsync {
	var cls *DropTargetAsync

	cret := xNewDropTargetAsync(FormatsVar, ActionsVar)

	if cret == 0 {
		return nil
	}
	cls = &DropTargetAsync{}
	cls.Ptr = cret
	return cls
}

var xDropTargetAsyncGetActions func(uintptr) gdk.DragAction

// Gets the actions that this drop target supports.
func (x *DropTargetAsync) GetActions() gdk.DragAction {

	cret := xDropTargetAsyncGetActions(x.GoPointer())
	return cret
}

var xDropTargetAsyncGetFormats func(uintptr) *gdk.ContentFormats

// Gets the data formats that this drop target accepts.
//
// If the result is %NULL, all formats are expected to be supported.
func (x *DropTargetAsync) GetFormats() *gdk.ContentFormats {

	cret := xDropTargetAsyncGetFormats(x.GoPointer())
	return cret
}

var xDropTargetAsyncRejectDrop func(uintptr, uintptr)

// Sets the @drop as not accepted on this drag site.
//
// This function should be used when delaying the decision
// on whether to accept a drag or not until after reading
// the data.
func (x *DropTargetAsync) RejectDrop(DropVar *gdk.Drop) {

	xDropTargetAsyncRejectDrop(x.GoPointer(), DropVar.GoPointer())

}

var xDropTargetAsyncSetActions func(uintptr, gdk.DragAction)

// Sets the actions that this drop target supports.
func (x *DropTargetAsync) SetActions(ActionsVar gdk.DragAction) {

	xDropTargetAsyncSetActions(x.GoPointer(), ActionsVar)

}

var xDropTargetAsyncSetFormats func(uintptr, *gdk.ContentFormats)

// Sets the data formats that this drop target will accept.
func (x *DropTargetAsync) SetFormats(FormatsVar *gdk.ContentFormats) {

	xDropTargetAsyncSetFormats(x.GoPointer(), FormatsVar)

}

func (c *DropTargetAsync) GoPointer() uintptr {
	return c.Ptr
}

func (c *DropTargetAsync) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted on the drop site when a drop operation is about to begin.
//
// If the drop is not accepted, %FALSE will be returned and the drop target
// will ignore the drop. If %TRUE is returned, the drop is accepted for now
// but may be rejected later via a call to [method@Gtk.DropTargetAsync.reject_drop]
// or ultimately by returning %FALSE from a [signal@Gtk.DropTargetAsync::drop]
// handler.
//
// The default handler for this signal decides whether to accept the drop
// based on the formats provided by the @drop.
//
// If the decision whether the drop will be accepted or rejected needs
// further processing, such as inspecting the data, this function should
// return %TRUE and proceed as is @drop was accepted and if it decides to
// reject the drop later, it should call [method@Gtk.DropTargetAsync.reject_drop].
func (x *DropTargetAsync) ConnectAccept(cb *func(DropTargetAsync, uintptr) bool) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "accept", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, DropVarp uintptr) bool {
		fa := DropTargetAsync{}
		fa.Ptr = clsPtr
		cbFn := *cb

		return cbFn(fa, DropVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "accept", cbRefPtr)
}

// Emitted on the drop site when the pointer enters the widget.
//
// It can be used to set up custom highlighting.
func (x *DropTargetAsync) ConnectDragEnter(cb *func(DropTargetAsync, uintptr, float64, float64) gdk.DragAction) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "drag-enter", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, DropVarp uintptr, XVarp float64, YVarp float64) gdk.DragAction {
		fa := DropTargetAsync{}
		fa.Ptr = clsPtr
		cbFn := *cb

		return cbFn(fa, DropVarp, XVarp, YVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "drag-enter", cbRefPtr)
}

// Emitted on the drop site when the pointer leaves the widget.
//
// Its main purpose it to undo things done in
// `GtkDropTargetAsync`::drag-enter.
func (x *DropTargetAsync) ConnectDragLeave(cb *func(DropTargetAsync, uintptr)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "drag-leave", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, DropVarp uintptr) {
		fa := DropTargetAsync{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, DropVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "drag-leave", cbRefPtr)
}

// Emitted while the pointer is moving over the drop target.
func (x *DropTargetAsync) ConnectDragMotion(cb *func(DropTargetAsync, uintptr, float64, float64) gdk.DragAction) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "drag-motion", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, DropVarp uintptr, XVarp float64, YVarp float64) gdk.DragAction {
		fa := DropTargetAsync{}
		fa.Ptr = clsPtr
		cbFn := *cb

		return cbFn(fa, DropVarp, XVarp, YVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "drag-motion", cbRefPtr)
}

// Emitted on the drop site when the user drops the data onto the widget.
//
// The signal handler must determine whether the pointer position is in a
// drop zone or not. If it is not in a drop zone, it returns %FALSE and no
// further processing is necessary.
//
// Otherwise, the handler returns %TRUE. In this case, this handler will
// accept the drop. The handler must ensure that [method@Gdk.Drop.finish]
// is called to let the source know that the drop is done. The call to
// [method@Gdk.Drop.finish] must only be done when all data has been received.
//
// To receive the data, use one of the read functions provided by
// [class@Gdk.Drop] such as [method@Gdk.Drop.read_async] or
// [method@Gdk.Drop.read_value_async].
func (x *DropTargetAsync) ConnectDrop(cb *func(DropTargetAsync, uintptr, float64, float64) bool) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "drop", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, DropVarp uintptr, XVarp float64, YVarp float64) bool {
		fa := DropTargetAsync{}
		fa.Ptr = clsPtr
		cbFn := *cb

		return cbFn(fa, DropVarp, XVarp, YVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "drop", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xDropTargetAsyncGLibType, lib, "gtk_drop_target_async_get_type")

	core.PuregoSafeRegister(&xNewDropTargetAsync, lib, "gtk_drop_target_async_new")

	core.PuregoSafeRegister(&xDropTargetAsyncGetActions, lib, "gtk_drop_target_async_get_actions")
	core.PuregoSafeRegister(&xDropTargetAsyncGetFormats, lib, "gtk_drop_target_async_get_formats")
	core.PuregoSafeRegister(&xDropTargetAsyncRejectDrop, lib, "gtk_drop_target_async_reject_drop")
	core.PuregoSafeRegister(&xDropTargetAsyncSetActions, lib, "gtk_drop_target_async_set_actions")
	core.PuregoSafeRegister(&xDropTargetAsyncSetFormats, lib, "gtk_drop_target_async_set_formats")

}
