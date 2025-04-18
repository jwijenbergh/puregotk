// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

type TitlebarGesture int

var xTitlebarGestureGLibType func() types.GType

func TitlebarGestureGLibType() types.GType {
	return xTitlebarGestureGLibType()
}

const (
	TitlebarGestureDoubleClickValue TitlebarGesture = 1

	TitlebarGestureRightClickValue TitlebarGesture = 2

	TitlebarGestureMiddleClickValue TitlebarGesture = 3
)

var xGlErrorQuark func() glib.Quark

func GlErrorQuark() glib.Quark {

	cret := xGlErrorQuark()
	return cret
}

var xTextureErrorQuark func() glib.Quark

func TextureErrorQuark() glib.Quark {

	cret := xTextureErrorQuark()
	return cret
}

var xVulkanErrorQuark func() glib.Quark

func VulkanErrorQuark() glib.Quark {

	cret := xVulkanErrorQuark()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xTitlebarGestureGLibType, lib, "gdk_titlebar_gesture_get_type")

	core.PuregoSafeRegister(&xGlErrorQuark, lib, "gdk_gl_error_quark")
	core.PuregoSafeRegister(&xTextureErrorQuark, lib, "gdk_texture_error_quark")
	core.PuregoSafeRegister(&xVulkanErrorQuark, lib, "gdk_vulkan_error_quark")

}
