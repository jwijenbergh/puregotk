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

// A function used by gtk_tree_selection_selected_foreach() to map all
// selected rows.  It will be called on every selected row in the view.
type TreeSelectionForeachFunc func(uintptr, *TreePath, *TreeIter, uintptr)

// A function used by gtk_tree_selection_set_select_function() to filter
// whether or not a row may be selected. It is called whenever a row's
// state might change.
//
// A return value of %TRUE indicates to @selection that it is okay to
// change the selection.
type TreeSelectionFunc func(uintptr, uintptr, *TreePath, bool, uintptr) bool

// The selection object for GtkTreeView
//
// The `GtkTreeSelection` object is a helper object to manage the selection
// for a `GtkTreeView` widget.  The `GtkTreeSelection` object is
// automatically created when a new `GtkTreeView` widget is created, and
// cannot exist independently of this widget.  The primary reason the
// `GtkTreeSelection` objects exists is for cleanliness of code and API.
// That is, there is no conceptual reason all these functions could not be
// methods on the `GtkTreeView` widget instead of a separate function.
//
// The `GtkTreeSelection` object is gotten from a `GtkTreeView` by calling
// gtk_tree_view_get_selection().  It can be manipulated to check the
// selection status of the tree, as well as select and deselect individual
// rows.  Selection is done completely view side.  As a result, multiple
// views of the same model can have completely different selections.
// Additionally, you cannot change the selection of a row on the model that
// is not currently displayed by the view without expanding its parents
// first.
//
// One of the important things to remember when monitoring the selection of
// a view is that the `GtkTreeSelection`::changed signal is mostly a hint.
// That is, it may only emit one signal when a range of rows is selected.
// Additionally, it may on occasion emit a `GtkTreeSelection`::changed signal
// when nothing has happened (mostly as a result of programmers calling
// select_row on an already selected row).
type TreeSelection struct {
	gobject.Object
}

var xTreeSelectionGLibType func() types.GType

func TreeSelectionGLibType() types.GType {
	return xTreeSelectionGLibType()
}

func TreeSelectionNewFromInternalPtr(ptr uintptr) *TreeSelection {
	cls := &TreeSelection{}
	cls.Ptr = ptr
	return cls
}

var xTreeSelectionCountSelectedRows func(uintptr) int

// Returns the number of rows that have been selected in @tree.
func (x *TreeSelection) CountSelectedRows() int {

	cret := xTreeSelectionCountSelectedRows(x.GoPointer())
	return cret
}

var xTreeSelectionGetMode func(uintptr) SelectionMode

// Gets the selection mode for @selection. See
// gtk_tree_selection_set_mode().
func (x *TreeSelection) GetMode() SelectionMode {

	cret := xTreeSelectionGetMode(x.GoPointer())
	return cret
}

var xTreeSelectionGetSelectFunction func(uintptr) uintptr

// Returns the current selection function.
func (x *TreeSelection) GetSelectFunction() uintptr {

	cret := xTreeSelectionGetSelectFunction(x.GoPointer())
	return cret
}

var xTreeSelectionGetSelected func(uintptr, *uintptr, *TreeIter) bool

// Sets @iter to the currently selected node if @selection is set to
// %GTK_SELECTION_SINGLE or %GTK_SELECTION_BROWSE.  @iter may be NULL if you
// just want to test if @selection has any selected nodes.  @model is filled
// with the current model as a convenience.  This function will not work if you
// use @selection is %GTK_SELECTION_MULTIPLE.
func (x *TreeSelection) GetSelected(ModelVar *TreeModel, IterVar *TreeIter) bool {

	cret := xTreeSelectionGetSelected(x.GoPointer(), gobject.ConvertPtr(ModelVar), IterVar)
	return cret
}

var xTreeSelectionGetSelectedRows func(uintptr, *uintptr) *glib.List

// Creates a list of path of all selected rows. Additionally, if you are
// planning on modifying the model after calling this function, you may
// want to convert the returned list into a list of `GtkTreeRowReference`s.
// To do this, you can use gtk_tree_row_reference_new().
//
// To free the return value, use:
// |[&lt;!-- language="C" --&gt;
// g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);
// ]|
func (x *TreeSelection) GetSelectedRows(ModelVar *TreeModel) *glib.List {

	cret := xTreeSelectionGetSelectedRows(x.GoPointer(), gobject.ConvertPtr(ModelVar))
	return cret
}

var xTreeSelectionGetTreeView func(uintptr) uintptr

// Returns the tree view associated with @selection.
func (x *TreeSelection) GetTreeView() *TreeView {
	var cls *TreeView

	cret := xTreeSelectionGetTreeView(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TreeView{}
	cls.Ptr = cret
	return cls
}

var xTreeSelectionGetUserData func(uintptr) uintptr

// Returns the user data for the selection function.
func (x *TreeSelection) GetUserData() uintptr {

	cret := xTreeSelectionGetUserData(x.GoPointer())
	return cret
}

var xTreeSelectionIterIsSelected func(uintptr, *TreeIter) bool

// Returns %TRUE if the row at @iter is currently selected.
func (x *TreeSelection) IterIsSelected(IterVar *TreeIter) bool {

	cret := xTreeSelectionIterIsSelected(x.GoPointer(), IterVar)
	return cret
}

var xTreeSelectionPathIsSelected func(uintptr, *TreePath) bool

// Returns %TRUE if the row pointed to by @path is currently selected.  If @path
// does not point to a valid location, %FALSE is returned
func (x *TreeSelection) PathIsSelected(PathVar *TreePath) bool {

	cret := xTreeSelectionPathIsSelected(x.GoPointer(), PathVar)
	return cret
}

var xTreeSelectionSelectAll func(uintptr)

// Selects all the nodes. @selection must be set to %GTK_SELECTION_MULTIPLE
// mode.
func (x *TreeSelection) SelectAll() {

	xTreeSelectionSelectAll(x.GoPointer())

}

var xTreeSelectionSelectIter func(uintptr, *TreeIter)

// Selects the specified iterator.
func (x *TreeSelection) SelectIter(IterVar *TreeIter) {

	xTreeSelectionSelectIter(x.GoPointer(), IterVar)

}

var xTreeSelectionSelectPath func(uintptr, *TreePath)

// Select the row at @path.
func (x *TreeSelection) SelectPath(PathVar *TreePath) {

	xTreeSelectionSelectPath(x.GoPointer(), PathVar)

}

var xTreeSelectionSelectRange func(uintptr, *TreePath, *TreePath)

// Selects a range of nodes, determined by @start_path and @end_path inclusive.
// @selection must be set to %GTK_SELECTION_MULTIPLE mode.
func (x *TreeSelection) SelectRange(StartPathVar *TreePath, EndPathVar *TreePath) {

	xTreeSelectionSelectRange(x.GoPointer(), StartPathVar, EndPathVar)

}

var xTreeSelectionSelectedForeach func(uintptr, uintptr, uintptr)

// Calls a function for each selected node. Note that you cannot modify
// the tree or selection from within this function. As a result,
// gtk_tree_selection_get_selected_rows() might be more useful.
func (x *TreeSelection) SelectedForeach(FuncVar *TreeSelectionForeachFunc, DataVar uintptr) {

	xTreeSelectionSelectedForeach(x.GoPointer(), glib.NewCallback(FuncVar), DataVar)

}

var xTreeSelectionSetMode func(uintptr, SelectionMode)

// Sets the selection mode of the @selection.  If the previous type was
// %GTK_SELECTION_MULTIPLE, then the anchor is kept selected, if it was
// previously selected.
func (x *TreeSelection) SetMode(TypeVar SelectionMode) {

	xTreeSelectionSetMode(x.GoPointer(), TypeVar)

}

var xTreeSelectionSetSelectFunction func(uintptr, uintptr, uintptr, uintptr)

// Sets the selection function.
//
// If set, this function is called before any node is selected or unselected,
// giving some control over which nodes are selected. The select function
// should return %TRUE if the state of the node may be toggled, and %FALSE
// if the state of the node should be left unchanged.
func (x *TreeSelection) SetSelectFunction(FuncVar *TreeSelectionFunc, DataVar uintptr, DestroyVar *glib.DestroyNotify) {

	xTreeSelectionSetSelectFunction(x.GoPointer(), glib.NewCallback(FuncVar), DataVar, glib.NewCallback(DestroyVar))

}

var xTreeSelectionUnselectAll func(uintptr)

// Unselects all the nodes.
func (x *TreeSelection) UnselectAll() {

	xTreeSelectionUnselectAll(x.GoPointer())

}

var xTreeSelectionUnselectIter func(uintptr, *TreeIter)

// Unselects the specified iterator.
func (x *TreeSelection) UnselectIter(IterVar *TreeIter) {

	xTreeSelectionUnselectIter(x.GoPointer(), IterVar)

}

var xTreeSelectionUnselectPath func(uintptr, *TreePath)

// Unselects the row at @path.
func (x *TreeSelection) UnselectPath(PathVar *TreePath) {

	xTreeSelectionUnselectPath(x.GoPointer(), PathVar)

}

var xTreeSelectionUnselectRange func(uintptr, *TreePath, *TreePath)

// Unselects a range of nodes, determined by @start_path and @end_path
// inclusive.
func (x *TreeSelection) UnselectRange(StartPathVar *TreePath, EndPathVar *TreePath) {

	xTreeSelectionUnselectRange(x.GoPointer(), StartPathVar, EndPathVar)

}

func (c *TreeSelection) GoPointer() uintptr {
	return c.Ptr
}

func (c *TreeSelection) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whenever the selection has (possibly) changed. Please note that
// this signal is mostly a hint.  It may only be emitted once when a range
// of rows are selected, and it may occasionally be emitted when nothing
// has happened.
func (x *TreeSelection) ConnectChanged(cb *func(TreeSelection)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := TreeSelection{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "changed", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xTreeSelectionGLibType, lib, "gtk_tree_selection_get_type")

	core.PuregoSafeRegister(&xTreeSelectionCountSelectedRows, lib, "gtk_tree_selection_count_selected_rows")
	core.PuregoSafeRegister(&xTreeSelectionGetMode, lib, "gtk_tree_selection_get_mode")
	core.PuregoSafeRegister(&xTreeSelectionGetSelectFunction, lib, "gtk_tree_selection_get_select_function")
	core.PuregoSafeRegister(&xTreeSelectionGetSelected, lib, "gtk_tree_selection_get_selected")
	core.PuregoSafeRegister(&xTreeSelectionGetSelectedRows, lib, "gtk_tree_selection_get_selected_rows")
	core.PuregoSafeRegister(&xTreeSelectionGetTreeView, lib, "gtk_tree_selection_get_tree_view")
	core.PuregoSafeRegister(&xTreeSelectionGetUserData, lib, "gtk_tree_selection_get_user_data")
	core.PuregoSafeRegister(&xTreeSelectionIterIsSelected, lib, "gtk_tree_selection_iter_is_selected")
	core.PuregoSafeRegister(&xTreeSelectionPathIsSelected, lib, "gtk_tree_selection_path_is_selected")
	core.PuregoSafeRegister(&xTreeSelectionSelectAll, lib, "gtk_tree_selection_select_all")
	core.PuregoSafeRegister(&xTreeSelectionSelectIter, lib, "gtk_tree_selection_select_iter")
	core.PuregoSafeRegister(&xTreeSelectionSelectPath, lib, "gtk_tree_selection_select_path")
	core.PuregoSafeRegister(&xTreeSelectionSelectRange, lib, "gtk_tree_selection_select_range")
	core.PuregoSafeRegister(&xTreeSelectionSelectedForeach, lib, "gtk_tree_selection_selected_foreach")
	core.PuregoSafeRegister(&xTreeSelectionSetMode, lib, "gtk_tree_selection_set_mode")
	core.PuregoSafeRegister(&xTreeSelectionSetSelectFunction, lib, "gtk_tree_selection_set_select_function")
	core.PuregoSafeRegister(&xTreeSelectionUnselectAll, lib, "gtk_tree_selection_unselect_all")
	core.PuregoSafeRegister(&xTreeSelectionUnselectIter, lib, "gtk_tree_selection_unselect_iter")
	core.PuregoSafeRegister(&xTreeSelectionUnselectPath, lib, "gtk_tree_selection_unselect_path")
	core.PuregoSafeRegister(&xTreeSelectionUnselectRange, lib, "gtk_tree_selection_unselect_range")

}
