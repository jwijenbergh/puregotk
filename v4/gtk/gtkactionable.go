// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The interface vtable for `GtkActionable`.
type ActionableInterface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *ActionableInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// The `GtkActionable` interface provides a convenient way of asscociating
// widgets with actions.
//
// It primarily consists of two properties: [property@Gtk.Actionable:action-name]
// and [property@Gtk.Actionable:action-target]. There are also some convenience
// APIs for setting these properties.
//
// The action will be looked up in action groups that are found among
// the widgets ancestors. Most commonly, these will be the actions with
// the “win.” or “app.” prefix that are associated with the
// `GtkApplicationWindow` or `GtkApplication`, but other action groups that
// are added with [method@Gtk.Widget.insert_action_group] will be consulted
// as well.
type Actionable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetActionName() string
	GetActionTargetValue() *glib.Variant
	SetActionName(ActionNameVar string)
	SetActionTarget(FormatStringVar string, varArgs ...interface{})
	SetActionTargetValue(TargetValueVar *glib.Variant)
	SetDetailedActionName(DetailedActionNameVar string)
}

var xActionableGLibType func() types.GType

func ActionableGLibType() types.GType {
	return xActionableGLibType()
}

type ActionableBase struct {
	Ptr uintptr
}

func (x *ActionableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *ActionableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the action name for @actionable.
func (x *ActionableBase) GetActionName() string {

	cret := XGtkActionableGetActionName(x.GoPointer())
	return cret
}

// Gets the current target value of @actionable.
func (x *ActionableBase) GetActionTargetValue() *glib.Variant {

	cret := XGtkActionableGetActionTargetValue(x.GoPointer())
	return cret
}

// Specifies the name of the action with which this widget should be
// associated.
//
// If @action_name is %NULL then the widget will be unassociated from
// any previous action.
//
// Usually this function is used when the widget is located (or will be
// located) within the hierarchy of a `GtkApplicationWindow`.
//
// Names are of the form “win.save” or “app.quit” for actions on the
// containing [class@ApplicationWindow] or its associated [class@Application],
// respectively. This is the same form used for actions in the [class@Gio.Menu]
// associated with the window.
func (x *ActionableBase) SetActionName(ActionNameVar string) {

	XGtkActionableSetActionName(x.GoPointer(), ActionNameVar)

}

// Sets the target of an actionable widget.
//
// This is a convenience function that calls [ctor@GLib.Variant.new] for
// @format_string and uses the result to call
// [method@Gtk.Actionable.set_action_target_value].
//
// If you are setting a string-valued target and want to set
// the action name at the same time, you can use
// [method@Gtk.Actionable.set_detailed_action_name].
func (x *ActionableBase) SetActionTarget(FormatStringVar string, varArgs ...interface{}) {

	XGtkActionableSetActionTarget(x.GoPointer(), FormatStringVar, varArgs...)

}

// Sets the target value of an actionable widget.
//
// If @target_value is %NULL then the target value is unset.
//
// The target value has two purposes. First, it is used as the parameter
// to activation of the action associated with the `GtkActionable` widget.
// Second, it is used to determine if the widget should be rendered as
// “active” — the widget is active if the state is equal to the given target.
//
// Consider the example of associating a set of buttons with a [iface@Gio.Action]
// with string state in a typical “radio button” situation. Each button
// will be associated with the same action, but with a different target
// value for that action. Clicking on a particular button will activate
// the action with the target of that button, which will typically cause
// the action’s state to change to that value. Since the action’s state
// is now equal to the target value of the button, the button will now
// be rendered as active (and the other buttons, with different targets,
// rendered inactive).
func (x *ActionableBase) SetActionTargetValue(TargetValueVar *glib.Variant) {

	XGtkActionableSetActionTargetValue(x.GoPointer(), TargetValueVar)

}

// Sets the action-name and associated string target value of an
// actionable widget.
//
// @detailed_action_name is a string in the format accepted by
// [func@Gio.Action.parse_detailed_name].
func (x *ActionableBase) SetDetailedActionName(DetailedActionNameVar string) {

	XGtkActionableSetDetailedActionName(x.GoPointer(), DetailedActionNameVar)

}

var XGtkActionableGetActionName func(uintptr) string
var XGtkActionableGetActionTargetValue func(uintptr) *glib.Variant
var XGtkActionableSetActionName func(uintptr, string)
var XGtkActionableSetActionTarget func(uintptr, string, ...interface{})
var XGtkActionableSetActionTargetValue func(uintptr, *glib.Variant)
var XGtkActionableSetDetailedActionName func(uintptr, string)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xActionableGLibType, lib, "gtk_actionable_get_type")

	core.PuregoSafeRegister(&XGtkActionableGetActionName, lib, "gtk_actionable_get_action_name")
	core.PuregoSafeRegister(&XGtkActionableGetActionTargetValue, lib, "gtk_actionable_get_action_target_value")
	core.PuregoSafeRegister(&XGtkActionableSetActionName, lib, "gtk_actionable_set_action_name")
	core.PuregoSafeRegister(&XGtkActionableSetActionTarget, lib, "gtk_actionable_set_action_target")
	core.PuregoSafeRegister(&XGtkActionableSetActionTargetValue, lib, "gtk_actionable_set_action_target_value")
	core.PuregoSafeRegister(&XGtkActionableSetDetailedActionName, lib, "gtk_actionable_set_detailed_action_name")

}
