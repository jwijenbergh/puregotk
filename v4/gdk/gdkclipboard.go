// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// The `GdkClipboard` object represents data shared between applications or
// inside an application.
//
// To get a `GdkClipboard` object, use [method@Gdk.Display.get_clipboard] or
// [method@Gdk.Display.get_primary_clipboard]. You can find out about the data
// that is currently available in a clipboard using
// [method@Gdk.Clipboard.get_formats].
//
// To make text or image data available in a clipboard, use
// [method@Gdk.Clipboard.set_text] or [method@Gdk.Clipboard.set_texture].
// For other data, you can use [method@Gdk.Clipboard.set_content], which
// takes a [class@Gdk.ContentProvider] object.
//
// To read textual or image data from a clipboard, use
// [method@Gdk.Clipboard.read_text_async] or
// [method@Gdk.Clipboard.read_texture_async]. For other data, use
// [method@Gdk.Clipboard.read_async], which provides a `GInputStream` object.
type Clipboard struct {
	gobject.Object
}

var xClipboardGLibType func() types.GType

func ClipboardGLibType() types.GType {
	return xClipboardGLibType()
}

func ClipboardNewFromInternalPtr(ptr uintptr) *Clipboard {
	cls := &Clipboard{}
	cls.Ptr = ptr
	return cls
}

var xClipboardGetContent func(uintptr) uintptr

// Returns the `GdkContentProvider` currently set on @clipboard.
//
// If the @clipboard is empty or its contents are not owned by the
// current process, %NULL will be returned.
func (x *Clipboard) GetContent() *ContentProvider {
	var cls *ContentProvider

	cret := xClipboardGetContent(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ContentProvider{}
	cls.Ptr = cret
	return cls
}

var xClipboardGetDisplay func(uintptr) uintptr

// Gets the `GdkDisplay` that the clipboard was created for.
func (x *Clipboard) GetDisplay() *Display {
	var cls *Display

	cret := xClipboardGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Display{}
	cls.Ptr = cret
	return cls
}

var xClipboardGetFormats func(uintptr) *ContentFormats

// Gets the formats that the clipboard can provide its current contents in.
func (x *Clipboard) GetFormats() *ContentFormats {

	cret := xClipboardGetFormats(x.GoPointer())
	return cret
}

var xClipboardIsLocal func(uintptr) bool

// Returns if the clipboard is local.
//
// A clipboard is considered local if it was last claimed
// by the running application.
//
// Note that [method@Gdk.Clipboard.get_content] may return %NULL
// even on a local clipboard. In this case the clipboard is empty.
func (x *Clipboard) IsLocal() bool {

	cret := xClipboardIsLocal(x.GoPointer())
	return cret
}

var xClipboardReadAsync func(uintptr, []string, int, uintptr, uintptr, uintptr)

// Asynchronously requests an input stream to read the @clipboard's
// contents from.
//
// When the operation is finished @callback will be called. You must then
// call [method@Gdk.Clipboard.read_finish] to get the result of the operation.
//
// The clipboard will choose the most suitable mime type from the given list
// to fulfill the request, preferring the ones listed first.
func (x *Clipboard) ReadAsync(MimeTypesVar []string, IoPriorityVar int, CancellableVar *gio.Cancellable, CallbackVar *gio.AsyncReadyCallback, UserDataVar uintptr) {

	xClipboardReadAsync(x.GoPointer(), MimeTypesVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xClipboardReadFinish func(uintptr, uintptr, string, **glib.Error) uintptr

// Finishes an asynchronous clipboard read.
//
// See [method@Gdk.Clipboard.read_async].
func (x *Clipboard) ReadFinish(ResultVar gio.AsyncResult, OutMimeTypeVar string) (*gio.InputStream, error) {
	var cls *gio.InputStream
	var cerr *glib.Error

	cret := xClipboardReadFinish(x.GoPointer(), ResultVar.GoPointer(), OutMimeTypeVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &gio.InputStream{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xClipboardReadTextAsync func(uintptr, uintptr, uintptr, uintptr)

// Asynchronously request the @clipboard contents converted to a string.
//
// When the operation is finished @callback will be called. You must then
// call [method@Gdk.Clipboard.read_text_finish] to get the result.
//
// This is a simple wrapper around [method@Gdk.Clipboard.read_value_async].
// Use that function or [method@Gdk.Clipboard.read_async] directly if you
// need more control over the operation.
func (x *Clipboard) ReadTextAsync(CancellableVar *gio.Cancellable, CallbackVar *gio.AsyncReadyCallback, UserDataVar uintptr) {

	xClipboardReadTextAsync(x.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xClipboardReadTextFinish func(uintptr, uintptr, **glib.Error) string

// Finishes an asynchronous clipboard read.
//
// See [method@Gdk.Clipboard.read_text_async].
func (x *Clipboard) ReadTextFinish(ResultVar gio.AsyncResult) (string, error) {
	var cerr *glib.Error

	cret := xClipboardReadTextFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xClipboardReadTextureAsync func(uintptr, uintptr, uintptr, uintptr)

// Asynchronously request the @clipboard contents converted to a `GdkPixbuf`.
//
// When the operation is finished @callback will be called. You must then
// call [method@Gdk.Clipboard.read_texture_finish] to get the result.
//
// This is a simple wrapper around [method@Gdk.Clipboard.read_value_async].
// Use that function or [method@Gdk.Clipboard.read_async] directly if you
// need more control over the operation.
func (x *Clipboard) ReadTextureAsync(CancellableVar *gio.Cancellable, CallbackVar *gio.AsyncReadyCallback, UserDataVar uintptr) {

	xClipboardReadTextureAsync(x.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xClipboardReadTextureFinish func(uintptr, uintptr, **glib.Error) uintptr

// Finishes an asynchronous clipboard read.
//
// See [method@Gdk.Clipboard.read_texture_async].
func (x *Clipboard) ReadTextureFinish(ResultVar gio.AsyncResult) (*Texture, error) {
	var cls *Texture
	var cerr *glib.Error

	cret := xClipboardReadTextureFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &Texture{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xClipboardReadValueAsync func(uintptr, types.GType, int, uintptr, uintptr, uintptr)

// Asynchronously request the @clipboard contents converted to the given
// @type.
//
// When the operation is finished @callback will be called. You must then call
// [method@Gdk.Clipboard.read_value_finish] to get the resulting `GValue`.
//
// For local clipboard contents that are available in the given `GType`,
// the value will be copied directly. Otherwise, GDK will try to use
// [func@content_deserialize_async] to convert the clipboard's data.
func (x *Clipboard) ReadValueAsync(TypeVar types.GType, IoPriorityVar int, CancellableVar *gio.Cancellable, CallbackVar *gio.AsyncReadyCallback, UserDataVar uintptr) {

	xClipboardReadValueAsync(x.GoPointer(), TypeVar, IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xClipboardReadValueFinish func(uintptr, uintptr, **glib.Error) *gobject.Value

// Finishes an asynchronous clipboard read.
//
// See [method@Gdk.Clipboard.read_value_async].
func (x *Clipboard) ReadValueFinish(ResultVar gio.AsyncResult) (*gobject.Value, error) {
	var cerr *glib.Error

	cret := xClipboardReadValueFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xClipboardSet func(uintptr, types.GType, ...interface{})

// Sets the clipboard to contain the value collected from the given varargs.
//
// Values should be passed the same way they are passed to other value
// collecting APIs, such as [`method@GObject.Object.set`] or
// [`func@GObject.signal_emit`].
//
// ```c
// gdk_clipboard_set (clipboard, GTK_TYPE_STRING, "Hello World");
//
// gdk_clipboard_set (clipboard, GDK_TYPE_TEXTURE, some_texture);
// ```
func (x *Clipboard) Set(TypeVar types.GType, varArgs ...interface{}) {

	xClipboardSet(x.GoPointer(), TypeVar, varArgs...)

}

var xClipboardSetContent func(uintptr, uintptr) bool

// Sets a new content provider on @clipboard.
//
// The clipboard will claim the `GdkDisplay`'s resources and advertise
// these new contents to other applications.
//
// In the rare case of a failure, this function will return %FALSE. The
// clipboard will then continue reporting its old contents and ignore
// @provider.
//
// If the contents are read by either an external application or the
// @clipboard's read functions, @clipboard will select the best format to
// transfer the contents and then request that format from @provider.
func (x *Clipboard) SetContent(ProviderVar *ContentProvider) bool {

	cret := xClipboardSetContent(x.GoPointer(), ProviderVar.GoPointer())
	return cret
}

var xClipboardSetText func(uintptr, string)

// Puts the given @text into the clipboard.
func (x *Clipboard) SetText(TextVar string) {

	xClipboardSetText(x.GoPointer(), TextVar)

}

var xClipboardSetTexture func(uintptr, uintptr)

// Puts the given @texture into the clipboard.
func (x *Clipboard) SetTexture(TextureVar *Texture) {

	xClipboardSetTexture(x.GoPointer(), TextureVar.GoPointer())

}

var xClipboardSetValist func(uintptr, types.GType, []interface{})

// Sets the clipboard to contain the value collected from the given @args.
func (x *Clipboard) SetValist(TypeVar types.GType, ArgsVar []interface{}) {

	xClipboardSetValist(x.GoPointer(), TypeVar, ArgsVar)

}

var xClipboardSetValue func(uintptr, *gobject.Value)

// Sets the @clipboard to contain the given @value.
func (x *Clipboard) SetValue(ValueVar *gobject.Value) {

	xClipboardSetValue(x.GoPointer(), ValueVar)

}

var xClipboardStoreAsync func(uintptr, int, uintptr, uintptr, uintptr)

// Asynchronously instructs the @clipboard to store its contents remotely.
//
// If the clipboard is not local, this function does nothing but report success.
//
// The @callback must call [method@Gdk.Clipboard.store_finish].
//
// The purpose of this call is to preserve clipboard contents beyond the
// lifetime of an application, so this function is typically called on
// exit. Depending on the platform, the functionality may not be available
// unless a "clipboard manager" is running.
//
// This function is called automatically when a [class@Gtk.Application] is
// shut down, so you likely don't need to call it.
func (x *Clipboard) StoreAsync(IoPriorityVar int, CancellableVar *gio.Cancellable, CallbackVar *gio.AsyncReadyCallback, UserDataVar uintptr) {

	xClipboardStoreAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xClipboardStoreFinish func(uintptr, uintptr, **glib.Error) bool

// Finishes an asynchronous clipboard store.
//
// See [method@Gdk.Clipboard.store_async].
func (x *Clipboard) StoreFinish(ResultVar gio.AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := xClipboardStoreFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func (c *Clipboard) GoPointer() uintptr {
	return c.Ptr
}

func (c *Clipboard) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the clipboard changes ownership.
func (x *Clipboard) ConnectChanged(cb *func(Clipboard)) uint32 {
	cbPtr := uintptr(unsafe.Pointer(cb))
	if cbRefPtr, ok := glib.GetCallback(cbPtr); ok {
		return gobject.SignalConnect(x.GoPointer(), "changed", cbRefPtr)
	}

	fcb := func(clsPtr uintptr) {
		fa := Clipboard{}
		fa.Ptr = clsPtr
		cbFn := *cb

		cbFn(fa)

	}
	cbRefPtr := purego.NewCallback(fcb)
	glib.SaveCallback(cbPtr, cbRefPtr)
	return gobject.SignalConnect(x.GoPointer(), "changed", cbRefPtr)
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xClipboardGLibType, lib, "gdk_clipboard_get_type")

	core.PuregoSafeRegister(&xClipboardGetContent, lib, "gdk_clipboard_get_content")
	core.PuregoSafeRegister(&xClipboardGetDisplay, lib, "gdk_clipboard_get_display")
	core.PuregoSafeRegister(&xClipboardGetFormats, lib, "gdk_clipboard_get_formats")
	core.PuregoSafeRegister(&xClipboardIsLocal, lib, "gdk_clipboard_is_local")
	core.PuregoSafeRegister(&xClipboardReadAsync, lib, "gdk_clipboard_read_async")
	core.PuregoSafeRegister(&xClipboardReadFinish, lib, "gdk_clipboard_read_finish")
	core.PuregoSafeRegister(&xClipboardReadTextAsync, lib, "gdk_clipboard_read_text_async")
	core.PuregoSafeRegister(&xClipboardReadTextFinish, lib, "gdk_clipboard_read_text_finish")
	core.PuregoSafeRegister(&xClipboardReadTextureAsync, lib, "gdk_clipboard_read_texture_async")
	core.PuregoSafeRegister(&xClipboardReadTextureFinish, lib, "gdk_clipboard_read_texture_finish")
	core.PuregoSafeRegister(&xClipboardReadValueAsync, lib, "gdk_clipboard_read_value_async")
	core.PuregoSafeRegister(&xClipboardReadValueFinish, lib, "gdk_clipboard_read_value_finish")
	core.PuregoSafeRegister(&xClipboardSet, lib, "gdk_clipboard_set")
	core.PuregoSafeRegister(&xClipboardSetContent, lib, "gdk_clipboard_set_content")
	core.PuregoSafeRegister(&xClipboardSetText, lib, "gdk_clipboard_set_text")
	core.PuregoSafeRegister(&xClipboardSetTexture, lib, "gdk_clipboard_set_texture")
	core.PuregoSafeRegister(&xClipboardSetValist, lib, "gdk_clipboard_set_valist")
	core.PuregoSafeRegister(&xClipboardSetValue, lib, "gdk_clipboard_set_value")
	core.PuregoSafeRegister(&xClipboardStoreAsync, lib, "gdk_clipboard_store_async")
	core.PuregoSafeRegister(&xClipboardStoreFinish, lib, "gdk_clipboard_store_finish")

}
