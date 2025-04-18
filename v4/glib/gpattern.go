// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// A GPatternSpec struct is the 'compiled' form of a pattern. This
// structure is opaque and its fields cannot be accessed directly.
type PatternSpec struct {
	_ structs.HostLayout
}

var xPatternSpecGLibType func() types.GType

func PatternSpecGLibType() types.GType {
	return xPatternSpecGLibType()
}

func (x *PatternSpec) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewPatternSpec func(string) *PatternSpec

// Compiles a pattern to a #GPatternSpec.
func NewPatternSpec(PatternVar string) *PatternSpec {

	cret := xNewPatternSpec(PatternVar)
	return cret
}

var xPatternSpecCopy func(uintptr) *PatternSpec

// Copies @pspec in a new #GPatternSpec.
func (x *PatternSpec) Copy() *PatternSpec {

	cret := xPatternSpecCopy(x.GoPointer())
	return cret
}

var xPatternSpecEqual func(uintptr, *PatternSpec) bool

// Compares two compiled pattern specs and returns whether they will
// match the same set of strings.
func (x *PatternSpec) Equal(Pspec2Var *PatternSpec) bool {

	cret := xPatternSpecEqual(x.GoPointer(), Pspec2Var)
	return cret
}

var xPatternSpecFree func(uintptr)

// Frees the memory allocated for the #GPatternSpec.
func (x *PatternSpec) Free() {

	xPatternSpecFree(x.GoPointer())

}

var xPatternSpecMatch func(uintptr, uint, string, string) bool

// Matches a string against a compiled pattern. Passing the correct
// length of the string given is mandatory. The reversed string can be
// omitted by passing %NULL, this is more efficient if the reversed
// version of the string to be matched is not at hand, as
// g_pattern_match() will only construct it if the compiled pattern
// requires reverse matches.
//
// Note that, if the user code will (possibly) match a string against a
// multitude of patterns containing wildcards, chances are high that
// some patterns will require a reversed string. In this case, it's
// more efficient to provide the reversed string to avoid multiple
// constructions thereof in the various calls to g_pattern_match().
//
// Note also that the reverse of a UTF-8 encoded string can in general
// not be obtained by g_strreverse(). This works only if the string
// does not contain any multibyte characters. GLib offers the
// g_utf8_strreverse() function to reverse UTF-8 encoded strings.
func (x *PatternSpec) Match(StringLengthVar uint, StringVar string, StringReversedVar string) bool {

	cret := xPatternSpecMatch(x.GoPointer(), StringLengthVar, StringVar, StringReversedVar)
	return cret
}

var xPatternSpecMatchString func(uintptr, string) bool

// Matches a string against a compiled pattern. If the string is to be
// matched against more than one pattern, consider using
// g_pattern_match() instead while supplying the reversed string.
func (x *PatternSpec) MatchString(StringVar string) bool {

	cret := xPatternSpecMatchString(x.GoPointer(), StringVar)
	return cret
}

var xPatternMatch func(*PatternSpec, uint, string, string) bool

// Matches a string against a compiled pattern. Passing the correct
// length of the string given is mandatory. The reversed string can be
// omitted by passing %NULL, this is more efficient if the reversed
// version of the string to be matched is not at hand, as
// g_pattern_match() will only construct it if the compiled pattern
// requires reverse matches.
//
// Note that, if the user code will (possibly) match a string against a
// multitude of patterns containing wildcards, chances are high that
// some patterns will require a reversed string. In this case, it's
// more efficient to provide the reversed string to avoid multiple
// constructions thereof in the various calls to g_pattern_match().
//
// Note also that the reverse of a UTF-8 encoded string can in general
// not be obtained by g_strreverse(). This works only if the string
// does not contain any multibyte characters. GLib offers the
// g_utf8_strreverse() function to reverse UTF-8 encoded strings.
func PatternMatch(PspecVar *PatternSpec, StringLengthVar uint, StringVar string, StringReversedVar string) bool {

	cret := xPatternMatch(PspecVar, StringLengthVar, StringVar, StringReversedVar)
	return cret
}

var xPatternMatchSimple func(string, string) bool

// Matches a string against a pattern given as a string. If this
// function is to be called in a loop, it's more efficient to compile
// the pattern once with g_pattern_spec_new() and call
// g_pattern_match_string() repeatedly.
func PatternMatchSimple(PatternVar string, StringVar string) bool {

	cret := xPatternMatchSimple(PatternVar, StringVar)
	return cret
}

var xPatternMatchString func(*PatternSpec, string) bool

// Matches a string against a compiled pattern. If the string is to be
// matched against more than one pattern, consider using
// g_pattern_match() instead while supplying the reversed string.
func PatternMatchString(PspecVar *PatternSpec, StringVar string) bool {

	cret := xPatternMatchString(PspecVar, StringVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xPatternMatch, lib, "g_pattern_match")
	core.PuregoSafeRegister(&xPatternMatchSimple, lib, "g_pattern_match_simple")
	core.PuregoSafeRegister(&xPatternMatchString, lib, "g_pattern_match_string")

	core.PuregoSafeRegister(&xPatternSpecGLibType, lib, "g_pattern_spec_get_type")

	core.PuregoSafeRegister(&xNewPatternSpec, lib, "g_pattern_spec_new")

	core.PuregoSafeRegister(&xPatternSpecCopy, lib, "g_pattern_spec_copy")
	core.PuregoSafeRegister(&xPatternSpecEqual, lib, "g_pattern_spec_equal")
	core.PuregoSafeRegister(&xPatternSpecFree, lib, "g_pattern_spec_free")
	core.PuregoSafeRegister(&xPatternSpecMatch, lib, "g_pattern_spec_match")
	core.PuregoSafeRegister(&xPatternSpecMatchString, lib, "g_pattern_spec_match_string")

}
