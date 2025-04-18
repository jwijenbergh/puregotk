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

// vtable for a #GDtlsClientConnection implementation.
type DtlsClientConnectionInterface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *DtlsClientConnectionInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GDtlsClientConnection is the client-side subclass of
// #GDtlsConnection, representing a client-side DTLS connection.
type DtlsClientConnection interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetAcceptedCas() *glib.List
	GetServerIdentity() *SocketConnectableBase
	GetValidationFlags() TlsCertificateFlags
	SetServerIdentity(IdentityVar SocketConnectable)
	SetValidationFlags(FlagsVar TlsCertificateFlags)
}

var xDtlsClientConnectionGLibType func() types.GType

func DtlsClientConnectionGLibType() types.GType {
	return xDtlsClientConnectionGLibType()
}

type DtlsClientConnectionBase struct {
	Ptr uintptr
}

func (x *DtlsClientConnectionBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *DtlsClientConnectionBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the list of distinguished names of the Certificate Authorities
// that the server will accept certificates from. This will be set
// during the TLS handshake if the server requests a certificate.
// Otherwise, it will be %NULL.
//
// Each item in the list is a #GByteArray which contains the complete
// subject DN of the certificate authority.
func (x *DtlsClientConnectionBase) GetAcceptedCas() *glib.List {

	cret := XGDtlsClientConnectionGetAcceptedCas(x.GoPointer())
	return cret
}

// Gets @conn's expected server identity
func (x *DtlsClientConnectionBase) GetServerIdentity() *SocketConnectableBase {
	var cls *SocketConnectableBase

	cret := XGDtlsClientConnectionGetServerIdentity(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &SocketConnectableBase{}
	cls.Ptr = cret
	return cls
}

// Gets @conn's validation flags
func (x *DtlsClientConnectionBase) GetValidationFlags() TlsCertificateFlags {

	cret := XGDtlsClientConnectionGetValidationFlags(x.GoPointer())
	return cret
}

// Sets @conn's expected server identity, which is used both to tell
// servers on virtual hosts which certificate to present, and also
// to let @conn know what name to look for in the certificate when
// performing %G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
func (x *DtlsClientConnectionBase) SetServerIdentity(IdentityVar SocketConnectable) {

	XGDtlsClientConnectionSetServerIdentity(x.GoPointer(), IdentityVar.GoPointer())

}

// Sets @conn's validation flags, to override the default set of
// checks performed when validating a server certificate. By default,
// %G_TLS_CERTIFICATE_VALIDATE_ALL is used.
func (x *DtlsClientConnectionBase) SetValidationFlags(FlagsVar TlsCertificateFlags) {

	XGDtlsClientConnectionSetValidationFlags(x.GoPointer(), FlagsVar)

}

var XGDtlsClientConnectionGetAcceptedCas func(uintptr) *glib.List
var XGDtlsClientConnectionGetServerIdentity func(uintptr) uintptr
var XGDtlsClientConnectionGetValidationFlags func(uintptr) TlsCertificateFlags
var XGDtlsClientConnectionSetServerIdentity func(uintptr, uintptr)
var XGDtlsClientConnectionSetValidationFlags func(uintptr, TlsCertificateFlags)

var xDtlsClientConnectionNew func(uintptr, uintptr, **glib.Error) uintptr

// Creates a new #GDtlsClientConnection wrapping @base_socket which is
// assumed to communicate with the server identified by @server_identity.
func DtlsClientConnectionNew(BaseSocketVar DatagramBased, ServerIdentityVar SocketConnectable) (*DtlsClientConnectionBase, error) {
	var cls *DtlsClientConnectionBase
	var cerr *glib.Error

	cret := xDtlsClientConnectionNew(BaseSocketVar.GoPointer(), ServerIdentityVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &DtlsClientConnectionBase{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xDtlsClientConnectionNew, lib, "g_dtls_client_connection_new")

	core.PuregoSafeRegister(&xDtlsClientConnectionGLibType, lib, "g_dtls_client_connection_get_type")

	core.PuregoSafeRegister(&XGDtlsClientConnectionGetAcceptedCas, lib, "g_dtls_client_connection_get_accepted_cas")
	core.PuregoSafeRegister(&XGDtlsClientConnectionGetServerIdentity, lib, "g_dtls_client_connection_get_server_identity")
	core.PuregoSafeRegister(&XGDtlsClientConnectionGetValidationFlags, lib, "g_dtls_client_connection_get_validation_flags")
	core.PuregoSafeRegister(&XGDtlsClientConnectionSetServerIdentity, lib, "g_dtls_client_connection_set_server_identity")
	core.PuregoSafeRegister(&XGDtlsClientConnectionSetValidationFlags, lib, "g_dtls_client_connection_set_validation_flags")

}
