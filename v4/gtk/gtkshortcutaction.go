// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Prototype for shortcuts based on user callbacks.
type ShortcutFunc func(uintptr, *glib.Variant, uintptr) bool

type ActivateActionClass struct {
	_ structs.HostLayout
}

func (x *ActivateActionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type CallbackActionClass struct {
	_ structs.HostLayout
}

func (x *CallbackActionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type MnemonicActionClass struct {
	_ structs.HostLayout
}

func (x *MnemonicActionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type NamedActionClass struct {
	_ structs.HostLayout
}

func (x *NamedActionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type NothingActionClass struct {
	_ structs.HostLayout
}

func (x *NothingActionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type ShortcutActionClass struct {
	_ structs.HostLayout
}

func (x *ShortcutActionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type SignalActionClass struct {
	_ structs.HostLayout
}

func (x *SignalActionClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// List of flags that can be passed to action activation.
//
// More flags may be added in the future.
type ShortcutActionFlags int

var xShortcutActionFlagsGLibType func() types.GType

func ShortcutActionFlagsGLibType() types.GType {
	return xShortcutActionFlagsGLibType()
}

const (

	// The action is the only
	//   action that can be activated. If this flag is not set,
	//   a future activation may select a different action.
	ShortcutActionExclusiveValue ShortcutActionFlags = 1
)

// A `GtkShortcutAction` that calls gtk_widget_activate().
type ActivateAction struct {
	ShortcutAction
}

var xActivateActionGLibType func() types.GType

func ActivateActionGLibType() types.GType {
	return xActivateActionGLibType()
}

func ActivateActionNewFromInternalPtr(ptr uintptr) *ActivateAction {
	cls := &ActivateAction{}
	cls.Ptr = ptr
	return cls
}

func (c *ActivateAction) GoPointer() uintptr {
	return c.Ptr
}

func (c *ActivateAction) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

var xActivateActionGet func() uintptr

// Gets the activate action.
//
// This is an action that calls gtk_widget_activate()
// on the given widget upon activation.
func ActivateActionGet() *ActivateAction {
	var cls *ActivateAction

	cret := xActivateActionGet()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ActivateAction{}
	cls.Ptr = cret
	return cls
}

// A `GtkShortcutAction` that invokes a callback.
type CallbackAction struct {
	ShortcutAction
}

var xCallbackActionGLibType func() types.GType

func CallbackActionGLibType() types.GType {
	return xCallbackActionGLibType()
}

func CallbackActionNewFromInternalPtr(ptr uintptr) *CallbackAction {
	cls := &CallbackAction{}
	cls.Ptr = ptr
	return cls
}

var xNewCallbackAction func(uintptr, uintptr, uintptr) uintptr

// Create a custom action that calls the given @callback when
// activated.
func NewCallbackAction(CallbackVar *ShortcutFunc, DataVar uintptr, DestroyVar *glib.DestroyNotify) *CallbackAction {
	var cls *CallbackAction

	cret := xNewCallbackAction(glib.NewCallback(CallbackVar), DataVar, glib.NewCallback(DestroyVar))

	if cret == 0 {
		return nil
	}
	cls = &CallbackAction{}
	cls.Ptr = cret
	return cls
}

func (c *CallbackAction) GoPointer() uintptr {
	return c.Ptr
}

func (c *CallbackAction) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `GtkShortcutAction` that calls gtk_widget_mnemonic_activate().
type MnemonicAction struct {
	ShortcutAction
}

var xMnemonicActionGLibType func() types.GType

func MnemonicActionGLibType() types.GType {
	return xMnemonicActionGLibType()
}

func MnemonicActionNewFromInternalPtr(ptr uintptr) *MnemonicAction {
	cls := &MnemonicAction{}
	cls.Ptr = ptr
	return cls
}

func (c *MnemonicAction) GoPointer() uintptr {
	return c.Ptr
}

func (c *MnemonicAction) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

var xMnemonicActionGet func() uintptr

// Gets the mnemonic action.
//
// This is an action that calls gtk_widget_mnemonic_activate()
// on the given widget upon activation.
func MnemonicActionGet() *MnemonicAction {
	var cls *MnemonicAction

	cret := xMnemonicActionGet()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &MnemonicAction{}
	cls.Ptr = cret
	return cls
}

// A `GtkShortcutAction` that activates an action by name.
type NamedAction struct {
	ShortcutAction
}

var xNamedActionGLibType func() types.GType

func NamedActionGLibType() types.GType {
	return xNamedActionGLibType()
}

func NamedActionNewFromInternalPtr(ptr uintptr) *NamedAction {
	cls := &NamedAction{}
	cls.Ptr = ptr
	return cls
}

var xNewNamedAction func(string) uintptr

// Creates an action that when activated, activates
// the named action on the widget.
//
// It also passes the given arguments to it.
//
// See [method@Gtk.Widget.insert_action_group] for
// how to add actions to widgets.
func NewNamedAction(NameVar string) *NamedAction {
	var cls *NamedAction

	cret := xNewNamedAction(NameVar)

	if cret == 0 {
		return nil
	}
	cls = &NamedAction{}
	cls.Ptr = cret
	return cls
}

var xNamedActionGetActionName func(uintptr) string

// Returns the name of the action that will be activated.
func (x *NamedAction) GetActionName() string {

	cret := xNamedActionGetActionName(x.GoPointer())
	return cret
}

func (c *NamedAction) GoPointer() uintptr {
	return c.Ptr
}

func (c *NamedAction) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `GtkShortcutAction` that does nothing.
type NothingAction struct {
	ShortcutAction
}

var xNothingActionGLibType func() types.GType

func NothingActionGLibType() types.GType {
	return xNothingActionGLibType()
}

func NothingActionNewFromInternalPtr(ptr uintptr) *NothingAction {
	cls := &NothingAction{}
	cls.Ptr = ptr
	return cls
}

func (c *NothingAction) GoPointer() uintptr {
	return c.Ptr
}

func (c *NothingAction) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

var xNothingActionGet func() uintptr

// Gets the nothing action.
//
// This is an action that does nothing and where
// activating it always fails.
func NothingActionGet() *NothingAction {
	var cls *NothingAction

	cret := xNothingActionGet()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &NothingAction{}
	cls.Ptr = cret
	return cls
}

// `GtkShortcutAction` encodes an action that can be triggered by a
// keyboard shortcut.
//
// `GtkShortcutActions` contain functions that allow easy presentation
// to end users as well as being printed for debugging.
//
// All `GtkShortcutActions` are immutable, you can only specify their
// properties during construction. If you want to change a action, you
// have to replace it with a new one. If you need to pass arguments to
// an action, these are specified by the higher-level `GtkShortcut` object.
//
// To activate a `GtkShortcutAction` manually, [method@Gtk.ShortcutAction.activate]
// can be called.
//
// GTK provides various actions:
//
//   - [class@Gtk.MnemonicAction]: a shortcut action that calls
//     gtk_widget_mnemonic_activate()
//   - [class@Gtk.CallbackAction]: a shortcut action that invokes
//     a given callback
//   - [class@Gtk.SignalAction]: a shortcut action that emits a
//     given signal
//   - [class@Gtk.ActivateAction]: a shortcut action that calls
//     gtk_widget_activate()
//   - [class@Gtk.NamedAction]: a shortcut action that calls
//     gtk_widget_activate_action()
//   - [class@Gtk.NothingAction]: a shortcut action that does nothing
type ShortcutAction struct {
	gobject.Object
}

var xShortcutActionGLibType func() types.GType

func ShortcutActionGLibType() types.GType {
	return xShortcutActionGLibType()
}

func ShortcutActionNewFromInternalPtr(ptr uintptr) *ShortcutAction {
	cls := &ShortcutAction{}
	cls.Ptr = ptr
	return cls
}

var xShortcutActionParseString func(string) uintptr

// Tries to parse the given string into an action.
//
// On success, the parsed action is returned. When parsing
// failed, %NULL is returned.
//
// The accepted strings are:
//
// - `nothing`, for `GtkNothingAction`
// - `activate`, for `GtkActivateAction`
// - `mnemonic-activate`, for `GtkMnemonicAction`
// - `action(NAME)`, for a `GtkNamedAction` for the action named `NAME`
// - `signal(NAME)`, for a `GtkSignalAction` for the signal `NAME`
func ShortcutActionParseString(StringVar string) *ShortcutAction {
	var cls *ShortcutAction

	cret := xShortcutActionParseString(StringVar)

	if cret == 0 {
		return nil
	}
	cls = &ShortcutAction{}
	cls.Ptr = cret
	return cls
}

var xShortcutActionActivate func(uintptr, ShortcutActionFlags, uintptr, *glib.Variant) bool

// Activates the action on the @widget with the given @args.
//
// Note that some actions ignore the passed in @flags, @widget or @args.
//
// Activation of an action can fail for various reasons. If the action
// is not supported by the @widget, if the @args don't match the action
// or if the activation otherwise had no effect, %FALSE will be returned.
func (x *ShortcutAction) Activate(FlagsVar ShortcutActionFlags, WidgetVar *Widget, ArgsVar *glib.Variant) bool {

	cret := xShortcutActionActivate(x.GoPointer(), FlagsVar, WidgetVar.GoPointer(), ArgsVar)
	return cret
}

var xShortcutActionPrint func(uintptr, *glib.String)

// Prints the given action into a string for the developer.
//
// This is meant for debugging and logging.
//
// The form of the representation may change at any time and is
// not guaranteed to stay identical.
func (x *ShortcutAction) Print(StringVar *glib.String) {

	xShortcutActionPrint(x.GoPointer(), StringVar)

}

var xShortcutActionToString func(uintptr) string

// Prints the given action into a human-readable string.
//
// This is a small wrapper around [method@Gtk.ShortcutAction.print]
// to help when debugging.
func (x *ShortcutAction) ToString() string {

	cret := xShortcutActionToString(x.GoPointer())
	return cret
}

func (c *ShortcutAction) GoPointer() uintptr {
	return c.Ptr
}

func (c *ShortcutAction) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `GtkShortcut`Action that emits a signal.
//
// Signals that are used in this way are referred to as keybinding signals,
// and they are expected to be defined with the %G_SIGNAL_ACTION flag.
type SignalAction struct {
	ShortcutAction
}

var xSignalActionGLibType func() types.GType

func SignalActionGLibType() types.GType {
	return xSignalActionGLibType()
}

func SignalActionNewFromInternalPtr(ptr uintptr) *SignalAction {
	cls := &SignalAction{}
	cls.Ptr = ptr
	return cls
}

var xNewSignalAction func(string) uintptr

// Creates an action that when activated, emits the given action signal
// on the provided widget.
//
// It will also unpack the args into arguments passed to the signal.
func NewSignalAction(SignalNameVar string) *SignalAction {
	var cls *SignalAction

	cret := xNewSignalAction(SignalNameVar)

	if cret == 0 {
		return nil
	}
	cls = &SignalAction{}
	cls.Ptr = cret
	return cls
}

var xSignalActionGetSignalName func(uintptr) string

// Returns the name of the signal that will be emitted.
func (x *SignalAction) GetSignalName() string {

	cret := xSignalActionGetSignalName(x.GoPointer())
	return cret
}

func (c *SignalAction) GoPointer() uintptr {
	return c.Ptr
}

func (c *SignalAction) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xShortcutActionFlagsGLibType, lib, "gtk_shortcut_action_flags_get_type")

	core.PuregoSafeRegister(&xActivateActionGLibType, lib, "gtk_activate_action_get_type")

	core.PuregoSafeRegister(&xActivateActionGet, lib, "gtk_activate_action_get")

	core.PuregoSafeRegister(&xCallbackActionGLibType, lib, "gtk_callback_action_get_type")

	core.PuregoSafeRegister(&xNewCallbackAction, lib, "gtk_callback_action_new")

	core.PuregoSafeRegister(&xMnemonicActionGLibType, lib, "gtk_mnemonic_action_get_type")

	core.PuregoSafeRegister(&xMnemonicActionGet, lib, "gtk_mnemonic_action_get")

	core.PuregoSafeRegister(&xNamedActionGLibType, lib, "gtk_named_action_get_type")

	core.PuregoSafeRegister(&xNewNamedAction, lib, "gtk_named_action_new")

	core.PuregoSafeRegister(&xNamedActionGetActionName, lib, "gtk_named_action_get_action_name")

	core.PuregoSafeRegister(&xNothingActionGLibType, lib, "gtk_nothing_action_get_type")

	core.PuregoSafeRegister(&xNothingActionGet, lib, "gtk_nothing_action_get")

	core.PuregoSafeRegister(&xShortcutActionGLibType, lib, "gtk_shortcut_action_get_type")

	core.PuregoSafeRegister(&xShortcutActionParseString, lib, "gtk_shortcut_action_parse_string")

	core.PuregoSafeRegister(&xShortcutActionActivate, lib, "gtk_shortcut_action_activate")
	core.PuregoSafeRegister(&xShortcutActionPrint, lib, "gtk_shortcut_action_print")
	core.PuregoSafeRegister(&xShortcutActionToString, lib, "gtk_shortcut_action_to_string")

	core.PuregoSafeRegister(&xSignalActionGLibType, lib, "gtk_signal_action_get_type")

	core.PuregoSafeRegister(&xNewSignalAction, lib, "gtk_signal_action_new")

	core.PuregoSafeRegister(&xSignalActionGetSignalName, lib, "gtk_signal_action_get_signal_name")

}
