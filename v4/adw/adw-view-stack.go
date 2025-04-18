// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ViewStackClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *ViewStackClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type ViewStackPageClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *ViewStackPageClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type ViewStackPagesClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *ViewStackPagesClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// A view container for [class@ViewSwitcher].
//
// `AdwViewStack` is a container which only shows one page at a time.
// It is typically used to hold an application's main views.
//
// It doesn't provide a way to transition between pages.
// Instead, a separate widget such as [class@ViewSwitcher] can be used with
// `AdwViewStack` to provide this functionality.
//
// `AdwViewStack` pages can have a title, an icon, an attention request, and a
// numbered badge that [class@ViewSwitcher] will use to let users identify which
// page is which. Set them using the [property@ViewStackPage:title],
// [property@ViewStackPage:icon-name],
// [property@ViewStackPage:needs-attention], and
// [property@ViewStackPage:badge-number] properties.
//
// Unlike [class@Gtk.Stack], transitions between views are not animated.
//
// `AdwViewStack` maintains a [class@ViewStackPage] object for each added child,
// which holds additional per-child properties. You obtain the
// [class@ViewStackPage] for a child with [method@ViewStack.get_page] and you
// can obtain a [iface@Gtk.SelectionModel] containing all the pages with
// [method@ViewStack.get_pages].
//
// ## AdwViewStack as GtkBuildable
//
// To set child-specific properties in a .ui file, create
// [class@ViewStackPage] objects explicitly, and set the child widget as a
// property on it:
//
// ```xml
//
//	&lt;object class="AdwViewStack" id="stack"&gt;
//	  &lt;child&gt;
//	    &lt;object class="AdwViewStackPage"&gt;
//	      &lt;property name="name"&gt;overview&lt;/property&gt;
//	      &lt;property name="title"&gt;Overview&lt;/property&gt;
//	      &lt;property name="child"&gt;
//	        &lt;object class="AdwStatusPage"&gt;
//	          &lt;property name="title"&gt;Welcome!&lt;/property&gt;
//	        &lt;/object&gt;
//	      &lt;/property&gt;
//	    &lt;/object&gt;
//	  &lt;/child&gt;
//	&lt;/object&gt;
//
// ```
//
// ## CSS nodes
//
// `AdwViewStack` has a single CSS node named `stack`.
//
// ## Accessibility
//
// `AdwViewStack` uses the `GTK_ACCESSIBLE_ROLE_TAB_PANEL` for the stack pages
// which are the accessible parent objects of the child widgets.
type ViewStack struct {
	gtk.Widget
}

var xViewStackGLibType func() types.GType

func ViewStackGLibType() types.GType {
	return xViewStackGLibType()
}

func ViewStackNewFromInternalPtr(ptr uintptr) *ViewStack {
	cls := &ViewStack{}
	cls.Ptr = ptr
	return cls
}

var xNewViewStack func() uintptr

// Creates a new `AdwViewStack`.
func NewViewStack() *ViewStack {
	var cls *ViewStack

	cret := xNewViewStack()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStack{}
	cls.Ptr = cret
	return cls
}

var xViewStackAdd func(uintptr, uintptr) uintptr

// Adds a child to @self.
func (x *ViewStack) Add(ChildVar *gtk.Widget) *ViewStackPage {
	var cls *ViewStackPage

	cret := xViewStackAdd(x.GoPointer(), ChildVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStackPage{}
	cls.Ptr = cret
	return cls
}

var xViewStackAddNamed func(uintptr, uintptr, string) uintptr

// Adds a child to @self.
//
// The child is identified by the @name.
func (x *ViewStack) AddNamed(ChildVar *gtk.Widget, NameVar string) *ViewStackPage {
	var cls *ViewStackPage

	cret := xViewStackAddNamed(x.GoPointer(), ChildVar.GoPointer(), NameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStackPage{}
	cls.Ptr = cret
	return cls
}

var xViewStackAddTitled func(uintptr, uintptr, string, string) uintptr

// Adds a child to @self.
//
// The child is identified by the @name. The @title will be used by
// [class@ViewSwitcher] to represent @child, so it should be short.
func (x *ViewStack) AddTitled(ChildVar *gtk.Widget, NameVar string, TitleVar string) *ViewStackPage {
	var cls *ViewStackPage

	cret := xViewStackAddTitled(x.GoPointer(), ChildVar.GoPointer(), NameVar, TitleVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStackPage{}
	cls.Ptr = cret
	return cls
}

var xViewStackAddTitledWithIcon func(uintptr, uintptr, string, string, string) uintptr

// Adds a child to @self.
//
// The child is identified by the @name. The @title and @icon_name will be used
// by [class@ViewSwitcher] to represent @child.
func (x *ViewStack) AddTitledWithIcon(ChildVar *gtk.Widget, NameVar string, TitleVar string, IconNameVar string) *ViewStackPage {
	var cls *ViewStackPage

	cret := xViewStackAddTitledWithIcon(x.GoPointer(), ChildVar.GoPointer(), NameVar, TitleVar, IconNameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStackPage{}
	cls.Ptr = cret
	return cls
}

var xViewStackGetChildByName func(uintptr, string) uintptr

// Finds the child with @name in @self.
func (x *ViewStack) GetChildByName(NameVar string) *gtk.Widget {
	var cls *gtk.Widget

	cret := xViewStackGetChildByName(x.GoPointer(), NameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xViewStackGetHhomogeneous func(uintptr) bool

// Gets whether @self is horizontally homogeneous.
func (x *ViewStack) GetHhomogeneous() bool {

	cret := xViewStackGetHhomogeneous(x.GoPointer())
	return cret
}

var xViewStackGetPage func(uintptr, uintptr) uintptr

// Gets the [class@ViewStackPage] object for @child.
func (x *ViewStack) GetPage(ChildVar *gtk.Widget) *ViewStackPage {
	var cls *ViewStackPage

	cret := xViewStackGetPage(x.GoPointer(), ChildVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStackPage{}
	cls.Ptr = cret
	return cls
}

var xViewStackGetPages func(uintptr) uintptr

// Returns a [iface@Gio.ListModel] that contains the pages of the stack.
//
// This can be used to keep an up-to-date view. The model also implements
// [iface@Gtk.SelectionModel] and can be used to track and change the visible
// page.
func (x *ViewStack) GetPages() *gtk.SelectionModelBase {
	var cls *gtk.SelectionModelBase

	cret := xViewStackGetPages(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &gtk.SelectionModelBase{}
	cls.Ptr = cret
	return cls
}

var xViewStackGetVhomogeneous func(uintptr) bool

// Gets whether @self is vertically homogeneous.
func (x *ViewStack) GetVhomogeneous() bool {

	cret := xViewStackGetVhomogeneous(x.GoPointer())
	return cret
}

var xViewStackGetVisibleChild func(uintptr) uintptr

// Gets the currently visible child of @self.
func (x *ViewStack) GetVisibleChild() *gtk.Widget {
	var cls *gtk.Widget

	cret := xViewStackGetVisibleChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xViewStackGetVisibleChildName func(uintptr) string

// Returns the name of the currently visible child of @self.
func (x *ViewStack) GetVisibleChildName() string {

	cret := xViewStackGetVisibleChildName(x.GoPointer())
	return cret
}

var xViewStackRemove func(uintptr, uintptr)

// Removes a child widget from @self.
func (x *ViewStack) Remove(ChildVar *gtk.Widget) {

	xViewStackRemove(x.GoPointer(), ChildVar.GoPointer())

}

var xViewStackSetHhomogeneous func(uintptr, bool)

// Sets @self to be horizontally homogeneous or not.
//
// If the stack is horizontally homogeneous, it allocates the same width for
// all children.
//
// If it's `FALSE`, the stack may change width when a different child becomes
// visible.
func (x *ViewStack) SetHhomogeneous(HhomogeneousVar bool) {

	xViewStackSetHhomogeneous(x.GoPointer(), HhomogeneousVar)

}

var xViewStackSetVhomogeneous func(uintptr, bool)

// Sets @self to be vertically homogeneous or not.
//
// If the stack is vertically homogeneous, it allocates the same height for
// all children.
//
// If it's `FALSE`, the stack may change height when a different child becomes
// visible.
func (x *ViewStack) SetVhomogeneous(VhomogeneousVar bool) {

	xViewStackSetVhomogeneous(x.GoPointer(), VhomogeneousVar)

}

var xViewStackSetVisibleChild func(uintptr, uintptr)

// Makes @child the visible child of @self.
func (x *ViewStack) SetVisibleChild(ChildVar *gtk.Widget) {

	xViewStackSetVisibleChild(x.GoPointer(), ChildVar.GoPointer())

}

var xViewStackSetVisibleChildName func(uintptr, string)

// Makes the child with @name visible.
//
// See [property@ViewStack:visible-child].
func (x *ViewStack) SetVisibleChildName(NameVar string) {

	xViewStackSetVisibleChildName(x.GoPointer(), NameVar)

}

func (c *ViewStack) GoPointer() uintptr {
	return c.Ptr
}

func (c *ViewStack) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ViewStack) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ViewStack) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ViewStack) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ViewStack) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *ViewStack) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewStack) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []gtk.AccessibleProperty, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *ViewStack) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewStack) UpdateRelationValue(NRelationsVar int, RelationsVar []gtk.AccessibleRelation, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *ViewStack) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewStack) UpdateStateValue(NStatesVar int, StatesVar []gtk.AccessibleState, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ViewStack) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// An auxiliary class used by [class@ViewStack].
type ViewStackPage struct {
	gobject.Object
}

var xViewStackPageGLibType func() types.GType

func ViewStackPageGLibType() types.GType {
	return xViewStackPageGLibType()
}

func ViewStackPageNewFromInternalPtr(ptr uintptr) *ViewStackPage {
	cls := &ViewStackPage{}
	cls.Ptr = ptr
	return cls
}

var xViewStackPageGetBadgeNumber func(uintptr) uint

// Gets the badge number for this page.
func (x *ViewStackPage) GetBadgeNumber() uint {

	cret := xViewStackPageGetBadgeNumber(x.GoPointer())
	return cret
}

var xViewStackPageGetChild func(uintptr) uintptr

// Gets the stack child to which @self belongs.
func (x *ViewStackPage) GetChild() *gtk.Widget {
	var cls *gtk.Widget

	cret := xViewStackPageGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xViewStackPageGetIconName func(uintptr) string

// Gets the icon name of the page.
func (x *ViewStackPage) GetIconName() string {

	cret := xViewStackPageGetIconName(x.GoPointer())
	return cret
}

var xViewStackPageGetName func(uintptr) string

// Gets the name of the page.
func (x *ViewStackPage) GetName() string {

	cret := xViewStackPageGetName(x.GoPointer())
	return cret
}

var xViewStackPageGetNeedsAttention func(uintptr) bool

// Gets whether the page requires the user attention.
func (x *ViewStackPage) GetNeedsAttention() bool {

	cret := xViewStackPageGetNeedsAttention(x.GoPointer())
	return cret
}

var xViewStackPageGetTitle func(uintptr) string

// Gets the page title.
func (x *ViewStackPage) GetTitle() string {

	cret := xViewStackPageGetTitle(x.GoPointer())
	return cret
}

var xViewStackPageGetUseUnderline func(uintptr) bool

// Gets whether underlines in the page title indicate mnemonics.
func (x *ViewStackPage) GetUseUnderline() bool {

	cret := xViewStackPageGetUseUnderline(x.GoPointer())
	return cret
}

var xViewStackPageGetVisible func(uintptr) bool

// Gets whether @self is visible in its `AdwViewStack`.
//
// This is independent from the [property@Gtk.Widget:visible]
// property of its widget.
func (x *ViewStackPage) GetVisible() bool {

	cret := xViewStackPageGetVisible(x.GoPointer())
	return cret
}

var xViewStackPageSetBadgeNumber func(uintptr, uint)

// Sets the badge number for this page.
//
// [class@ViewSwitcher] can display it as a badge next to the page icon. It is
// commonly used to display a number of unread items within the page.
//
// It can be used together with [property@ViewStack{age}:needs-attention].
func (x *ViewStackPage) SetBadgeNumber(BadgeNumberVar uint) {

	xViewStackPageSetBadgeNumber(x.GoPointer(), BadgeNumberVar)

}

var xViewStackPageSetIconName func(uintptr, string)

// Sets the icon name of the page.
func (x *ViewStackPage) SetIconName(IconNameVar string) {

	xViewStackPageSetIconName(x.GoPointer(), IconNameVar)

}

var xViewStackPageSetName func(uintptr, string)

// Sets the name of the page.
func (x *ViewStackPage) SetName(NameVar string) {

	xViewStackPageSetName(x.GoPointer(), NameVar)

}

var xViewStackPageSetNeedsAttention func(uintptr, bool)

// Sets whether the page requires the user attention.
//
// [class@ViewSwitcher] will display it as a dot next to the page icon.
func (x *ViewStackPage) SetNeedsAttention(NeedsAttentionVar bool) {

	xViewStackPageSetNeedsAttention(x.GoPointer(), NeedsAttentionVar)

}

var xViewStackPageSetTitle func(uintptr, string)

// Sets the page title.
func (x *ViewStackPage) SetTitle(TitleVar string) {

	xViewStackPageSetTitle(x.GoPointer(), TitleVar)

}

var xViewStackPageSetUseUnderline func(uintptr, bool)

// Sets whether underlines in the page title indicate mnemonics.
func (x *ViewStackPage) SetUseUnderline(UseUnderlineVar bool) {

	xViewStackPageSetUseUnderline(x.GoPointer(), UseUnderlineVar)

}

var xViewStackPageSetVisible func(uintptr, bool)

// Sets whether @page is visible in its `AdwViewStack`.
//
// This is independent from the [property@Gtk.Widget:visible] property of
// [property@ViewStackPage:child].
func (x *ViewStackPage) SetVisible(VisibleVar bool) {

	xViewStackPageSetVisible(x.GoPointer(), VisibleVar)

}

func (c *ViewStackPage) GoPointer() uintptr {
	return c.Ptr
}

func (c *ViewStackPage) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ViewStackPage) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ViewStackPage) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ViewStackPage) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ViewStackPage) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *ViewStackPage) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewStackPage) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []gtk.AccessibleProperty, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *ViewStackPage) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewStackPage) UpdateRelationValue(NRelationsVar int, RelationsVar []gtk.AccessibleRelation, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *ViewStackPage) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ViewStackPage) UpdateStateValue(NStatesVar int, StatesVar []gtk.AccessibleState, ValuesVar []gobject.Value) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// An auxiliary class used by [class@ViewStack].
//
// See [property@ViewStack:pages].
type ViewStackPages struct {
	gobject.Object
}

var xViewStackPagesGLibType func() types.GType

func ViewStackPagesGLibType() types.GType {
	return xViewStackPagesGLibType()
}

func ViewStackPagesNewFromInternalPtr(ptr uintptr) *ViewStackPages {
	cls := &ViewStackPages{}
	cls.Ptr = ptr
	return cls
}

var xViewStackPagesGetSelectedPage func(uintptr) uintptr

// Gets the [class@ViewStackPage] for the visible child of a view stack
//
// Gets the [class@ViewStackPage] for the visible child of the associated stack.
//
// Returns `NULL` if there's no selected page.
func (x *ViewStackPages) GetSelectedPage() *ViewStackPage {
	var cls *ViewStackPage

	cret := xViewStackPagesGetSelectedPage(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ViewStackPage{}
	cls.Ptr = cret
	return cls
}

var xViewStackPagesSetSelectedPage func(uintptr, uintptr)

// Sets the visible child in the associated [class@ViewStack].
//
// See [property@ViewStack:visible-child].
func (x *ViewStackPages) SetSelectedPage(PageVar *ViewStackPage) {

	xViewStackPagesSetSelectedPage(x.GoPointer(), PageVar.GoPointer())

}

func (c *ViewStackPages) GoPointer() uintptr {
	return c.Ptr
}

func (c *ViewStackPages) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// See also: g_list_model_get_n_items()
func (x *ViewStackPages) GetItem(PositionVar uint) uintptr {

	cret := gio.XGListModelGetItem(x.GoPointer(), PositionVar)
	return cret
}

// Gets the type of the items in @list.
//
// All items returned from g_list_model_get_item() are of the type
// returned by this function, or a subtype, or if the type is an
// interface, they are an implementation of that interface.
//
// The item type of a #GListModel can not change during the life of the
// model.
func (x *ViewStackPages) GetItemType() types.GType {

	cret := gio.XGListModelGetItemType(x.GoPointer())
	return cret
}

// Gets the number of items in @list.
//
// Depending on the model implementation, calling this function may be
// less efficient than iterating the list with increasing values for
// @position until g_list_model_get_item() returns %NULL.
func (x *ViewStackPages) GetNItems() uint {

	cret := gio.XGListModelGetNItems(x.GoPointer())
	return cret
}

// Get the item at @position.
//
// If @position is greater than the number of items in @list, %NULL is
// returned.
//
// %NULL is never returned for an index that is smaller than the length
// of the list.
//
// This function is meant to be used by language bindings in place
// of g_list_model_get_item().
//
// See also: g_list_model_get_n_items()
func (x *ViewStackPages) GetObject(PositionVar uint) *gobject.Object {
	var cls *gobject.Object

	cret := gio.XGListModelGetObject(x.GoPointer(), PositionVar)

	if cret == 0 {
		return nil
	}
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

// Emits the #GListModel::items-changed signal on @list.
//
// This function should only be called by classes implementing
// #GListModel. It has to be called after the internal representation
// of @list has been updated, because handlers connected to this signal
// might query the new state of the list.
//
// Implementations must only make changes to the model (as visible to
// its consumer) in places that will not cause problems for that
// consumer.  For models that are driven directly by a write API (such
// as #GListStore), changes can be reported in response to uses of that
// API.  For models that represent remote data, changes should only be
// made from a fresh mainloop dispatch.  It is particularly not
// permitted to make changes in response to a call to the #GListModel
// consumer API.
//
// Stated another way: in general, it is assumed that code making a
// series of accesses to the model via the API, without returning to the
// mainloop, and without calling other code, will continue to view the
// same contents of the model.
func (x *ViewStackPages) ItemsChanged(PositionVar uint, RemovedVar uint, AddedVar uint) {

	gio.XGListModelItemsChanged(x.GoPointer(), PositionVar, RemovedVar, AddedVar)

}

// Gets the set containing all currently selected items in the model.
//
// This function may be slow, so if you are only interested in single item,
// consider using [method@Gtk.SelectionModel.is_selected] or if you are only
// interested in a few, consider [method@Gtk.SelectionModel.get_selection_in_range].
func (x *ViewStackPages) GetSelection() *gtk.Bitset {

	cret := gtk.XGtkSelectionModelGetSelection(x.GoPointer())
	return cret
}

// Gets the set of selected items in a range.
//
// This function is an optimization for
// [method@Gtk.SelectionModel.get_selection] when you are only
// interested in part of the model's selected state. A common use
// case is in response to the [signal@Gtk.SelectionModel::selection-changed]
// signal.
func (x *ViewStackPages) GetSelectionInRange(PositionVar uint, NItemsVar uint) *gtk.Bitset {

	cret := gtk.XGtkSelectionModelGetSelectionInRange(x.GoPointer(), PositionVar, NItemsVar)
	return cret
}

// Checks if the given item is selected.
func (x *ViewStackPages) IsSelected(PositionVar uint) bool {

	cret := gtk.XGtkSelectionModelIsSelected(x.GoPointer(), PositionVar)
	return cret
}

// Requests to select all items in the model.
func (x *ViewStackPages) SelectAll() bool {

	cret := gtk.XGtkSelectionModelSelectAll(x.GoPointer())
	return cret
}

// Requests to select an item in the model.
func (x *ViewStackPages) SelectItem(PositionVar uint, UnselectRestVar bool) bool {

	cret := gtk.XGtkSelectionModelSelectItem(x.GoPointer(), PositionVar, UnselectRestVar)
	return cret
}

// Requests to select a range of items in the model.
func (x *ViewStackPages) SelectRange(PositionVar uint, NItemsVar uint, UnselectRestVar bool) bool {

	cret := gtk.XGtkSelectionModelSelectRange(x.GoPointer(), PositionVar, NItemsVar, UnselectRestVar)
	return cret
}

// Helper function for implementations of `GtkSelectionModel`.
//
// Call this when a the selection changes to emit the
// [signal@Gtk.SelectionModel::selection-changed] signal.
func (x *ViewStackPages) SelectionChanged(PositionVar uint, NItemsVar uint) {

	gtk.XGtkSelectionModelSelectionChanged(x.GoPointer(), PositionVar, NItemsVar)

}

// Make selection changes.
//
// This is the most advanced selection updating method that allows
// the most fine-grained control over selection changes. If you can,
// you should try the simpler versions, as implementations are more
// likely to implement support for those.
//
// Requests that the selection state of all positions set in @mask
// be updated to the respective value in the @selected bitmask.
//
// In pseudocode, it would look something like this:
//
// ```c
// for (i = 0; i &lt; n_items; i++)
//
//	{
//	  // don't change values not in the mask
//	  if (!gtk_bitset_contains (mask, i))
//	    continue;
//
//	  if (gtk_bitset_contains (selected, i))
//	    select_item (i);
//	  else
//	    unselect_item (i);
//	}
//
// gtk_selection_model_selection_changed (model,
//
//	first_changed_item,
//	n_changed_items);
//
// ```
//
// @mask and @selected must not be modified. They may refer to the
// same bitset, which would mean that every item in the set should
// be selected.
func (x *ViewStackPages) SetSelection(SelectedVar *gtk.Bitset, MaskVar *gtk.Bitset) bool {

	cret := gtk.XGtkSelectionModelSetSelection(x.GoPointer(), SelectedVar, MaskVar)
	return cret
}

// Requests to unselect all items in the model.
func (x *ViewStackPages) UnselectAll() bool {

	cret := gtk.XGtkSelectionModelUnselectAll(x.GoPointer())
	return cret
}

// Requests to unselect an item in the model.
func (x *ViewStackPages) UnselectItem(PositionVar uint) bool {

	cret := gtk.XGtkSelectionModelUnselectItem(x.GoPointer(), PositionVar)
	return cret
}

// Requests to unselect a range of items in the model.
func (x *ViewStackPages) UnselectRange(PositionVar uint, NItemsVar uint) bool {

	cret := gtk.XGtkSelectionModelUnselectRange(x.GoPointer(), PositionVar, NItemsVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xViewStackGLibType, lib, "adw_view_stack_get_type")

	core.PuregoSafeRegister(&xNewViewStack, lib, "adw_view_stack_new")

	core.PuregoSafeRegister(&xViewStackAdd, lib, "adw_view_stack_add")
	core.PuregoSafeRegister(&xViewStackAddNamed, lib, "adw_view_stack_add_named")
	core.PuregoSafeRegister(&xViewStackAddTitled, lib, "adw_view_stack_add_titled")
	core.PuregoSafeRegister(&xViewStackAddTitledWithIcon, lib, "adw_view_stack_add_titled_with_icon")
	core.PuregoSafeRegister(&xViewStackGetChildByName, lib, "adw_view_stack_get_child_by_name")
	core.PuregoSafeRegister(&xViewStackGetHhomogeneous, lib, "adw_view_stack_get_hhomogeneous")
	core.PuregoSafeRegister(&xViewStackGetPage, lib, "adw_view_stack_get_page")
	core.PuregoSafeRegister(&xViewStackGetPages, lib, "adw_view_stack_get_pages")
	core.PuregoSafeRegister(&xViewStackGetVhomogeneous, lib, "adw_view_stack_get_vhomogeneous")
	core.PuregoSafeRegister(&xViewStackGetVisibleChild, lib, "adw_view_stack_get_visible_child")
	core.PuregoSafeRegister(&xViewStackGetVisibleChildName, lib, "adw_view_stack_get_visible_child_name")
	core.PuregoSafeRegister(&xViewStackRemove, lib, "adw_view_stack_remove")
	core.PuregoSafeRegister(&xViewStackSetHhomogeneous, lib, "adw_view_stack_set_hhomogeneous")
	core.PuregoSafeRegister(&xViewStackSetVhomogeneous, lib, "adw_view_stack_set_vhomogeneous")
	core.PuregoSafeRegister(&xViewStackSetVisibleChild, lib, "adw_view_stack_set_visible_child")
	core.PuregoSafeRegister(&xViewStackSetVisibleChildName, lib, "adw_view_stack_set_visible_child_name")

	core.PuregoSafeRegister(&xViewStackPageGLibType, lib, "adw_view_stack_page_get_type")

	core.PuregoSafeRegister(&xViewStackPageGetBadgeNumber, lib, "adw_view_stack_page_get_badge_number")
	core.PuregoSafeRegister(&xViewStackPageGetChild, lib, "adw_view_stack_page_get_child")
	core.PuregoSafeRegister(&xViewStackPageGetIconName, lib, "adw_view_stack_page_get_icon_name")
	core.PuregoSafeRegister(&xViewStackPageGetName, lib, "adw_view_stack_page_get_name")
	core.PuregoSafeRegister(&xViewStackPageGetNeedsAttention, lib, "adw_view_stack_page_get_needs_attention")
	core.PuregoSafeRegister(&xViewStackPageGetTitle, lib, "adw_view_stack_page_get_title")
	core.PuregoSafeRegister(&xViewStackPageGetUseUnderline, lib, "adw_view_stack_page_get_use_underline")
	core.PuregoSafeRegister(&xViewStackPageGetVisible, lib, "adw_view_stack_page_get_visible")
	core.PuregoSafeRegister(&xViewStackPageSetBadgeNumber, lib, "adw_view_stack_page_set_badge_number")
	core.PuregoSafeRegister(&xViewStackPageSetIconName, lib, "adw_view_stack_page_set_icon_name")
	core.PuregoSafeRegister(&xViewStackPageSetName, lib, "adw_view_stack_page_set_name")
	core.PuregoSafeRegister(&xViewStackPageSetNeedsAttention, lib, "adw_view_stack_page_set_needs_attention")
	core.PuregoSafeRegister(&xViewStackPageSetTitle, lib, "adw_view_stack_page_set_title")
	core.PuregoSafeRegister(&xViewStackPageSetUseUnderline, lib, "adw_view_stack_page_set_use_underline")
	core.PuregoSafeRegister(&xViewStackPageSetVisible, lib, "adw_view_stack_page_set_visible")

	core.PuregoSafeRegister(&xViewStackPagesGLibType, lib, "adw_view_stack_pages_get_type")

	core.PuregoSafeRegister(&xViewStackPagesGetSelectedPage, lib, "adw_view_stack_pages_get_selected_page")
	core.PuregoSafeRegister(&xViewStackPagesSetSelectedPage, lib, "adw_view_stack_pages_set_selected_page")

}
