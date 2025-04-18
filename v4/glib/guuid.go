// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xUuidStringIsValid func(string) bool

// Parses the string @str and verify if it is a UUID.
//
// The function accepts the following syntax:
//
// - simple forms (e.g. `f81d4fae-7dec-11d0-a765-00a0c91e6bf6`)
//
// Note that hyphens are required within the UUID string itself,
// as per the aforementioned RFC.
func UuidStringIsValid(StrVar string) bool {

	cret := xUuidStringIsValid(StrVar)
	return cret
}

var xUuidStringRandom func() string

// Generates a random UUID (RFC 4122 version 4) as a string. It has the same
// randomness guarantees as #GRand, so must not be used for cryptographic
// purposes such as key generation, nonces, salts or one-time pads.
func UuidStringRandom() string {

	cret := xUuidStringRandom()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xUuidStringIsValid, lib, "g_uuid_string_is_valid")
	core.PuregoSafeRegister(&xUuidStringRandom, lib, "g_uuid_string_random")

}
