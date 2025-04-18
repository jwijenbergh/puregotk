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

type GridViewClass struct {
	_ structs.HostLayout
}

func (x *GridViewClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkGridView` presents a large dynamic grid of items.
//
// `GtkGridView` uses its factory to generate one child widget for each
// visible item and shows them in a grid. The orientation of the grid view
// determines if the grid reflows vertically or horizontally.
//
// `GtkGridView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// [property@Gtk.GridView:enable-rubberband].
//
// To learn more about the list widget framework, see the
// [overview](section-list-widget.html).
//
// # CSS nodes
//
// ```
// gridview
// ├── child[.activatable]
// │
// ├── child[.activatable]
// │
// ┊
// ╰── [rubberband]
// ```
//
// `GtkGridView` uses a single CSS node with name `gridview`. Each child uses
// a single CSS node with name `child`. If the [property@Gtk.ListItem:activatable]
// property is set, the corresponding row will have the `.activatable` style
// class. For rubberband selection, a subnode with name `rubberband` is used.
//
// # Accessibility
//
// `GtkGridView` uses the %GTK_ACCESSIBLE_ROLE_GRID role, and the items
// use the %GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type GridView struct {
	ListBase
}

var xGridViewGLibType func() types.GType

func GridViewGLibType() types.GType {
	return xGridViewGLibType()
}

func GridViewNewFromInternalPtr(ptr uintptr) *GridView {
	cls := &GridView{}
	cls.Ptr = ptr
	return cls
}

var xNewGridView func(uintptr, uintptr) uintptr

// Creates a new `GtkGridView` that uses the given @factory for
// mapping items to widgets.
//
// The function takes ownership of the
// arguments, so you can write code like
// ```c
// grid_view = gtk_grid_view_new (create_model (),
//
//	gtk_builder_list_item_factory_new_from_resource ("/resource.ui"));
//
// ```
func NewGridView(ModelVar SelectionModel, FactoryVar *ListItemFactory) *GridView {
	var cls *GridView

	cret := xNewGridView(ModelVar.GoPointer(), FactoryVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &GridView{}
	cls.Ptr = cret
	return cls
}

var xGridViewGetEnableRubberband func(uintptr) bool

// Returns whether rows can be selected by dragging with the mouse.
func (x *GridView) GetEnableRubberband() bool {

	cret := xGridViewGetEnableRubberband(x.GoPointer())
	return cret
}

var xGridViewGetFactory func(uintptr) uintptr

// Gets the factory that's currently used to populate list items.
func (x *GridView) GetFactory() *ListItemFactory {
	var cls *ListItemFactory

	cret := xGridViewGetFactory(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ListItemFactory{}
	cls.Ptr = cret
	return cls
}

var xGridViewGetMaxColumns func(uintptr) uint

// Gets the maximum number of columns that the grid will use.
func (x *GridView) GetMaxColumns() uint {

	cret := xGridViewGetMaxColumns(x.GoPointer())
	return cret
}

var xGridViewGetMinColumns func(uintptr) uint

// Gets the minimum number of columns that the grid will use.
func (x *GridView) GetMinColumns() uint {

	cret := xGridViewGetMinColumns(x.GoPointer())
	return cret
}

var xGridViewGetModel func(uintptr) uintptr

// Gets the model that's currently used to read the items displayed.
func (x *GridView) GetModel() *SelectionModelBase {
	var cls *SelectionModelBase

	cret := xGridViewGetModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &SelectionModelBase{}
	cls.Ptr = cret
	return cls
}

var xGridViewGetSingleClickActivate func(uintptr) bool

// Returns whether items will be activated on single click and
// selected on hover.
func (x *GridView) GetSingleClickActivate() bool {

	cret := xGridViewGetSingleClickActivate(x.GoPointer())
	return cret
}

var xGridViewSetEnableRubberband func(uintptr, bool)

// Sets whether selections can be changed by dragging with the mouse.
func (x *GridView) SetEnableRubberband(EnableRubberbandVar bool) {

	xGridViewSetEnableRubberband(x.GoPointer(), EnableRubberbandVar)

}

var xGridViewSetFactory func(uintptr, uintptr)

// Sets the `GtkListItemFactory` to use for populating list items.
func (x *GridView) SetFactory(FactoryVar *ListItemFactory) {

	xGridViewSetFactory(x.GoPointer(), FactoryVar.GoPointer())

}

var xGridViewSetMaxColumns func(uintptr, uint)

// Sets the maximum number of columns to use.
//
// This number must be at least 1.
//
// If @max_columns is smaller than the minimum set via
// [method@Gtk.GridView.set_min_columns], that value is used instead.
func (x *GridView) SetMaxColumns(MaxColumnsVar uint) {

	xGridViewSetMaxColumns(x.GoPointer(), MaxColumnsVar)

}

var xGridViewSetMinColumns func(uintptr, uint)

// Sets the minimum number of columns to use.
//
// This number must be at least 1.
//
// If @min_columns is smaller than the minimum set via
// [method@Gtk.GridView.set_max_columns], that value is ignored.
func (x *GridView) SetMinColumns(MinColumnsVar uint) {

	xGridViewSetMinColumns(x.GoPointer(), MinColumnsVar)

}

var xGridViewSetModel func(uintptr, uintptr)

// Sets the imodel to use.
//
// This must be a [iface@Gtk.SelectionModel].
func (x *GridView) SetModel(ModelVar SelectionModel) {

	xGridViewSetModel(x.GoPointer(), ModelVar.GoPointer())

}

var xGridViewSetSingleClickActivate func(uintptr, bool)

// Sets whether items should be activated on single click and
// selected on hover.
func (x *GridView) SetSingleClickActivate(SingleClickActivateVar bool) {

	xGridViewSetSingleClickActivate(x.GoPointer(), SingleClickActivateVar)

}

func (c *GridView) GoPointer() uintptr {
	return c.Ptr
}

func (c *GridView) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a cell has been activated by the user,
// usually via activating the GtkGridView|list.activate-item action.
//
// This allows for a convenient way to handle activation in a gridview.
// See [property@Gtk.ListItem:activatable] for details on how to use
// this signal.
func (x *GridView) ConnectActivate(cb *func(GridView, uint)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, PositionVarp uint) {
		fa := GridView{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, PositionVarp)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "activate", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *GridView) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *GridView) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *GridView) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *GridView) ResetState(StateVar AccessibleState) {

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
func (x *GridView) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *GridView) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *GridView) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *GridView) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *GridView) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *GridView) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *GridView) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *GridView) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *GridView) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

// Returns the size of a non-scrolling border around the
// outside of the scrollable.
//
// An example for this would be treeview headers. GTK can use
// this information to display overlaid graphics, like the
// overshoot indication, at the right position.
func (x *GridView) GetBorder(BorderVar *Border) bool {

	cret := XGtkScrollableGetBorder(x.GoPointer(), BorderVar)
	return cret
}

// Retrieves the `GtkAdjustment` used for horizontal scrolling.
func (x *GridView) GetHadjustment() *Adjustment {
	var cls *Adjustment

	cret := XGtkScrollableGetHadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

// Gets the horizontal `GtkScrollablePolicy`.
func (x *GridView) GetHscrollPolicy() ScrollablePolicy {

	cret := XGtkScrollableGetHscrollPolicy(x.GoPointer())
	return cret
}

// Retrieves the `GtkAdjustment` used for vertical scrolling.
func (x *GridView) GetVadjustment() *Adjustment {
	var cls *Adjustment

	cret := XGtkScrollableGetVadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

// Gets the vertical `GtkScrollablePolicy`.
func (x *GridView) GetVscrollPolicy() ScrollablePolicy {

	cret := XGtkScrollableGetVscrollPolicy(x.GoPointer())
	return cret
}

// Sets the horizontal adjustment of the `GtkScrollable`.
func (x *GridView) SetHadjustment(HadjustmentVar *Adjustment) {

	XGtkScrollableSetHadjustment(x.GoPointer(), HadjustmentVar.GoPointer())

}

// Sets the `GtkScrollablePolicy`.
//
// The policy determines whether horizontal scrolling should start
// below the minimum width or below the natural width.
func (x *GridView) SetHscrollPolicy(PolicyVar ScrollablePolicy) {

	XGtkScrollableSetHscrollPolicy(x.GoPointer(), PolicyVar)

}

// Sets the vertical adjustment of the `GtkScrollable`.
func (x *GridView) SetVadjustment(VadjustmentVar *Adjustment) {

	XGtkScrollableSetVadjustment(x.GoPointer(), VadjustmentVar.GoPointer())

}

// Sets the `GtkScrollablePolicy`.
//
// The policy determines whether vertical scrolling should start
// below the minimum height or below the natural height.
func (x *GridView) SetVscrollPolicy(PolicyVar ScrollablePolicy) {

	XGtkScrollableSetVscrollPolicy(x.GoPointer(), PolicyVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xGridViewGLibType, lib, "gtk_grid_view_get_type")

	core.PuregoSafeRegister(&xNewGridView, lib, "gtk_grid_view_new")

	core.PuregoSafeRegister(&xGridViewGetEnableRubberband, lib, "gtk_grid_view_get_enable_rubberband")
	core.PuregoSafeRegister(&xGridViewGetFactory, lib, "gtk_grid_view_get_factory")
	core.PuregoSafeRegister(&xGridViewGetMaxColumns, lib, "gtk_grid_view_get_max_columns")
	core.PuregoSafeRegister(&xGridViewGetMinColumns, lib, "gtk_grid_view_get_min_columns")
	core.PuregoSafeRegister(&xGridViewGetModel, lib, "gtk_grid_view_get_model")
	core.PuregoSafeRegister(&xGridViewGetSingleClickActivate, lib, "gtk_grid_view_get_single_click_activate")
	core.PuregoSafeRegister(&xGridViewSetEnableRubberband, lib, "gtk_grid_view_set_enable_rubberband")
	core.PuregoSafeRegister(&xGridViewSetFactory, lib, "gtk_grid_view_set_factory")
	core.PuregoSafeRegister(&xGridViewSetMaxColumns, lib, "gtk_grid_view_set_max_columns")
	core.PuregoSafeRegister(&xGridViewSetMinColumns, lib, "gtk_grid_view_set_min_columns")
	core.PuregoSafeRegister(&xGridViewSetModel, lib, "gtk_grid_view_set_model")
	core.PuregoSafeRegister(&xGridViewSetSingleClickActivate, lib, "gtk_grid_view_set_single_click_activate")

}
