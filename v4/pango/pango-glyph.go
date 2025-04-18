// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The `PangoGlyphGeometry` structure contains width and positioning
// information for a single glyph.
//
// Note that @width is not guaranteed to be the same as the glyph
// extents. Kerning and other positioning applied during shaping will
// affect both the @width and the @x_offset for the glyphs in the
// glyph string that results from shaping.
//
// The information in this struct is intended for rendering the glyphs,
// as follows:
//
// 1. Assume the current point is (x, y)
// 2. Render the current glyph at (x + x_offset, y + y_offset),
// 3. Advance the current point to (x + width, y)
// 4. Render the next glyph
type GlyphGeometry struct {
	_ structs.HostLayout

	Width GlyphUnit

	XOffset GlyphUnit

	YOffset GlyphUnit
}

func (x *GlyphGeometry) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `PangoGlyphInfo` structure represents a single glyph with
// positioning information and visual attributes.
type GlyphInfo struct {
	_ structs.HostLayout

	Glyph Glyph

	Geometry uintptr

	Attr uintptr
}

func (x *GlyphInfo) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A `PangoGlyphString` is used to store strings of glyphs with geometry
// and visual attribute information.
//
// The storage for the glyph information is owned by the structure
// which simplifies memory management.
type GlyphString struct {
	_ structs.HostLayout

	NumGlyphs int

	Glyphs []GlyphInfo

	LogClusters int

	Space int
}

var xGlyphStringGLibType func() types.GType

func GlyphStringGLibType() types.GType {
	return xGlyphStringGLibType()
}

func (x *GlyphString) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewGlyphString func() *GlyphString

// Create a new `PangoGlyphString`.
func NewGlyphString() *GlyphString {

	cret := xNewGlyphString()
	return cret
}

var xGlyphStringCopy func(uintptr) *GlyphString

// Copy a glyph string and associated storage.
func (x *GlyphString) Copy() *GlyphString {

	cret := xGlyphStringCopy(x.GoPointer())
	return cret
}

var xGlyphStringExtents func(uintptr, uintptr, *Rectangle, *Rectangle)

// Compute the logical and ink extents of a glyph string.
//
// See the documentation for [method@Pango.Font.get_glyph_extents] for details
// about the interpretation of the rectangles.
//
// Examples of logical (red) and ink (green) rects:
//
// ![](rects1.png) ![](rects2.png)
func (x *GlyphString) Extents(FontVar *Font, InkRectVar *Rectangle, LogicalRectVar *Rectangle) {

	xGlyphStringExtents(x.GoPointer(), FontVar.GoPointer(), InkRectVar, LogicalRectVar)

}

var xGlyphStringExtentsRange func(uintptr, int, int, uintptr, *Rectangle, *Rectangle)

// Computes the extents of a sub-portion of a glyph string.
//
// The extents are relative to the start of the glyph string range
// (the origin of their coordinate system is at the start of the range,
// not at the start of the entire glyph string).
func (x *GlyphString) ExtentsRange(StartVar int, EndVar int, FontVar *Font, InkRectVar *Rectangle, LogicalRectVar *Rectangle) {

	xGlyphStringExtentsRange(x.GoPointer(), StartVar, EndVar, FontVar.GoPointer(), InkRectVar, LogicalRectVar)

}

var xGlyphStringFree func(uintptr)

// Free a glyph string and associated storage.
func (x *GlyphString) Free() {

	xGlyphStringFree(x.GoPointer())

}

var xGlyphStringGetLogicalWidths func(uintptr, string, int, int, []int)

// Given a `PangoGlyphString` and corresponding text, determine the width
// corresponding to each character.
//
// When multiple characters compose a single cluster, the width of the
// entire cluster is divided equally among the characters.
//
// See also [method@Pango.GlyphItem.get_logical_widths].
func (x *GlyphString) GetLogicalWidths(TextVar string, LengthVar int, EmbeddingLevelVar int, LogicalWidthsVar []int) {

	xGlyphStringGetLogicalWidths(x.GoPointer(), TextVar, LengthVar, EmbeddingLevelVar, LogicalWidthsVar)

}

var xGlyphStringGetWidth func(uintptr) int

// Computes the logical width of the glyph string.
//
// This can also be computed using [method@Pango.GlyphString.extents].
// However, since this only computes the width, it's much faster. This
// is in fact only a convenience function that computes the sum of
// @geometry.width for each glyph in the @glyphs.
func (x *GlyphString) GetWidth() int {

	cret := xGlyphStringGetWidth(x.GoPointer())
	return cret
}

var xGlyphStringIndexToX func(uintptr, string, int, *Analysis, int, bool, int)

// Converts from character position to x position.
//
// The X position is measured from the left edge of the run.
// Character positions are obtained using font metrics for ligatures
// where available, and computed by dividing up each cluster
// into equal portions, otherwise.
//
// &lt;picture&gt;
//
//	&lt;source srcset="glyphstring-positions-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img alt="Glyph positions" src="glyphstring-positions-light.png"&gt;
//
// &lt;/picture&gt;
func (x *GlyphString) IndexToX(TextVar string, LengthVar int, AnalysisVar *Analysis, IndexVar int, TrailingVar bool, XPosVar int) {

	xGlyphStringIndexToX(x.GoPointer(), TextVar, LengthVar, AnalysisVar, IndexVar, TrailingVar, XPosVar)

}

var xGlyphStringIndexToXFull func(uintptr, string, int, *Analysis, *LogAttr, int, bool, int)

// Converts from character position to x position.
//
// This variant of [method@Pango.GlyphString.index_to_x] additionally
// accepts a `PangoLogAttr` array. The grapheme boundary information
// in it can be used to disambiguate positioning inside some complex
// clusters.
func (x *GlyphString) IndexToXFull(TextVar string, LengthVar int, AnalysisVar *Analysis, AttrsVar *LogAttr, IndexVar int, TrailingVar bool, XPosVar int) {

	xGlyphStringIndexToXFull(x.GoPointer(), TextVar, LengthVar, AnalysisVar, AttrsVar, IndexVar, TrailingVar, XPosVar)

}

var xGlyphStringSetSize func(uintptr, int)

// Resize a glyph string to the given length.
func (x *GlyphString) SetSize(NewLenVar int) {

	xGlyphStringSetSize(x.GoPointer(), NewLenVar)

}

var xGlyphStringXToIndex func(uintptr, string, int, *Analysis, int, int, int)

// Convert from x offset to character position.
//
// Character positions are computed by dividing up each cluster into
// equal portions. In scripts where positioning within a cluster is
// not allowed (such as Thai), the returned value may not be a valid
// cursor position; the caller must combine the result with the logical
// attributes for the text to compute the valid cursor position.
func (x *GlyphString) XToIndex(TextVar string, LengthVar int, AnalysisVar *Analysis, XPosVar int, IndexVar int, TrailingVar int) {

	xGlyphStringXToIndex(x.GoPointer(), TextVar, LengthVar, AnalysisVar, XPosVar, IndexVar, TrailingVar)

}

// A `PangoGlyphVisAttr` structure communicates information between
// the shaping and rendering phases.
//
// Currently, it contains cluster start and color information.
// More attributes may be added in the future.
//
// Clusters are stored in visual order, within the cluster, glyphs
// are always ordered in logical order, since visual order is meaningless;
// that is, in Arabic text, accent glyphs follow the glyphs for the
// base character.
type GlyphVisAttr struct {
	_ structs.HostLayout

	IsClusterStart uint

	IsColor uint
}

func (x *GlyphVisAttr) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// The `PangoGlyphUnit` type is used to store dimensions within
// Pango.
//
// Dimensions are stored in 1/PANGO_SCALE of a device unit.
// (A device unit might be a pixel for screen display, or
// a point on a printer.) PANGO_SCALE is currently 1024, and
// may change in the future (unlikely though), but you should not
// depend on its exact value.
//
// The PANGO_PIXELS() macro can be used to convert from glyph units
// into device units with correct rounding.
type GlyphUnit = int32

// Flags influencing the shaping process.
//
// `PangoShapeFlags` can be passed to [func@Pango.shape_with_flags].
type ShapeFlags int

var xShapeFlagsGLibType func() types.GType

func ShapeFlagsGLibType() types.GType {
	return xShapeFlagsGLibType()
}

const (

	// Default value
	ShapeNoneValue ShapeFlags = 0
	// Round glyph positions and widths to whole device units
	//   This option should be set if the target renderer can't do subpixel positioning of glyphs
	ShapeRoundPositionsValue ShapeFlags = 1
)

var xShape func(string, int, *Analysis, *GlyphString)

// Convert the characters in @text into glyphs.
//
// Given a segment of text and the corresponding `PangoAnalysis` structure
// returned from [func@Pango.itemize], convert the characters into glyphs. You
// may also pass in only a substring of the item from [func@Pango.itemize].
//
// It is recommended that you use [func@Pango.shape_full] instead, since
// that API allows for shaping interaction happening across text item
// boundaries.
//
// Note that the extra attributes in the @analyis that is returned from
// [func@Pango.itemize] have indices that are relative to the entire paragraph,
// so you need to subtract the item offset from their indices before
// calling [func@Pango.shape].
func Shape(TextVar string, LengthVar int, AnalysisVar *Analysis, GlyphsVar *GlyphString) {

	xShape(TextVar, LengthVar, AnalysisVar, GlyphsVar)

}

var xShapeFull func(string, int, string, int, *Analysis, *GlyphString)

// Convert the characters in @text into glyphs.
//
// Given a segment of text and the corresponding `PangoAnalysis` structure
// returned from [func@Pango.itemize], convert the characters into glyphs.
// You may also pass in only a substring of the item from [func@Pango.itemize].
//
// This is similar to [func@Pango.shape], except it also can optionally take
// the full paragraph text as input, which will then be used to perform
// certain cross-item shaping interactions. If you have access to the broader
// text of which @item_text is part of, provide the broader text as
// @paragraph_text. If @paragraph_text is %NULL, item text is used instead.
//
// Note that the extra attributes in the @analyis that is returned from
// [func@Pango.itemize] have indices that are relative to the entire paragraph,
// so you do not pass the full paragraph text as @paragraph_text, you need
// to subtract the item offset from their indices before calling
// [func@Pango.shape_full].
func ShapeFull(ItemTextVar string, ItemLengthVar int, ParagraphTextVar string, ParagraphLengthVar int, AnalysisVar *Analysis, GlyphsVar *GlyphString) {

	xShapeFull(ItemTextVar, ItemLengthVar, ParagraphTextVar, ParagraphLengthVar, AnalysisVar, GlyphsVar)

}

var xShapeItem func(*Item, string, int, *LogAttr, *GlyphString, ShapeFlags)

// Convert the characters in @item into glyphs.
//
// This is similar to [func@Pango.shape_with_flags], except it takes a
// `PangoItem` instead of separate @item_text and @analysis arguments.
// It also takes @log_attrs, which may be used in implementing text
// transforms.
//
// Note that the extra attributes in the @analyis that is returned from
// [func@Pango.itemize] have indices that are relative to the entire paragraph,
// so you do not pass the full paragraph text as @paragraph_text, you need
// to subtract the item offset from their indices before calling
// [func@Pango.shape_with_flags].
func ShapeItem(ItemVar *Item, ParagraphTextVar string, ParagraphLengthVar int, LogAttrsVar *LogAttr, GlyphsVar *GlyphString, FlagsVar ShapeFlags) {

	xShapeItem(ItemVar, ParagraphTextVar, ParagraphLengthVar, LogAttrsVar, GlyphsVar, FlagsVar)

}

var xShapeWithFlags func(string, int, string, int, *Analysis, *GlyphString, ShapeFlags)

// Convert the characters in @text into glyphs.
//
// Given a segment of text and the corresponding `PangoAnalysis` structure
// returned from [func@Pango.itemize], convert the characters into glyphs.
// You may also pass in only a substring of the item from [func@Pango.itemize].
//
// This is similar to [func@Pango.shape_full], except it also takes flags
// that can influence the shaping process.
//
// Note that the extra attributes in the @analyis that is returned from
// [func@Pango.itemize] have indices that are relative to the entire paragraph,
// so you do not pass the full paragraph text as @paragraph_text, you need
// to subtract the item offset from their indices before calling
// [func@Pango.shape_with_flags].
func ShapeWithFlags(ItemTextVar string, ItemLengthVar int, ParagraphTextVar string, ParagraphLengthVar int, AnalysisVar *Analysis, GlyphsVar *GlyphString, FlagsVar ShapeFlags) {

	xShapeWithFlags(ItemTextVar, ItemLengthVar, ParagraphTextVar, ParagraphLengthVar, AnalysisVar, GlyphsVar, FlagsVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xShapeFlagsGLibType, lib, "pango_shape_flags_get_type")

	core.PuregoSafeRegister(&xShape, lib, "pango_shape")
	core.PuregoSafeRegister(&xShapeFull, lib, "pango_shape_full")
	core.PuregoSafeRegister(&xShapeItem, lib, "pango_shape_item")
	core.PuregoSafeRegister(&xShapeWithFlags, lib, "pango_shape_with_flags")

	core.PuregoSafeRegister(&xGlyphStringGLibType, lib, "pango_glyph_string_get_type")

	core.PuregoSafeRegister(&xNewGlyphString, lib, "pango_glyph_string_new")

	core.PuregoSafeRegister(&xGlyphStringCopy, lib, "pango_glyph_string_copy")
	core.PuregoSafeRegister(&xGlyphStringExtents, lib, "pango_glyph_string_extents")
	core.PuregoSafeRegister(&xGlyphStringExtentsRange, lib, "pango_glyph_string_extents_range")
	core.PuregoSafeRegister(&xGlyphStringFree, lib, "pango_glyph_string_free")
	core.PuregoSafeRegister(&xGlyphStringGetLogicalWidths, lib, "pango_glyph_string_get_logical_widths")
	core.PuregoSafeRegister(&xGlyphStringGetWidth, lib, "pango_glyph_string_get_width")
	core.PuregoSafeRegister(&xGlyphStringIndexToX, lib, "pango_glyph_string_index_to_x")
	core.PuregoSafeRegister(&xGlyphStringIndexToXFull, lib, "pango_glyph_string_index_to_x_full")
	core.PuregoSafeRegister(&xGlyphStringSetSize, lib, "pango_glyph_string_set_size")
	core.PuregoSafeRegister(&xGlyphStringXToIndex, lib, "pango_glyph_string_x_to_index")

}
