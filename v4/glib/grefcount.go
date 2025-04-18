// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xAtomicRefCountCompare func(int, int) bool

// Atomically compares the current value of @arc with @val.
func AtomicRefCountCompare(ArcVar int, ValVar int) bool {

	cret := xAtomicRefCountCompare(ArcVar, ValVar)
	return cret
}

var xAtomicRefCountDec func(int) bool

// Atomically decreases the reference count.
//
// If %TRUE is returned, the reference count reached 0. After this point, @arc
// is an undefined state and must be reinitialized with
// g_atomic_ref_count_init() to be used again.
func AtomicRefCountDec(ArcVar int) bool {

	cret := xAtomicRefCountDec(ArcVar)
	return cret
}

var xAtomicRefCountInc func(int)

// Atomically increases the reference count.
func AtomicRefCountInc(ArcVar int) {

	xAtomicRefCountInc(ArcVar)

}

var xAtomicRefCountInit func(int)

// Initializes a reference count variable to 1.
func AtomicRefCountInit(ArcVar int) {

	xAtomicRefCountInit(ArcVar)

}

var xRefCountCompare func(int, int) bool

// Compares the current value of @rc with @val.
func RefCountCompare(RcVar int, ValVar int) bool {

	cret := xRefCountCompare(RcVar, ValVar)
	return cret
}

var xRefCountDec func(int) bool

// Decreases the reference count.
//
// If %TRUE is returned, the reference count reached 0. After this point, @rc
// is an undefined state and must be reinitialized with
// g_ref_count_init() to be used again.
func RefCountDec(RcVar int) bool {

	cret := xRefCountDec(RcVar)
	return cret
}

var xRefCountInc func(int)

// Increases the reference count.
func RefCountInc(RcVar int) {

	xRefCountInc(RcVar)

}

var xRefCountInit func(int)

// Initializes a reference count variable to 1.
func RefCountInit(RcVar int) {

	xRefCountInit(RcVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAtomicRefCountCompare, lib, "g_atomic_ref_count_compare")
	core.PuregoSafeRegister(&xAtomicRefCountDec, lib, "g_atomic_ref_count_dec")
	core.PuregoSafeRegister(&xAtomicRefCountInc, lib, "g_atomic_ref_count_inc")
	core.PuregoSafeRegister(&xAtomicRefCountInit, lib, "g_atomic_ref_count_init")
	core.PuregoSafeRegister(&xRefCountCompare, lib, "g_ref_count_compare")
	core.PuregoSafeRegister(&xRefCountDec, lib, "g_ref_count_dec")
	core.PuregoSafeRegister(&xRefCountInc, lib, "g_ref_count_inc")
	core.PuregoSafeRegister(&xRefCountInit, lib, "g_ref_count_init")

}
