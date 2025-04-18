// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Class structure for #GSocketAddressEnumerator.
type SocketAddressEnumeratorClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *SocketAddressEnumeratorClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GSocketAddressEnumerator is an enumerator type for #GSocketAddress
// instances. It is returned by enumeration functions such as
// g_socket_connectable_enumerate(), which returns a #GSocketAddressEnumerator
// to list each #GSocketAddress which could be used to connect to that
// #GSocketConnectable.
//
// Enumeration is typically a blocking operation, so the asynchronous methods
// g_socket_address_enumerator_next_async() and
// g_socket_address_enumerator_next_finish() should be used where possible.
//
// Each #GSocketAddressEnumerator can only be enumerated once. Once
// g_socket_address_enumerator_next() has returned %NULL, further
// enumeration with that #GSocketAddressEnumerator is not possible, and it can
// be unreffed.
type SocketAddressEnumerator struct {
	gobject.Object
}

var xSocketAddressEnumeratorGLibType func() types.GType

func SocketAddressEnumeratorGLibType() types.GType {
	return xSocketAddressEnumeratorGLibType()
}

func SocketAddressEnumeratorNewFromInternalPtr(ptr uintptr) *SocketAddressEnumerator {
	cls := &SocketAddressEnumerator{}
	cls.Ptr = ptr
	return cls
}

var xSocketAddressEnumeratorNext func(uintptr, uintptr, **glib.Error) uintptr

// Retrieves the next #GSocketAddress from @enumerator. Note that this
// may block for some amount of time. (Eg, a #GNetworkAddress may need
// to do a DNS lookup before it can return an address.) Use
// g_socket_address_enumerator_next_async() if you need to avoid
// blocking.
//
// If @enumerator is expected to yield addresses, but for some reason
// is unable to (eg, because of a DNS error), then the first call to
// g_socket_address_enumerator_next() will return an appropriate error
// in *@error. However, if the first call to
// g_socket_address_enumerator_next() succeeds, then any further
// internal errors (other than @cancellable being triggered) will be
// ignored.
func (x *SocketAddressEnumerator) Next(CancellableVar *Cancellable) (*SocketAddress, error) {
	var cls *SocketAddress
	var cerr *glib.Error

	cret := xSocketAddressEnumeratorNext(x.GoPointer(), CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &SocketAddress{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xSocketAddressEnumeratorNextAsync func(uintptr, uintptr, uintptr, uintptr)

// Asynchronously retrieves the next #GSocketAddress from @enumerator
// and then calls @callback, which must call
// g_socket_address_enumerator_next_finish() to get the result.
//
// It is an error to call this multiple times before the previous callback has finished.
func (x *SocketAddressEnumerator) NextAsync(CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	xSocketAddressEnumeratorNextAsync(x.GoPointer(), CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

var xSocketAddressEnumeratorNextFinish func(uintptr, uintptr, **glib.Error) uintptr

// Retrieves the result of a completed call to
// g_socket_address_enumerator_next_async(). See
// g_socket_address_enumerator_next() for more information about
// error handling.
func (x *SocketAddressEnumerator) NextFinish(ResultVar AsyncResult) (*SocketAddress, error) {
	var cls *SocketAddress
	var cerr *glib.Error

	cret := xSocketAddressEnumeratorNextFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &SocketAddress{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

func (c *SocketAddressEnumerator) GoPointer() uintptr {
	return c.Ptr
}

func (c *SocketAddressEnumerator) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSocketAddressEnumeratorGLibType, lib, "g_socket_address_enumerator_get_type")

	core.PuregoSafeRegister(&xSocketAddressEnumeratorNext, lib, "g_socket_address_enumerator_next")
	core.PuregoSafeRegister(&xSocketAddressEnumeratorNextAsync, lib, "g_socket_address_enumerator_next_async")
	core.PuregoSafeRegister(&xSocketAddressEnumeratorNextFinish, lib, "g_socket_address_enumerator_next_finish")

}
