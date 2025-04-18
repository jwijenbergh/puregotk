// Package graphene was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package graphene

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The location and size of a rectangle region.
//
// The width and height of a #graphene_rect_t can be negative; for instance,
// a #graphene_rect_t with an origin of [ 0, 0 ] and a size of [ 10, 10 ] is
// equivalent to a #graphene_rect_t with an origin of [ 10, 10 ] and a size
// of [ -10, -10 ].
//
// Application code can normalize rectangles using graphene_rect_normalize();
// this function will ensure that the width and height of a rectangle are
// positive values. All functions taking a #graphene_rect_t as an argument
// will internally operate on a normalized copy; all functions returning a
// #graphene_rect_t will always return a normalized rectangle.
type Rect struct {
	_ structs.HostLayout

	Origin uintptr

	Size uintptr
}

var xRectGLibType func() types.GType

func RectGLibType() types.GType {
	return xRectGLibType()
}

func (x *Rect) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xRectContainsPoint func(uintptr, *Point) bool

// Checks whether a #graphene_rect_t contains the given coordinates.
func (x *Rect) ContainsPoint(PVar *Point) bool {

	cret := xRectContainsPoint(x.GoPointer(), PVar)
	return cret
}

var xRectContainsRect func(uintptr, *Rect) bool

// Checks whether a #graphene_rect_t fully contains the given
// rectangle.
func (x *Rect) ContainsRect(BVar *Rect) bool {

	cret := xRectContainsRect(x.GoPointer(), BVar)
	return cret
}

var xRectEqual func(uintptr, *Rect) bool

// Checks whether the two given rectangle are equal.
func (x *Rect) Equal(BVar *Rect) bool {

	cret := xRectEqual(x.GoPointer(), BVar)
	return cret
}

var xRectExpand func(uintptr, *Point, *Rect)

// Expands a #graphene_rect_t to contain the given #graphene_point_t.
func (x *Rect) Expand(PVar *Point, ResVar *Rect) {

	xRectExpand(x.GoPointer(), PVar, ResVar)

}

var xRectFree func(uintptr)

// Frees the resources allocated by graphene_rect_alloc().
func (x *Rect) Free() {

	xRectFree(x.GoPointer())

}

var xRectGetArea func(uintptr) float32

// Compute the area of given normalized rectangle.
func (x *Rect) GetArea() float32 {

	cret := xRectGetArea(x.GoPointer())
	return cret
}

var xRectGetBottomLeft func(uintptr, *Point)

// Retrieves the coordinates of the bottom-left corner of the given rectangle.
func (x *Rect) GetBottomLeft(PVar *Point) {

	xRectGetBottomLeft(x.GoPointer(), PVar)

}

var xRectGetBottomRight func(uintptr, *Point)

// Retrieves the coordinates of the bottom-right corner of the given rectangle.
func (x *Rect) GetBottomRight(PVar *Point) {

	xRectGetBottomRight(x.GoPointer(), PVar)

}

var xRectGetCenter func(uintptr, *Point)

// Retrieves the coordinates of the center of the given rectangle.
func (x *Rect) GetCenter(PVar *Point) {

	xRectGetCenter(x.GoPointer(), PVar)

}

var xRectGetHeight func(uintptr) float32

// Retrieves the normalized height of the given rectangle.
func (x *Rect) GetHeight() float32 {

	cret := xRectGetHeight(x.GoPointer())
	return cret
}

var xRectGetTopLeft func(uintptr, *Point)

// Retrieves the coordinates of the top-left corner of the given rectangle.
func (x *Rect) GetTopLeft(PVar *Point) {

	xRectGetTopLeft(x.GoPointer(), PVar)

}

var xRectGetTopRight func(uintptr, *Point)

// Retrieves the coordinates of the top-right corner of the given rectangle.
func (x *Rect) GetTopRight(PVar *Point) {

	xRectGetTopRight(x.GoPointer(), PVar)

}

var xRectGetVertices func(uintptr, [4]Vec2)

// Computes the four vertices of a #graphene_rect_t.
func (x *Rect) GetVertices(VerticesVar [4]Vec2) {

	xRectGetVertices(x.GoPointer(), VerticesVar)

}

var xRectGetWidth func(uintptr) float32

// Retrieves the normalized width of the given rectangle.
func (x *Rect) GetWidth() float32 {

	cret := xRectGetWidth(x.GoPointer())
	return cret
}

var xRectGetX func(uintptr) float32

// Retrieves the normalized X coordinate of the origin of the given
// rectangle.
func (x *Rect) GetX() float32 {

	cret := xRectGetX(x.GoPointer())
	return cret
}

var xRectGetY func(uintptr) float32

// Retrieves the normalized Y coordinate of the origin of the given
// rectangle.
func (x *Rect) GetY() float32 {

	cret := xRectGetY(x.GoPointer())
	return cret
}

var xRectInit func(uintptr, float32, float32, float32, float32) *Rect

// Initializes the given #graphene_rect_t with the given values.
//
// This function will implicitly normalize the #graphene_rect_t
// before returning.
func (x *Rect) Init(XVar float32, YVar float32, WidthVar float32, HeightVar float32) *Rect {

	cret := xRectInit(x.GoPointer(), XVar, YVar, WidthVar, HeightVar)
	return cret
}

var xRectInitFromRect func(uintptr, *Rect) *Rect

// Initializes @r using the given @src rectangle.
//
// This function will implicitly normalize the #graphene_rect_t
// before returning.
func (x *Rect) InitFromRect(SrcVar *Rect) *Rect {

	cret := xRectInitFromRect(x.GoPointer(), SrcVar)
	return cret
}

var xRectInset func(uintptr, float32, float32) *Rect

// Changes the given rectangle to be smaller, or larger depending on the
// given inset parameters.
//
// To create an inset rectangle, use positive @d_x or @d_y values; to
// create a larger, encompassing rectangle, use negative @d_x or @d_y
// values.
//
// The origin of the rectangle is offset by @d_x and @d_y, while the size
// is adjusted by `(2 * @d_x, 2 * @d_y)`. If @d_x and @d_y are positive
// values, the size of the rectangle is decreased; if @d_x and @d_y are
// negative values, the size of the rectangle is increased.
//
// If the size of the resulting inset rectangle has a negative width or
// height then the size will be set to zero.
func (x *Rect) Inset(DXVar float32, DYVar float32) *Rect {

	cret := xRectInset(x.GoPointer(), DXVar, DYVar)
	return cret
}

var xRectInsetR func(uintptr, float32, float32, *Rect)

// Changes the given rectangle to be smaller, or larger depending on the
// given inset parameters.
//
// To create an inset rectangle, use positive @d_x or @d_y values; to
// create a larger, encompassing rectangle, use negative @d_x or @d_y
// values.
//
// The origin of the rectangle is offset by @d_x and @d_y, while the size
// is adjusted by `(2 * @d_x, 2 * @d_y)`. If @d_x and @d_y are positive
// values, the size of the rectangle is decreased; if @d_x and @d_y are
// negative values, the size of the rectangle is increased.
//
// If the size of the resulting inset rectangle has a negative width or
// height then the size will be set to zero.
func (x *Rect) InsetR(DXVar float32, DYVar float32, ResVar *Rect) {

	xRectInsetR(x.GoPointer(), DXVar, DYVar, ResVar)

}

var xRectInterpolate func(uintptr, *Rect, float64, *Rect)

// Linearly interpolates the origin and size of the two given
// rectangles.
func (x *Rect) Interpolate(BVar *Rect, FactorVar float64, ResVar *Rect) {

	xRectInterpolate(x.GoPointer(), BVar, FactorVar, ResVar)

}

var xRectIntersection func(uintptr, *Rect, *Rect) bool

// Computes the intersection of the two given rectangles.
//
// ![](rectangle-intersection.png)
//
// The intersection in the image above is the blue outline.
//
// If the two rectangles do not intersect, @res will contain
// a degenerate rectangle with origin in (0, 0) and a size of 0.
func (x *Rect) Intersection(BVar *Rect, ResVar *Rect) bool {

	cret := xRectIntersection(x.GoPointer(), BVar, ResVar)
	return cret
}

var xRectNormalize func(uintptr) *Rect

// Normalizes the passed rectangle.
//
// This function ensures that the size of the rectangle is made of
// positive values, and that the origin is the top-left corner of
// the rectangle.
func (x *Rect) Normalize() *Rect {

	cret := xRectNormalize(x.GoPointer())
	return cret
}

var xRectNormalizeR func(uintptr, *Rect)

// Normalizes the passed rectangle.
//
// This function ensures that the size of the rectangle is made of
// positive values, and that the origin is in the top-left corner
// of the rectangle.
func (x *Rect) NormalizeR(ResVar *Rect) {

	xRectNormalizeR(x.GoPointer(), ResVar)

}

var xRectOffset func(uintptr, float32, float32) *Rect

// Offsets the origin by @d_x and @d_y.
//
// The size of the rectangle is unchanged.
func (x *Rect) Offset(DXVar float32, DYVar float32) *Rect {

	cret := xRectOffset(x.GoPointer(), DXVar, DYVar)
	return cret
}

var xRectOffsetR func(uintptr, float32, float32, *Rect)

// Offsets the origin of the given rectangle by @d_x and @d_y.
//
// The size of the rectangle is left unchanged.
func (x *Rect) OffsetR(DXVar float32, DYVar float32, ResVar *Rect) {

	xRectOffsetR(x.GoPointer(), DXVar, DYVar, ResVar)

}

var xRectRound func(uintptr, *Rect)

// Rounds the origin and size of the given rectangle to
// their nearest integer values; the rounding is guaranteed
// to be large enough to have an area bigger or equal to the
// original rectangle, but might not fully contain its extents.
// Use graphene_rect_round_extents() in case you need to round
// to a rectangle that covers fully the original one.
//
// This function is the equivalent of calling `floor` on
// the coordinates of the origin, and `ceil` on the size.
func (x *Rect) Round(ResVar *Rect) {

	xRectRound(x.GoPointer(), ResVar)

}

var xRectRoundExtents func(uintptr, *Rect)

// Rounds the origin of the given rectangle to its nearest
// integer value and and recompute the size so that the
// rectangle is large enough to contain all the conrners
// of the original rectangle.
//
// This function is the equivalent of calling `floor` on
// the coordinates of the origin, and recomputing the size
// calling `ceil` on the bottom-right coordinates.
//
// If you want to be sure that the rounded rectangle
// completely covers the area that was covered by the
// original rectangle — i.e. you want to cover the area
// including all its corners — this function will make sure
// that the size is recomputed taking into account the ceiling
// of the coordinates of the bottom-right corner.
// If the difference between the original coordinates and the
// coordinates of the rounded rectangle is greater than the
// difference between the original size and and the rounded
// size, then the move of the origin would not be compensated
// by a move in the anti-origin, leaving the corners of the
// original rectangle outside the rounded one.
func (x *Rect) RoundExtents(ResVar *Rect) {

	xRectRoundExtents(x.GoPointer(), ResVar)

}

var xRectRoundToPixel func(uintptr) *Rect

// Rounds the origin and the size of the given rectangle to
// their nearest integer values; the rounding is guaranteed
// to be large enough to contain the original rectangle.
func (x *Rect) RoundToPixel() *Rect {

	cret := xRectRoundToPixel(x.GoPointer())
	return cret
}

var xRectScale func(uintptr, float32, float32, *Rect)

// Scales the size and origin of a rectangle horizontaly by @s_h,
// and vertically by @s_v. The result @res is normalized.
func (x *Rect) Scale(SHVar float32, SVVar float32, ResVar *Rect) {

	xRectScale(x.GoPointer(), SHVar, SVVar, ResVar)

}

var xRectUnion func(uintptr, *Rect, *Rect)

// Computes the union of the two given rectangles.
//
// ![](rectangle-union.png)
//
// The union in the image above is the blue outline.
func (x *Rect) Union(BVar *Rect, ResVar *Rect) {

	xRectUnion(x.GoPointer(), BVar, ResVar)

}

var xRectAlloc func() *Rect

// Allocates a new #graphene_rect_t.
//
// The contents of the returned rectangle are undefined.
func RectAlloc() *Rect {

	cret := xRectAlloc()
	return cret
}

var xRectZero func() *Rect

// Returns a degenerate rectangle with origin fixed at (0, 0) and
// a size of 0, 0.
func RectZero() *Rect {

	cret := xRectZero()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GRAPHENE"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xRectAlloc, lib, "graphene_rect_alloc")
	core.PuregoSafeRegister(&xRectZero, lib, "graphene_rect_zero")

	core.PuregoSafeRegister(&xRectGLibType, lib, "graphene_rect_get_type")

	core.PuregoSafeRegister(&xRectContainsPoint, lib, "graphene_rect_contains_point")
	core.PuregoSafeRegister(&xRectContainsRect, lib, "graphene_rect_contains_rect")
	core.PuregoSafeRegister(&xRectEqual, lib, "graphene_rect_equal")
	core.PuregoSafeRegister(&xRectExpand, lib, "graphene_rect_expand")
	core.PuregoSafeRegister(&xRectFree, lib, "graphene_rect_free")
	core.PuregoSafeRegister(&xRectGetArea, lib, "graphene_rect_get_area")
	core.PuregoSafeRegister(&xRectGetBottomLeft, lib, "graphene_rect_get_bottom_left")
	core.PuregoSafeRegister(&xRectGetBottomRight, lib, "graphene_rect_get_bottom_right")
	core.PuregoSafeRegister(&xRectGetCenter, lib, "graphene_rect_get_center")
	core.PuregoSafeRegister(&xRectGetHeight, lib, "graphene_rect_get_height")
	core.PuregoSafeRegister(&xRectGetTopLeft, lib, "graphene_rect_get_top_left")
	core.PuregoSafeRegister(&xRectGetTopRight, lib, "graphene_rect_get_top_right")
	core.PuregoSafeRegister(&xRectGetVertices, lib, "graphene_rect_get_vertices")
	core.PuregoSafeRegister(&xRectGetWidth, lib, "graphene_rect_get_width")
	core.PuregoSafeRegister(&xRectGetX, lib, "graphene_rect_get_x")
	core.PuregoSafeRegister(&xRectGetY, lib, "graphene_rect_get_y")
	core.PuregoSafeRegister(&xRectInit, lib, "graphene_rect_init")
	core.PuregoSafeRegister(&xRectInitFromRect, lib, "graphene_rect_init_from_rect")
	core.PuregoSafeRegister(&xRectInset, lib, "graphene_rect_inset")
	core.PuregoSafeRegister(&xRectInsetR, lib, "graphene_rect_inset_r")
	core.PuregoSafeRegister(&xRectInterpolate, lib, "graphene_rect_interpolate")
	core.PuregoSafeRegister(&xRectIntersection, lib, "graphene_rect_intersection")
	core.PuregoSafeRegister(&xRectNormalize, lib, "graphene_rect_normalize")
	core.PuregoSafeRegister(&xRectNormalizeR, lib, "graphene_rect_normalize_r")
	core.PuregoSafeRegister(&xRectOffset, lib, "graphene_rect_offset")
	core.PuregoSafeRegister(&xRectOffsetR, lib, "graphene_rect_offset_r")
	core.PuregoSafeRegister(&xRectRound, lib, "graphene_rect_round")
	core.PuregoSafeRegister(&xRectRoundExtents, lib, "graphene_rect_round_extents")
	core.PuregoSafeRegister(&xRectRoundToPixel, lib, "graphene_rect_round_to_pixel")
	core.PuregoSafeRegister(&xRectScale, lib, "graphene_rect_scale")
	core.PuregoSafeRegister(&xRectUnion, lib, "graphene_rect_union")

}
