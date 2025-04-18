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

type SimpleProxyResolverClass struct {
	_ structs.HostLayout

	ParentClass uintptr
}

func (x *SimpleProxyResolverClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type SimpleProxyResolverPrivate struct {
	_ structs.HostLayout
}

func (x *SimpleProxyResolverPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GSimpleProxyResolver is a simple #GProxyResolver implementation
// that handles a single default proxy, multiple URI-scheme-specific
// proxies, and a list of hosts that proxies should not be used for.
//
// #GSimpleProxyResolver is never the default proxy resolver, but it
// can be used as the base class for another proxy resolver
// implementation, or it can be created and used manually, such as
// with g_socket_client_set_proxy_resolver().
type SimpleProxyResolver struct {
	gobject.Object
}

var xSimpleProxyResolverGLibType func() types.GType

func SimpleProxyResolverGLibType() types.GType {
	return xSimpleProxyResolverGLibType()
}

func SimpleProxyResolverNewFromInternalPtr(ptr uintptr) *SimpleProxyResolver {
	cls := &SimpleProxyResolver{}
	cls.Ptr = ptr
	return cls
}

var xSimpleProxyResolverSetDefaultProxy func(uintptr, string)

// Sets the default proxy on @resolver, to be used for any URIs that
// don't match #GSimpleProxyResolver:ignore-hosts or a proxy set
// via g_simple_proxy_resolver_set_uri_proxy().
//
// If @default_proxy starts with "socks://",
// #GSimpleProxyResolver will treat it as referring to all three of
// the socks5, socks4a, and socks4 proxy types.
func (x *SimpleProxyResolver) SetDefaultProxy(DefaultProxyVar string) {

	xSimpleProxyResolverSetDefaultProxy(x.GoPointer(), DefaultProxyVar)

}

var xSimpleProxyResolverSetIgnoreHosts func(uintptr, []string)

// Sets the list of ignored hosts.
//
// See #GSimpleProxyResolver:ignore-hosts for more details on how the
// @ignore_hosts argument is interpreted.
func (x *SimpleProxyResolver) SetIgnoreHosts(IgnoreHostsVar []string) {

	xSimpleProxyResolverSetIgnoreHosts(x.GoPointer(), IgnoreHostsVar)

}

var xSimpleProxyResolverSetUriProxy func(uintptr, string, string)

// Adds a URI-scheme-specific proxy to @resolver; URIs whose scheme
// matches @uri_scheme (and which don't match
// #GSimpleProxyResolver:ignore-hosts) will be proxied via @proxy.
//
// As with #GSimpleProxyResolver:default-proxy, if @proxy starts with
// "socks://", #GSimpleProxyResolver will treat it
// as referring to all three of the socks5, socks4a, and socks4 proxy
// types.
func (x *SimpleProxyResolver) SetUriProxy(UriSchemeVar string, ProxyVar string) {

	xSimpleProxyResolverSetUriProxy(x.GoPointer(), UriSchemeVar, ProxyVar)

}

func (c *SimpleProxyResolver) GoPointer() uintptr {
	return c.Ptr
}

func (c *SimpleProxyResolver) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Checks if @resolver can be used on this system. (This is used
// internally; g_proxy_resolver_get_default() will only return a proxy
// resolver that returns %TRUE for this method.)
func (x *SimpleProxyResolver) IsSupported() bool {

	cret := XGProxyResolverIsSupported(x.GoPointer())
	return cret
}

// Looks into the system proxy configuration to determine what proxy,
// if any, to use to connect to @uri. The returned proxy URIs are of
// the form `&lt;protocol&gt;://[user[:password]@]host:port` or
// `direct://`, where &lt;protocol&gt; could be http, rtsp, socks
// or other proxying protocol.
//
// If you don't know what network protocol is being used on the
// socket, you should use `none` as the URI protocol.
// In this case, the resolver might still return a generic proxy type
// (such as SOCKS), but would not return protocol-specific proxy types
// (such as http).
//
// `direct://` is used when no proxy is needed.
// Direct connection should not be attempted unless it is part of the
// returned array of proxies.
func (x *SimpleProxyResolver) Lookup(UriVar string, CancellableVar *Cancellable) ([]string, error) {
	var cerr *glib.Error

	cret := XGProxyResolverLookup(x.GoPointer(), UriVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Asynchronous lookup of proxy. See g_proxy_resolver_lookup() for more
// details.
func (x *SimpleProxyResolver) LookupAsync(UriVar string, CancellableVar *Cancellable, CallbackVar *AsyncReadyCallback, UserDataVar uintptr) {

	XGProxyResolverLookupAsync(x.GoPointer(), UriVar, CancellableVar.GoPointer(), glib.NewCallback(CallbackVar), UserDataVar)

}

// Call this function to obtain the array of proxy URIs when
// g_proxy_resolver_lookup_async() is complete. See
// g_proxy_resolver_lookup() for more details.
func (x *SimpleProxyResolver) LookupFinish(ResultVar AsyncResult) ([]string, error) {
	var cerr *glib.Error

	cret := XGProxyResolverLookupFinish(x.GoPointer(), ResultVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xSimpleProxyResolverNew func(string, []string) uintptr

// Creates a new #GSimpleProxyResolver. See
// #GSimpleProxyResolver:default-proxy and
// #GSimpleProxyResolver:ignore-hosts for more details on how the
// arguments are interpreted.
func SimpleProxyResolverNew(DefaultProxyVar string, IgnoreHostsVar []string) *ProxyResolverBase {
	var cls *ProxyResolverBase

	cret := xSimpleProxyResolverNew(DefaultProxyVar, IgnoreHostsVar)

	if cret == 0 {
		return nil
	}
	cls = &ProxyResolverBase{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSimpleProxyResolverGLibType, lib, "g_simple_proxy_resolver_get_type")

	core.PuregoSafeRegister(&xSimpleProxyResolverSetDefaultProxy, lib, "g_simple_proxy_resolver_set_default_proxy")
	core.PuregoSafeRegister(&xSimpleProxyResolverSetIgnoreHosts, lib, "g_simple_proxy_resolver_set_ignore_hosts")
	core.PuregoSafeRegister(&xSimpleProxyResolverSetUriProxy, lib, "g_simple_proxy_resolver_set_uri_proxy")

	core.PuregoSafeRegister(&xSimpleProxyResolverNew, lib, "g_simple_proxy_resolver_new")

}
