// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type CarouselIndicatorLinesClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *CarouselIndicatorLinesClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A lines indicator for [class@Carousel].
//
// &lt;picture&gt;
//
//	&lt;source srcset="carousel-indicator-dots-lines.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="carousel-indicator-lines.png" alt="carousel-indicator-lines"&gt;
//
// &lt;/picture&gt;
//
// The `AdwCarouselIndicatorLines` widget shows a set of lines for each page of
// a given [class@Carousel]. The carousel's active page is shown as another line
// that moves between them to match the carousel's position.
//
// See also [class@CarouselIndicatorDots].
//
// ## CSS nodes
//
// `AdwCarouselIndicatorLines` has a single CSS node with name
// `carouselindicatorlines`.
type CarouselIndicatorLines struct {
	gtk.Widget
}

var xCarouselIndicatorLinesGLibType func() types.GType

func CarouselIndicatorLinesGLibType() types.GType {
	return xCarouselIndicatorLinesGLibType()
}

func CarouselIndicatorLinesNewFromInternalPtr(ptr uintptr) *CarouselIndicatorLines {
	cls := &CarouselIndicatorLines{}
	cls.Ptr = ptr
	return cls
}

var xNewCarouselIndicatorLines func() uintptr

// Creates a new `AdwCarouselIndicatorLines`.
func NewCarouselIndicatorLines() *CarouselIndicatorLines {
	var cls *CarouselIndicatorLines

	cret := xNewCarouselIndicatorLines()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CarouselIndicatorLines{}
	cls.Ptr = cret
	return cls
}

var xCarouselIndicatorLinesGetCarousel func(uintptr) uintptr

// Gets the displayed carousel.
func (x *CarouselIndicatorLines) GetCarousel() *Carousel {
	var cls *Carousel

	cret := xCarouselIndicatorLinesGetCarousel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Carousel{}
	cls.Ptr = cret
	return cls
}

var xCarouselIndicatorLinesSetCarousel func(uintptr, uintptr)

// Sets the displayed carousel.
func (x *CarouselIndicatorLines) SetCarousel(CarouselVar *Carousel) {

	xCarouselIndicatorLinesSetCarousel(x.GoPointer(), CarouselVar.GoPointer())

}

func (c *CarouselIndicatorLines) GoPointer() uintptr {
	return c.Ptr
}

func (c *CarouselIndicatorLines) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *CarouselIndicatorLines) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *CarouselIndicatorLines) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *CarouselIndicatorLines) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *CarouselIndicatorLines) ResetState(StateVar gtk.AccessibleState) {

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
func (x *CarouselIndicatorLines) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CarouselIndicatorLines) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []gtk.AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *CarouselIndicatorLines) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CarouselIndicatorLines) UpdateRelationValue(NRelationsVar int, RelationsVar []gtk.AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *CarouselIndicatorLines) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CarouselIndicatorLines) UpdateStateValue(NStatesVar int, StatesVar []gtk.AccessibleState, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *CarouselIndicatorLines) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *CarouselIndicatorLines) GetOrientation() gtk.Orientation {

	cret := gtk.XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *CarouselIndicatorLines) SetOrientation(OrientationVar gtk.Orientation) {

	gtk.XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xCarouselIndicatorLinesGLibType, lib, "adw_carousel_indicator_lines_get_type")

	core.PuregoSafeRegister(&xNewCarouselIndicatorLines, lib, "adw_carousel_indicator_lines_new")

	core.PuregoSafeRegister(&xCarouselIndicatorLinesGetCarousel, lib, "adw_carousel_indicator_lines_get_carousel")
	core.PuregoSafeRegister(&xCarouselIndicatorLinesSetCarousel, lib, "adw_carousel_indicator_lines_set_carousel")

}
