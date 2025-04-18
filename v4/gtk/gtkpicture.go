// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gdkpixbuf"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

type PictureClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *PictureClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// The `GtkPicture` widget displays a `GdkPaintable`.
//
// ![An example GtkPicture](picture.png)
//
// Many convenience functions are provided to make pictures simple to use.
// For example, if you want to load an image from a file, and then display
// it, there’s a convenience function to do this:
//
// ```c
// GtkWidget *widget = gtk_picture_new_for_filename ("myfile.png");
// ```
//
// If the file isn’t loaded successfully, the picture will contain a
// “broken image” icon similar to that used in many web browsers.
// If you want to handle errors in loading the file yourself,
// for example by displaying an error message, then load the image with
// [ctor@Gdk.Texture.new_from_file], then create the `GtkPicture` with
// [ctor@Gtk.Picture.new_for_paintable].
//
// Sometimes an application will want to avoid depending on external data
// files, such as image files. See the documentation of `GResource` for details.
// In this case, [ctor@Gtk.Picture.new_for_resource] and
// [method@Gtk.Picture.set_resource] should be used.
//
// `GtkPicture` displays an image at its natural size. See [class@Gtk.Image]
// if you want to display a fixed-size image, such as an icon.
//
// ## Sizing the paintable
//
// You can influence how the paintable is displayed inside the `GtkPicture`
// by changing [property@Gtk.Picture:content-fit]. See [enum@Gtk.ContentFit]
// for details. [property@Gtk.Picture:can-shrink] can be unset to make sure
// that paintables are never made smaller than their ideal size - but
// be careful if you do not know the size of the paintable in use (like
// when displaying user-loaded images). This can easily cause the picture to
// grow larger than the screen. And [property@GtkWidget:halign] and
// [property@GtkWidget:valign] can be used to make sure the paintable doesn't
// fill all available space but is instead displayed at its original size.
//
// ## CSS nodes
//
// `GtkPicture` has a single CSS node with the name `picture`.
//
// ## Accessibility
//
// `GtkPicture` uses the `GTK_ACCESSIBLE_ROLE_IMG` role.
type Picture struct {
	Widget
}

var xPictureGLibType func() types.GType

func PictureGLibType() types.GType {
	return xPictureGLibType()
}

func PictureNewFromInternalPtr(ptr uintptr) *Picture {
	cls := &Picture{}
	cls.Ptr = ptr
	return cls
}

var xNewPicture func() uintptr

// Creates a new empty `GtkPicture` widget.
func NewPicture() *Picture {
	var cls *Picture

	cret := xNewPicture()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Picture{}
	cls.Ptr = cret
	return cls
}

var xNewPictureForFile func(uintptr) uintptr

// Creates a new `GtkPicture` displaying the given @file.
//
// If the file isn’t found or can’t be loaded, the resulting
// `GtkPicture` is empty.
//
// If you need to detect failures to load the file, use
// [ctor@Gdk.Texture.new_from_file] to load the file yourself,
// then create the `GtkPicture` from the texture.
func NewPictureForFile(FileVar gio.File) *Picture {
	var cls *Picture

	cret := xNewPictureForFile(FileVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Picture{}
	cls.Ptr = cret
	return cls
}

var xNewPictureForFilename func(string) uintptr

// Creates a new `GtkPicture` displaying the file @filename.
//
// This is a utility function that calls [ctor@Gtk.Picture.new_for_file].
// See that function for details.
func NewPictureForFilename(FilenameVar string) *Picture {
	var cls *Picture

	cret := xNewPictureForFilename(FilenameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Picture{}
	cls.Ptr = cret
	return cls
}

var xNewPictureForPaintable func(uintptr) uintptr

// Creates a new `GtkPicture` displaying @paintable.
//
// The `GtkPicture` will track changes to the @paintable and update
// its size and contents in response to it.
func NewPictureForPaintable(PaintableVar gdk.Paintable) *Picture {
	var cls *Picture

	cret := xNewPictureForPaintable(PaintableVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Picture{}
	cls.Ptr = cret
	return cls
}

var xNewPictureForPixbuf func(uintptr) uintptr

// Creates a new `GtkPicture` displaying @pixbuf.
//
// This is a utility function that calls [ctor@Gtk.Picture.new_for_paintable],
// See that function for details.
//
// The pixbuf must not be modified after passing it to this function.
func NewPictureForPixbuf(PixbufVar *gdkpixbuf.Pixbuf) *Picture {
	var cls *Picture

	cret := xNewPictureForPixbuf(PixbufVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Picture{}
	cls.Ptr = cret
	return cls
}

var xNewPictureForResource func(string) uintptr

// Creates a new `GtkPicture` displaying the resource at @resource_path.
//
// This is a utility function that calls [ctor@Gtk.Picture.new_for_file].
// See that function for details.
func NewPictureForResource(ResourcePathVar string) *Picture {
	var cls *Picture

	cret := xNewPictureForResource(ResourcePathVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Picture{}
	cls.Ptr = cret
	return cls
}

var xPictureGetAlternativeText func(uintptr) string

// Gets the alternative textual description of the picture.
//
// The returned string will be %NULL if the picture cannot be described textually.
func (x *Picture) GetAlternativeText() string {

	cret := xPictureGetAlternativeText(x.GoPointer())
	return cret
}

var xPictureGetCanShrink func(uintptr) bool

// Returns whether the `GtkPicture` respects its contents size.
func (x *Picture) GetCanShrink() bool {

	cret := xPictureGetCanShrink(x.GoPointer())
	return cret
}

var xPictureGetContentFit func(uintptr) ContentFit

// Returns the fit mode for the content of the `GtkPicture`.
//
// See [enum@Gtk.ContentFit] for details.
func (x *Picture) GetContentFit() ContentFit {

	cret := xPictureGetContentFit(x.GoPointer())
	return cret
}

var xPictureGetFile func(uintptr) uintptr

// Gets the `GFile` currently displayed if @self is displaying a file.
//
// If @self is not displaying a file, for example when
// [method@Gtk.Picture.set_paintable] was used, then %NULL is returned.
func (x *Picture) GetFile() *gio.FileBase {
	var cls *gio.FileBase

	cret := xPictureGetFile(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.FileBase{}
	cls.Ptr = cret
	return cls
}

var xPictureGetKeepAspectRatio func(uintptr) bool

// Returns whether the `GtkPicture` preserves its contents aspect ratio.
func (x *Picture) GetKeepAspectRatio() bool {

	cret := xPictureGetKeepAspectRatio(x.GoPointer())
	return cret
}

var xPictureGetPaintable func(uintptr) uintptr

// Gets the `GdkPaintable` being displayed by the `GtkPicture`.
func (x *Picture) GetPaintable() *gdk.PaintableBase {
	var cls *gdk.PaintableBase

	cret := xPictureGetPaintable(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.PaintableBase{}
	cls.Ptr = cret
	return cls
}

var xPictureSetAlternativeText func(uintptr, string)

// Sets an alternative textual description for the picture contents.
//
// It is equivalent to the "alt" attribute for images on websites.
//
// This text will be made available to accessibility tools.
//
// If the picture cannot be described textually, set this property to %NULL.
func (x *Picture) SetAlternativeText(AlternativeTextVar string) {

	xPictureSetAlternativeText(x.GoPointer(), AlternativeTextVar)

}

var xPictureSetCanShrink func(uintptr, bool)

// If set to %TRUE, the @self can be made smaller than its contents.
//
// The contents will then be scaled down when rendering.
//
// If you want to still force a minimum size manually, consider using
// [method@Gtk.Widget.set_size_request].
//
// Also of note is that a similar function for growing does not exist
// because the grow behavior can be controlled via
// [method@Gtk.Widget.set_halign] and [method@Gtk.Widget.set_valign].
func (x *Picture) SetCanShrink(CanShrinkVar bool) {

	xPictureSetCanShrink(x.GoPointer(), CanShrinkVar)

}

var xPictureSetContentFit func(uintptr, ContentFit)

// Sets how the content should be resized to fit the `GtkPicture`.
//
// See [enum@Gtk.ContentFit] for details.
func (x *Picture) SetContentFit(ContentFitVar ContentFit) {

	xPictureSetContentFit(x.GoPointer(), ContentFitVar)

}

var xPictureSetFile func(uintptr, uintptr)

// Makes @self load and display @file.
//
// See [ctor@Gtk.Picture.new_for_file] for details.
func (x *Picture) SetFile(FileVar gio.File) {

	xPictureSetFile(x.GoPointer(), FileVar.GoPointer())

}

var xPictureSetFilename func(uintptr, string)

// Makes @self load and display the given @filename.
//
// This is a utility function that calls [method@Gtk.Picture.set_file].
func (x *Picture) SetFilename(FilenameVar string) {

	xPictureSetFilename(x.GoPointer(), FilenameVar)

}

var xPictureSetKeepAspectRatio func(uintptr, bool)

// If set to %TRUE, the @self will render its contents according to
// their aspect ratio.
//
// That means that empty space may show up at the top/bottom or
// left/right of @self.
//
// If set to %FALSE or if the contents provide no aspect ratio,
// the contents will be stretched over the picture's whole area.
func (x *Picture) SetKeepAspectRatio(KeepAspectRatioVar bool) {

	xPictureSetKeepAspectRatio(x.GoPointer(), KeepAspectRatioVar)

}

var xPictureSetPaintable func(uintptr, uintptr)

// Makes @self display the given @paintable.
//
// If @paintable is %NULL, nothing will be displayed.
//
// See [ctor@Gtk.Picture.new_for_paintable] for details.
func (x *Picture) SetPaintable(PaintableVar gdk.Paintable) {

	xPictureSetPaintable(x.GoPointer(), PaintableVar.GoPointer())

}

var xPictureSetPixbuf func(uintptr, uintptr)

// Sets a `GtkPicture` to show a `GdkPixbuf`.
//
// See [ctor@Gtk.Picture.new_for_pixbuf] for details.
//
// This is a utility function that calls [method@Gtk.Picture.set_paintable].
func (x *Picture) SetPixbuf(PixbufVar *gdkpixbuf.Pixbuf) {

	xPictureSetPixbuf(x.GoPointer(), PixbufVar.GoPointer())

}

var xPictureSetResource func(uintptr, string)

// Makes @self load and display the resource at the given
// @resource_path.
//
// This is a utility function that calls [method@Gtk.Picture.set_file].
func (x *Picture) SetResource(ResourcePathVar string) {

	xPictureSetResource(x.GoPointer(), ResourcePathVar)

}

func (c *Picture) GoPointer() uintptr {
	return c.Ptr
}

func (c *Picture) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Picture) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Picture) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Picture) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Picture) ResetState(StateVar AccessibleState) {

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
func (x *Picture) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Picture) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *Picture) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Picture) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *Picture) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Picture) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Picture) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xPictureGLibType, lib, "gtk_picture_get_type")

	core.PuregoSafeRegister(&xNewPicture, lib, "gtk_picture_new")
	core.PuregoSafeRegister(&xNewPictureForFile, lib, "gtk_picture_new_for_file")
	core.PuregoSafeRegister(&xNewPictureForFilename, lib, "gtk_picture_new_for_filename")
	core.PuregoSafeRegister(&xNewPictureForPaintable, lib, "gtk_picture_new_for_paintable")
	core.PuregoSafeRegister(&xNewPictureForPixbuf, lib, "gtk_picture_new_for_pixbuf")
	core.PuregoSafeRegister(&xNewPictureForResource, lib, "gtk_picture_new_for_resource")

	core.PuregoSafeRegister(&xPictureGetAlternativeText, lib, "gtk_picture_get_alternative_text")
	core.PuregoSafeRegister(&xPictureGetCanShrink, lib, "gtk_picture_get_can_shrink")
	core.PuregoSafeRegister(&xPictureGetContentFit, lib, "gtk_picture_get_content_fit")
	core.PuregoSafeRegister(&xPictureGetFile, lib, "gtk_picture_get_file")
	core.PuregoSafeRegister(&xPictureGetKeepAspectRatio, lib, "gtk_picture_get_keep_aspect_ratio")
	core.PuregoSafeRegister(&xPictureGetPaintable, lib, "gtk_picture_get_paintable")
	core.PuregoSafeRegister(&xPictureSetAlternativeText, lib, "gtk_picture_set_alternative_text")
	core.PuregoSafeRegister(&xPictureSetCanShrink, lib, "gtk_picture_set_can_shrink")
	core.PuregoSafeRegister(&xPictureSetContentFit, lib, "gtk_picture_set_content_fit")
	core.PuregoSafeRegister(&xPictureSetFile, lib, "gtk_picture_set_file")
	core.PuregoSafeRegister(&xPictureSetFilename, lib, "gtk_picture_set_filename")
	core.PuregoSafeRegister(&xPictureSetKeepAspectRatio, lib, "gtk_picture_set_keep_aspect_ratio")
	core.PuregoSafeRegister(&xPictureSetPaintable, lib, "gtk_picture_set_paintable")
	core.PuregoSafeRegister(&xPictureSetPixbuf, lib, "gtk_picture_set_pixbuf")
	core.PuregoSafeRegister(&xPictureSetResource, lib, "gtk_picture_set_resource")

}
