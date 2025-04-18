// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type DialogClass struct {
	_ structs.HostLayout

	ParentClass uintptr

	Padding [4]uintptr
}

func (x *DialogClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Describes the available presentation modes for [class@Dialog].
//
// New values may be added to this enumeration over time.
//
// See [property@Dialog:presentation-mode].
type DialogPresentationMode int

var xDialogPresentationModeGLibType func() types.GType

func DialogPresentationModeGLibType() types.GType {
	return xDialogPresentationModeGLibType()
}

const (

	// Switch between `ADW_DIALOG_FLOATING` and
	//   `ADW_DIALOG_BOTTOM_SHEET` depending on available size.
	DialogAutoValue DialogPresentationMode = 0
	// Present dialog as a centered floating window.
	DialogFloatingValue DialogPresentationMode = 1
	// Present dialog as a bottom sheet.
	DialogBottomSheetValue DialogPresentationMode = 2
)

// An adaptive dialog container.
//
// &lt;picture&gt;
//
//	&lt;source srcset="dialog-floating-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="dialog-floating.png" alt="dialog-floating"&gt;
//
// &lt;/picture&gt;
// &lt;picture&gt;
//
//	&lt;source srcset="dialog-bottom-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="dialog-bottom.png" alt="dialog-bottom"&gt;
//
// &lt;/picture&gt;
//
// `AdwDialog` is similar to a window, but is shown within another window. It
// can be used with [class@Window] and [class@ApplicationWindow], use
// [method@Dialog.present] to show it.
//
// `AdwDialog` is not resizable. Use the [property@Dialog:content-width] and
// [property@Dialog:content-height] properties to set its size, or set
// [property@Dialog:follows-content-size] to `TRUE` to make the dialog track the
// content's size as it changes. `AdwDialog` can never be larger than its parent
// window.
//
// `AdwDialog` can be presented as a centered floating window or a bottom sheet.
// By default it's automatic depending on the available size.
// [property@Dialog:presentation-mode] can be used to change that.
//
// `AdwDialog` can be closed via [method@Dialog.close].
//
// When presented as a bottom sheet, `AdwDialog` can also be closed via swiping
// it down.
//
// The [property@Dialog:can-close] can be used to prevent closing. In that case,
// [signal@Dialog::close-attempt] gets emitted instead.
//
// Use [method@Dialog.force_close] to close the dialog even when `can-close` is set to
// `FALSE`.
//
// `AdwDialog` is transient and doesn't integrate with the window below it, for
// example it's not possible to collapse it into a bottom bar. See
// [class@BottomSheet] for persistent and more tightly integrated bottom sheets.
//
// ## Header Bar Integration
//
// When placed inside an `AdwDialog`, [class@HeaderBar] will display the dialog
// title instead of window title. It will also adjust the decoration layout to
// ensure it always has a close button and nothing else. Set
// [property@HeaderBar:show-start-title-buttons] and
// [property@HeaderBar:show-end-title-buttons] to `FALSE` to remove it if it's
// unwanted.
//
// ## Breakpoints
//
// `AdwDialog` can be used with [class@Breakpoint] the same way as
// [class@BreakpointBin]. Refer to that widget's documentation for details.
//
// Like `AdwBreakpointBin`, if breakpoints are used, `AdwDialog` doesn't have a
// minimum size, and [property@Gtk.Widget:width-request] and
// [property@Gtk.Widget:height-request] properties must be set manually.
type Dialog struct {
	gtk.Widget
}

var xDialogGLibType func() types.GType

func DialogGLibType() types.GType {
	return xDialogGLibType()
}

func DialogNewFromInternalPtr(ptr uintptr) *Dialog {
	cls := &Dialog{}
	cls.Ptr = ptr
	return cls
}

var xNewDialog func() uintptr

// Creates a new `AdwDialog`.
func NewDialog() *Dialog {
	var cls *Dialog

	cret := xNewDialog()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Dialog{}
	cls.Ptr = cret
	return cls
}

var xDialogAddBreakpoint func(uintptr, uintptr)

// Adds @breakpoint to @self.
func (x *Dialog) AddBreakpoint(BreakpointVar *Breakpoint) {

	xDialogAddBreakpoint(x.GoPointer(), BreakpointVar.GoPointer())

}

var xDialogClose func(uintptr) bool

// Attempts to close @self.
//
// If the [property@Dialog:can-close] property is set to `FALSE`, the
// [signal@Dialog::close-attempt] signal is emitted.
//
// See also: [method@Dialog.force_close].
func (x *Dialog) Close() bool {

	cret := xDialogClose(x.GoPointer())
	return cret
}

var xDialogForceClose func(uintptr)

// Closes @self.
//
// Unlike [method@Dialog.close], it succeeds even if [property@Dialog:can-close]
// is set to `FALSE`.
func (x *Dialog) ForceClose() {

	xDialogForceClose(x.GoPointer())

}

var xDialogGetCanClose func(uintptr) bool

// Gets whether @self can be closed.
func (x *Dialog) GetCanClose() bool {

	cret := xDialogGetCanClose(x.GoPointer())
	return cret
}

var xDialogGetChild func(uintptr) uintptr

// Gets the child widget of @self.
func (x *Dialog) GetChild() *gtk.Widget {
	var cls *gtk.Widget

	cret := xDialogGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xDialogGetContentHeight func(uintptr) int

// Gets the height of the dialog's contents.
func (x *Dialog) GetContentHeight() int {

	cret := xDialogGetContentHeight(x.GoPointer())
	return cret
}

var xDialogGetContentWidth func(uintptr) int

// Gets the width of the dialog's contents.
func (x *Dialog) GetContentWidth() int {

	cret := xDialogGetContentWidth(x.GoPointer())
	return cret
}

var xDialogGetCurrentBreakpoint func(uintptr) uintptr

// Gets the current breakpoint.
func (x *Dialog) GetCurrentBreakpoint() *Breakpoint {
	var cls *Breakpoint

	cret := xDialogGetCurrentBreakpoint(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Breakpoint{}
	cls.Ptr = cret
	return cls
}

var xDialogGetDefaultWidget func(uintptr) uintptr

// Gets the default widget for @self.
func (x *Dialog) GetDefaultWidget() *gtk.Widget {
	var cls *gtk.Widget

	cret := xDialogGetDefaultWidget(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xDialogGetFocus func(uintptr) uintptr

// Gets the focus widget for @self.
func (x *Dialog) GetFocus() *gtk.Widget {
	var cls *gtk.Widget

	cret := xDialogGetFocus(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xDialogGetFollowsContentSize func(uintptr) bool

// Gets whether to size content of @self automatically.
func (x *Dialog) GetFollowsContentSize() bool {

	cret := xDialogGetFollowsContentSize(x.GoPointer())
	return cret
}

var xDialogGetPresentationMode func(uintptr) DialogPresentationMode

// Gets presentation mode for @self.
func (x *Dialog) GetPresentationMode() DialogPresentationMode {

	cret := xDialogGetPresentationMode(x.GoPointer())
	return cret
}

var xDialogGetTitle func(uintptr) string

// Gets the title of @self.
func (x *Dialog) GetTitle() string {

	cret := xDialogGetTitle(x.GoPointer())
	return cret
}

var xDialogPresent func(uintptr, uintptr)

// Presents @self within @parent's window.
//
// If @self is already shown, raises it to the top instead.
//
// If the window is an [class@Window] or [class@ApplicationWindow], the dialog
// will be shown within it. Otherwise, it will be a separate window.
func (x *Dialog) Present(ParentVar *gtk.Widget) {

	xDialogPresent(x.GoPointer(), ParentVar.GoPointer())

}

var xDialogSetCanClose func(uintptr, bool)

// Sets whether @self can be closed.
//
// If set to `FALSE`, the close button, shortcuts and
// [method@Dialog.close] will result in [signal@Dialog::close-attempt] being
// emitted instead, and bottom sheet close swipe will be disabled.
// [method@Dialog.force_close] still works.
func (x *Dialog) SetCanClose(CanCloseVar bool) {

	xDialogSetCanClose(x.GoPointer(), CanCloseVar)

}

var xDialogSetChild func(uintptr, uintptr)

// Sets the child widget of @self.
func (x *Dialog) SetChild(ChildVar *gtk.Widget) {

	xDialogSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xDialogSetContentHeight func(uintptr, int)

// Sets the height of the dialog's contents.
//
// Set it to -1 to reset it to the content's natural height.
//
// See also: [property@Gtk.Window:default-height]
func (x *Dialog) SetContentHeight(ContentHeightVar int) {

	xDialogSetContentHeight(x.GoPointer(), ContentHeightVar)

}

var xDialogSetContentWidth func(uintptr, int)

// Sets the width of the dialog's contents.
//
// Set it to -1 to reset it to the content's natural width.
//
// See also: [property@Gtk.Window:default-width]
func (x *Dialog) SetContentWidth(ContentWidthVar int) {

	xDialogSetContentWidth(x.GoPointer(), ContentWidthVar)

}

var xDialogSetDefaultWidget func(uintptr, uintptr)

// Sets the default widget for @self.
//
// It's activated when the user presses Enter.
func (x *Dialog) SetDefaultWidget(DefaultWidgetVar *gtk.Widget) {

	xDialogSetDefaultWidget(x.GoPointer(), DefaultWidgetVar.GoPointer())

}

var xDialogSetFocus func(uintptr, uintptr)

// Sets the focus widget for @self.
//
// If @focus is not the current focus widget, and is focusable, sets it as the
// focus widget for the dialog.
//
// If focus is `NULL`, unsets the focus widget for this dialog. To set the focus
// to a particular widget in the dialog, it is usually more convenient to use
// [method@Gtk.Widget.grab_focus] instead of this function.
func (x *Dialog) SetFocus(FocusVar *gtk.Widget) {

	xDialogSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

var xDialogSetFollowsContentSize func(uintptr, bool)

// Sets whether to size content of @self automatically.
//
// If set to `TRUE`, always use the content's natural size instead of
// [property@Dialog:content-width] and [property@Dialog:content-height]. If
// the content resizes, the dialog will immediately resize as well.
//
// See also: [property@Gtk.Window:resizable]
func (x *Dialog) SetFollowsContentSize(FollowsContentSizeVar bool) {

	xDialogSetFollowsContentSize(x.GoPointer(), FollowsContentSizeVar)

}

var xDialogSetPresentationMode func(uintptr, DialogPresentationMode)

// Sets presentation mode for @self.
//
// When set to `ADW_DIALOG_AUTO`, the dialog appears as a bottom sheet when the
// following condition is met: `max-width: 450px or max-height: 360px`, and as a
// floating window otherwise.
//
// Set it to `ADW_DIALOG_FLOATING` or `ADW_DIALOG_BOTTOM_SHEET` to always
// present it a floating window or a bottom sheet respectively, regardless of
// available size.
//
// Presentation mode does nothing for dialogs presented as a window.
func (x *Dialog) SetPresentationMode(PresentationModeVar DialogPresentationMode) {

	xDialogSetPresentationMode(x.GoPointer(), PresentationModeVar)

}

var xDialogSetTitle func(uintptr, string)

// Sets the title of @self.
func (x *Dialog) SetTitle(TitleVar string) {

	xDialogSetTitle(x.GoPointer(), TitleVar)

}

func (c *Dialog) GoPointer() uintptr {
	return c.Ptr
}

func (c *Dialog) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the close button or shortcut is used, or
// [method@Dialog.close] is called while [property@Dialog:can-close] is set to
// `FALSE`.
func (x *Dialog) ConnectCloseAttempt(cb *func(Dialog)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "close-attempt", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Dialog{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "close-attempt", cbRefPtr)
}

// Emitted when the dialog is successfully closed.
func (x *Dialog) ConnectClosed(cb *func(Dialog)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "closed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Dialog{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "closed", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Dialog) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Dialog) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Dialog) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Dialog) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *Dialog) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Dialog) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []gtk.AccessibleProperty, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *Dialog) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Dialog) UpdateRelationValue(NRelationsVar int, RelationsVar []gtk.AccessibleRelation, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *Dialog) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Dialog) UpdateStateValue(NStatesVar int, StatesVar []gtk.AccessibleState, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Dialog) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xDialogPresentationModeGLibType, lib, "adw_dialog_presentation_mode_get_type")

	core.PuregoSafeRegister(&xDialogGLibType, lib, "adw_dialog_get_type")

	core.PuregoSafeRegister(&xNewDialog, lib, "adw_dialog_new")

	core.PuregoSafeRegister(&xDialogAddBreakpoint, lib, "adw_dialog_add_breakpoint")
	core.PuregoSafeRegister(&xDialogClose, lib, "adw_dialog_close")
	core.PuregoSafeRegister(&xDialogForceClose, lib, "adw_dialog_force_close")
	core.PuregoSafeRegister(&xDialogGetCanClose, lib, "adw_dialog_get_can_close")
	core.PuregoSafeRegister(&xDialogGetChild, lib, "adw_dialog_get_child")
	core.PuregoSafeRegister(&xDialogGetContentHeight, lib, "adw_dialog_get_content_height")
	core.PuregoSafeRegister(&xDialogGetContentWidth, lib, "adw_dialog_get_content_width")
	core.PuregoSafeRegister(&xDialogGetCurrentBreakpoint, lib, "adw_dialog_get_current_breakpoint")
	core.PuregoSafeRegister(&xDialogGetDefaultWidget, lib, "adw_dialog_get_default_widget")
	core.PuregoSafeRegister(&xDialogGetFocus, lib, "adw_dialog_get_focus")
	core.PuregoSafeRegister(&xDialogGetFollowsContentSize, lib, "adw_dialog_get_follows_content_size")
	core.PuregoSafeRegister(&xDialogGetPresentationMode, lib, "adw_dialog_get_presentation_mode")
	core.PuregoSafeRegister(&xDialogGetTitle, lib, "adw_dialog_get_title")
	core.PuregoSafeRegister(&xDialogPresent, lib, "adw_dialog_present")
	core.PuregoSafeRegister(&xDialogSetCanClose, lib, "adw_dialog_set_can_close")
	core.PuregoSafeRegister(&xDialogSetChild, lib, "adw_dialog_set_child")
	core.PuregoSafeRegister(&xDialogSetContentHeight, lib, "adw_dialog_set_content_height")
	core.PuregoSafeRegister(&xDialogSetContentWidth, lib, "adw_dialog_set_content_width")
	core.PuregoSafeRegister(&xDialogSetDefaultWidget, lib, "adw_dialog_set_default_widget")
	core.PuregoSafeRegister(&xDialogSetFocus, lib, "adw_dialog_set_focus")
	core.PuregoSafeRegister(&xDialogSetFollowsContentSize, lib, "adw_dialog_set_follows_content_size")
	core.PuregoSafeRegister(&xDialogSetPresentationMode, lib, "adw_dialog_set_presentation_mode")
	core.PuregoSafeRegister(&xDialogSetTitle, lib, "adw_dialog_set_title")

}
