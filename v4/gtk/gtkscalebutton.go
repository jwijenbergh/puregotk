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

type ScaleButtonClass struct {
	_ structs.HostLayout

	ParentClass uintptr

	Padding [8]uintptr
}

func (x *ScaleButtonClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkScaleButton` provides a button which pops up a scale widget.
//
// This kind of widget is commonly used for volume controls in multimedia
// applications, and GTK provides a [class@Gtk.VolumeButton] subclass that
// is tailored for this use case.
//
// # CSS nodes
//
// `GtkScaleButton` has a single CSS node with name button. To differentiate
// it from a plain `GtkButton`, it gets the .scale style class.
type ScaleButton struct {
	Widget
}

var xScaleButtonGLibType func() types.GType

func ScaleButtonGLibType() types.GType {
	return xScaleButtonGLibType()
}

func ScaleButtonNewFromInternalPtr(ptr uintptr) *ScaleButton {
	cls := &ScaleButton{}
	cls.Ptr = ptr
	return cls
}

var xNewScaleButton func(float64, float64, float64, []string) uintptr

// Creates a `GtkScaleButton`.
//
// The new scale button has a range between @min and @max,
// with a stepping of @step.
func NewScaleButton(MinVar float64, MaxVar float64, StepVar float64, IconsVar []string) *ScaleButton {
	var cls *ScaleButton

	cret := xNewScaleButton(MinVar, MaxVar, StepVar, IconsVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ScaleButton{}
	cls.Ptr = cret
	return cls
}

var xScaleButtonGetAdjustment func(uintptr) uintptr

// Gets the `GtkAdjustment` associated with the `GtkScaleButton`’s scale.
//
// See [method@Gtk.Range.get_adjustment] for details.
func (x *ScaleButton) GetAdjustment() *Adjustment {
	var cls *Adjustment

	cret := xScaleButtonGetAdjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

var xScaleButtonGetMinusButton func(uintptr) uintptr

// Retrieves the minus button of the `GtkScaleButton`.
func (x *ScaleButton) GetMinusButton() *Button {
	var cls *Button

	cret := xScaleButtonGetMinusButton(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Button{}
	cls.Ptr = cret
	return cls
}

var xScaleButtonGetPlusButton func(uintptr) uintptr

// Retrieves the plus button of the `GtkScaleButton.`
func (x *ScaleButton) GetPlusButton() *Button {
	var cls *Button

	cret := xScaleButtonGetPlusButton(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Button{}
	cls.Ptr = cret
	return cls
}

var xScaleButtonGetPopup func(uintptr) uintptr

// Retrieves the popup of the `GtkScaleButton`.
func (x *ScaleButton) GetPopup() *Widget {
	var cls *Widget

	cret := xScaleButtonGetPopup(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xScaleButtonGetValue func(uintptr) float64

// Gets the current value of the scale button.
func (x *ScaleButton) GetValue() float64 {

	cret := xScaleButtonGetValue(x.GoPointer())
	return cret
}

var xScaleButtonSetAdjustment func(uintptr, uintptr)

// Sets the `GtkAdjustment` to be used as a model
// for the `GtkScaleButton`’s scale.
//
// See [method@Gtk.Range.set_adjustment] for details.
func (x *ScaleButton) SetAdjustment(AdjustmentVar *Adjustment) {

	xScaleButtonSetAdjustment(x.GoPointer(), AdjustmentVar.GoPointer())

}

var xScaleButtonSetIcons func(uintptr, []string)

// Sets the icons to be used by the scale button.
func (x *ScaleButton) SetIcons(IconsVar []string) {

	xScaleButtonSetIcons(x.GoPointer(), IconsVar)

}

var xScaleButtonSetValue func(uintptr, float64)

// Sets the current value of the scale.
//
// If the value is outside the minimum or maximum range values,
// it will be clamped to fit inside them.
//
// The scale button emits the [signal@Gtk.ScaleButton::value-changed]
// signal if the value changes.
func (x *ScaleButton) SetValue(ValueVar float64) {

	xScaleButtonSetValue(x.GoPointer(), ValueVar)

}

func (c *ScaleButton) GoPointer() uintptr {
	return c.Ptr
}

func (c *ScaleButton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to dismiss the popup.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is &lt;kbd&gt;Escape&lt;/kbd&gt;.
func (x *ScaleButton) ConnectPopdown(cb *func(ScaleButton)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "popdown", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := ScaleButton{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "popdown", cbRefPtr)
}

// Emitted to popup the scale widget.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default bindings for this signal are &lt;kbd&gt;Space&lt;/kbd&gt;,
// &lt;kbd&gt;Enter&lt;/kbd&gt; and &lt;kbd&gt;Return&lt;/kbd&gt;.
func (x *ScaleButton) ConnectPopup(cb *func(ScaleButton)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "popup", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := ScaleButton{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "popup", cbRefPtr)
}

// Emitted when the value field has changed.
func (x *ScaleButton) ConnectValueChanged(cb *func(ScaleButton, float64)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "value-changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ValueVarp float64) {
		fa := ScaleButton{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ValueVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "value-changed", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ScaleButton) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ScaleButton) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ScaleButton) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ScaleButton) ResetState(StateVar AccessibleState) {

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
func (x *ScaleButton) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScaleButton) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *ScaleButton) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScaleButton) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *ScaleButton) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScaleButton) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ScaleButton) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *ScaleButton) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *ScaleButton) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xScaleButtonGLibType, lib, "gtk_scale_button_get_type")

	core.PuregoSafeRegister(&xNewScaleButton, lib, "gtk_scale_button_new")

	core.PuregoSafeRegister(&xScaleButtonGetAdjustment, lib, "gtk_scale_button_get_adjustment")
	core.PuregoSafeRegister(&xScaleButtonGetMinusButton, lib, "gtk_scale_button_get_minus_button")
	core.PuregoSafeRegister(&xScaleButtonGetPlusButton, lib, "gtk_scale_button_get_plus_button")
	core.PuregoSafeRegister(&xScaleButtonGetPopup, lib, "gtk_scale_button_get_popup")
	core.PuregoSafeRegister(&xScaleButtonGetValue, lib, "gtk_scale_button_get_value")
	core.PuregoSafeRegister(&xScaleButtonSetAdjustment, lib, "gtk_scale_button_set_adjustment")
	core.PuregoSafeRegister(&xScaleButtonSetIcons, lib, "gtk_scale_button_set_icons")
	core.PuregoSafeRegister(&xScaleButtonSetValue, lib, "gtk_scale_button_set_value")

}
