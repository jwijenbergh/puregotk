// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

var xSrvTargetListSort func(*glib.List) *glib.List

// Sorts @targets in place according to the algorithm in RFC 2782.
func SrvTargetListSort(TargetsVar *glib.List) *glib.List {

	cret := xSrvTargetListSort(TargetsVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSrvTargetListSort, lib, "g_srv_target_list_sort")

}
