// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// `GtkFileFilter` filters files by name or mime type.
//
// `GtkFileFilter` can be used to restrict the files being shown in a
// `GtkFileChooser`. Files can be filtered based on their name (with
// [method@Gtk.FileFilter.add_pattern] or [method@Gtk.FileFilter.add_suffix])
// or on their mime type (with [method@Gtk.FileFilter.add_mime_type]).
//
// Filtering by mime types handles aliasing and subclassing of mime
// types; e.g. a filter for text/plain also matches a file with mime
// type application/rtf, since application/rtf is a subclass of
// text/plain. Note that `GtkFileFilter` allows wildcards for the
// subtype of a mime type, so you can e.g. filter for image/\*.
//
// Normally, file filters are used by adding them to a `GtkFileChooser`
// (see [method@Gtk.FileChooser.add_filter]), but it is also possible to
// manually use a file filter on any [class@Gtk.FilterListModel] containing
// `GFileInfo` objects.
//
// # GtkFileFilter as GtkBuildable
//
// The `GtkFileFilter` implementation of the `GtkBuildable` interface
// supports adding rules using the `&lt;mime-types&gt;` and `&lt;patterns&gt;` and
// `&lt;suffixes&gt;` elements and listing the rules within. Specifying a
// `&lt;mime-type&gt;` or `&lt;pattern&gt;` or `&lt;suffix&gt;` has the same effect as
// as calling
// [method@Gtk.FileFilter.add_mime_type] or
// [method@Gtk.FileFilter.add_pattern] or
// [method@Gtk.FileFilter.add_suffix].
//
// An example of a UI definition fragment specifying `GtkFileFilter`
// rules:
// ```xml
// &lt;object class="GtkFileFilter"&gt;
//
//	&lt;property name="name" translatable="yes"&gt;Text and Images&lt;/property&gt;
//	&lt;mime-types&gt;
//	  &lt;mime-type&gt;text/plain&lt;/mime-type&gt;
//	  &lt;mime-type&gt;image/ *&lt;/mime-type&gt;
//	&lt;/mime-types&gt;
//	&lt;patterns&gt;
//	  &lt;pattern&gt;*.txt&lt;/pattern&gt;
//	&lt;/patterns&gt;
//	&lt;suffixes&gt;
//	  &lt;suffix&gt;png&lt;/suffix&gt;
//	&lt;/suffixes&gt;
//
// &lt;/object&gt;
// ```
type FileFilter struct {
	Filter
}

var xFileFilterGLibType func() types.GType

func FileFilterGLibType() types.GType {
	return xFileFilterGLibType()
}

func FileFilterNewFromInternalPtr(ptr uintptr) *FileFilter {
	cls := &FileFilter{}
	cls.Ptr = ptr
	return cls
}

var xNewFileFilter func() uintptr

// Creates a new `GtkFileFilter` with no rules added to it.
//
// Such a filter doesn’t accept any files, so is not
// particularly useful until you add rules with
// [method@Gtk.FileFilter.add_mime_type],
// [method@Gtk.FileFilter.add_pattern],
// [method@Gtk.FileFilter.add_suffix] or
// [method@Gtk.FileFilter.add_pixbuf_formats].
//
// To create a filter that accepts any file, use:
// ```c
// GtkFileFilter *filter = gtk_file_filter_new ();
// gtk_file_filter_add_pattern (filter, "*");
// ```
func NewFileFilter() *FileFilter {
	var cls *FileFilter

	cret := xNewFileFilter()

	if cret == 0 {
		return nil
	}
	cls = &FileFilter{}
	cls.Ptr = cret
	return cls
}

var xNewFileFilterFromGvariant func(*glib.Variant) uintptr

// Deserialize a file filter from a `GVariant`.
//
// The variant must be in the format produced by
// [method@Gtk.FileFilter.to_gvariant].
func NewFileFilterFromGvariant(VariantVar *glib.Variant) *FileFilter {
	var cls *FileFilter

	cret := xNewFileFilterFromGvariant(VariantVar)

	if cret == 0 {
		return nil
	}
	cls = &FileFilter{}
	cls.Ptr = cret
	return cls
}

var xFileFilterAddMimeType func(uintptr, string)

// Adds a rule allowing a given mime type to @filter.
func (x *FileFilter) AddMimeType(MimeTypeVar string) {

	xFileFilterAddMimeType(x.GoPointer(), MimeTypeVar)

}

var xFileFilterAddPattern func(uintptr, string)

// Adds a rule allowing a shell style glob to a filter.
//
// Note that it depends on the platform whether pattern
// matching ignores case or not. On Windows, it does, on
// other platforms, it doesn't.
func (x *FileFilter) AddPattern(PatternVar string) {

	xFileFilterAddPattern(x.GoPointer(), PatternVar)

}

var xFileFilterAddPixbufFormats func(uintptr)

// Adds a rule allowing image files in the formats supported
// by GdkPixbuf.
//
// This is equivalent to calling [method@Gtk.FileFilter.add_mime_type]
// for all the supported mime types.
func (x *FileFilter) AddPixbufFormats() {

	xFileFilterAddPixbufFormats(x.GoPointer())

}

var xFileFilterAddSuffix func(uintptr, string)

// Adds a suffix match rule to a filter.
//
// This is similar to adding a match for the pattern
// "*.@suffix".
//
// In contrast to pattern matches, suffix matches
// are *always* case-insensitive.
func (x *FileFilter) AddSuffix(SuffixVar string) {

	xFileFilterAddSuffix(x.GoPointer(), SuffixVar)

}

var xFileFilterGetAttributes func(uintptr) []string

// Gets the attributes that need to be filled in for the `GFileInfo`
// passed to this filter.
//
// This function will not typically be used by applications;
// it is intended principally for use in the implementation
// of `GtkFileChooser`.
func (x *FileFilter) GetAttributes() []string {

	cret := xFileFilterGetAttributes(x.GoPointer())
	return cret
}

var xFileFilterGetName func(uintptr) string

// Gets the human-readable name for the filter.
//
// See [method@Gtk.FileFilter.set_name].
func (x *FileFilter) GetName() string {

	cret := xFileFilterGetName(x.GoPointer())
	return cret
}

var xFileFilterSetName func(uintptr, string)

// Sets a human-readable name of the filter.
//
// This is the string that will be displayed in the file chooser
// if there is a selectable list of filters.
func (x *FileFilter) SetName(NameVar string) {

	xFileFilterSetName(x.GoPointer(), NameVar)

}

var xFileFilterToGvariant func(uintptr) *glib.Variant

// Serialize a file filter to an `a{sv}` variant.
func (x *FileFilter) ToGvariant() *glib.Variant {

	cret := xFileFilterToGvariant(x.GoPointer())
	return cret
}

func (c *FileFilter) GoPointer() uintptr {
	return c.Ptr
}

func (c *FileFilter) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *FileFilter) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFileFilterGLibType, lib, "gtk_file_filter_get_type")

	core.PuregoSafeRegister(&xNewFileFilter, lib, "gtk_file_filter_new")
	core.PuregoSafeRegister(&xNewFileFilterFromGvariant, lib, "gtk_file_filter_new_from_gvariant")

	core.PuregoSafeRegister(&xFileFilterAddMimeType, lib, "gtk_file_filter_add_mime_type")
	core.PuregoSafeRegister(&xFileFilterAddPattern, lib, "gtk_file_filter_add_pattern")
	core.PuregoSafeRegister(&xFileFilterAddPixbufFormats, lib, "gtk_file_filter_add_pixbuf_formats")
	core.PuregoSafeRegister(&xFileFilterAddSuffix, lib, "gtk_file_filter_add_suffix")
	core.PuregoSafeRegister(&xFileFilterGetAttributes, lib, "gtk_file_filter_get_attributes")
	core.PuregoSafeRegister(&xFileFilterGetName, lib, "gtk_file_filter_get_name")
	core.PuregoSafeRegister(&xFileFilterSetName, lib, "gtk_file_filter_set_name")
	core.PuregoSafeRegister(&xFileFilterToGvariant, lib, "gtk_file_filter_to_gvariant")

}
