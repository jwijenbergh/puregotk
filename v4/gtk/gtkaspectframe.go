// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// `GtkAspectFrame` preserves the aspect ratio of its child.
//
// The frame can respect the aspect ratio of the child widget,
// or use its own aspect ratio.
//
// # CSS nodes
//
// `GtkAspectFrame` uses a CSS node with name `frame`.
type AspectFrame struct {
	Widget
}

var xAspectFrameGLibType func() types.GType

func AspectFrameGLibType() types.GType {
	return xAspectFrameGLibType()
}

func AspectFrameNewFromInternalPtr(ptr uintptr) *AspectFrame {
	cls := &AspectFrame{}
	cls.Ptr = ptr
	return cls
}

var xNewAspectFrame func(float32, float32, float32, bool) uintptr

// Create a new `GtkAspectFrame`.
func NewAspectFrame(XalignVar float32, YalignVar float32, RatioVar float32, ObeyChildVar bool) *AspectFrame {
	var cls *AspectFrame

	cret := xNewAspectFrame(XalignVar, YalignVar, RatioVar, ObeyChildVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &AspectFrame{}
	cls.Ptr = cret
	return cls
}

var xAspectFrameGetChild func(uintptr) uintptr

// Gets the child widget of @self.
func (x *AspectFrame) GetChild() *Widget {
	var cls *Widget

	cret := xAspectFrameGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xAspectFrameGetObeyChild func(uintptr) bool

// Returns whether the child's size request should override
// the set aspect ratio of the `GtkAspectFrame`.
func (x *AspectFrame) GetObeyChild() bool {

	cret := xAspectFrameGetObeyChild(x.GoPointer())
	return cret
}

var xAspectFrameGetRatio func(uintptr) float32

// Returns the desired aspect ratio of the child.
func (x *AspectFrame) GetRatio() float32 {

	cret := xAspectFrameGetRatio(x.GoPointer())
	return cret
}

var xAspectFrameGetXalign func(uintptr) float32

// Returns the horizontal alignment of the child within the
// allocation of the `GtkAspectFrame`.
func (x *AspectFrame) GetXalign() float32 {

	cret := xAspectFrameGetXalign(x.GoPointer())
	return cret
}

var xAspectFrameGetYalign func(uintptr) float32

// Returns the vertical alignment of the child within the
// allocation of the `GtkAspectFrame`.
func (x *AspectFrame) GetYalign() float32 {

	cret := xAspectFrameGetYalign(x.GoPointer())
	return cret
}

var xAspectFrameSetChild func(uintptr, uintptr)

// Sets the child widget of @self.
func (x *AspectFrame) SetChild(ChildVar *Widget) {

	xAspectFrameSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xAspectFrameSetObeyChild func(uintptr, bool)

// Sets whether the aspect ratio of the child's size
// request should override the set aspect ratio of
// the `GtkAspectFrame`.
func (x *AspectFrame) SetObeyChild(ObeyChildVar bool) {

	xAspectFrameSetObeyChild(x.GoPointer(), ObeyChildVar)

}

var xAspectFrameSetRatio func(uintptr, float32)

// Sets the desired aspect ratio of the child.
func (x *AspectFrame) SetRatio(RatioVar float32) {

	xAspectFrameSetRatio(x.GoPointer(), RatioVar)

}

var xAspectFrameSetXalign func(uintptr, float32)

// Sets the horizontal alignment of the child within the allocation
// of the `GtkAspectFrame`.
func (x *AspectFrame) SetXalign(XalignVar float32) {

	xAspectFrameSetXalign(x.GoPointer(), XalignVar)

}

var xAspectFrameSetYalign func(uintptr, float32)

// Sets the vertical alignment of the child within the allocation
// of the `GtkAspectFrame`.
func (x *AspectFrame) SetYalign(YalignVar float32) {

	xAspectFrameSetYalign(x.GoPointer(), YalignVar)

}

func (c *AspectFrame) GoPointer() uintptr {
	return c.Ptr
}

func (c *AspectFrame) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *AspectFrame) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *AspectFrame) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *AspectFrame) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *AspectFrame) ResetState(StateVar AccessibleState) {

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
func (x *AspectFrame) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AspectFrame) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *AspectFrame) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AspectFrame) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *AspectFrame) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *AspectFrame) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *AspectFrame) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAspectFrameGLibType, lib, "gtk_aspect_frame_get_type")

	core.PuregoSafeRegister(&xNewAspectFrame, lib, "gtk_aspect_frame_new")

	core.PuregoSafeRegister(&xAspectFrameGetChild, lib, "gtk_aspect_frame_get_child")
	core.PuregoSafeRegister(&xAspectFrameGetObeyChild, lib, "gtk_aspect_frame_get_obey_child")
	core.PuregoSafeRegister(&xAspectFrameGetRatio, lib, "gtk_aspect_frame_get_ratio")
	core.PuregoSafeRegister(&xAspectFrameGetXalign, lib, "gtk_aspect_frame_get_xalign")
	core.PuregoSafeRegister(&xAspectFrameGetYalign, lib, "gtk_aspect_frame_get_yalign")
	core.PuregoSafeRegister(&xAspectFrameSetChild, lib, "gtk_aspect_frame_set_child")
	core.PuregoSafeRegister(&xAspectFrameSetObeyChild, lib, "gtk_aspect_frame_set_obey_child")
	core.PuregoSafeRegister(&xAspectFrameSetRatio, lib, "gtk_aspect_frame_set_ratio")
	core.PuregoSafeRegister(&xAspectFrameSetXalign, lib, "gtk_aspect_frame_set_xalign")
	core.PuregoSafeRegister(&xAspectFrameSetYalign, lib, "gtk_aspect_frame_set_yalign")

}
