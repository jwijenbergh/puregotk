// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ClampLayoutClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *ClampLayoutClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A layout manager constraining its children to a given size.
//
// &lt;picture&gt;
//
//	&lt;source srcset="clamp-wide-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="clamp-wide.png" alt="clamp-wide"&gt;
//
// &lt;/picture&gt;
// &lt;picture&gt;
//
//	&lt;source srcset="clamp-narrow-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="clamp-narrow.png" alt="clamp-narrow"&gt;
//
// &lt;/picture&gt;
//
// `AdwClampLayout` constraints the size of the widgets it contains to a given
// maximum size. It will constrain the width if it is horizontal, or the height
// if it is vertical. The expansion of the children from their minimum to their
// maximum size is eased out for a smooth transition.
//
// If a child requires more than the requested maximum size, it will be
// allocated the minimum size it can fit in instead.
//
// `AdwClampLayout` can scale with the text scale factor, use the
// [property@ClampLayout:unit] property to enable that behavior.
//
// See also: [class@Clamp], [class@ClampScrollable].
type ClampLayout struct {
	gtk.LayoutManager
}

var xClampLayoutGLibType func() types.GType

func ClampLayoutGLibType() types.GType {
	return xClampLayoutGLibType()
}

func ClampLayoutNewFromInternalPtr(ptr uintptr) *ClampLayout {
	cls := &ClampLayout{}
	cls.Ptr = ptr
	return cls
}

var xNewClampLayout func() uintptr

// Creates a new `AdwClampLayout`.
func NewClampLayout() *ClampLayout {
	var cls *ClampLayout

	cret := xNewClampLayout()

	if cret == 0 {
		return nil
	}
	cls = &ClampLayout{}
	cls.Ptr = cret
	return cls
}

var xClampLayoutGetMaximumSize func(uintptr) int

// Gets the maximum size allocated to the children.
func (x *ClampLayout) GetMaximumSize() int {

	cret := xClampLayoutGetMaximumSize(x.GoPointer())
	return cret
}

var xClampLayoutGetTighteningThreshold func(uintptr) int

// Gets the size above which the children are clamped.
func (x *ClampLayout) GetTighteningThreshold() int {

	cret := xClampLayoutGetTighteningThreshold(x.GoPointer())
	return cret
}

var xClampLayoutGetUnit func(uintptr) LengthUnit

// Gets the length unit for maximum size and tightening threshold.
func (x *ClampLayout) GetUnit() LengthUnit {

	cret := xClampLayoutGetUnit(x.GoPointer())
	return cret
}

var xClampLayoutSetMaximumSize func(uintptr, int)

// Sets the maximum size allocated to the children.
//
// It is the width if the layout is horizontal, or the height if it is vertical.
func (x *ClampLayout) SetMaximumSize(MaximumSizeVar int) {

	xClampLayoutSetMaximumSize(x.GoPointer(), MaximumSizeVar)

}

var xClampLayoutSetTighteningThreshold func(uintptr, int)

// Sets the size above which the children are clamped.
//
// Starting from this size, the layout will tighten its grip on the children,
// slowly allocating less and less of the available size up to the maximum
// allocated size. Below that threshold and below the maximum size, the children
// will be allocated all the available size.
//
// If the threshold is greater than the maximum size to allocate to the
// children, they will be allocated the whole size up to the maximum. If the
// threshold is lower than the minimum size to allocate to the children, that
// size will be used as the tightening threshold.
//
// Effectively, tightening the grip on a child before it reaches its maximum
// size makes transitions to and from the maximum size smoother when resizing.
func (x *ClampLayout) SetTighteningThreshold(TighteningThresholdVar int) {

	xClampLayoutSetTighteningThreshold(x.GoPointer(), TighteningThresholdVar)

}

var xClampLayoutSetUnit func(uintptr, LengthUnit)

// Sets the length unit for maximum size and tightening threshold.
//
// Allows the sizes to vary depending on the text scale factor.
func (x *ClampLayout) SetUnit(UnitVar LengthUnit) {

	xClampLayoutSetUnit(x.GoPointer(), UnitVar)

}

func (c *ClampLayout) GoPointer() uintptr {
	return c.Ptr
}

func (c *ClampLayout) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the orientation of the @orientable.
func (x *ClampLayout) GetOrientation() gtk.Orientation {

	cret := gtk.XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *ClampLayout) SetOrientation(OrientationVar gtk.Orientation) {

	gtk.XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xClampLayoutGLibType, lib, "adw_clamp_layout_get_type")

	core.PuregoSafeRegister(&xNewClampLayout, lib, "adw_clamp_layout_new")

	core.PuregoSafeRegister(&xClampLayoutGetMaximumSize, lib, "adw_clamp_layout_get_maximum_size")
	core.PuregoSafeRegister(&xClampLayoutGetTighteningThreshold, lib, "adw_clamp_layout_get_tightening_threshold")
	core.PuregoSafeRegister(&xClampLayoutGetUnit, lib, "adw_clamp_layout_get_unit")
	core.PuregoSafeRegister(&xClampLayoutSetMaximumSize, lib, "adw_clamp_layout_set_maximum_size")
	core.PuregoSafeRegister(&xClampLayoutSetTighteningThreshold, lib, "adw_clamp_layout_set_tightening_threshold")
	core.PuregoSafeRegister(&xClampLayoutSetUnit, lib, "adw_clamp_layout_set_unit")

}
