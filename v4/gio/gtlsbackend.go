// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Provides an interface for describing TLS-related types.
type TlsBackendInterface struct {
	_ structs.HostLayout

	GIface uintptr
}

func (x *TlsBackendInterface) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// TLS (Transport Layer Security, aka SSL) and DTLS backend.
type TlsBackend interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetCertificateType() types.GType
	GetClientConnectionType() types.GType
	GetDefaultDatabase() *TlsDatabase
	GetDtlsClientConnectionType() types.GType
	GetDtlsServerConnectionType() types.GType
	GetFileDatabaseType() types.GType
	GetServerConnectionType() types.GType
	SetDefaultDatabase(DatabaseVar *TlsDatabase)
	SupportsDtls() bool
	SupportsTls() bool
}

var xTlsBackendGLibType func() types.GType

func TlsBackendGLibType() types.GType {
	return xTlsBackendGLibType()
}

type TlsBackendBase struct {
	Ptr uintptr
}

func (x *TlsBackendBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *TlsBackendBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Gets the #GType of @backend's #GTlsCertificate implementation.
func (x *TlsBackendBase) GetCertificateType() types.GType {

	cret := XGTlsBackendGetCertificateType(x.GoPointer())
	return cret
}

// Gets the #GType of @backend's #GTlsClientConnection implementation.
func (x *TlsBackendBase) GetClientConnectionType() types.GType {

	cret := XGTlsBackendGetClientConnectionType(x.GoPointer())
	return cret
}

// Gets the default #GTlsDatabase used to verify TLS connections.
func (x *TlsBackendBase) GetDefaultDatabase() *TlsDatabase {
	var cls *TlsDatabase

	cret := XGTlsBackendGetDefaultDatabase(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &TlsDatabase{}
	cls.Ptr = cret
	return cls
}

// Gets the #GType of @backend’s #GDtlsClientConnection implementation.
func (x *TlsBackendBase) GetDtlsClientConnectionType() types.GType {

	cret := XGTlsBackendGetDtlsClientConnectionType(x.GoPointer())
	return cret
}

// Gets the #GType of @backend’s #GDtlsServerConnection implementation.
func (x *TlsBackendBase) GetDtlsServerConnectionType() types.GType {

	cret := XGTlsBackendGetDtlsServerConnectionType(x.GoPointer())
	return cret
}

// Gets the #GType of @backend's #GTlsFileDatabase implementation.
func (x *TlsBackendBase) GetFileDatabaseType() types.GType {

	cret := XGTlsBackendGetFileDatabaseType(x.GoPointer())
	return cret
}

// Gets the #GType of @backend's #GTlsServerConnection implementation.
func (x *TlsBackendBase) GetServerConnectionType() types.GType {

	cret := XGTlsBackendGetServerConnectionType(x.GoPointer())
	return cret
}

// Set the default #GTlsDatabase used to verify TLS connections
//
// Any subsequent call to g_tls_backend_get_default_database() will return
// the database set in this call.  Existing databases and connections are not
// modified.
//
// Setting a %NULL default database will reset to using the system default
// database as if g_tls_backend_set_default_database() had never been called.
func (x *TlsBackendBase) SetDefaultDatabase(DatabaseVar *TlsDatabase) {

	XGTlsBackendSetDefaultDatabase(x.GoPointer(), DatabaseVar.GoPointer())

}

// Checks if DTLS is supported. DTLS support may not be available even if TLS
// support is available, and vice-versa.
func (x *TlsBackendBase) SupportsDtls() bool {

	cret := XGTlsBackendSupportsDtls(x.GoPointer())
	return cret
}

// Checks if TLS is supported; if this returns %FALSE for the default
// #GTlsBackend, it means no "real" TLS backend is available.
func (x *TlsBackendBase) SupportsTls() bool {

	cret := XGTlsBackendSupportsTls(x.GoPointer())
	return cret
}

var XGTlsBackendGetCertificateType func(uintptr) types.GType
var XGTlsBackendGetClientConnectionType func(uintptr) types.GType
var XGTlsBackendGetDefaultDatabase func(uintptr) uintptr
var XGTlsBackendGetDtlsClientConnectionType func(uintptr) types.GType
var XGTlsBackendGetDtlsServerConnectionType func(uintptr) types.GType
var XGTlsBackendGetFileDatabaseType func(uintptr) types.GType
var XGTlsBackendGetServerConnectionType func(uintptr) types.GType
var XGTlsBackendSetDefaultDatabase func(uintptr, uintptr)
var XGTlsBackendSupportsDtls func(uintptr) bool
var XGTlsBackendSupportsTls func(uintptr) bool

const (
	// Extension point for TLS functionality via #GTlsBackend.
	// See [Extending GIO][extending-gio].
	TLS_BACKEND_EXTENSION_POINT_NAME string = "gio-tls-backend"
)

var xTlsBackendGetDefault func() uintptr

// Gets the default #GTlsBackend for the system.
func TlsBackendGetDefault() *TlsBackendBase {
	var cls *TlsBackendBase

	cret := xTlsBackendGetDefault()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TlsBackendBase{}
	cls.Ptr = cret
	return cls
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xTlsBackendGetDefault, lib, "g_tls_backend_get_default")

	core.PuregoSafeRegister(&xTlsBackendGLibType, lib, "g_tls_backend_get_type")

	core.PuregoSafeRegister(&XGTlsBackendGetCertificateType, lib, "g_tls_backend_get_certificate_type")
	core.PuregoSafeRegister(&XGTlsBackendGetClientConnectionType, lib, "g_tls_backend_get_client_connection_type")
	core.PuregoSafeRegister(&XGTlsBackendGetDefaultDatabase, lib, "g_tls_backend_get_default_database")
	core.PuregoSafeRegister(&XGTlsBackendGetDtlsClientConnectionType, lib, "g_tls_backend_get_dtls_client_connection_type")
	core.PuregoSafeRegister(&XGTlsBackendGetDtlsServerConnectionType, lib, "g_tls_backend_get_dtls_server_connection_type")
	core.PuregoSafeRegister(&XGTlsBackendGetFileDatabaseType, lib, "g_tls_backend_get_file_database_type")
	core.PuregoSafeRegister(&XGTlsBackendGetServerConnectionType, lib, "g_tls_backend_get_server_connection_type")
	core.PuregoSafeRegister(&XGTlsBackendSetDefaultDatabase, lib, "g_tls_backend_set_default_database")
	core.PuregoSafeRegister(&XGTlsBackendSupportsDtls, lib, "g_tls_backend_supports_dtls")
	core.PuregoSafeRegister(&XGTlsBackendSupportsTls, lib, "g_tls_backend_supports_tls")

}
