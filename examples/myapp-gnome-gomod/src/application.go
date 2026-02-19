package main

import (
	"runtime"
	"unsafe"

	"github.com/jwijenbergh/puregotk/v4/adw"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

var (
	gTypeApplication gobject.Type
)

type Application struct {
	adw.Application

	window *Window
}

func NewApplication(FirstPropertyNameVar string, varArgs ...interface{}) Application {
	obj := gobject.NewObject(gTypeApplication, FirstPropertyNameVar, varArgs...)

	var v Application
	obj.Cast(&v)

	return v
}

func init() {
	var appClassInit gobject.ClassInitFunc = func(tc *gobject.TypeClass, u uintptr) {
		objClass := (*gobject.ObjectClass)(unsafe.Pointer(tc))

		objClass.OverrideConstructed(func(o *gobject.Object) {
			parentObjClass := (*gobject.ObjectClass)(unsafe.Pointer(tc.PeekParent()))

			parentObjClass.GetConstructed()(o)

			var parent adw.Application
			o.Cast(&parent)

			app := &Application{
				Application: parent,
			}

			var pinner runtime.Pinner
			pinner.Pin(app)

			var cleanupCallback glib.DestroyNotify = func(data uintptr) {
				pinner.Unpin()
			}
			o.SetDataFull(dataKeyGoInstance, uintptr(unsafe.Pointer(app)), &cleanupCallback)
		})

		applicationClass := (*gio.ApplicationClass)(unsafe.Pointer(tc))

		applicationClass.OverrideActivate(func(a *gio.Application) {
			myApp := (*Application)(unsafe.Pointer(a.GetData(dataKeyGoInstance)))

			if myApp.window != nil {
				myApp.window.ApplicationWindow.Present()
				return
			}

			var app gtk.Application
			a.Cast(&app)

			obj := NewWindow("application", app)

			myApp.window = (*Window)(unsafe.Pointer(obj.GetData(dataKeyGoInstance)))

			myApp.Application.AddWindow(&myApp.window.ApplicationWindow.Window)
			myApp.window.ApplicationWindow.Present()
		})
	}

	var appInstanceInit gobject.InstanceInitFunc = func(ti *gobject.TypeInstance, tc *gobject.TypeClass) {}

	var appParentQuery gobject.TypeQuery
	gobject.NewTypeQuery(adw.ApplicationGLibType(), &appParentQuery)

	gTypeApplication = gobject.TypeRegisterStaticSimple(
		appParentQuery.Type,
		"MyAppGnomeGomodApplication",
		appParentQuery.ClassSize,
		&appClassInit,
		appParentQuery.InstanceSize,
		&appInstanceInit,
		0,
	)
}
