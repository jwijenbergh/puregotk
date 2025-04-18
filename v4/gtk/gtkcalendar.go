// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// `GtkCalendar` is a widget that displays a Gregorian calendar, one month
// at a time.
//
// ![An example GtkCalendar](calendar.png)
//
// A `GtkCalendar` can be created with [ctor@Gtk.Calendar.new].
//
// The date that is currently displayed can be altered with
// [method@Gtk.Calendar.select_day].
//
// To place a visual marker on a particular day, use
// [method@Gtk.Calendar.mark_day] and to remove the marker,
// [method@Gtk.Calendar.unmark_day]. Alternative, all
// marks can be cleared with [method@Gtk.Calendar.clear_marks].
//
// The selected date can be retrieved from a `GtkCalendar` using
// [method@Gtk.Calendar.get_date].
//
// Users should be aware that, although the Gregorian calendar is the
// legal calendar in most countries, it was adopted progressively
// between 1582 and 1929. Display before these dates is likely to be
// historically incorrect.
//
// # CSS nodes
//
// ```
// calendar.view
// ├── header
// │   ├── button
// │   ├── stack.month
// │   ├── button
// │   ├── button
// │   ├── label.year
// │   ╰── button
// ╰── grid
//
//	╰── label[.day-name][.week-number][.day-number][.other-month][.today]
//
// ```
//
// `GtkCalendar` has a main node with name calendar. It contains a subnode
// called header containing the widgets for switching between years and months.
//
// The grid subnode contains all day labels, including week numbers on the left
// (marked with the .week-number css class) and day names on top (marked with the
// .day-name css class).
//
// Day labels that belong to the previous or next month get the .other-month
// style class. The label of the current day get the .today style class.
//
// Marked day labels get the :selected state assigned.
type Calendar struct {
	Widget
}

var xCalendarGLibType func() types.GType

func CalendarGLibType() types.GType {
	return xCalendarGLibType()
}

func CalendarNewFromInternalPtr(ptr uintptr) *Calendar {
	cls := &Calendar{}
	cls.Ptr = ptr
	return cls
}

var xNewCalendar func() uintptr

// Creates a new calendar, with the current date being selected.
func NewCalendar() *Calendar {
	var cls *Calendar

	cret := xNewCalendar()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Calendar{}
	cls.Ptr = cret
	return cls
}

var xCalendarClearMarks func(uintptr)

// Remove all visual markers.
func (x *Calendar) ClearMarks() {

	xCalendarClearMarks(x.GoPointer())

}

var xCalendarGetDate func(uintptr) *glib.DateTime

// Returns a `GDateTime` representing the shown
// year, month and the selected day.
//
// The returned date is in the local time zone.
func (x *Calendar) GetDate() *glib.DateTime {

	cret := xCalendarGetDate(x.GoPointer())
	return cret
}

var xCalendarGetDayIsMarked func(uintptr, uint) bool

// Returns if the @day of the @calendar is already marked.
func (x *Calendar) GetDayIsMarked(DayVar uint) bool {

	cret := xCalendarGetDayIsMarked(x.GoPointer(), DayVar)
	return cret
}

var xCalendarGetShowDayNames func(uintptr) bool

// Returns whether @self is currently showing the names
// of the week days.
//
// This is the value of the [property@Gtk.Calendar:show-day-names]
// property.
func (x *Calendar) GetShowDayNames() bool {

	cret := xCalendarGetShowDayNames(x.GoPointer())
	return cret
}

var xCalendarGetShowHeading func(uintptr) bool

// Returns whether @self is currently showing the heading.
//
// This is the value of the [property@Gtk.Calendar:show-heading]
// property.
func (x *Calendar) GetShowHeading() bool {

	cret := xCalendarGetShowHeading(x.GoPointer())
	return cret
}

var xCalendarGetShowWeekNumbers func(uintptr) bool

// Returns whether @self is showing week numbers right
// now.
//
// This is the value of the [property@Gtk.Calendar:show-week-numbers]
// property.
func (x *Calendar) GetShowWeekNumbers() bool {

	cret := xCalendarGetShowWeekNumbers(x.GoPointer())
	return cret
}

var xCalendarMarkDay func(uintptr, uint)

// Places a visual marker on a particular day.
func (x *Calendar) MarkDay(DayVar uint) {

	xCalendarMarkDay(x.GoPointer(), DayVar)

}

var xCalendarSelectDay func(uintptr, *glib.DateTime)

// Switches to @date's year and month and select its day.
func (x *Calendar) SelectDay(DateVar *glib.DateTime) {

	xCalendarSelectDay(x.GoPointer(), DateVar)

}

var xCalendarSetShowDayNames func(uintptr, bool)

// Sets whether the calendar shows day names.
func (x *Calendar) SetShowDayNames(ValueVar bool) {

	xCalendarSetShowDayNames(x.GoPointer(), ValueVar)

}

var xCalendarSetShowHeading func(uintptr, bool)

// Sets whether the calendar should show a heading.
//
// The heading contains the current year and month as well as
// buttons for changing both.
func (x *Calendar) SetShowHeading(ValueVar bool) {

	xCalendarSetShowHeading(x.GoPointer(), ValueVar)

}

var xCalendarSetShowWeekNumbers func(uintptr, bool)

// Sets whether week numbers are shown in the calendar.
func (x *Calendar) SetShowWeekNumbers(ValueVar bool) {

	xCalendarSetShowWeekNumbers(x.GoPointer(), ValueVar)

}

var xCalendarUnmarkDay func(uintptr, uint)

// Removes the visual marker from a particular day.
func (x *Calendar) UnmarkDay(DayVar uint) {

	xCalendarUnmarkDay(x.GoPointer(), DayVar)

}

func (c *Calendar) GoPointer() uintptr {
	return c.Ptr
}

func (c *Calendar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the user selects a day.
func (x *Calendar) ConnectDaySelected(cb *func(Calendar)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "day-selected", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Calendar{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "day-selected", cbRefPtr)
}

// Emitted when the user switched to the next month.
func (x *Calendar) ConnectNextMonth(cb *func(Calendar)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "next-month", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Calendar{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "next-month", cbRefPtr)
}

// Emitted when user switched to the next year.
func (x *Calendar) ConnectNextYear(cb *func(Calendar)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "next-year", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Calendar{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "next-year", cbRefPtr)
}

// Emitted when the user switched to the previous month.
func (x *Calendar) ConnectPrevMonth(cb *func(Calendar)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "prev-month", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Calendar{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "prev-month", cbRefPtr)
}

// Emitted when user switched to the previous year.
func (x *Calendar) ConnectPrevYear(cb *func(Calendar)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "prev-year", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Calendar{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "prev-year", cbRefPtr)
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Calendar) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Calendar) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Calendar) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Calendar) ResetState(StateVar AccessibleState) {

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
func (x *Calendar) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Calendar) UpdatePropertyValue(NPropertiesVar int, PropertiesVar []AccessibleProperty, ValuesVar []gobject.Value) {

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
func (x *Calendar) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Calendar) UpdateRelationValue(NRelationsVar int, RelationsVar []AccessibleRelation, ValuesVar []gobject.Value) {

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
func (x *Calendar) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Calendar) UpdateStateValue(NStatesVar int, StatesVar []AccessibleState, ValuesVar []gobject.Value) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Calendar) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xCalendarGLibType, lib, "gtk_calendar_get_type")

	core.PuregoSafeRegister(&xNewCalendar, lib, "gtk_calendar_new")

	core.PuregoSafeRegister(&xCalendarClearMarks, lib, "gtk_calendar_clear_marks")
	core.PuregoSafeRegister(&xCalendarGetDate, lib, "gtk_calendar_get_date")
	core.PuregoSafeRegister(&xCalendarGetDayIsMarked, lib, "gtk_calendar_get_day_is_marked")
	core.PuregoSafeRegister(&xCalendarGetShowDayNames, lib, "gtk_calendar_get_show_day_names")
	core.PuregoSafeRegister(&xCalendarGetShowHeading, lib, "gtk_calendar_get_show_heading")
	core.PuregoSafeRegister(&xCalendarGetShowWeekNumbers, lib, "gtk_calendar_get_show_week_numbers")
	core.PuregoSafeRegister(&xCalendarMarkDay, lib, "gtk_calendar_mark_day")
	core.PuregoSafeRegister(&xCalendarSelectDay, lib, "gtk_calendar_select_day")
	core.PuregoSafeRegister(&xCalendarSetShowDayNames, lib, "gtk_calendar_set_show_day_names")
	core.PuregoSafeRegister(&xCalendarSetShowHeading, lib, "gtk_calendar_set_show_heading")
	core.PuregoSafeRegister(&xCalendarSetShowWeekNumbers, lib, "gtk_calendar_set_show_week_numbers")
	core.PuregoSafeRegister(&xCalendarUnmarkDay, lib, "gtk_calendar_unmark_day")

}
