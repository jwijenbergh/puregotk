// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

// The `GtkFontChooserWidget` widget lets the user select a font.
//
// It is used in the `GtkFontChooserDialog` widget to provide a
// dialog for selecting fonts.
//
// To set the font which is initially selected, use
// [method@Gtk.FontChooser.set_font] or [method@Gtk.FontChooser.set_font_desc].
//
// To get the selected font use [method@Gtk.FontChooser.get_font] or
// [method@Gtk.FontChooser.get_font_desc].
//
// To change the text which is shown in the preview area, use
// [method@Gtk.FontChooser.set_preview_text].
//
// # CSS nodes
//
// `GtkFontChooserWidget` has a single CSS node with name fontchooser.
type FontChooserWidget struct {
	Widget
}

var xFontChooserWidgetGLibType func() types.GType

func FontChooserWidgetGLibType() types.GType {
	return xFontChooserWidgetGLibType()
}

func FontChooserWidgetNewFromInternalPtr(ptr uintptr) *FontChooserWidget {
	cls := &FontChooserWidget{}
	cls.Ptr = ptr
	return cls
}

var xNewFontChooserWidget func() uintptr

// Creates a new `GtkFontChooserWidget`.
func NewFontChooserWidget() *FontChooserWidget {
	var cls *FontChooserWidget

	cret := xNewFontChooserWidget()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &FontChooserWidget{}
	cls.Ptr = cret
	return cls
}

func (c *FontChooserWidget) GoPointer() uintptr {
	return c.Ptr
}

func (c *FontChooserWidget) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *FontChooserWidget) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *FontChooserWidget) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *FontChooserWidget) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *FontChooserWidget) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *FontChooserWidget) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FontChooserWidget) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *FontChooserWidget) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FontChooserWidget) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *FontChooserWidget) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *FontChooserWidget) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *FontChooserWidget) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// [method@Gtk.FontChooser.set_font], as the font chooser widget may
// normalize font names and thus return a string with a different
// structure. For example, “Helvetica Italic Bold 12” could be
// normalized to “Helvetica Bold Italic 12”.
//
// Use [method@Pango.FontDescription.equal] if you want to compare two
// font descriptions.
func (x *FontChooserWidget) GetFont() string {

	cret := XGtkFontChooserGetFont(x.GoPointer())
	return cret
}

// Gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// [method@Gtk.FontChooser.set_font], as the font chooser widget may
// normalize font names and thus return a string with a different
// structure. For example, “Helvetica Italic Bold 12” could be
// normalized to “Helvetica Bold Italic 12”.
//
// Use [method@Pango.FontDescription.equal] if you want to compare two
// font descriptions.
func (x *FontChooserWidget) GetFontDesc() *pango.FontDescription {

	cret := XGtkFontChooserGetFontDesc(x.GoPointer())
	return cret
}

// Gets the `PangoFontFace` representing the selected font group
// details (i.e. family, slant, weight, width, etc).
//
// If the selected font is not installed, returns %NULL.
func (x *FontChooserWidget) GetFontFace() *pango.FontFace {
	var cls *pango.FontFace

	cret := XGtkFontChooserGetFontFace(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &pango.FontFace{}
	cls.Ptr = cret
	return cls
}

// Gets the `PangoFontFamily` representing the selected font family.
//
// Font families are a collection of font faces.
//
// If the selected font is not installed, returns %NULL.
func (x *FontChooserWidget) GetFontFamily() *pango.FontFamily {
	var cls *pango.FontFamily

	cret := XGtkFontChooserGetFontFamily(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &pango.FontFamily{}
	cls.Ptr = cret
	return cls
}

// Gets the currently-selected font features.
//
// The format of the returned string is compatible with the
// [CSS font-feature-settings property](https://www.w3.org/TR/css-fonts-4/#font-rend-desc).
// It can be passed to [func@Pango.AttrFontFeatures.new].
func (x *FontChooserWidget) GetFontFeatures() string {

	cret := XGtkFontChooserGetFontFeatures(x.GoPointer())
	return cret
}

// Gets the custom font map of this font chooser widget,
// or %NULL if it does not have one.
func (x *FontChooserWidget) GetFontMap() *pango.FontMap {
	var cls *pango.FontMap

	cret := XGtkFontChooserGetFontMap(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &pango.FontMap{}
	cls.Ptr = cret
	return cls
}

// The selected font size.
func (x *FontChooserWidget) GetFontSize() int {

	cret := XGtkFontChooserGetFontSize(x.GoPointer())
	return cret
}

// Gets the language that is used for font features.
func (x *FontChooserWidget) GetLanguage() string {

	cret := XGtkFontChooserGetLanguage(x.GoPointer())
	return cret
}

// Returns the current level of granularity for selecting fonts.
func (x *FontChooserWidget) GetLevel() FontChooserLevel {

	cret := XGtkFontChooserGetLevel(x.GoPointer())
	return cret
}

// Gets the text displayed in the preview area.
func (x *FontChooserWidget) GetPreviewText() string {

	cret := XGtkFontChooserGetPreviewText(x.GoPointer())
	return cret
}

// Returns whether the preview entry is shown or not.
func (x *FontChooserWidget) GetShowPreviewEntry() bool {

	cret := XGtkFontChooserGetShowPreviewEntry(x.GoPointer())
	return cret
}

// Adds a filter function that decides which fonts to display
// in the font chooser.
func (x *FontChooserWidget) SetFilterFunc(FilterVar *FontFilterFunc, UserDataVar uintptr, DestroyVar *glib.DestroyNotify) {

	XGtkFontChooserSetFilterFunc(x.GoPointer(), glib.NewCallback(FilterVar), UserDataVar, glib.NewCallback(DestroyVar))

}

// Sets the currently-selected font.
func (x *FontChooserWidget) SetFont(FontnameVar string) {

	XGtkFontChooserSetFont(x.GoPointer(), FontnameVar)

}

// Sets the currently-selected font from @font_desc.
func (x *FontChooserWidget) SetFontDesc(FontDescVar *pango.FontDescription) {

	XGtkFontChooserSetFontDesc(x.GoPointer(), FontDescVar)

}

// Sets a custom font map to use for this font chooser widget.
//
// A custom font map can be used to present application-specific
// fonts instead of or in addition to the normal system fonts.
//
// ```c
// FcConfig *config;
// PangoFontMap *fontmap;
//
// config = FcInitLoadConfigAndFonts ();
// FcConfigAppFontAddFile (config, my_app_font_file);
//
// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
// gtk_font_chooser_set_font_map (font_chooser, fontmap);
// ```
//
// Note that other GTK widgets will only be able to use the
// application-specific font if it is present in the font map they use:
//
// ```c
// context = gtk_widget_get_pango_context (label);
// pango_context_set_font_map (context, fontmap);
// ```
func (x *FontChooserWidget) SetFontMap(FontmapVar *pango.FontMap) {

	XGtkFontChooserSetFontMap(x.GoPointer(), FontmapVar.GoPointer())

}

// Sets the language to use for font features.
func (x *FontChooserWidget) SetLanguage(LanguageVar string) {

	XGtkFontChooserSetLanguage(x.GoPointer(), LanguageVar)

}

// Sets the desired level of granularity for selecting fonts.
func (x *FontChooserWidget) SetLevel(LevelVar FontChooserLevel) {

	XGtkFontChooserSetLevel(x.GoPointer(), LevelVar)

}

// Sets the text displayed in the preview area.
//
// The @text is used to show how the selected font looks.
func (x *FontChooserWidget) SetPreviewText(TextVar string) {

	XGtkFontChooserSetPreviewText(x.GoPointer(), TextVar)

}

// Shows or hides the editable preview entry.
func (x *FontChooserWidget) SetShowPreviewEntry(ShowPreviewEntryVar bool) {

	XGtkFontChooserSetShowPreviewEntry(x.GoPointer(), ShowPreviewEntryVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFontChooserWidgetGLibType, lib, "gtk_font_chooser_widget_get_type")

	core.PuregoSafeRegister(&xNewFontChooserWidget, lib, "gtk_font_chooser_widget_new")

}
