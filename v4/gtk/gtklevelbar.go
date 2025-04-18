// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

const (
	// The name used for the stock full offset included by `GtkLevelBar`.
	LEVEL_BAR_OFFSET_FULL string = "full"
	// The name used for the stock high offset included by `GtkLevelBar`.
	LEVEL_BAR_OFFSET_HIGH string = "high"
	// The name used for the stock low offset included by `GtkLevelBar`.
	LEVEL_BAR_OFFSET_LOW string = "low"
)

// `GtkLevelBar` is a widget that can be used as a level indicator.
//
// Typical use cases are displaying the strength of a password, or
// showing the charge level of a battery.
//
// ![An example GtkLevelBar](levelbar.png)
//
// Use [method@Gtk.LevelBar.set_value] to set the current value, and
// [method@Gtk.LevelBar.add_offset_value] to set the value offsets at which
// the bar will be considered in a different state. GTK will add a few
// offsets by default on the level bar: %GTK_LEVEL_BAR_OFFSET_LOW,
// %GTK_LEVEL_BAR_OFFSET_HIGH and %GTK_LEVEL_BAR_OFFSET_FULL, with
// values 0.25, 0.75 and 1.0 respectively.
//
// Note that it is your responsibility to update preexisting offsets
// when changing the minimum or maximum value. GTK will simply clamp
// them to the new range.
//
// ## Adding a custom offset on the bar
//
// ```c
// static GtkWidget *
// create_level_bar (void)
//
//	{
//	  GtkWidget *widget;
//	  GtkLevelBar *bar;
//
//	  widget = gtk_level_bar_new ();
//	  bar = GTK_LEVEL_BAR (widget);
//
//	  // This changes the value of the default low offset
//
//	  gtk_level_bar_add_offset_value (bar,
//	                                  GTK_LEVEL_BAR_OFFSET_LOW,
//	                                  0.10);
//
//	  // This adds a new offset to the bar; the application will
//	  // be able to change its color CSS like this:
//	  //
//	  // levelbar block.my-offset {
//	  //   background-color: magenta;
//	  //   border-style: solid;
//	  //   border-color: black;
//	  //   border-style: 1px;
//	  // }
//
//	  gtk_level_bar_add_offset_value (bar, "my-offset", 0.60);
//
//	  return widget;
//	}
//
// ```
//
// The default interval of values is between zero and one, but it’s possible
// to modify the interval using [method@Gtk.LevelBar.set_min_value] and
// [method@Gtk.LevelBar.set_max_value]. The value will be always drawn in
// proportion to the admissible interval, i.e. a value of 15 with a specified
// interval between 10 and 20 is equivalent to a value of 0.5 with an interval
// between 0 and 1. When %GTK_LEVEL_BAR_MODE_DISCRETE is used, the bar level
// is rendered as a finite number of separated blocks instead of a single one.
// The number of blocks that will be rendered is equal to the number of units
// specified by the admissible interval.
//
// For instance, to build a bar rendered with five blocks, it’s sufficient to
// set the minimum value to 0 and the maximum value to 5 after changing the
// indicator mode to discrete.
//
// # GtkLevelBar as GtkBuildable
//
// The `GtkLevelBar` implementation of the `GtkBuildable` interface supports a
// custom &lt;offsets&gt; element, which can contain any number of &lt;offset&gt; elements,
// each of which must have name and value attributes.
//
// # CSS nodes
//
// ```
// levelbar[.discrete]
// ╰── trough
//
//	├── block.filled.level-name
//	┊
//	├── block.empty
//	┊
//
// ```
//
// `GtkLevelBar` has a main CSS node with name levelbar and one of the style
// classes .discrete or .continuous and a subnode with name trough. Below the
// trough node are a number of nodes with name block and style class .filled
// or .empty. In continuous mode, there is exactly one node of each, in discrete
// mode, the number of filled and unfilled nodes corresponds to blocks that are
// drawn. The block.filled nodes also get a style class .level-name corresponding
// to the level for the current value.
//
// In horizontal orientation, the nodes are always arranged from left to right,
// regardless of text direction.
//
// # Accessibility
//
// `GtkLevelBar` uses the %GTK_ACCESSIBLE_ROLE_METER role.
type LevelBar struct {
	Widget
}

var xLevelBarGLibType func() types.GType

func LevelBarGLibType() types.GType {
	return xLevelBarGLibType()
}

func LevelBarNewFromInternalPtr(ptr uintptr) *LevelBar {
	cls := &LevelBar{}
	cls.Ptr = ptr
	return cls
}

var xNewLevelBar func() uintptr

// Creates a new `GtkLevelBar`.
func NewLevelBar() *LevelBar {
	var cls *LevelBar

	cret := xNewLevelBar()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &LevelBar{}
	cls.Ptr = cret
	return cls
}

var xNewLevelBarForInterval func(float64, float64) uintptr

// Creates a new `GtkLevelBar` for the specified interval.
func NewLevelBarForInterval(MinValueVar float64, MaxValueVar float64) *LevelBar {
	var cls *LevelBar

	cret := xNewLevelBarForInterval(MinValueVar, MaxValueVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &LevelBar{}
	cls.Ptr = cret
	return cls
}

var xLevelBarAddOffsetValue func(uintptr, string, float64)

// Adds a new offset marker on @self at the position specified by @value.
//
// When the bar value is in the interval topped by @value (or between @value
// and [property@Gtk.LevelBar:max-value] in case the offset is the last one
// on the bar) a style class named `level-`@name will be applied
// when rendering the level bar fill.
//
// If another offset marker named @name exists, its value will be
// replaced by @value.
func (x *LevelBar) AddOffsetValue(NameVar string, ValueVar float64) {

	xLevelBarAddOffsetValue(x.GoPointer(), NameVar, ValueVar)

}

var xLevelBarGetInverted func(uintptr) bool

// Returns whether the levelbar is inverted.
func (x *LevelBar) GetInverted() bool {

	cret := xLevelBarGetInverted(x.GoPointer())
	return cret
}

var xLevelBarGetMaxValue func(uintptr) float64

// Returns the `max-value` of the `GtkLevelBar`.
func (x *LevelBar) GetMaxValue() float64 {

	cret := xLevelBarGetMaxValue(x.GoPointer())
	return cret
}

var xLevelBarGetMinValue func(uintptr) float64

// Returns the `min-value of the `GtkLevelBar`.
func (x *LevelBar) GetMinValue() float64 {

	cret := xLevelBarGetMinValue(x.GoPointer())
	return cret
}

var xLevelBarGetMode func(uintptr) LevelBarMode

// Returns the `mode` of the `GtkLevelBar`.
func (x *LevelBar) GetMode() LevelBarMode {

	cret := xLevelBarGetMode(x.GoPointer())
	return cret
}

var xLevelBarGetOffsetValue func(uintptr, string, float64) bool

// Fetches the value specified for the offset marker @name in @self.
func (x *LevelBar) GetOffsetValue(NameVar string, ValueVar float64) bool {

	cret := xLevelBarGetOffsetValue(x.GoPointer(), NameVar, ValueVar)
	return cret
}

var xLevelBarGetValue func(uintptr) float64

// Returns the `value` of the `GtkLevelBar`.
func (x *LevelBar) GetValue() float64 {

	cret := xLevelBarGetValue(x.GoPointer())
	return cret
}

var xLevelBarRemoveOffsetValue func(uintptr, string)

// Removes an offset marker from a `GtkLevelBar`.
//
// The marker must have been previously added with
// [method@Gtk.LevelBar.add_offset_value].
func (x *LevelBar) RemoveOffsetValue(NameVar string) {

	xLevelBarRemoveOffsetValue(x.GoPointer(), NameVar)

}

var xLevelBarSetInverted func(uintptr, bool)

// Sets whether the `GtkLevelBar` is inverted.
func (x *LevelBar) SetInverted(InvertedVar bool) {

	xLevelBarSetInverted(x.GoPointer(), InvertedVar)

}

var xLevelBarSetMaxValue func(uintptr, float64)

// Sets the `max-value` of the `GtkLevelBar`.
//
// You probably want to update preexisting level offsets after calling
// this function.
func (x *LevelBar) SetMaxValue(ValueVar float64) {

	xLevelBarSetMaxValue(x.GoPointer(), ValueVar)

}

var xLevelBarSetMinValue func(uintptr, float64)

// Sets the `min-value` of the `GtkLevelBar`.
//
// You probably want to update preexisting level offsets after calling
// this function.
func (x *LevelBar) SetMinValue(ValueVar float64) {

	xLevelBarSetMinValue(x.GoPointer(), ValueVar)

}

var xLevelBarSetMode func(uintptr, LevelBarMode)

// Sets the `mode` of the `GtkLevelBar`.
func (x *LevelBar) SetMode(ModeVar LevelBarMode) {

	xLevelBarSetMode(x.GoPointer(), ModeVar)

}

var xLevelBarSetValue func(uintptr, float64)

// Sets the value of the `GtkLevelBar`.
func (x *LevelBar) SetValue(ValueVar float64) {

	xLevelBarSetValue(x.GoPointer(), ValueVar)

}

func (c *LevelBar) GoPointer() uintptr {
	return c.Ptr
}

func (c *LevelBar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when an offset specified on the bar changes value.
//
// This typically is the result of a [method@Gtk.LevelBar.add_offset_value]
// call.
//
// The signal supports detailed connections; you can connect to the
// detailed signal "changed::x" in order to only receive callbacks when
// the value of offset "x" changes.
func (x *LevelBar) ConnectOffsetChanged(cb *func(LevelBar, string)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "offset-changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, NameVarp string) {
		fa := LevelBar{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, NameVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "offset-changed", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *LevelBar) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *LevelBar) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *LevelBar) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *LevelBar) ResetState(StateVar AccessibleState) {

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
func (x *LevelBar) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *LevelBar) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *LevelBar) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *LevelBar) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *LevelBar) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *LevelBar) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *LevelBar) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *LevelBar) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *LevelBar) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xLevelBarGLibType, lib, "gtk_level_bar_get_type")

	core.PuregoSafeRegister(&xNewLevelBar, lib, "gtk_level_bar_new")
	core.PuregoSafeRegister(&xNewLevelBarForInterval, lib, "gtk_level_bar_new_for_interval")

	core.PuregoSafeRegister(&xLevelBarAddOffsetValue, lib, "gtk_level_bar_add_offset_value")
	core.PuregoSafeRegister(&xLevelBarGetInverted, lib, "gtk_level_bar_get_inverted")
	core.PuregoSafeRegister(&xLevelBarGetMaxValue, lib, "gtk_level_bar_get_max_value")
	core.PuregoSafeRegister(&xLevelBarGetMinValue, lib, "gtk_level_bar_get_min_value")
	core.PuregoSafeRegister(&xLevelBarGetMode, lib, "gtk_level_bar_get_mode")
	core.PuregoSafeRegister(&xLevelBarGetOffsetValue, lib, "gtk_level_bar_get_offset_value")
	core.PuregoSafeRegister(&xLevelBarGetValue, lib, "gtk_level_bar_get_value")
	core.PuregoSafeRegister(&xLevelBarRemoveOffsetValue, lib, "gtk_level_bar_remove_offset_value")
	core.PuregoSafeRegister(&xLevelBarSetInverted, lib, "gtk_level_bar_set_inverted")
	core.PuregoSafeRegister(&xLevelBarSetMaxValue, lib, "gtk_level_bar_set_max_value")
	core.PuregoSafeRegister(&xLevelBarSetMinValue, lib, "gtk_level_bar_set_min_value")
	core.PuregoSafeRegister(&xLevelBarSetMode, lib, "gtk_level_bar_set_mode")
	core.PuregoSafeRegister(&xLevelBarSetValue, lib, "gtk_level_bar_set_value")

}
