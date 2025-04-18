// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

type MultiSorterClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *MultiSorterClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkMultiSorter` combines multiple sorters by trying them
// in turn.
//
// If the first sorter compares two items as equal,
// the second is tried next, and so on.
type MultiSorter struct {
	Sorter
}

var xMultiSorterGLibType func() types.GType

func MultiSorterGLibType() types.GType {
	return xMultiSorterGLibType()
}

func MultiSorterNewFromInternalPtr(ptr uintptr) *MultiSorter {
	cls := &MultiSorter{}
	cls.Ptr = ptr
	return cls
}

var xNewMultiSorter func() uintptr

// Creates a new multi sorter.
//
// This sorter compares items by trying each of the sorters
// in turn, until one returns non-zero. In particular, if
// no sorter has been added to it, it will always compare
// items as equal.
func NewMultiSorter() *MultiSorter {
	var cls *MultiSorter

	cret := xNewMultiSorter()

	if cret == 0 {
		return nil
	}
	cls = &MultiSorter{}
	cls.Ptr = cret
	return cls
}

var xMultiSorterAppend func(uintptr, uintptr)

// Add @sorter to @self to use for sorting at the end.
//
// @self will consult all existing sorters before it will
// sort with the given @sorter.
func (x *MultiSorter) Append(SorterVar *Sorter) {

	xMultiSorterAppend(x.GoPointer(), SorterVar.GoPointer())

}

var xMultiSorterRemove func(uintptr, uint)

// Removes the sorter at the given @position from the list of sorter
// used by @self.
//
// If @position is larger than the number of sorters, nothing happens.
func (x *MultiSorter) Remove(PositionVar uint) {

	xMultiSorterRemove(x.GoPointer(), PositionVar)

}

func (c *MultiSorter) GoPointer() uintptr {
	return c.Ptr
}

func (c *MultiSorter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// See also: g_list_model_get_n_items()
func (x *MultiSorter) GetItem(PositionVar uint) uintptr {

	cret := gio.XGListModelGetItem(x.GoPointer(), PositionVar)
	return cret
}

// Gets the type of the items in @list.
//
// All items returned from g_list_model_get_item() are of the type
// returned by this function, or a subtype, or if the type is an
// interface, they are an implementation of that interface.
//
// The item type of a #GListModel can not change during the life of the
// model.
func (x *MultiSorter) GetItemType() types.GType {

	cret := gio.XGListModelGetItemType(x.GoPointer())
	return cret
}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
func (x *MultiSorter) GetNItems() uint {

	cret := gio.XGListModelGetNItems(x.GoPointer())
	return cret
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// This function is meant to be used by language bindings in place
// of g_list_model_get_item().
//
// See also: g_list_model_get_n_items()
func (x *MultiSorter) GetObject(PositionVar uint) *gobject.Object {
	var cls *gobject.Object

	cret := gio.XGListModelGetObject(x.GoPointer(), PositionVar)

	if cret == 0 {
		return nil
	}
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

// Emits the #GListModel::items-changed signal on @list.
//
// This function should only be called by classes implementing
// #GListModel. It has to be called after the internal representation
// of @list has been updated, because handlers connected to this signal
// might query the new state of the list.
//
// Implementations must only make changes to the model (as visible to
// its consumer) in places that will not cause problems for that
// consumer.  For models that are driven directly by a write API (such
// as #GListStore), changes can be reported in response to uses of that
// API.  For models that represent remote data, changes should only be
// made from a fresh mainloop dispatch.  It is particularly not
// permitted to make changes in response to a call to the #GListModel
// consumer API.
//
// Stated another way: in general, it is assumed that code making a
// series of accesses to the model via the API, without returning to the
// mainloop, and without calling other code, will continue to view the
// same contents of the model.
func (x *MultiSorter) ItemsChanged(PositionVar uint, RemovedVar uint, AddedVar uint) {

	gio.XGListModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *MultiSorter) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xMultiSorterGLibType, lib, "gtk_multi_sorter_get_type")

	core.PuregoSafeRegister(&xNewMultiSorter, lib, "gtk_multi_sorter_new")

	core.PuregoSafeRegister(&xMultiSorterAppend, lib, "gtk_multi_sorter_append")
	core.PuregoSafeRegister(&xMultiSorterRemove, lib, "gtk_multi_sorter_remove")

}
