// Package pangocairo was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pangocairo

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/cairo"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

// Function type for rendering attributes of type %PANGO_ATTR_SHAPE
// with Pango's Cairo renderer.
type ShapeRendererFunc func(*cairo.Context, *pango.AttrShape, bool, uintptr)

// `PangoCairoFont` is an interface exported by fonts for
// use with Cairo.
//
// The actual type of the font will depend on the particular
// font technology Cairo was compiled to use.
type Font interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetScaledFont() *cairo.ScaledFont
}

var xFontGLibType func() types.GType

func FontGLibType() types.GType {
	return xFontGLibType()
}

type FontBase struct {
	Ptr uintptr
}

func (x *FontBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *FontBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the `cairo_scaled_font_t` used by @font.
// The scaled font can be referenced and kept using
// cairo_scaled_font_reference().
func (x *FontBase) GetScaledFont() *cairo.ScaledFont {

	cret := XPangoCairoFontGetScaledFont(x.GoPointer())
	return cret
}

var XPangoCairoFontGetScaledFont func(uintptr) *cairo.ScaledFont

// `PangoCairoFontMap` is an interface exported by font maps for
// use with Cairo.
//
// The actual type of the font map will depend on the particular
// font technology Cairo was compiled to use.
type FontMap interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	CreateContext() *pango.Context
	GetFontType() cairo.FontType
	GetResolution() float64
	SetDefault()
	SetResolution(DpiVar float64)
}

var xFontMapGLibType func() types.GType

func FontMapGLibType() types.GType {
	return xFontMapGLibType()
}

type FontMapBase struct {
	Ptr uintptr
}

func (x *FontMapBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *FontMapBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Create a `PangoContext` for the given fontmap.
func (x *FontMapBase) CreateContext() *pango.Context {
	var cls *pango.Context

	cret := XPangoCairoFontMapCreateContext(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &pango.Context{}
	cls.Ptr = cret
	return cls
}

// Gets the type of Cairo font backend that @fontmap uses.
func (x *FontMapBase) GetFontType() cairo.FontType {

	cret := XPangoCairoFontMapGetFontType(x.GoPointer())
	return cret
}

// Gets the resolution for the fontmap.
//
// See [method@PangoCairo.FontMap.set_resolution].
func (x *FontMapBase) GetResolution() float64 {

	cret := XPangoCairoFontMapGetResolution(x.GoPointer())
	return cret
}

// Sets a default `PangoCairoFontMap` to use with Cairo.
//
// This can be used to change the Cairo font backend that the
// default fontmap uses for example. The old default font map
// is unreffed and the new font map referenced.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread.
// This function only changes the default fontmap for
// the current thread. Default fontmaps of existing threads
// are not changed. Default fontmaps of any new threads will
// still be created using [func@PangoCairo.FontMap.new].
//
// A value of %NULL for @fontmap will cause the current default
// font map to be released and a new default font map to be created
// on demand, using [func@PangoCairo.FontMap.new].
func (x *FontMapBase) SetDefault() {

	XPangoCairoFontMapSetDefault(x.GoPointer())

}

// Sets the resolution for the fontmap.
//
// This is a scale factor between
// points specified in a `PangoFontDescription` and Cairo units. The
// default value is 96, meaning that a 10 point font will be 13
// units high. (10 * 96. / 72. = 13.3).
func (x *FontMapBase) SetResolution(DpiVar float64) {

	XPangoCairoFontMapSetResolution(x.GoPointer(), DpiVar)

}

var XPangoCairoFontMapCreateContext func(uintptr) uintptr
var XPangoCairoFontMapGetFontType func(uintptr) cairo.FontType
var XPangoCairoFontMapGetResolution func(uintptr) float64
var XPangoCairoFontMapSetDefault func(uintptr)
var XPangoCairoFontMapSetResolution func(uintptr, float64)

var xContextGetFontOptions func(uintptr) *cairo.FontOptions

// Retrieves any font rendering options previously set with
// [func@PangoCairo.context_set_font_options].
//
// This function does not report options that are derived from
// the target surface by [func@update_context].
func ContextGetFontOptions(ContextVar *pango.Context) *cairo.FontOptions {

	cret := xContextGetFontOptions(ContextVar.GoPointer())
	return cret
}

var xContextGetResolution func(uintptr) float64

// Gets the resolution for the context.
//
// See [func@PangoCairo.context_set_resolution]
func ContextGetResolution(ContextVar *pango.Context) float64 {

	cret := xContextGetResolution(ContextVar.GoPointer())
	return cret
}

var xContextGetShapeRenderer func(uintptr, uintptr) uintptr

// Sets callback function for context to use for rendering attributes
// of type %PANGO_ATTR_SHAPE.
//
// See `PangoCairoShapeRendererFunc` for details.
//
// Retrieves callback function and associated user data for rendering
// attributes of type %PANGO_ATTR_SHAPE as set by
// [func@PangoCairo.context_set_shape_renderer], if any.
func ContextGetShapeRenderer(ContextVar *pango.Context, DataVar uintptr) uintptr {

	cret := xContextGetShapeRenderer(ContextVar.GoPointer(), DataVar)
	return cret
}

var xContextSetFontOptions func(uintptr, *cairo.FontOptions)

// Sets the font options used when rendering text with this context.
//
// These options override any options that [func@update_context]
// derives from the target surface.
func ContextSetFontOptions(ContextVar *pango.Context, OptionsVar *cairo.FontOptions) {

	xContextSetFontOptions(ContextVar.GoPointer(), OptionsVar)

}

var xContextSetResolution func(uintptr, float64)

// Sets the resolution for the context.
//
// This is a scale factor between points specified in a `PangoFontDescription`
// and Cairo units. The default value is 96, meaning that a 10 point font will
// be 13 units high. (10 * 96. / 72. = 13.3).
func ContextSetResolution(ContextVar *pango.Context, DpiVar float64) {

	xContextSetResolution(ContextVar.GoPointer(), DpiVar)

}

var xContextSetShapeRenderer func(uintptr, uintptr, uintptr, uintptr)

// Sets callback function for context to use for rendering attributes
// of type %PANGO_ATTR_SHAPE.
//
// See `PangoCairoShapeRendererFunc` for details.
func ContextSetShapeRenderer(ContextVar *pango.Context, FuncVar *ShapeRendererFunc, DataVar uintptr, DnotifyVar *glib.DestroyNotify) {

	xContextSetShapeRenderer(ContextVar.GoPointer(), glib.NewCallback(FuncVar), DataVar, glib.NewCallback(DnotifyVar))

}

var xCreateContext func(*cairo.Context) uintptr

// Creates a context object set up to match the current transformation
// and target surface of the Cairo context.
//
// This context can then be
// used to create a layout using [ctor@Pango.Layout.new].
//
// This function is a convenience function that creates a context using
// the default font map, then updates it to @cr. If you just need to
// create a layout for use with @cr and do not need to access `PangoContext`
// directly, you can use [func@create_layout] instead.
func CreateContext(CrVar *cairo.Context) *pango.Context {
	var cls *pango.Context

	cret := xCreateContext(CrVar)

	if cret == 0 {
		return nil
	}
	cls = &pango.Context{}
	cls.Ptr = cret
	return cls
}

var xCreateLayout func(*cairo.Context) uintptr

// Creates a layout object set up to match the current transformation
// and target surface of the Cairo context.
//
// This layout can then be used for text measurement with functions
// like [method@Pango.Layout.get_size] or drawing with functions like
// [func@show_layout]. If you change the transformation or target
// surface for @cr, you need to call [func@update_layout].
//
// This function is the most convenient way to use Cairo with Pango,
// however it is slightly inefficient since it creates a separate
// `PangoContext` object for each layout. This might matter in an
// application that was laying out large amounts of text.
func CreateLayout(CrVar *cairo.Context) *pango.Layout {
	var cls *pango.Layout

	cret := xCreateLayout(CrVar)

	if cret == 0 {
		return nil
	}
	cls = &pango.Layout{}
	cls.Ptr = cret
	return cls
}

var xErrorUnderlinePath func(*cairo.Context, float64, float64, float64, float64)

// Add a squiggly line to the current path in the specified cairo context that
// approximately covers the given rectangle in the style of an underline used
// to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ErrorUnderlinePath(CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xErrorUnderlinePath(CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xFontMapGetDefault func() uintptr

// Gets a default `PangoCairoFontMap` to use with Cairo.
//
// Note that the type of the returned object will depend on the
// particular font backend Cairo was compiled to use; you generally
// should only use the `PangoFontMap` and `PangoCairoFontMap`
// interfaces on the returned object.
//
// The default Cairo fontmap can be changed by using
// [method@PangoCairo.FontMap.set_default]. This can be used to
// change the Cairo font backend that the default fontmap uses
// for example.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread.
// Each thread gets its own default fontmap. In this way, PangoCairo
// can be used safely from multiple threads.
func FontMapGetDefault() *pango.FontMap {
	var cls *pango.FontMap

	cret := xFontMapGetDefault()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &pango.FontMap{}
	cls.Ptr = cret
	return cls
}

var xFontMapNew func() uintptr

// Creates a new `PangoCairoFontMap` object.
//
// A fontmap is used to cache information about available fonts,
// and holds certain global parameters such as the resolution.
// In most cases, you can use `func@PangoCairo.font_map_get_default]
// instead.
//
// Note that the type of the returned object will depend
// on the particular font backend Cairo was compiled to use;
// You generally should only use the `PangoFontMap` and
// `PangoCairoFontMap` interfaces on the returned object.
//
// You can override the type of backend returned by using an
// environment variable %PANGOCAIRO_BACKEND. Supported types,
// based on your build, are fc (fontconfig), win32, and coretext.
// If requested type is not available, NULL is returned. Ie.
// this is only useful for testing, when at least two backends
// are compiled in.
func FontMapNew() *pango.FontMap {
	var cls *pango.FontMap

	cret := xFontMapNew()

	if cret == 0 {
		return nil
	}
	cls = &pango.FontMap{}
	cls.Ptr = cret
	return cls
}

var xFontMapNewForFontType func(cairo.FontType) uintptr

// Creates a new `PangoCairoFontMap` object of the type suitable
// to be used with cairo font backend of type @fonttype.
//
// In most cases one should simply use [func@PangoCairo.FontMap.new], or
// in fact in most of those cases, just use [func@PangoCairo.FontMap.get_default].
func FontMapNewForFontType(FonttypeVar cairo.FontType) *pango.FontMap {
	var cls *pango.FontMap

	cret := xFontMapNewForFontType(FonttypeVar)

	if cret == 0 {
		return nil
	}
	cls = &pango.FontMap{}
	cls.Ptr = cret
	return cls
}

var xGlyphStringPath func(*cairo.Context, uintptr, *pango.GlyphString)

// Adds the glyphs in @glyphs to the current path in the specified
// cairo context.
//
// The origin of the glyphs (the left edge of the baseline)
// will be at the current point of the cairo context.
func GlyphStringPath(CrVar *cairo.Context, FontVar *pango.Font, GlyphsVar *pango.GlyphString) {

	xGlyphStringPath(CrVar, FontVar.GoPointer(), GlyphsVar)

}

var xLayoutLinePath func(*cairo.Context, *pango.LayoutLine)

// Adds the text in `PangoLayoutLine` to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will be
// at the current point of the cairo context.
func LayoutLinePath(CrVar *cairo.Context, LineVar *pango.LayoutLine) {

	xLayoutLinePath(CrVar, LineVar)

}

var xLayoutPath func(*cairo.Context, uintptr)

// Adds the text in a `PangoLayout` to the current path in the
// specified cairo context.
//
// The top-left corner of the `PangoLayout` will be at the
// current point of the cairo context.
func LayoutPath(CrVar *cairo.Context, LayoutVar *pango.Layout) {

	xLayoutPath(CrVar, LayoutVar.GoPointer())

}

var xShowErrorUnderline func(*cairo.Context, float64, float64, float64, float64)

// Draw a squiggly line in the specified cairo context that approximately
// covers the given rectangle in the style of an underline used to indicate a
// spelling error.
//
// The width of the underline is rounded to an integer
// number of up/down segments and the resulting rectangle is centered in the
// original rectangle.
func ShowErrorUnderline(CrVar *cairo.Context, XVar float64, YVar float64, WidthVar float64, HeightVar float64) {

	xShowErrorUnderline(CrVar, XVar, YVar, WidthVar, HeightVar)

}

var xShowGlyphItem func(*cairo.Context, string, *pango.GlyphItem)

// Draws the glyphs in @glyph_item in the specified cairo context,
//
// embedding the text associated with the glyphs in the output if the
// output format supports it (PDF for example), otherwise it acts
// similar to [func@show_glyph_string].
//
// The origin of the glyphs (the left edge of the baseline) will
// be drawn at the current point of the cairo context.
//
// Note that @text is the start of the text for layout, which is then
// indexed by `glyph_item-&gt;item-&gt;offset`.
func ShowGlyphItem(CrVar *cairo.Context, TextVar string, GlyphItemVar *pango.GlyphItem) {

	xShowGlyphItem(CrVar, TextVar, GlyphItemVar)

}

var xShowGlyphString func(*cairo.Context, uintptr, *pango.GlyphString)

// Draws the glyphs in @glyphs in the specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will
// be drawn at the current point of the cairo context.
func ShowGlyphString(CrVar *cairo.Context, FontVar *pango.Font, GlyphsVar *pango.GlyphString) {

	xShowGlyphString(CrVar, FontVar.GoPointer(), GlyphsVar)

}

var xShowLayout func(*cairo.Context, uintptr)

// Draws a `PangoLayout` in the specified cairo context.
//
// The top-left corner of the `PangoLayout` will be drawn
// at the current point of the cairo context.
func ShowLayout(CrVar *cairo.Context, LayoutVar *pango.Layout) {

	xShowLayout(CrVar, LayoutVar.GoPointer())

}

var xShowLayoutLine func(*cairo.Context, *pango.LayoutLine)

// Draws a `PangoLayoutLine` in the specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will
// be drawn at the current point of the cairo context.
func ShowLayoutLine(CrVar *cairo.Context, LineVar *pango.LayoutLine) {

	xShowLayoutLine(CrVar, LineVar)

}

var xUpdateContext func(*cairo.Context, uintptr)

// Updates a `PangoContext` previously created for use with Cairo to
// match the current transformation and target surface of a Cairo
// context.
//
// If any layouts have been created for the context, it's necessary
// to call [method@Pango.Layout.context_changed] on those layouts.
func UpdateContext(CrVar *cairo.Context, ContextVar *pango.Context) {

	xUpdateContext(CrVar, ContextVar.GoPointer())

}

var xUpdateLayout func(*cairo.Context, uintptr)

// Updates the private `PangoContext` of a `PangoLayout` created with
// [func@create_layout] to match the current transformation and target
// surface of a Cairo context.
func UpdateLayout(CrVar *cairo.Context, LayoutVar *pango.Layout) {

	xUpdateLayout(CrVar, LayoutVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGOCAIRO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xContextGetFontOptions, lib, "pango_cairo_context_get_font_options")
	core.PuregoSafeRegister(&xContextGetResolution, lib, "pango_cairo_context_get_resolution")
	core.PuregoSafeRegister(&xContextGetShapeRenderer, lib, "pango_cairo_context_get_shape_renderer")
	core.PuregoSafeRegister(&xContextSetFontOptions, lib, "pango_cairo_context_set_font_options")
	core.PuregoSafeRegister(&xContextSetResolution, lib, "pango_cairo_context_set_resolution")
	core.PuregoSafeRegister(&xContextSetShapeRenderer, lib, "pango_cairo_context_set_shape_renderer")
	core.PuregoSafeRegister(&xCreateContext, lib, "pango_cairo_create_context")
	core.PuregoSafeRegister(&xCreateLayout, lib, "pango_cairo_create_layout")
	core.PuregoSafeRegister(&xErrorUnderlinePath, lib, "pango_cairo_error_underline_path")
	core.PuregoSafeRegister(&xFontMapGetDefault, lib, "pango_cairo_font_map_get_default")
	core.PuregoSafeRegister(&xFontMapNew, lib, "pango_cairo_font_map_new")
	core.PuregoSafeRegister(&xFontMapNewForFontType, lib, "pango_cairo_font_map_new_for_font_type")
	core.PuregoSafeRegister(&xGlyphStringPath, lib, "pango_cairo_glyph_string_path")
	core.PuregoSafeRegister(&xLayoutLinePath, lib, "pango_cairo_layout_line_path")
	core.PuregoSafeRegister(&xLayoutPath, lib, "pango_cairo_layout_path")
	core.PuregoSafeRegister(&xShowErrorUnderline, lib, "pango_cairo_show_error_underline")
	core.PuregoSafeRegister(&xShowGlyphItem, lib, "pango_cairo_show_glyph_item")
	core.PuregoSafeRegister(&xShowGlyphString, lib, "pango_cairo_show_glyph_string")
	core.PuregoSafeRegister(&xShowLayout, lib, "pango_cairo_show_layout")
	core.PuregoSafeRegister(&xShowLayoutLine, lib, "pango_cairo_show_layout_line")
	core.PuregoSafeRegister(&xUpdateContext, lib, "pango_cairo_update_context")
	core.PuregoSafeRegister(&xUpdateLayout, lib, "pango_cairo_update_layout")

	core.PuregoSafeRegister(&xFontGLibType, lib, "pango_cairo_font_get_type")

	core.PuregoSafeRegister(&XPangoCairoFontGetScaledFont, lib, "pango_cairo_font_get_scaled_font")

	core.PuregoSafeRegister(&xFontMapGLibType, lib, "pango_cairo_font_map_get_type")

	core.PuregoSafeRegister(&XPangoCairoFontMapCreateContext, lib, "pango_cairo_font_map_create_context")
	core.PuregoSafeRegister(&XPangoCairoFontMapGetFontType, lib, "pango_cairo_font_map_get_font_type")
	core.PuregoSafeRegister(&XPangoCairoFontMapGetResolution, lib, "pango_cairo_font_map_get_resolution")
	core.PuregoSafeRegister(&XPangoCairoFontMapSetDefault, lib, "pango_cairo_font_map_set_default")
	core.PuregoSafeRegister(&XPangoCairoFontMapSetResolution, lib, "pango_cairo_font_map_set_resolution")

}
