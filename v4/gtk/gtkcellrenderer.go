// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

type CellRendererClass struct {
	_ structs.HostLayout

	ParentClass uintptr

	Padding [8]uintptr
}

func (x *CellRendererClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type CellRendererClassPrivate struct {
	_ structs.HostLayout
}

func (x *CellRendererClassPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type CellRendererPrivate struct {
	_ structs.HostLayout
}

func (x *CellRendererPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Tells how a cell is to be rendered.
type CellRendererState int

var xCellRendererStateGLibType func() types.GType

func CellRendererStateGLibType() types.GType {
	return xCellRendererStateGLibType()
}

const (

	// The cell is currently selected, and
	//  probably has a selection colored background to render to.
	CellRendererSelectedValue CellRendererState = 1
	// The mouse is hovering over the cell.
	CellRendererPrelitValue CellRendererState = 2
	// The cell is drawn in an insensitive manner
	CellRendererInsensitiveValue CellRendererState = 4
	// The cell is in a sorted row
	CellRendererSortedValue CellRendererState = 8
	// The cell is in the focus row.
	CellRendererFocusedValue CellRendererState = 16
	// The cell is in a row that can be expanded
	CellRendererExpandableValue CellRendererState = 32
	// The cell is in a row that is expanded
	CellRendererExpandedValue CellRendererState = 64
)

// Identifies how the user can interact with a particular cell.
type CellRendererMode int

var xCellRendererModeGLibType func() types.GType

func CellRendererModeGLibType() types.GType {
	return xCellRendererModeGLibType()
}

const (

	// The cell is just for display
	//  and cannot be interacted with.  Note that this doesn’t mean that eg. the
	//  row being drawn can’t be selected -- just that a particular element of
	//  it cannot be individually modified.
	CellRendererModeInertValue CellRendererMode = 0
	// The cell can be clicked.
	CellRendererModeActivatableValue CellRendererMode = 1
	// The cell can be edited or otherwise modified.
	CellRendererModeEditableValue CellRendererMode = 2
)

// An object for rendering a single cell
//
// The `GtkCellRenderer` is a base class of a set of objects used for
// rendering a cell to a `cairo_t`.  These objects are used primarily by
// the `GtkTreeView` widget, though they aren’t tied to them in any
// specific way.  It is worth noting that `GtkCellRenderer` is not a
// `GtkWidget` and cannot be treated as such.
//
// The primary use of a `GtkCellRenderer` is for drawing a certain graphical
// elements on a `cairo_t`. Typically, one cell renderer is used to
// draw many cells on the screen.  To this extent, it isn’t expected that a
// CellRenderer keep any permanent state around.  Instead, any state is set
// just prior to use using `GObject`s property system.  Then, the
// cell is measured using gtk_cell_renderer_get_preferred_size(). Finally, the cell
// is rendered in the correct location using gtk_cell_renderer_snapshot().
//
// There are a number of rules that must be followed when writing a new
// `GtkCellRenderer`.  First and foremost, it’s important that a certain set
// of properties will always yield a cell renderer of the same size,
// barring a style change. The `GtkCellRenderer` also has a number of
// generic properties that are expected to be honored by all children.
//
// Beyond merely rendering a cell, cell renderers can optionally
// provide active user interface elements. A cell renderer can be
// “activatable” like `GtkCellRenderer`Toggle,
// which toggles when it gets activated by a mouse click, or it can be
// “editable” like `GtkCellRenderer`Text, which
// allows the user to edit the text using a widget implementing the
// `GtkCellEditable` interface, e.g. `GtkEntry`.
// To make a cell renderer activatable or editable, you have to
// implement the `GtkCellRenderer`Class.activate or
// `GtkCellRenderer`Class.start_editing virtual functions, respectively.
//
// Many properties of `GtkCellRenderer` and its subclasses have a
// corresponding “set” property, e.g. “cell-background-set” corresponds
// to “cell-background”. These “set” properties reflect whether a property
// has been set or not. You should not set them independently.
type CellRenderer struct {
	gobject.InitiallyUnowned
}

var xCellRendererGLibType func() types.GType

func CellRendererGLibType() types.GType {
	return xCellRendererGLibType()
}

func CellRendererNewFromInternalPtr(ptr uintptr) *CellRenderer {
	cls := &CellRenderer{}
	cls.Ptr = ptr
	return cls
}

var xCellRendererActivate func(uintptr, uintptr, uintptr, string, *gdk.Rectangle, *gdk.Rectangle, CellRendererState) bool

// Passes an activate event to the cell renderer for possible processing.
// Some cell renderers may use events; for example, `GtkCellRendererToggle`
// toggles when it gets a mouse click.
func (x *CellRenderer) Activate(EventVar *gdk.Event, WidgetVar *Widget, PathVar string, BackgroundAreaVar *gdk.Rectangle, CellAreaVar *gdk.Rectangle, FlagsVar CellRendererState) bool {

	cret := xCellRendererActivate(x.GoPointer(), EventVar.GoPointer(), WidgetVar.GoPointer(), PathVar, BackgroundAreaVar, CellAreaVar, FlagsVar)
	return cret
}

var xCellRendererGetAlignedArea func(uintptr, uintptr, CellRendererState, *gdk.Rectangle, *gdk.Rectangle)

// Gets the aligned area used by @cell inside @cell_area. Used for finding
// the appropriate edit and focus rectangle.
func (x *CellRenderer) GetAlignedArea(WidgetVar *Widget, FlagsVar CellRendererState, CellAreaVar *gdk.Rectangle, AlignedAreaVar *gdk.Rectangle) {

	xCellRendererGetAlignedArea(x.GoPointer(), WidgetVar.GoPointer(), FlagsVar, CellAreaVar, AlignedAreaVar)

}

var xCellRendererGetAlignment func(uintptr, float32, float32)

// Fills in @xalign and @yalign with the appropriate values of @cell.
func (x *CellRenderer) GetAlignment(XalignVar float32, YalignVar float32) {

	xCellRendererGetAlignment(x.GoPointer(), XalignVar, YalignVar)

}

var xCellRendererGetFixedSize func(uintptr, int, int)

// Fills in @width and @height with the appropriate size of @cell.
func (x *CellRenderer) GetFixedSize(WidthVar int, HeightVar int) {

	xCellRendererGetFixedSize(x.GoPointer(), WidthVar, HeightVar)

}

var xCellRendererGetIsExpanded func(uintptr) bool

// Checks whether the given `GtkCellRenderer` is expanded.
func (x *CellRenderer) GetIsExpanded() bool {

	cret := xCellRendererGetIsExpanded(x.GoPointer())
	return cret
}

var xCellRendererGetIsExpander func(uintptr) bool

// Checks whether the given `GtkCellRenderer` is an expander.
func (x *CellRenderer) GetIsExpander() bool {

	cret := xCellRendererGetIsExpander(x.GoPointer())
	return cret
}

var xCellRendererGetPadding func(uintptr, int, int)

// Fills in @xpad and @ypad with the appropriate values of @cell.
func (x *CellRenderer) GetPadding(XpadVar int, YpadVar int) {

	xCellRendererGetPadding(x.GoPointer(), XpadVar, YpadVar)

}

var xCellRendererGetPreferredHeight func(uintptr, uintptr, int, int)

// Retrieves a renderer’s natural size when rendered to @widget.
func (x *CellRenderer) GetPreferredHeight(WidgetVar *Widget, MinimumSizeVar int, NaturalSizeVar int) {

	xCellRendererGetPreferredHeight(x.GoPointer(), WidgetVar.GoPointer(), MinimumSizeVar, NaturalSizeVar)

}

var xCellRendererGetPreferredHeightForWidth func(uintptr, uintptr, int, int, int)

// Retrieves a cell renderers’s minimum and natural height if it were rendered to
// @widget with the specified @width.
func (x *CellRenderer) GetPreferredHeightForWidth(WidgetVar *Widget, WidthVar int, MinimumHeightVar int, NaturalHeightVar int) {

	xCellRendererGetPreferredHeightForWidth(x.GoPointer(), WidgetVar.GoPointer(), WidthVar, MinimumHeightVar, NaturalHeightVar)

}

var xCellRendererGetPreferredSize func(uintptr, uintptr, *Requisition, *Requisition)

// Retrieves the minimum and natural size of a cell taking
// into account the widget’s preference for height-for-width management.
func (x *CellRenderer) GetPreferredSize(WidgetVar *Widget, MinimumSizeVar *Requisition, NaturalSizeVar *Requisition) {

	xCellRendererGetPreferredSize(x.GoPointer(), WidgetVar.GoPointer(), MinimumSizeVar, NaturalSizeVar)

}

var xCellRendererGetPreferredWidth func(uintptr, uintptr, int, int)

// Retrieves a renderer’s natural size when rendered to @widget.
func (x *CellRenderer) GetPreferredWidth(WidgetVar *Widget, MinimumSizeVar int, NaturalSizeVar int) {

	xCellRendererGetPreferredWidth(x.GoPointer(), WidgetVar.GoPointer(), MinimumSizeVar, NaturalSizeVar)

}

var xCellRendererGetPreferredWidthForHeight func(uintptr, uintptr, int, int, int)

// Retrieves a cell renderers’s minimum and natural width if it were rendered to
// @widget with the specified @height.
func (x *CellRenderer) GetPreferredWidthForHeight(WidgetVar *Widget, HeightVar int, MinimumWidthVar int, NaturalWidthVar int) {

	xCellRendererGetPreferredWidthForHeight(x.GoPointer(), WidgetVar.GoPointer(), HeightVar, MinimumWidthVar, NaturalWidthVar)

}

var xCellRendererGetRequestMode func(uintptr) SizeRequestMode

// Gets whether the cell renderer prefers a height-for-width layout
// or a width-for-height layout.
func (x *CellRenderer) GetRequestMode() SizeRequestMode {

	cret := xCellRendererGetRequestMode(x.GoPointer())
	return cret
}

var xCellRendererGetSensitive func(uintptr) bool

// Returns the cell renderer’s sensitivity.
func (x *CellRenderer) GetSensitive() bool {

	cret := xCellRendererGetSensitive(x.GoPointer())
	return cret
}

var xCellRendererGetState func(uintptr, uintptr, CellRendererState) StateFlags

// Translates the cell renderer state to `GtkStateFlags`,
// based on the cell renderer and widget sensitivity, and
// the given `GtkCellRenderer`State.
func (x *CellRenderer) GetState(WidgetVar *Widget, CellStateVar CellRendererState) StateFlags {

	cret := xCellRendererGetState(x.GoPointer(), WidgetVar.GoPointer(), CellStateVar)
	return cret
}

var xCellRendererGetVisible func(uintptr) bool

// Returns the cell renderer’s visibility.
func (x *CellRenderer) GetVisible() bool {

	cret := xCellRendererGetVisible(x.GoPointer())
	return cret
}

var xCellRendererIsActivatable func(uintptr) bool

// Checks whether the cell renderer can do something when activated.
func (x *CellRenderer) IsActivatable() bool {

	cret := xCellRendererIsActivatable(x.GoPointer())
	return cret
}

var xCellRendererSetAlignment func(uintptr, float32, float32)

// Sets the renderer’s alignment within its available space.
func (x *CellRenderer) SetAlignment(XalignVar float32, YalignVar float32) {

	xCellRendererSetAlignment(x.GoPointer(), XalignVar, YalignVar)

}

var xCellRendererSetFixedSize func(uintptr, int, int)

// Sets the renderer size to be explicit, independent of the properties set.
func (x *CellRenderer) SetFixedSize(WidthVar int, HeightVar int) {

	xCellRendererSetFixedSize(x.GoPointer(), WidthVar, HeightVar)

}

var xCellRendererSetIsExpanded func(uintptr, bool)

// Sets whether the given `GtkCellRenderer` is expanded.
func (x *CellRenderer) SetIsExpanded(IsExpandedVar bool) {

	xCellRendererSetIsExpanded(x.GoPointer(), IsExpandedVar)

}

var xCellRendererSetIsExpander func(uintptr, bool)

// Sets whether the given `GtkCellRenderer` is an expander.
func (x *CellRenderer) SetIsExpander(IsExpanderVar bool) {

	xCellRendererSetIsExpander(x.GoPointer(), IsExpanderVar)

}

var xCellRendererSetPadding func(uintptr, int, int)

// Sets the renderer’s padding.
func (x *CellRenderer) SetPadding(XpadVar int, YpadVar int) {

	xCellRendererSetPadding(x.GoPointer(), XpadVar, YpadVar)

}

var xCellRendererSetSensitive func(uintptr, bool)

// Sets the cell renderer’s sensitivity.
func (x *CellRenderer) SetSensitive(SensitiveVar bool) {

	xCellRendererSetSensitive(x.GoPointer(), SensitiveVar)

}

var xCellRendererSetVisible func(uintptr, bool)

// Sets the cell renderer’s visibility.
func (x *CellRenderer) SetVisible(VisibleVar bool) {

	xCellRendererSetVisible(x.GoPointer(), VisibleVar)

}

var xCellRendererSnapshot func(uintptr, uintptr, uintptr, *gdk.Rectangle, *gdk.Rectangle, CellRendererState)

// Invokes the virtual render function of the `GtkCellRenderer`. The three
// passed-in rectangles are areas in @cr. Most renderers will draw within
// @cell_area; the xalign, yalign, xpad, and ypad fields of the `GtkCellRenderer`
// should be honored with respect to @cell_area. @background_area includes the
// blank space around the cell, and also the area containing the tree expander;
// so the @background_area rectangles for all cells tile to cover the entire
// @window.
func (x *CellRenderer) Snapshot(SnapshotVar *Snapshot, WidgetVar *Widget, BackgroundAreaVar *gdk.Rectangle, CellAreaVar *gdk.Rectangle, FlagsVar CellRendererState) {

	xCellRendererSnapshot(x.GoPointer(), SnapshotVar.GoPointer(), WidgetVar.GoPointer(), BackgroundAreaVar, CellAreaVar, FlagsVar)

}

var xCellRendererStartEditing func(uintptr, uintptr, uintptr, string, *gdk.Rectangle, *gdk.Rectangle, CellRendererState) uintptr

// Starts editing the contents of this @cell, through a new `GtkCellEditable`
// widget created by the `GtkCellRenderer`Class.start_editing virtual function.
func (x *CellRenderer) StartEditing(EventVar *gdk.Event, WidgetVar *Widget, PathVar string, BackgroundAreaVar *gdk.Rectangle, CellAreaVar *gdk.Rectangle, FlagsVar CellRendererState) *CellEditableBase {
	var cls *CellEditableBase

	cret := xCellRendererStartEditing(x.GoPointer(), EventVar.GoPointer(), WidgetVar.GoPointer(), PathVar, BackgroundAreaVar, CellAreaVar, FlagsVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellEditableBase{}
	cls.Ptr = cret
	return cls
}

var xCellRendererStopEditing func(uintptr, bool)

// Informs the cell renderer that the editing is stopped.
// If @canceled is %TRUE, the cell renderer will emit the
// `GtkCellRenderer`::editing-canceled signal.
//
// This function should be called by cell renderer implementations
// in response to the `GtkCellEditable::editing-done` signal of
// `GtkCellEditable`.
func (x *CellRenderer) StopEditing(CanceledVar bool) {

	xCellRendererStopEditing(x.GoPointer(), CanceledVar)

}

func (c *CellRenderer) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellRenderer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// This signal gets emitted when the user cancels the process of editing a
// cell.  For example, an editable cell renderer could be written to cancel
// editing when the user presses Escape.
//
// See also: gtk_cell_renderer_stop_editing().
func (x *CellRenderer) ConnectEditingCanceled(cb *func(CellRenderer)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "editing-canceled", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := CellRenderer{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "editing-canceled", cbRefPtr)
}

// This signal gets emitted when a cell starts to be edited.
// The intended use of this signal is to do special setup
// on @editable, e.g. adding a `GtkEntryCompletion` or setting
// up additional columns in a `GtkComboBox`.
//
// See gtk_cell_editable_start_editing() for information on the lifecycle of
// the @editable and a way to do setup that doesn’t depend on the @renderer.
//
// Note that GTK doesn't guarantee that cell renderers will
// continue to use the same kind of widget for editing in future
// releases, therefore you should check the type of @editable
// before doing any specific setup, as in the following example:
// |[&lt;!-- language="C" --&gt;
// static void
// text_editing_started (GtkCellRenderer *cell,
//
//	GtkCellEditable *editable,
//	const char      *path,
//	gpointer         data)
//
//	{
//	  if (GTK_IS_ENTRY (editable))
//	    {
//	      GtkEntry *entry = GTK_ENTRY (editable);
//
//	      // ... create a GtkEntryCompletion
//
//	      gtk_entry_set_completion (entry, completion);
//	    }
//	}
//
// ]|
func (x *CellRenderer) ConnectEditingStarted(cb *func(CellRenderer, uintptr, string)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "editing-started", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, EditableVarp uintptr, PathVarp string) {
		fa := CellRenderer{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, EditableVarp, PathVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "editing-started", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xCellRendererStateGLibType, lib, "gtk_cell_renderer_state_get_type")

	core.PuregoSafeRegister(&xCellRendererModeGLibType, lib, "gtk_cell_renderer_mode_get_type")

	core.PuregoSafeRegister(&xCellRendererGLibType, lib, "gtk_cell_renderer_get_type")

	core.PuregoSafeRegister(&xCellRendererActivate, lib, "gtk_cell_renderer_activate")
	core.PuregoSafeRegister(&xCellRendererGetAlignedArea, lib, "gtk_cell_renderer_get_aligned_area")
	core.PuregoSafeRegister(&xCellRendererGetAlignment, lib, "gtk_cell_renderer_get_alignment")
	core.PuregoSafeRegister(&xCellRendererGetFixedSize, lib, "gtk_cell_renderer_get_fixed_size")
	core.PuregoSafeRegister(&xCellRendererGetIsExpanded, lib, "gtk_cell_renderer_get_is_expanded")
	core.PuregoSafeRegister(&xCellRendererGetIsExpander, lib, "gtk_cell_renderer_get_is_expander")
	core.PuregoSafeRegister(&xCellRendererGetPadding, lib, "gtk_cell_renderer_get_padding")
	core.PuregoSafeRegister(&xCellRendererGetPreferredHeight, lib, "gtk_cell_renderer_get_preferred_height")
	core.PuregoSafeRegister(&xCellRendererGetPreferredHeightForWidth, lib, "gtk_cell_renderer_get_preferred_height_for_width")
	core.PuregoSafeRegister(&xCellRendererGetPreferredSize, lib, "gtk_cell_renderer_get_preferred_size")
	core.PuregoSafeRegister(&xCellRendererGetPreferredWidth, lib, "gtk_cell_renderer_get_preferred_width")
	core.PuregoSafeRegister(&xCellRendererGetPreferredWidthForHeight, lib, "gtk_cell_renderer_get_preferred_width_for_height")
	core.PuregoSafeRegister(&xCellRendererGetRequestMode, lib, "gtk_cell_renderer_get_request_mode")
	core.PuregoSafeRegister(&xCellRendererGetSensitive, lib, "gtk_cell_renderer_get_sensitive")
	core.PuregoSafeRegister(&xCellRendererGetState, lib, "gtk_cell_renderer_get_state")
	core.PuregoSafeRegister(&xCellRendererGetVisible, lib, "gtk_cell_renderer_get_visible")
	core.PuregoSafeRegister(&xCellRendererIsActivatable, lib, "gtk_cell_renderer_is_activatable")
	core.PuregoSafeRegister(&xCellRendererSetAlignment, lib, "gtk_cell_renderer_set_alignment")
	core.PuregoSafeRegister(&xCellRendererSetFixedSize, lib, "gtk_cell_renderer_set_fixed_size")
	core.PuregoSafeRegister(&xCellRendererSetIsExpanded, lib, "gtk_cell_renderer_set_is_expanded")
	core.PuregoSafeRegister(&xCellRendererSetIsExpander, lib, "gtk_cell_renderer_set_is_expander")
	core.PuregoSafeRegister(&xCellRendererSetPadding, lib, "gtk_cell_renderer_set_padding")
	core.PuregoSafeRegister(&xCellRendererSetSensitive, lib, "gtk_cell_renderer_set_sensitive")
	core.PuregoSafeRegister(&xCellRendererSetVisible, lib, "gtk_cell_renderer_set_visible")
	core.PuregoSafeRegister(&xCellRendererSnapshot, lib, "gtk_cell_renderer_snapshot")
	core.PuregoSafeRegister(&xCellRendererStartEditing, lib, "gtk_cell_renderer_start_editing")
	core.PuregoSafeRegister(&xCellRendererStopEditing, lib, "gtk_cell_renderer_stop_editing")

}
