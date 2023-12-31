// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type Type = uint

var xStrvGetType func() []interface{}

func StrvGetType() []interface{} {

	cret := xStrvGetType()
	return cret
}

var xVariantGetGtype func() []interface{}

func VariantGetGtype() []interface{} {

	cret := xVariantGetGtype()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xStrvGetType, lib, "g_strv_get_type")
	core.PuregoSafeRegister(&xVariantGetGtype, lib, "g_variant_get_gtype")

}
