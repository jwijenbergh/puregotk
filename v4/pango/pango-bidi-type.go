// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// `PangoBidiType` represents the bidirectional character
// type of a Unicode character.
//
// The values in this enumeration are specified by the
// [Unicode bidirectional algorithm](http://www.unicode.org/reports/tr9/).
type BidiType int

var xBidiTypeGLibType func() types.GType

func BidiTypeGLibType() types.GType {
	return xBidiTypeGLibType()
}

const (

	// Left-to-Right
	BidiTypeLValue BidiType = 0
	// Left-to-Right Embedding
	BidiTypeLreValue BidiType = 1
	// Left-to-Right Override
	BidiTypeLroValue BidiType = 2
	// Right-to-Left
	BidiTypeRValue BidiType = 3
	// Right-to-Left Arabic
	BidiTypeAlValue BidiType = 4
	// Right-to-Left Embedding
	BidiTypeRleValue BidiType = 5
	// Right-to-Left Override
	BidiTypeRloValue BidiType = 6
	// Pop Directional Format
	BidiTypePdfValue BidiType = 7
	// European Number
	BidiTypeEnValue BidiType = 8
	// European Number Separator
	BidiTypeEsValue BidiType = 9
	// European Number Terminator
	BidiTypeEtValue BidiType = 10
	// Arabic Number
	BidiTypeAnValue BidiType = 11
	// Common Number Separator
	BidiTypeCsValue BidiType = 12
	// Nonspacing Mark
	BidiTypeNsmValue BidiType = 13
	// Boundary Neutral
	BidiTypeBnValue BidiType = 14
	// Paragraph Separator
	BidiTypeBValue BidiType = 15
	// Segment Separator
	BidiTypeSValue BidiType = 16
	// Whitespace
	BidiTypeWsValue BidiType = 17
	// Other Neutrals
	BidiTypeOnValue BidiType = 18
	// Left-to-Right isolate. Since 1.48.6
	BidiTypeLriValue BidiType = 19
	// Right-to-Left isolate. Since 1.48.6
	BidiTypeRliValue BidiType = 20
	// First strong isolate. Since 1.48.6
	BidiTypeFsiValue BidiType = 21
	// Pop directional isolate. Since 1.48.6
	BidiTypePdiValue BidiType = 22
)

var xBidiTypeForUnichar func(uint32) BidiType

// Determines the bidirectional type of a character.
//
// The bidirectional type is specified in the Unicode Character Database.
//
// A simplified version of this function is available as [func@unichar_direction].
func BidiTypeForUnichar(ChVar uint32) BidiType {

	cret := xBidiTypeForUnichar(ChVar)
	return cret
}

var xFindBaseDir func(string, int) Direction

// Searches a string the first character that has a strong
// direction, according to the Unicode bidirectional algorithm.
func FindBaseDir(TextVar string, LengthVar int) Direction {

	cret := xFindBaseDir(TextVar, LengthVar)
	return cret
}

var xGetMirrorChar func(uint32, uint32) bool

// Returns the mirrored character of a Unicode character.
//
// Mirror characters are determined by the Unicode mirrored property.
func GetMirrorChar(ChVar uint32, MirroredChVar uint32) bool {

	cret := xGetMirrorChar(ChVar, MirroredChVar)
	return cret
}

var xUnicharDirection func(uint32) Direction

// Determines the inherent direction of a character.
//
// The inherent direction is either `PANGO_DIRECTION_LTR`, `PANGO_DIRECTION_RTL`,
// or `PANGO_DIRECTION_NEUTRAL`.
//
// This function is useful to categorize characters into left-to-right
// letters, right-to-left letters, and everything else. If full Unicode
// bidirectional type of a character is needed, [func@Pango.BidiType.for_unichar]
// can be used instead.
func UnicharDirection(ChVar uint32) Direction {

	cret := xUnicharDirection(ChVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xBidiTypeGLibType, lib, "pango_bidi_type_get_type")

	core.PuregoSafeRegister(&xBidiTypeForUnichar, lib, "pango_bidi_type_for_unichar")
	core.PuregoSafeRegister(&xFindBaseDir, lib, "pango_find_base_dir")
	core.PuregoSafeRegister(&xGetMirrorChar, lib, "pango_get_mirror_char")
	core.PuregoSafeRegister(&xUnicharDirection, lib, "pango_unichar_direction")

}
