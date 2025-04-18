// Package gobject was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gobject

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

var xSourceSetClosure func(*glib.Source, *Closure)

// Set the callback for a source as a #GClosure.
//
// If the source is not one of the standard GLib types, the @closure_callback
// and @closure_marshal fields of the #GSourceFuncs structure must have been
// filled in with pointers to appropriate functions.
func SourceSetClosure(SourceVar *glib.Source, ClosureVar *Closure) {

	xSourceSetClosure(SourceVar, ClosureVar)

}

var xSourceSetDummyCallback func(*glib.Source)

// Sets a dummy callback for @source. The callback will do nothing, and
// if the source expects a #gboolean return value, it will return %TRUE.
// (If the source expects any other type of return value, it will return
// a 0/%NULL value; whatever g_value_init() initializes a #GValue to for
// that type.)
//
// If the source is not one of the standard GLib types, the
// @closure_callback and @closure_marshal fields of the #GSourceFuncs
// structure must have been filled in with pointers to appropriate
// functions.
func SourceSetDummyCallback(SourceVar *glib.Source) {

	xSourceSetDummyCallback(SourceVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GOBJECT"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSourceSetClosure, lib, "g_source_set_closure")
	core.PuregoSafeRegister(&xSourceSetDummyCallback, lib, "g_source_set_dummy_callback")

}
