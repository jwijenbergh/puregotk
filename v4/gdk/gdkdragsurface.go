// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The `GdkDragSurfaceInterface` implementation is private to GDK.
type DragSurfaceInterface struct {
	_ structs.HostLayout
}

func (x *DragSurfaceInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `GdkDragSurface` is an interface for surfaces used during DND.
type DragSurface interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	Present(WidthVar int, HeightVar int) bool
}

var xDragSurfaceGLibType func() types.GType

func DragSurfaceGLibType() types.GType {
	return xDragSurfaceGLibType()
}

type DragSurfaceBase struct {
	Ptr uintptr
}

func (x *DragSurfaceBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *DragSurfaceBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Present @drag_surface.
func (x *DragSurfaceBase) Present(WidthVar int, HeightVar int) bool {

	cret := XGdkDragSurfacePresent(x.GoPointer(), WidthVar, HeightVar)
	return cret
}

var XGdkDragSurfacePresent func(uintptr, int, int) bool

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xDragSurfaceGLibType, lib, "gdk_drag_surface_get_type")

	core.PuregoSafeRegister(&XGdkDragSurfacePresent, lib, "gdk_drag_surface_present")

}
