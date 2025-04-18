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

// The virtual table for `GtkSorter`.
type SorterClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *SorterClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Describes changes in a sorter in more detail and allows users
// to optimize resorting.
type SorterChange int

var xSorterChangeGLibType func() types.GType

func SorterChangeGLibType() types.GType {
	return xSorterChangeGLibType()
}

const (

	// The sorter change cannot be described
	//   by any of the other enumeration values
	SorterChangeDifferentValue SorterChange = 0
	// The sort order was inverted. Comparisons
	//   that returned %GTK_ORDERING_SMALLER now return %GTK_ORDERING_LARGER
	//   and vice versa. Other comparisons return the same values as before.
	SorterChangeInvertedValue SorterChange = 1
	// The sorter is less strict: Comparisons
	//   may now return %GTK_ORDERING_EQUAL that did not do so before.
	SorterChangeLessStrictValue SorterChange = 2
	// The sorter is more strict: Comparisons
	//   that did return %GTK_ORDERING_EQUAL may not do so anymore.
	SorterChangeMoreStrictValue SorterChange = 3
)

// Describes the type of order that a `GtkSorter` may produce.
type SorterOrder int

var xSorterOrderGLibType func() types.GType

func SorterOrderGLibType() types.GType {
	return xSorterOrderGLibType()
}

const (

	// A partial order. Any `GtkOrdering` is possible.
	SorterOrderPartialValue SorterOrder = 0
	// No order, all elements are considered equal.
	//   gtk_sorter_compare() will only return %GTK_ORDERING_EQUAL.
	SorterOrderNoneValue SorterOrder = 1
	// A total order. gtk_sorter_compare() will only
	//   return %GTK_ORDERING_EQUAL if an item is compared with itself. Two
	//   different items will never cause this value to be returned.
	SorterOrderTotalValue SorterOrder = 2
)

// `GtkSorter` is an object to describe sorting criteria.
//
// Its primary user is [class@Gtk.SortListModel]
//
// The model will use a sorter to determine the order in which
// its items should appear by calling [method@Gtk.Sorter.compare]
// for pairs of items.
//
// Sorters may change their sorting behavior through their lifetime.
// In that case, they will emit the [signal@Gtk.Sorter::changed] signal
// to notify that the sort order is no longer valid and should be updated
// by calling gtk_sorter_compare() again.
//
// GTK provides various pre-made sorter implementations for common sorting
// operations. [class@Gtk.ColumnView] has built-in support for sorting lists
// via the [property@Gtk.ColumnViewColumn:sorter] property, where the user can
// change the sorting by clicking on list headers.
//
// Of course, in particular for large lists, it is also possible to subclass
// `GtkSorter` and provide one's own sorter.
type Sorter struct {
	gobject.Object
}

var xSorterGLibType func() types.GType

func SorterGLibType() types.GType {
	return xSorterGLibType()
}

func SorterNewFromInternalPtr(ptr uintptr) *Sorter {
	cls := &Sorter{}
	cls.Ptr = ptr
	return cls
}

var xSorterChanged func(uintptr, SorterChange)

// Notifies all users of the sorter that it has changed.
//
// This emits the [signal@Gtk.Sorter::changed] signal. Users
// of the sorter should then update the sort order via
// [method@Gtk.Sorter.compare].
//
// Depending on the @change parameter, it may be possible to
// update the sort order without a full resorting. Refer to
// the [enum@Gtk.SorterChange] documentation for details.
//
// This function is intended for implementors of `GtkSorter`
// subclasses and should not be called from other functions.
func (x *Sorter) Changed(ChangeVar SorterChange) {

	xSorterChanged(x.GoPointer(), ChangeVar)

}

var xSorterCompare func(uintptr, uintptr, uintptr) Ordering

// Compares two given items according to the sort order implemented
// by the sorter.
//
// Sorters implement a partial order:
//
//   - It is reflexive, ie a = a
//   - It is antisymmetric, ie if a &lt; b and b &lt; a, then a = b
//   - It is transitive, ie given any 3 items with a ≤ b and b ≤ c,
//     then a ≤ c
//
// The sorter may signal it conforms to additional constraints
// via the return value of [method@Gtk.Sorter.get_order].
func (x *Sorter) Compare(Item1Var *gobject.Object, Item2Var *gobject.Object) Ordering {

	cret := xSorterCompare(x.GoPointer(), Item1Var.GoPointer(), Item2Var.GoPointer())
	return cret
}

var xSorterGetOrder func(uintptr) SorterOrder

// Gets the order that @self conforms to.
//
// See [enum@Gtk.SorterOrder] for details
// of the possible return values.
//
// This function is intended to allow optimizations.
func (x *Sorter) GetOrder() SorterOrder {

	cret := xSorterGetOrder(x.GoPointer())
	return cret
}

func (c *Sorter) GoPointer() uintptr {
	return c.Ptr
}

func (c *Sorter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whenever the sorter changed.
//
// Users of the sorter should then update the sort order
// again via gtk_sorter_compare().
//
// [class@Gtk.SortListModel] handles this signal automatically.
//
// Depending on the @change parameter, it may be possible to update
// the sort order without a full resorting. Refer to the
// [enum@Gtk.SorterChange] documentation for details.
func (x *Sorter) ConnectChanged(cb *func(Sorter, SorterChange)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr, ChangeVarp SorterChange) {
		fa := Sorter{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa, ChangeVarp)

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

	core.PuregoSafeRegister(&xSorterChangeGLibType, lib, "gtk_sorter_change_get_type")

	core.PuregoSafeRegister(&xSorterOrderGLibType, lib, "gtk_sorter_order_get_type")

	core.PuregoSafeRegister(&xSorterGLibType, lib, "gtk_sorter_get_type")

	core.PuregoSafeRegister(&xSorterChanged, lib, "gtk_sorter_changed")
	core.PuregoSafeRegister(&xSorterCompare, lib, "gtk_sorter_compare")
	core.PuregoSafeRegister(&xSorterGetOrder, lib, "gtk_sorter_get_order")

}
