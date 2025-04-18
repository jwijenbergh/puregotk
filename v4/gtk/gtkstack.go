// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Possible transitions between pages in a `GtkStack` widget.
//
// New values may be added to this enumeration over time.
type StackTransitionType int

var xStackTransitionTypeGLibType func() types.GType

func StackTransitionTypeGLibType() types.GType {
	return xStackTransitionTypeGLibType()
}

const (

	// No transition
	StackTransitionTypeNoneValue StackTransitionType = 0
	// A cross-fade
	StackTransitionTypeCrossfadeValue StackTransitionType = 1
	// Slide from left to right
	StackTransitionTypeSlideRightValue StackTransitionType = 2
	// Slide from right to left
	StackTransitionTypeSlideLeftValue StackTransitionType = 3
	// Slide from bottom up
	StackTransitionTypeSlideUpValue StackTransitionType = 4
	// Slide from top down
	StackTransitionTypeSlideDownValue StackTransitionType = 5
	// Slide from left or right according to the children order
	StackTransitionTypeSlideLeftRightValue StackTransitionType = 6
	// Slide from top down or bottom up according to the order
	StackTransitionTypeSlideUpDownValue StackTransitionType = 7
	// Cover the old page by sliding up
	StackTransitionTypeOverUpValue StackTransitionType = 8
	// Cover the old page by sliding down
	StackTransitionTypeOverDownValue StackTransitionType = 9
	// Cover the old page by sliding to the left
	StackTransitionTypeOverLeftValue StackTransitionType = 10
	// Cover the old page by sliding to the right
	StackTransitionTypeOverRightValue StackTransitionType = 11
	// Uncover the new page by sliding up
	StackTransitionTypeUnderUpValue StackTransitionType = 12
	// Uncover the new page by sliding down
	StackTransitionTypeUnderDownValue StackTransitionType = 13
	// Uncover the new page by sliding to the left
	StackTransitionTypeUnderLeftValue StackTransitionType = 14
	// Uncover the new page by sliding to the right
	StackTransitionTypeUnderRightValue StackTransitionType = 15
	// Cover the old page sliding up or uncover the new page sliding down, according to order
	StackTransitionTypeOverUpDownValue StackTransitionType = 16
	// Cover the old page sliding down or uncover the new page sliding up, according to order
	StackTransitionTypeOverDownUpValue StackTransitionType = 17
	// Cover the old page sliding left or uncover the new page sliding right, according to order
	StackTransitionTypeOverLeftRightValue StackTransitionType = 18
	// Cover the old page sliding right or uncover the new page sliding left, according to order
	StackTransitionTypeOverRightLeftValue StackTransitionType = 19
	// Pretend the pages are sides of a cube and rotate that cube to the left
	StackTransitionTypeRotateLeftValue StackTransitionType = 20
	// Pretend the pages are sides of a cube and rotate that cube to the right
	StackTransitionTypeRotateRightValue StackTransitionType = 21
	// Pretend the pages are sides of a cube and rotate that cube to the left or right according to the children order
	StackTransitionTypeRotateLeftRightValue StackTransitionType = 22
)

// `GtkStack` is a container which only shows one of its children
// at a time.
//
// In contrast to `GtkNotebook`, `GtkStack` does not provide a means
// for users to change the visible child. Instead, a separate widget
// such as [class@Gtk.StackSwitcher] or [class@Gtk.StackSidebar] can
// be used with `GtkStack` to provide this functionality.
//
// Transitions between pages can be animated as slides or fades. This
// can be controlled with [method@Gtk.Stack.set_transition_type].
// These animations respect the [property@Gtk.Settings:gtk-enable-animations]
// setting.
//
// `GtkStack` maintains a [class@Gtk.StackPage] object for each added
// child, which holds additional per-child properties. You
// obtain the `GtkStackPage` for a child with [method@Gtk.Stack.get_page]
// and you can obtain a `GtkSelectionModel` containing all the pages
// with [method@Gtk.Stack.get_pages].
//
// # GtkStack as GtkBuildable
//
// To set child-specific properties in a .ui file, create `GtkStackPage`
// objects explicitly, and set the child widget as a property on it:
//
// ```xml
//
//	&lt;object class="GtkStack" id="stack"&gt;
//	  &lt;child&gt;
//	    &lt;object class="GtkStackPage"&gt;
//	      &lt;property name="name"&gt;page1&lt;/property&gt;
//	      &lt;property name="title"&gt;In the beginning…&lt;/property&gt;
//	      &lt;property name="child"&gt;
//	        &lt;object class="GtkLabel"&gt;
//	          &lt;property name="label"&gt;It was dark&lt;/property&gt;
//	        &lt;/object&gt;
//	      &lt;/property&gt;
//	    &lt;/object&gt;
//	  &lt;/child&gt;
//
// ```
//
// # CSS nodes
//
// `GtkStack` has a single CSS node named stack.
//
// # Accessibility
//
// `GtkStack` uses the %GTK_ACCESSIBLE_ROLE_TAB_PANEL for the stack
// pages, which are the accessible parent objects of the child widgets.
type Stack struct {
	Widget
}

var xStackGLibType func() types.GType

func StackGLibType() types.GType {
	return xStackGLibType()
}

func StackNewFromInternalPtr(ptr uintptr) *Stack {
	cls := &Stack{}
	cls.Ptr = ptr
	return cls
}

var xNewStack func() uintptr

// Creates a new `GtkStack`.
func NewStack() *Stack {
	var cls *Stack

	cret := xNewStack()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Stack{}
	cls.Ptr = cret
	return cls
}

var xStackAddChild func(uintptr, uintptr) uintptr

// Adds a child to @stack.
func (x *Stack) AddChild(ChildVar *Widget) *StackPage {
	var cls *StackPage

	cret := xStackAddChild(x.GoPointer(), ChildVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &StackPage{}
	cls.Ptr = cret
	return cls
}

var xStackAddNamed func(uintptr, uintptr, string) uintptr

// Adds a child to @stack.
//
// The child is identified by the @name.
func (x *Stack) AddNamed(ChildVar *Widget, NameVar string) *StackPage {
	var cls *StackPage

	cret := xStackAddNamed(x.GoPointer(), ChildVar.GoPointer(), NameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &StackPage{}
	cls.Ptr = cret
	return cls
}

var xStackAddTitled func(uintptr, uintptr, string, string) uintptr

// Adds a child to @stack.
//
// The child is identified by the @name. The @title
// will be used by `GtkStackSwitcher` to represent
// @child in a tab bar, so it should be short.
func (x *Stack) AddTitled(ChildVar *Widget, NameVar string, TitleVar string) *StackPage {
	var cls *StackPage

	cret := xStackAddTitled(x.GoPointer(), ChildVar.GoPointer(), NameVar, TitleVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &StackPage{}
	cls.Ptr = cret
	return cls
}

var xStackGetChildByName func(uintptr, string) uintptr

// Finds the child with the name given as the argument.
//
// Returns %NULL if there is no child with this name.
func (x *Stack) GetChildByName(NameVar string) *Widget {
	var cls *Widget

	cret := xStackGetChildByName(x.GoPointer(), NameVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xStackGetHhomogeneous func(uintptr) bool

// Gets whether @stack is horizontally homogeneous.
func (x *Stack) GetHhomogeneous() bool {

	cret := xStackGetHhomogeneous(x.GoPointer())
	return cret
}

var xStackGetInterpolateSize func(uintptr) bool

// Returns whether the `GtkStack` is set up to interpolate between
// the sizes of children on page switch.
func (x *Stack) GetInterpolateSize() bool {

	cret := xStackGetInterpolateSize(x.GoPointer())
	return cret
}

var xStackGetPage func(uintptr, uintptr) uintptr

// Returns the `GtkStackPage` object for @child.
func (x *Stack) GetPage(ChildVar *Widget) *StackPage {
	var cls *StackPage

	cret := xStackGetPage(x.GoPointer(), ChildVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &StackPage{}
	cls.Ptr = cret
	return cls
}

var xStackGetPages func(uintptr) uintptr

// Returns a `GListModel` that contains the pages of the stack.
//
// This can be used to keep an up-to-date view. The model also
// implements [iface@Gtk.SelectionModel] and can be used to track
// and modify the visible page.
func (x *Stack) GetPages() *SelectionModelBase {
	var cls *SelectionModelBase

	cret := xStackGetPages(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &SelectionModelBase{}
	cls.Ptr = cret
	return cls
}

var xStackGetTransitionDuration func(uintptr) uint

// Returns the amount of time (in milliseconds) that
// transitions between pages in @stack will take.
func (x *Stack) GetTransitionDuration() uint {

	cret := xStackGetTransitionDuration(x.GoPointer())
	return cret
}

var xStackGetTransitionRunning func(uintptr) bool

// Returns whether the @stack is currently in a transition from one page to
// another.
func (x *Stack) GetTransitionRunning() bool {

	cret := xStackGetTransitionRunning(x.GoPointer())
	return cret
}

var xStackGetTransitionType func(uintptr) StackTransitionType

// Gets the type of animation that will be used
// for transitions between pages in @stack.
func (x *Stack) GetTransitionType() StackTransitionType {

	cret := xStackGetTransitionType(x.GoPointer())
	return cret
}

var xStackGetVhomogeneous func(uintptr) bool

// Gets whether @stack is vertically homogeneous.
func (x *Stack) GetVhomogeneous() bool {

	cret := xStackGetVhomogeneous(x.GoPointer())
	return cret
}

var xStackGetVisibleChild func(uintptr) uintptr

// Gets the currently visible child of @stack.
//
// Returns %NULL if there are no visible children.
func (x *Stack) GetVisibleChild() *Widget {
	var cls *Widget

	cret := xStackGetVisibleChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xStackGetVisibleChildName func(uintptr) string

// Returns the name of the currently visible child of @stack.
//
// Returns %NULL if there is no visible child.
func (x *Stack) GetVisibleChildName() string {

	cret := xStackGetVisibleChildName(x.GoPointer())
	return cret
}

var xStackRemove func(uintptr, uintptr)

// Removes a child widget from @stack.
func (x *Stack) Remove(ChildVar *Widget) {

	xStackRemove(x.GoPointer(), ChildVar.GoPointer())

}

var xStackSetHhomogeneous func(uintptr, bool)

// Sets the `GtkStack` to be horizontally homogeneous or not.
//
// If it is homogeneous, the `GtkStack` will request the same
// width for all its children. If it isn't, the stack
// may change width when a different child becomes visible.
func (x *Stack) SetHhomogeneous(HhomogeneousVar bool) {

	xStackSetHhomogeneous(x.GoPointer(), HhomogeneousVar)

}

var xStackSetInterpolateSize func(uintptr, bool)

// Sets whether or not @stack will interpolate its size when
// changing the visible child.
//
// If the [property@Gtk.Stack:interpolate-size] property is set
// to %TRUE, @stack will interpolate its size between the current
// one and the one it'll take after changing the visible child,
// according to the set transition duration.
func (x *Stack) SetInterpolateSize(InterpolateSizeVar bool) {

	xStackSetInterpolateSize(x.GoPointer(), InterpolateSizeVar)

}

var xStackSetTransitionDuration func(uintptr, uint)

// Sets the duration that transitions between pages in @stack
// will take.
func (x *Stack) SetTransitionDuration(DurationVar uint) {

	xStackSetTransitionDuration(x.GoPointer(), DurationVar)

}

var xStackSetTransitionType func(uintptr, StackTransitionType)

// Sets the type of animation that will be used for
// transitions between pages in @stack.
//
// Available types include various kinds of fades and slides.
//
// The transition type can be changed without problems
// at runtime, so it is possible to change the animation
// based on the page that is about to become current.
func (x *Stack) SetTransitionType(TransitionVar StackTransitionType) {

	xStackSetTransitionType(x.GoPointer(), TransitionVar)

}

var xStackSetVhomogeneous func(uintptr, bool)

// Sets the `GtkStack` to be vertically homogeneous or not.
//
// If it is homogeneous, the `GtkStack` will request the same
// height for all its children. If it isn't, the stack
// may change height when a different child becomes visible.
func (x *Stack) SetVhomogeneous(VhomogeneousVar bool) {

	xStackSetVhomogeneous(x.GoPointer(), VhomogeneousVar)

}

var xStackSetVisibleChild func(uintptr, uintptr)

// Makes @child the visible child of @stack.
//
// If @child is different from the currently visible child,
// the transition between the two will be animated with the
// current transition type of @stack.
//
// Note that the @child widget has to be visible itself
// (see [method@Gtk.Widget.show]) in order to become the visible
// child of @stack.
func (x *Stack) SetVisibleChild(ChildVar *Widget) {

	xStackSetVisibleChild(x.GoPointer(), ChildVar.GoPointer())

}

var xStackSetVisibleChildFull func(uintptr, string, StackTransitionType)

// Makes the child with the given name visible.
//
// Note that the child widget has to be visible itself
// (see [method@Gtk.Widget.show]) in order to become the visible
// child of @stack.
func (x *Stack) SetVisibleChildFull(NameVar string, TransitionVar StackTransitionType) {

	xStackSetVisibleChildFull(x.GoPointer(), NameVar, TransitionVar)

}

var xStackSetVisibleChildName func(uintptr, string)

// Makes the child with the given name visible.
//
// If @child is different from the currently visible child,
// the transition between the two will be animated with the
// current transition type of @stack.
//
// Note that the child widget has to be visible itself
// (see [method@Gtk.Widget.show]) in order to become the visible
// child of @stack.
func (x *Stack) SetVisibleChildName(NameVar string) {

	xStackSetVisibleChildName(x.GoPointer(), NameVar)

}

func (c *Stack) GoPointer() uintptr {
	return c.Ptr
}

func (c *Stack) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Stack) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Stack) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Stack) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Stack) ResetState(StateVar AccessibleState) {

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
func (x *Stack) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Stack) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *Stack) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Stack) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *Stack) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Stack) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Stack) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// `GtkStackPage` is an auxiliary class used by `GtkStack`.
type StackPage struct {
	gobject.Object
}

var xStackPageGLibType func() types.GType

func StackPageGLibType() types.GType {
	return xStackPageGLibType()
}

func StackPageNewFromInternalPtr(ptr uintptr) *StackPage {
	cls := &StackPage{}
	cls.Ptr = ptr
	return cls
}

var xStackPageGetChild func(uintptr) uintptr

// Returns the stack child to which @self belongs.
func (x *StackPage) GetChild() *Widget {
	var cls *Widget

	cret := xStackPageGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xStackPageGetIconName func(uintptr) string

// Returns the icon name of the page.
func (x *StackPage) GetIconName() string {

	cret := xStackPageGetIconName(x.GoPointer())
	return cret
}

var xStackPageGetName func(uintptr) string

// Returns the name of the page.
func (x *StackPage) GetName() string {

	cret := xStackPageGetName(x.GoPointer())
	return cret
}

var xStackPageGetNeedsAttention func(uintptr) bool

// Returns whether the page is marked as “needs attention”.
func (x *StackPage) GetNeedsAttention() bool {

	cret := xStackPageGetNeedsAttention(x.GoPointer())
	return cret
}

var xStackPageGetTitle func(uintptr) string

// Gets the page title.
func (x *StackPage) GetTitle() string {

	cret := xStackPageGetTitle(x.GoPointer())
	return cret
}

var xStackPageGetUseUnderline func(uintptr) bool

// Gets whether underlines in the page title indicate mnemonics.
func (x *StackPage) GetUseUnderline() bool {

	cret := xStackPageGetUseUnderline(x.GoPointer())
	return cret
}

var xStackPageGetVisible func(uintptr) bool

// Returns whether @page is visible in its `GtkStack`.
//
// This is independent from the [property@Gtk.Widget:visible]
// property of its widget.
func (x *StackPage) GetVisible() bool {

	cret := xStackPageGetVisible(x.GoPointer())
	return cret
}

var xStackPageSetIconName func(uintptr, string)

// Sets the icon name of the page.
func (x *StackPage) SetIconName(SettingVar string) {

	xStackPageSetIconName(x.GoPointer(), SettingVar)

}

var xStackPageSetName func(uintptr, string)

// Sets the name of the page.
func (x *StackPage) SetName(SettingVar string) {

	xStackPageSetName(x.GoPointer(), SettingVar)

}

var xStackPageSetNeedsAttention func(uintptr, bool)

// Sets whether the page is marked as “needs attention”.
func (x *StackPage) SetNeedsAttention(SettingVar bool) {

	xStackPageSetNeedsAttention(x.GoPointer(), SettingVar)

}

var xStackPageSetTitle func(uintptr, string)

// Sets the page title.
func (x *StackPage) SetTitle(SettingVar string) {

	xStackPageSetTitle(x.GoPointer(), SettingVar)

}

var xStackPageSetUseUnderline func(uintptr, bool)

// Sets whether underlines in the page title indicate mnemonics.
func (x *StackPage) SetUseUnderline(SettingVar bool) {

	xStackPageSetUseUnderline(x.GoPointer(), SettingVar)

}

var xStackPageSetVisible func(uintptr, bool)

// Sets whether @page is visible in its `GtkStack`.
func (x *StackPage) SetVisible(VisibleVar bool) {

	xStackPageSetVisible(x.GoPointer(), VisibleVar)

}

func (c *StackPage) GoPointer() uintptr {
	return c.Ptr
}

func (c *StackPage) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *StackPage) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *StackPage) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *StackPage) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *StackPage) ResetState(StateVar AccessibleState) {

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
func (x *StackPage) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *StackPage) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *StackPage) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *StackPage) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *StackPage) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *StackPage) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xStackTransitionTypeGLibType, lib, "gtk_stack_transition_type_get_type")

	core.PuregoSafeRegister(&xStackGLibType, lib, "gtk_stack_get_type")

	core.PuregoSafeRegister(&xNewStack, lib, "gtk_stack_new")

	core.PuregoSafeRegister(&xStackAddChild, lib, "gtk_stack_add_child")
	core.PuregoSafeRegister(&xStackAddNamed, lib, "gtk_stack_add_named")
	core.PuregoSafeRegister(&xStackAddTitled, lib, "gtk_stack_add_titled")
	core.PuregoSafeRegister(&xStackGetChildByName, lib, "gtk_stack_get_child_by_name")
	core.PuregoSafeRegister(&xStackGetHhomogeneous, lib, "gtk_stack_get_hhomogeneous")
	core.PuregoSafeRegister(&xStackGetInterpolateSize, lib, "gtk_stack_get_interpolate_size")
	core.PuregoSafeRegister(&xStackGetPage, lib, "gtk_stack_get_page")
	core.PuregoSafeRegister(&xStackGetPages, lib, "gtk_stack_get_pages")
	core.PuregoSafeRegister(&xStackGetTransitionDuration, lib, "gtk_stack_get_transition_duration")
	core.PuregoSafeRegister(&xStackGetTransitionRunning, lib, "gtk_stack_get_transition_running")
	core.PuregoSafeRegister(&xStackGetTransitionType, lib, "gtk_stack_get_transition_type")
	core.PuregoSafeRegister(&xStackGetVhomogeneous, lib, "gtk_stack_get_vhomogeneous")
	core.PuregoSafeRegister(&xStackGetVisibleChild, lib, "gtk_stack_get_visible_child")
	core.PuregoSafeRegister(&xStackGetVisibleChildName, lib, "gtk_stack_get_visible_child_name")
	core.PuregoSafeRegister(&xStackRemove, lib, "gtk_stack_remove")
	core.PuregoSafeRegister(&xStackSetHhomogeneous, lib, "gtk_stack_set_hhomogeneous")
	core.PuregoSafeRegister(&xStackSetInterpolateSize, lib, "gtk_stack_set_interpolate_size")
	core.PuregoSafeRegister(&xStackSetTransitionDuration, lib, "gtk_stack_set_transition_duration")
	core.PuregoSafeRegister(&xStackSetTransitionType, lib, "gtk_stack_set_transition_type")
	core.PuregoSafeRegister(&xStackSetVhomogeneous, lib, "gtk_stack_set_vhomogeneous")
	core.PuregoSafeRegister(&xStackSetVisibleChild, lib, "gtk_stack_set_visible_child")
	core.PuregoSafeRegister(&xStackSetVisibleChildFull, lib, "gtk_stack_set_visible_child_full")
	core.PuregoSafeRegister(&xStackSetVisibleChildName, lib, "gtk_stack_set_visible_child_name")

	core.PuregoSafeRegister(&xStackPageGLibType, lib, "gtk_stack_page_get_type")

	core.PuregoSafeRegister(&xStackPageGetChild, lib, "gtk_stack_page_get_child")
	core.PuregoSafeRegister(&xStackPageGetIconName, lib, "gtk_stack_page_get_icon_name")
	core.PuregoSafeRegister(&xStackPageGetName, lib, "gtk_stack_page_get_name")
	core.PuregoSafeRegister(&xStackPageGetNeedsAttention, lib, "gtk_stack_page_get_needs_attention")
	core.PuregoSafeRegister(&xStackPageGetTitle, lib, "gtk_stack_page_get_title")
	core.PuregoSafeRegister(&xStackPageGetUseUnderline, lib, "gtk_stack_page_get_use_underline")
	core.PuregoSafeRegister(&xStackPageGetVisible, lib, "gtk_stack_page_get_visible")
	core.PuregoSafeRegister(&xStackPageSetIconName, lib, "gtk_stack_page_set_icon_name")
	core.PuregoSafeRegister(&xStackPageSetName, lib, "gtk_stack_page_set_name")
	core.PuregoSafeRegister(&xStackPageSetNeedsAttention, lib, "gtk_stack_page_set_needs_attention")
	core.PuregoSafeRegister(&xStackPageSetTitle, lib, "gtk_stack_page_set_title")
	core.PuregoSafeRegister(&xStackPageSetUseUnderline, lib, "gtk_stack_page_set_use_underline")
	core.PuregoSafeRegister(&xStackPageSetVisible, lib, "gtk_stack_page_set_visible")

}
