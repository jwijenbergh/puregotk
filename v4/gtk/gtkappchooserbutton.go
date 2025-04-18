// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The `GtkAppChooserButton` lets the user select an application.
//
// ![An example GtkAppChooserButton](appchooserbutton.png)
//
// Initially, a `GtkAppChooserButton` selects the first application
// in its list, which will either be the most-recently used application
// or, if [property@Gtk.AppChooserButton:show-default-item] is %TRUE, the
// default application.
//
// The list of applications shown in a `GtkAppChooserButton` includes
// the recommended applications for the given content type. When
// [property@Gtk.AppChooserButton:show-default-item] is set, the default
// application is also included. To let the user chooser other applications,
// you can set the [property@Gtk.AppChooserButton:show-dialog-item] property,
// which allows to open a full [class@Gtk.AppChooserDialog].
//
// It is possible to add custom items to the list, using
// [method@Gtk.AppChooserButton.append_custom_item]. These items cause
// the [signal@Gtk.AppChooserButton::custom-item-activated] signal to be
// emitted when they are selected.
//
// To track changes in the selected application, use the
// [signal@Gtk.AppChooserButton::changed] signal.
//
// # CSS nodes
//
// `GtkAppChooserButton` has a single CSS node with the name “appchooserbutton”.
type AppChooserButton struct {
	Widget
}

var xAppChooserButtonGLibType func() types.GType

func AppChooserButtonGLibType() types.GType {
	return xAppChooserButtonGLibType()
}

func AppChooserButtonNewFromInternalPtr(ptr uintptr) *AppChooserButton {
	cls := &AppChooserButton{}
	cls.Ptr = ptr
	return cls
}

var xNewAppChooserButton func(string) uintptr

// Creates a new `GtkAppChooserButton` for applications
// that can handle content of the given type.
func NewAppChooserButton(ContentTypeVar string) *AppChooserButton {
	var cls *AppChooserButton

	cret := xNewAppChooserButton(ContentTypeVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &AppChooserButton{}
	cls.Ptr = cret
	return cls
}

var xAppChooserButtonAppendCustomItem func(uintptr, string, string, uintptr)

// Appends a custom item to the list of applications that is shown
// in the popup.
//
// The item name must be unique per-widget. Clients can use the
// provided name as a detail for the
// [signal@Gtk.AppChooserButton::custom-item-activated] signal, to add a
// callback for the activation of a particular custom item in the list.
//
// See also [method@Gtk.AppChooserButton.append_separator].
func (x *AppChooserButton) AppendCustomItem(NameVar string, LabelVar string, IconVar gio.Icon) {

	xAppChooserButtonAppendCustomItem(x.GoPointer(), NameVar, LabelVar, IconVar.GoPointer())

}

var xAppChooserButtonAppendSeparator func(uintptr)

// Appends a separator to the list of applications that is shown
// in the popup.
func (x *AppChooserButton) AppendSeparator() {

	xAppChooserButtonAppendSeparator(x.GoPointer())

}

var xAppChooserButtonGetHeading func(uintptr) string

// Returns the text to display at the top of the dialog.
func (x *AppChooserButton) GetHeading() string {

	cret := xAppChooserButtonGetHeading(x.GoPointer())
	return cret
}

var xAppChooserButtonGetModal func(uintptr) bool

// Gets whether the dialog is modal.
func (x *AppChooserButton) GetModal() bool {

	cret := xAppChooserButtonGetModal(x.GoPointer())
	return cret
}

var xAppChooserButtonGetShowDefaultItem func(uintptr) bool

// Returns whether the dropdown menu should show the default
// application at the top.
func (x *AppChooserButton) GetShowDefaultItem() bool {

	cret := xAppChooserButtonGetShowDefaultItem(x.GoPointer())
	return cret
}

var xAppChooserButtonGetShowDialogItem func(uintptr) bool

// Returns whether the dropdown menu shows an item
// for a `GtkAppChooserDialog`.
func (x *AppChooserButton) GetShowDialogItem() bool {

	cret := xAppChooserButtonGetShowDialogItem(x.GoPointer())
	return cret
}

var xAppChooserButtonSetActiveCustomItem func(uintptr, string)

// Selects a custom item.
//
// See [method@Gtk.AppChooserButton.append_custom_item].
//
// Use [method@Gtk.AppChooser.refresh] to bring the selection
// to its initial state.
func (x *AppChooserButton) SetActiveCustomItem(NameVar string) {

	xAppChooserButtonSetActiveCustomItem(x.GoPointer(), NameVar)

}

var xAppChooserButtonSetHeading func(uintptr, string)

// Sets the text to display at the top of the dialog.
//
// If the heading is not set, the dialog displays a default text.
func (x *AppChooserButton) SetHeading(HeadingVar string) {

	xAppChooserButtonSetHeading(x.GoPointer(), HeadingVar)

}

var xAppChooserButtonSetModal func(uintptr, bool)

// Sets whether the dialog should be modal.
func (x *AppChooserButton) SetModal(ModalVar bool) {

	xAppChooserButtonSetModal(x.GoPointer(), ModalVar)

}

var xAppChooserButtonSetShowDefaultItem func(uintptr, bool)

// Sets whether the dropdown menu of this button should show the
// default application for the given content type at top.
func (x *AppChooserButton) SetShowDefaultItem(SettingVar bool) {

	xAppChooserButtonSetShowDefaultItem(x.GoPointer(), SettingVar)

}

var xAppChooserButtonSetShowDialogItem func(uintptr, bool)

// Sets whether the dropdown menu of this button should show an
// entry to trigger a `GtkAppChooserDialog`.
func (x *AppChooserButton) SetShowDialogItem(SettingVar bool) {

	xAppChooserButtonSetShowDialogItem(x.GoPointer(), SettingVar)

}

func (c *AppChooserButton) GoPointer() uintptr {
	return c.Ptr
}

func (c *AppChooserButton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to when the button is activated.
//
// The `::activate` signal on `GtkAppChooserButton` is an action signal and
// emitting it causes the button to pop up its dialog.
func (x *AppChooserButton) ConnectActivate(cb *func(AppChooserButton)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := AppChooserButton{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
}

// Emitted when the active application changes.
func (x *AppChooserButton) ConnectChanged(cb *func(AppChooserButton)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := AppChooserButton{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "changed", cbRefPtr)
}

// Emitted when a custom item is activated.
//
// Use [method@Gtk.AppChooserButton.append_custom_item],
// to add custom items.
func (x *AppChooserButton) ConnectCustomItemActivated(cb *func(AppChooserButton, string)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "custom-item-activated", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ItemNameVarp string) {
		fa := AppChooserButton{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ItemNameVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "custom-item-activated", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *AppChooserButton) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *AppChooserButton) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *AppChooserButton) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *AppChooserButton) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *AppChooserButton) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AppChooserButton) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *AppChooserButton) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AppChooserButton) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *AppChooserButton) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AppChooserButton) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Returns the currently selected application.
func (x *AppChooserButton) GetAppInfo() *gio.AppInfoBase {
	var cls *gio.AppInfoBase

	cret := XGtkAppChooserGetAppInfo(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gio.AppInfoBase{}
	cls.Ptr = cret
	return cls
}

// Returns the content type for which the `GtkAppChooser`
// shows applications.
func (x *AppChooserButton) GetContentType() string {

	cret := XGtkAppChooserGetContentType(x.GoPointer())
	return cret
}

// Reloads the list of applications.
func (x *AppChooserButton) Refresh() {

	XGtkAppChooserRefresh(x.GoPointer())

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *AppChooserButton) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAppChooserButtonGLibType, lib, "gtk_app_chooser_button_get_type")

	core.PuregoSafeRegister(&xNewAppChooserButton, lib, "gtk_app_chooser_button_new")

	core.PuregoSafeRegister(&xAppChooserButtonAppendCustomItem, lib, "gtk_app_chooser_button_append_custom_item")
	core.PuregoSafeRegister(&xAppChooserButtonAppendSeparator, lib, "gtk_app_chooser_button_append_separator")
	core.PuregoSafeRegister(&xAppChooserButtonGetHeading, lib, "gtk_app_chooser_button_get_heading")
	core.PuregoSafeRegister(&xAppChooserButtonGetModal, lib, "gtk_app_chooser_button_get_modal")
	core.PuregoSafeRegister(&xAppChooserButtonGetShowDefaultItem, lib, "gtk_app_chooser_button_get_show_default_item")
	core.PuregoSafeRegister(&xAppChooserButtonGetShowDialogItem, lib, "gtk_app_chooser_button_get_show_dialog_item")
	core.PuregoSafeRegister(&xAppChooserButtonSetActiveCustomItem, lib, "gtk_app_chooser_button_set_active_custom_item")
	core.PuregoSafeRegister(&xAppChooserButtonSetHeading, lib, "gtk_app_chooser_button_set_heading")
	core.PuregoSafeRegister(&xAppChooserButtonSetModal, lib, "gtk_app_chooser_button_set_modal")
	core.PuregoSafeRegister(&xAppChooserButtonSetShowDefaultItem, lib, "gtk_app_chooser_button_set_show_default_item")
	core.PuregoSafeRegister(&xAppChooserButtonSetShowDialogItem, lib, "gtk_app_chooser_button_set_show_dialog_item")

}
