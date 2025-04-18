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

type StringSorterClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *StringSorterClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GtkStringSorter` is a `GtkSorter` that compares strings.
//
// It does the comparison in a linguistically correct way using the
// current locale by normalizing Unicode strings and possibly case-folding
// them before performing the comparison.
//
// To obtain the strings to compare, this sorter evaluates a
// [class@Gtk.Expression].
type StringSorter struct {
	Sorter
}

var xStringSorterGLibType func() types.GType

func StringSorterGLibType() types.GType {
	return xStringSorterGLibType()
}

func StringSorterNewFromInternalPtr(ptr uintptr) *StringSorter {
	cls := &StringSorter{}
	cls.Ptr = ptr
	return cls
}

var xNewStringSorter func(uintptr) uintptr

// Creates a new string sorter that compares items using the given
// @expression.
//
// Unless an expression is set on it, this sorter will always
// compare items as invalid.
func NewStringSorter(ExpressionVar *Expression) *StringSorter {
	var cls *StringSorter

	cret := xNewStringSorter(ExpressionVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &StringSorter{}
	cls.Ptr = cret
	return cls
}

var xStringSorterGetExpression func(uintptr) uintptr

// Gets the expression that is evaluated to obtain strings from items.
func (x *StringSorter) GetExpression() *Expression {
	var cls *Expression

	cret := xStringSorterGetExpression(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Expression{}
	cls.Ptr = cret
	return cls
}

var xStringSorterGetIgnoreCase func(uintptr) bool

// Gets whether the sorter ignores case differences.
func (x *StringSorter) GetIgnoreCase() bool {

	cret := xStringSorterGetIgnoreCase(x.GoPointer())
	return cret
}

var xStringSorterSetExpression func(uintptr, uintptr)

// Sets the expression that is evaluated to obtain strings from items.
//
// The expression must have the type %G_TYPE_STRING.
func (x *StringSorter) SetExpression(ExpressionVar *Expression) {

	xStringSorterSetExpression(x.GoPointer(), ExpressionVar.GoPointer())

}

var xStringSorterSetIgnoreCase func(uintptr, bool)

// Sets whether the sorter will ignore case differences.
func (x *StringSorter) SetIgnoreCase(IgnoreCaseVar bool) {

	xStringSorterSetIgnoreCase(x.GoPointer(), IgnoreCaseVar)

}

func (c *StringSorter) GoPointer() uintptr {
	return c.Ptr
}

func (c *StringSorter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xStringSorterGLibType, lib, "gtk_string_sorter_get_type")

	core.PuregoSafeRegister(&xNewStringSorter, lib, "gtk_string_sorter_new")

	core.PuregoSafeRegister(&xStringSorterGetExpression, lib, "gtk_string_sorter_get_expression")
	core.PuregoSafeRegister(&xStringSorterGetIgnoreCase, lib, "gtk_string_sorter_get_ignore_case")
	core.PuregoSafeRegister(&xStringSorterSetExpression, lib, "gtk_string_sorter_set_expression")
	core.PuregoSafeRegister(&xStringSorterSetIgnoreCase, lib, "gtk_string_sorter_set_ignore_case")

}
