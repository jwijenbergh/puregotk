package main

import (
	"runtime"
	"unsafe"

	"github.com/jwijenbergh/puregotk/v4/adw"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

import "C"

var (
	gTypeMyLibGtkMesonMainApplicationWindow gobject.Type
)

//export mylib_gtk_meson_main_application_window_get_type
func mylib_gtk_meson_main_application_window_get_type() C.ulong {
	return C.ulong(gTypeMyLibGtkMesonMainApplicationWindow)
}

type myLibGtkMesonMainApplicationWindow struct {
	adw.ApplicationWindow

	buttonTest   *gtk.Button
	toastOverlay *adw.ToastOverlay
}

func init() {
	var classInit gobject.ClassInitFunc = func(tc *gobject.TypeClass, u uintptr) {
		typeClass := (*gtk.WidgetClass)(unsafe.Pointer(tc))
		typeClass.SetTemplateFromResource(resourceWindowUIPath)

		typeClass.BindTemplateChildFull("button_test", false, 0)
		typeClass.BindTemplateChildFull("toast_overlay", false, 0)

		gobject.SignalNewv(
			"button-test-clicked",
			gTypeMyLibGtkMesonMainApplicationWindow,
			gobject.GSignalRunFirstValue,
			nil,
			nil,
			0,
			nil,
			types.GType(gobject.TypeNoneVal),
			0,
			nil,
		)

		objClass := (*gobject.ObjectClass)(unsafe.Pointer(tc))

		objClass.OverrideConstructed(func(o *gobject.Object) {
			parentObjClass := (*gobject.ObjectClass)(unsafe.Pointer(tc.PeekParent()))

			parentObjClass.GetConstructed()(o)

			var parent adw.ApplicationWindow
			o.Cast(&parent)

			parent.InitTemplate()

			var buttonTest gtk.Button
			parent.Widget.GetTemplateChild(
				gTypeMyLibGtkMesonMainApplicationWindow,
				"button_test",
			).Cast(&buttonTest)

			var toastOverlay adw.ToastOverlay
			parent.Widget.GetTemplateChild(
				gTypeMyLibGtkMesonMainApplicationWindow,
				"toast_overlay",
			).Cast(&toastOverlay)

			w := &myLibGtkMesonMainApplicationWindow{
				ApplicationWindow: parent,
				buttonTest:        &buttonTest,
				toastOverlay:      &toastOverlay,
			}

			var pinner runtime.Pinner
			pinner.Pin(w)

			var cleanupCallback glib.DestroyNotify = func(data uintptr) {
				pinner.Unpin()
			}
			o.SetDataFull(dataKeyGoInstance, uintptr(unsafe.Pointer(w)), &cleanupCallback)

			onButtonTestClicked := func(gtk.Button) {
				gobject.SignalEmit(
					o,
					gobject.SignalLookup("button-test-clicked", gTypeMyLibGtkMesonMainApplicationWindow),
					0,
				)
			}

			buttonTest.ConnectClicked(&onButtonTestClicked)
		})

		objClass.OverrideSetProperty(func(o *gobject.Object, u uint32, v *gobject.Value, ps *gobject.ParamSpec) {
			switch u {
			case propertyIdTestButtonSensitive:
				w := (*myLibGtkMesonMainApplicationWindow)(unsafe.Pointer(o.GetData(dataKeyGoInstance)))

				w.buttonTest.SetSensitive(v.GetBoolean())
			}
		})

		objClass.OverrideGetProperty(func(o *gobject.Object, u uint32, v *gobject.Value, ps *gobject.ParamSpec) {
			switch u {
			case propertyIdTestButtonSensitive:
				w := (*myLibGtkMesonMainApplicationWindow)(unsafe.Pointer(o.GetData(dataKeyGoInstance)))

				v.SetBoolean(w.buttonTest.IsSensitive())
			}
		})

		pspec := gobject.NewParamSpecBoolean(
			"test-button-sensitive",
			"Test Button Sensitive",
			"Whether the test button is sensitive",
			true,
			gobject.GParamReadwriteValue,
		)

		objClass.InstallProperty(propertyIdTestButtonSensitive, pspec)
	}

	var instanceInit gobject.InstanceInitFunc = func(ti *gobject.TypeInstance, tc *gobject.TypeClass) {}

	var parentQuery gobject.TypeQuery
	gobject.NewTypeQuery(adw.ApplicationWindowGLibType(), &parentQuery)

	gTypeMyLibGtkMesonMainApplicationWindow = gobject.TypeRegisterStaticSimple(
		parentQuery.Type,
		"MyLibGtkMesonMainApplicationWindow",
		parentQuery.ClassSize,
		&classInit,
		parentQuery.InstanceSize,
		&instanceInit,
		0,
	)
}

func (w *myLibGtkMesonMainApplicationWindow) showToast(message string) {
	w.toastOverlay.AddToast(adw.NewToast(message))
}

//export mylib_gtk_meson_main_application_window_show_toast
func mylib_gtk_meson_main_application_window_show_toast(window unsafe.Pointer, message *C.char) {
	w := (*myLibGtkMesonMainApplicationWindow)(unsafe.Pointer(gobject.ObjectNewFromInternalPtr(uintptr(window)).GetData(dataKeyGoInstance)))

	w.showToast(C.GoString(message))
}
