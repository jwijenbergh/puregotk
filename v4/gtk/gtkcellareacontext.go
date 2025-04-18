// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

type CellAreaContextClass struct {
	_ structs.HostLayout

	ParentClass uintptr

	Padding [8]uintptr
}

func (x *CellAreaContextClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type CellAreaContextPrivate struct {
	_ structs.HostLayout
}

func (x *CellAreaContextPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Stores geometrical information for a series of rows in a GtkCellArea
//
// The `GtkCellAreaContext` object is created by a given `GtkCellArea`
// implementation via its `GtkCellAreaClass.create_context()` virtual
// method and is used to store cell sizes and alignments for a series of
// `GtkTreeModel` rows that are requested and rendered in the same context.
//
// `GtkCellLayout` widgets can create any number of contexts in which to
// request and render groups of data rows. However, it’s important that the
// same context which was used to request sizes for a given `GtkTreeModel`
// row also be used for the same row when calling other `GtkCellArea` APIs
// such as gtk_cell_area_render() and gtk_cell_area_event().
type CellAreaContext struct {
	gobject.Object
}

var xCellAreaContextGLibType func() types.GType

func CellAreaContextGLibType() types.GType {
	return xCellAreaContextGLibType()
}

func CellAreaContextNewFromInternalPtr(ptr uintptr) *CellAreaContext {
	cls := &CellAreaContext{}
	cls.Ptr = ptr
	return cls
}

var xCellAreaContextAllocate func(uintptr, int, int)

// Allocates a width and/or a height for all rows which are to be
// rendered with @context.
//
// Usually allocation is performed only horizontally or sometimes
// vertically since a group of rows are usually rendered side by
// side vertically or horizontally and share either the same width
// or the same height. Sometimes they are allocated in both horizontal
// and vertical orientations producing a homogeneous effect of the
// rows. This is generally the case for `GtkTreeView` when
// `GtkTreeView:fixed-height-mode` is enabled.
func (x *CellAreaContext) Allocate(WidthVar int, HeightVar int) {

	xCellAreaContextAllocate(x.GoPointer(), WidthVar, HeightVar)

}

var xCellAreaContextGetAllocation func(uintptr, int, int)

// Fetches the current allocation size for @context.
//
// If the context was not allocated in width or height, or if the
// context was recently reset with gtk_cell_area_context_reset(),
// the returned value will be -1.
func (x *CellAreaContext) GetAllocation(WidthVar int, HeightVar int) {

	xCellAreaContextGetAllocation(x.GoPointer(), WidthVar, HeightVar)

}

var xCellAreaContextGetArea func(uintptr) uintptr

// Fetches the `GtkCellArea` this @context was created by.
//
// This is generally unneeded by layouting widgets; however,
// it is important for the context implementation itself to
// fetch information about the area it is being used for.
//
// For instance at `GtkCellAreaContextClass.allocate()` time
// it’s important to know details about any cell spacing
// that the `GtkCellArea` is configured with in order to
// compute a proper allocation.
func (x *CellAreaContext) GetArea() *CellArea {
	var cls *CellArea

	cret := xCellAreaContextGetArea(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellArea{}
	cls.Ptr = cret
	return cls
}

var xCellAreaContextGetPreferredHeight func(uintptr, int, int)

// Gets the accumulative preferred height for all rows which have been
// requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a `GtkCellArea`, the returned values are 0.
func (x *CellAreaContext) GetPreferredHeight(MinimumHeightVar int, NaturalHeightVar int) {

	xCellAreaContextGetPreferredHeight(x.GoPointer(), MinimumHeightVar, NaturalHeightVar)

}

var xCellAreaContextGetPreferredHeightForWidth func(uintptr, int, int, int)

// Gets the accumulative preferred height for @width for all rows
// which have been requested for the same said @width with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a `GtkCellArea`, the returned values are -1.
func (x *CellAreaContext) GetPreferredHeightForWidth(WidthVar int, MinimumHeightVar int, NaturalHeightVar int) {

	xCellAreaContextGetPreferredHeightForWidth(x.GoPointer(), WidthVar, MinimumHeightVar, NaturalHeightVar)

}

var xCellAreaContextGetPreferredWidth func(uintptr, int, int)

// Gets the accumulative preferred width for all rows which have been
// requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a `GtkCellArea`, the returned values are 0.
func (x *CellAreaContext) GetPreferredWidth(MinimumWidthVar int, NaturalWidthVar int) {

	xCellAreaContextGetPreferredWidth(x.GoPointer(), MinimumWidthVar, NaturalWidthVar)

}

var xCellAreaContextGetPreferredWidthForHeight func(uintptr, int, int, int)

// Gets the accumulative preferred width for @height for all rows which
// have been requested for the same said @height with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever
// requesting the size of a `GtkCellArea`, the returned values are -1.
func (x *CellAreaContext) GetPreferredWidthForHeight(HeightVar int, MinimumWidthVar int, NaturalWidthVar int) {

	xCellAreaContextGetPreferredWidthForHeight(x.GoPointer(), HeightVar, MinimumWidthVar, NaturalWidthVar)

}

var xCellAreaContextPushPreferredHeight func(uintptr, int, int)

// Causes the minimum and/or natural height to grow if the new
// proposed sizes exceed the current minimum and natural height.
//
// This is used by `GtkCellAreaContext` implementations during
// the request process over a series of `GtkTreeModel` rows to
// progressively push the requested height over a series of
// gtk_cell_area_get_preferred_height() requests.
func (x *CellAreaContext) PushPreferredHeight(MinimumHeightVar int, NaturalHeightVar int) {

	xCellAreaContextPushPreferredHeight(x.GoPointer(), MinimumHeightVar, NaturalHeightVar)

}

var xCellAreaContextPushPreferredWidth func(uintptr, int, int)

// Causes the minimum and/or natural width to grow if the new
// proposed sizes exceed the current minimum and natural width.
//
// This is used by `GtkCellAreaContext` implementations during
// the request process over a series of `GtkTreeModel` rows to
// progressively push the requested width over a series of
// gtk_cell_area_get_preferred_width() requests.
func (x *CellAreaContext) PushPreferredWidth(MinimumWidthVar int, NaturalWidthVar int) {

	xCellAreaContextPushPreferredWidth(x.GoPointer(), MinimumWidthVar, NaturalWidthVar)

}

var xCellAreaContextReset func(uintptr)

// Resets any previously cached request and allocation
// data.
//
// When underlying `GtkTreeModel` data changes its
// important to reset the context if the content
// size is allowed to shrink. If the content size
// is only allowed to grow (this is usually an option
// for views rendering large data stores as a measure
// of optimization), then only the row that changed
// or was inserted needs to be (re)requested with
// gtk_cell_area_get_preferred_width().
//
// When the new overall size of the context requires
// that the allocated size changes (or whenever this
// allocation changes at all), the variable row
// sizes need to be re-requested for every row.
//
// For instance, if the rows are displayed all with
// the same width from top to bottom then a change
// in the allocated width necessitates a recalculation
// of all the displayed row heights using
// gtk_cell_area_get_preferred_height_for_width().
func (x *CellAreaContext) Reset() {

	xCellAreaContextReset(x.GoPointer())

}

func (c *CellAreaContext) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellAreaContext) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xCellAreaContextGLibType, lib, "gtk_cell_area_context_get_type")

	core.PuregoSafeRegister(&xCellAreaContextAllocate, lib, "gtk_cell_area_context_allocate")
	core.PuregoSafeRegister(&xCellAreaContextGetAllocation, lib, "gtk_cell_area_context_get_allocation")
	core.PuregoSafeRegister(&xCellAreaContextGetArea, lib, "gtk_cell_area_context_get_area")
	core.PuregoSafeRegister(&xCellAreaContextGetPreferredHeight, lib, "gtk_cell_area_context_get_preferred_height")
	core.PuregoSafeRegister(&xCellAreaContextGetPreferredHeightForWidth, lib, "gtk_cell_area_context_get_preferred_height_for_width")
	core.PuregoSafeRegister(&xCellAreaContextGetPreferredWidth, lib, "gtk_cell_area_context_get_preferred_width")
	core.PuregoSafeRegister(&xCellAreaContextGetPreferredWidthForHeight, lib, "gtk_cell_area_context_get_preferred_width_for_height")
	core.PuregoSafeRegister(&xCellAreaContextPushPreferredHeight, lib, "gtk_cell_area_context_push_preferred_height")
	core.PuregoSafeRegister(&xCellAreaContextPushPreferredWidth, lib, "gtk_cell_area_context_push_preferred_width")
	core.PuregoSafeRegister(&xCellAreaContextReset, lib, "gtk_cell_area_context_reset")

}
