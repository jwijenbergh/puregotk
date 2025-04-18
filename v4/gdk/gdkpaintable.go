// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The list of functions that can be implemented for the `GdkPaintable`
// interface.
//
// Note that apart from the [vfunc@Gdk.Paintable.snapshot] function,
// no virtual function of this interface is mandatory to implement, though it
// is a good idea to implement [vfunc@Gdk.Paintable.get_current_image]
// for non-static paintables and [vfunc@Gdk.Paintable.get_flags] if the
// image is not dynamic as the default implementation returns no flags and
// that will make the implementation likely quite slow.
type PaintableInterface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *PaintableInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// `GdkPaintable` is a simple interface used by GTK to represent content that
// can be painted.
//
// The content of a `GdkPaintable` can be painted anywhere at any size
// without requiring any sort of layout. The interface is inspired by
// similar concepts elsewhere, such as
// [ClutterContent](https://developer.gnome.org/clutter/stable/ClutterContent.html),
// [HTML/CSS Paint Sources](https://www.w3.org/TR/css-images-4/#paint-source),
// or [SVG Paint Servers](https://www.w3.org/TR/SVG2/pservers.html).
//
// A `GdkPaintable` can be snapshot at any time and size using
// [method@Gdk.Paintable.snapshot]. How the paintable interprets that size and
// if it scales or centers itself into the given rectangle is implementation
// defined, though if you are implementing a `GdkPaintable` and don't know what
// to do, it is suggested that you scale your paintable ignoring any potential
// aspect ratio.
//
// The contents that a `GdkPaintable` produces may depend on the [class@GdkSnapshot]
// passed to it. For example, paintables may decide to use more detailed images
// on higher resolution screens or when OpenGL is available. A `GdkPaintable`
// will however always produce the same output for the same snapshot.
//
// A `GdkPaintable` may change its contents, meaning that it will now produce
// a different output with the same snapshot. Once that happens, it will call
// [method@Gdk.Paintable.invalidate_contents] which will emit the
// [signal@GdkPaintable::invalidate-contents] signal. If a paintable is known
// to never change its contents, it will set the %GDK_PAINTABLE_STATIC_CONTENTS
// flag. If a consumer cannot deal with changing contents, it may call
// [method@Gdk.Paintable.get_current_image] which will return a static
// paintable and use that.
//
// A paintable can report an intrinsic (or preferred) size or aspect ratio it
// wishes to be rendered at, though it doesn't have to. Consumers of the interface
// can use this information to layout thepaintable appropriately. Just like the
// contents, the size of a paintable can change. A paintable will indicate this
// by calling [method@Gdk.Paintable.invalidate_size] which will emit the
// [signal@GdkPaintable::invalidate-size] signal. And just like for contents,
// if a paintable is known to never change its size, it will set the
// %GDK_PAINTABLE_STATIC_SIZE flag.
//
// Besides API for applications, there are some functions that are only
// useful for implementing subclasses and should not be used by applications:
// [method@Gdk.Paintable.invalidate_contents],
// [method@Gdk.Paintable.invalidate_size],
// [func@Gdk.Paintable.new_empty].
type Paintable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	ComputeConcreteSize(SpecifiedWidthVar float64, SpecifiedHeightVar float64, DefaultWidthVar float64, DefaultHeightVar float64, ConcreteWidthVar float64, ConcreteHeightVar float64)
	GetCurrentImage() *PaintableBase
	GetFlags() PaintableFlags
	GetIntrinsicAspectRatio() float64
	GetIntrinsicHeight() int
	GetIntrinsicWidth() int
	InvalidateContents()
	InvalidateSize()
	Snapshot(SnapshotVar *Snapshot, WidthVar float64, HeightVar float64)
}

var xPaintableGLibType func() types.GType

func PaintableGLibType() types.GType {
	return xPaintableGLibType()
}

type PaintableBase struct {
	Ptr uintptr
}

func (x *PaintableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *PaintableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Compute a concrete size for the `GdkPaintable`.
//
// Applies the sizing algorithm outlined in the
// [CSS Image spec](https://drafts.csswg.org/css-images-3/#default-sizing)
// to the given @paintable. See that link for more details.
//
// It is not necessary to call this function when both @specified_width
// and @specified_height are known, but it is useful to call this
// function in GtkWidget:measure implementations to compute the
// other dimension when only one dimension is given.
func (x *PaintableBase) ComputeConcreteSize(SpecifiedWidthVar float64, SpecifiedHeightVar float64, DefaultWidthVar float64, DefaultHeightVar float64, ConcreteWidthVar float64, ConcreteHeightVar float64) {

	XGdkPaintableComputeConcreteSize(x.GoPointer(), SpecifiedWidthVar, SpecifiedHeightVar, DefaultWidthVar, DefaultHeightVar, ConcreteWidthVar, ConcreteHeightVar)

}

// Gets an immutable paintable for the current contents displayed by @paintable.
//
// This is useful when you want to retain the current state of an animation,
// for example to take a screenshot of a running animation.
//
// If the @paintable is already immutable, it will return itself.
func (x *PaintableBase) GetCurrentImage() *PaintableBase {
	var cls *PaintableBase

	cret := XGdkPaintableGetCurrentImage(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &PaintableBase{}
	cls.Ptr = cret
	return cls
}

// Get flags for the paintable.
//
// This is oftentimes useful for optimizations.
//
// See [flags@Gdk.PaintableFlags] for the flags and what they mean.
func (x *PaintableBase) GetFlags() PaintableFlags {

	cret := XGdkPaintableGetFlags(x.GoPointer())
	return cret
}

// Gets the preferred aspect ratio the @paintable would like to be displayed at.
//
// The aspect ratio is the width divided by the height, so a value of 0.5
// means that the @paintable prefers to be displayed twice as high as it
// is wide. Consumers of this interface can use this to preserve aspect
// ratio when displaying the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// Usually when a @paintable returns nonzero values from
// [method@Gdk.Paintable.get_intrinsic_width] and
// [method@Gdk.Paintable.get_intrinsic_height] the aspect ratio
// should conform to those values, though that is not required.
//
// If the @paintable does not have a preferred aspect ratio,
// it returns 0. Negative values are never returned.
func (x *PaintableBase) GetIntrinsicAspectRatio() float64 {

	cret := XGdkPaintableGetIntrinsicAspectRatio(x.GoPointer())
	return cret
}

// Gets the preferred height the @paintable would like to be displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred height, it returns 0.
// Negative values are never returned.
func (x *PaintableBase) GetIntrinsicHeight() int {

	cret := XGdkPaintableGetIntrinsicHeight(x.GoPointer())
	return cret
}

// Gets the preferred width the @paintable would like to be displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred width, it returns 0.
// Negative values are never returned.
func (x *PaintableBase) GetIntrinsicWidth() int {

	cret := XGdkPaintableGetIntrinsicWidth(x.GoPointer())
	return cret
}

// Called by implementations of `GdkPaintable` to invalidate their contents.
//
// Unless the contents are invalidated, implementations must guarantee that
// multiple calls of [method@Gdk.Paintable.snapshot] produce the same output.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-contents]
// signal.
//
// If a @paintable reports the %GDK_PAINTABLE_STATIC_CONTENTS flag,
// it must not call this function.
func (x *PaintableBase) InvalidateContents() {

	XGdkPaintableInvalidateContents(x.GoPointer())

}

// Called by implementations of `GdkPaintable` to invalidate their size.
//
// As long as the size is not invalidated, @paintable must return the same
// values for its intrinsic width, height and aspect ratio.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-size]
// signal.
//
// If a @paintable reports the %GDK_PAINTABLE_STATIC_SIZE flag,
// it must not call this function.
func (x *PaintableBase) InvalidateSize() {

	XGdkPaintableInvalidateSize(x.GoPointer())

}

// Snapshots the given paintable with the given @width and @height.
//
// The paintable is drawn at the current (0,0) offset of the @snapshot.
// If @width and @height are not larger than zero, this function will
// do nothing.
func (x *PaintableBase) Snapshot(SnapshotVar *Snapshot, WidthVar float64, HeightVar float64) {

	XGdkPaintableSnapshot(x.GoPointer(), SnapshotVar.GoPointer(), WidthVar, HeightVar)

}

var XGdkPaintableComputeConcreteSize func(uintptr, float64, float64, float64, float64, float64, float64)
var XGdkPaintableGetCurrentImage func(uintptr) uintptr
var XGdkPaintableGetFlags func(uintptr) PaintableFlags
var XGdkPaintableGetIntrinsicAspectRatio func(uintptr) float64
var XGdkPaintableGetIntrinsicHeight func(uintptr) int
var XGdkPaintableGetIntrinsicWidth func(uintptr) int
var XGdkPaintableInvalidateContents func(uintptr)
var XGdkPaintableInvalidateSize func(uintptr)
var XGdkPaintableSnapshot func(uintptr, uintptr, float64, float64)

// Flags about a paintable object.
//
// Implementations use these for optimizations such as caching.
type PaintableFlags int

var xPaintableFlagsGLibType func() types.GType

func PaintableFlagsGLibType() types.GType {
	return xPaintableFlagsGLibType()
}

const (

	// The size is immutable.
	//   The [signal@GdkPaintable::invalidate-size] signal will never be
	//   emitted.
	PaintableStaticSizeValue PaintableFlags = 1
	// The content is immutable.
	//   The [signal@GdkPaintable::invalidate-contents] signal will never be
	//   emitted.
	PaintableStaticContentsValue PaintableFlags = 2
)

var xPaintableNewEmpty func(int, int) uintptr

// Returns a paintable that has the given intrinsic size and draws nothing.
//
// This is often useful for implementing the
// [vfunc@Gdk.Paintable.get_current_image] virtual function
// when the paintable is in an incomplete state (like a
// [class@Gtk.MediaStream] before receiving the first frame).
func PaintableNewEmpty(IntrinsicWidthVar int, IntrinsicHeightVar int) *PaintableBase {
	var cls *PaintableBase

	cret := xPaintableNewEmpty(IntrinsicWidthVar, IntrinsicHeightVar)

	if cret == 0 {
		return nil
	}
	cls = &PaintableBase{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xPaintableFlagsGLibType, lib, "gdk_paintable_flags_get_type")

	core.PuregoSafeRegister(&xPaintableNewEmpty, lib, "gdk_paintable_new_empty")

	core.PuregoSafeRegister(&xPaintableGLibType, lib, "gdk_paintable_get_type")

	core.PuregoSafeRegister(&XGdkPaintableComputeConcreteSize, lib, "gdk_paintable_compute_concrete_size")
	core.PuregoSafeRegister(&XGdkPaintableGetCurrentImage, lib, "gdk_paintable_get_current_image")
	core.PuregoSafeRegister(&XGdkPaintableGetFlags, lib, "gdk_paintable_get_flags")
	core.PuregoSafeRegister(&XGdkPaintableGetIntrinsicAspectRatio, lib, "gdk_paintable_get_intrinsic_aspect_ratio")
	core.PuregoSafeRegister(&XGdkPaintableGetIntrinsicHeight, lib, "gdk_paintable_get_intrinsic_height")
	core.PuregoSafeRegister(&XGdkPaintableGetIntrinsicWidth, lib, "gdk_paintable_get_intrinsic_width")
	core.PuregoSafeRegister(&XGdkPaintableInvalidateContents, lib, "gdk_paintable_invalidate_contents")
	core.PuregoSafeRegister(&XGdkPaintableInvalidateSize, lib, "gdk_paintable_invalidate_size")
	core.PuregoSafeRegister(&XGdkPaintableSnapshot, lib, "gdk_paintable_snapshot")

}
