// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Interface for icons that can be loaded as a stream.
type LoadableIconIface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *LoadableIconIface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Extends the #GIcon interface and adds the ability to
// load icons from streams.
type LoadableIcon interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	Load(SizeVar int, TypeVar string, CancellableVar *Cancellable) *InputStream
	LoadAsync(SizeVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr)
	LoadFinish(ResVar AsyncResult, TypeVar string) *InputStream
}

var xLoadableIconGLibType func() types.GType

func LoadableIconGLibType() types.GType {
	return xLoadableIconGLibType()
}

type LoadableIconBase struct {
	Ptr uintptr
}

func (x *LoadableIconBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *LoadableIconBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Loads a loadable icon. For the asynchronous version of this function,
// see g_loadable_icon_load_async().
func (x *LoadableIconBase) Load(SizeVar int, TypeVar string, CancellableVar *Cancellable) (*InputStream, error) {
	var cls *InputStream
	var cerr *glib.Error

	cret := XGLoadableIconLoad(x.GoPointer(), SizeVar, TypeVar, CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &InputStream{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

// Loads an icon asynchronously. To finish this function, see
// g_loadable_icon_load_finish(). For the synchronous, blocking
// version of this function, see g_loadable_icon_load().
func (x *LoadableIconBase) LoadAsync(SizeVar int, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGLoadableIconLoadAsync(x.GoPointer(), SizeVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Finishes an asynchronous icon load started in g_loadable_icon_load_async().
func (x *LoadableIconBase) LoadFinish(ResVar AsyncResult, TypeVar string) (*InputStream, error) {
	var cls *InputStream
	var cerr *glib.Error

	cret := XGLoadableIconLoadFinish(x.GoPointer(), ResVar.GoPointer(), TypeVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &InputStream{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var XGLoadableIconLoad func(uintptr, int, string, uintptr, **glib.Error) uintptr
var XGLoadableIconLoadAsync func(uintptr, int, uintptr, uintptr, uintptr)
var XGLoadableIconLoadFinish func(uintptr, uintptr, string, **glib.Error) uintptr

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xLoadableIconGLibType, lib, "g_loadable_icon_get_type")

	core.PuregoSafeRegister(&XGLoadableIconLoad, lib, "g_loadable_icon_load")
	core.PuregoSafeRegister(&XGLoadableIconLoadAsync, lib, "g_loadable_icon_load_async")
	core.PuregoSafeRegister(&XGLoadableIconLoadFinish, lib, "g_loadable_icon_load_finish")

}
