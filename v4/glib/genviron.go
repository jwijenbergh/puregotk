// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xEnvironGetenv func([]string, string) string

// Returns the value of the environment variable @variable in the
// provided list @envp.
func EnvironGetenv(EnvpVar []string, VariableVar string) string {

	cret := xEnvironGetenv(EnvpVar, VariableVar)
	return cret
}

var xEnvironSetenv func([]string, string, string, bool) []string

// Sets the environment variable @variable in the provided list
// @envp to @value.
func EnvironSetenv(EnvpVar []string, VariableVar string, ValueVar string, OverwriteVar bool) []string {

	cret := xEnvironSetenv(EnvpVar, VariableVar, ValueVar, OverwriteVar)
	return cret
}

var xEnvironUnsetenv func([]string, string) []string

// Removes the environment variable @variable from the provided
// environment @envp.
func EnvironUnsetenv(EnvpVar []string, VariableVar string) []string {

	cret := xEnvironUnsetenv(EnvpVar, VariableVar)
	return cret
}

var xGetEnviron func() []string

// Gets the list of environment variables for the current process.
//
// The list is %NULL terminated and each item in the list is of the
// form 'NAME=VALUE'.
//
// This is equivalent to direct access to the 'environ' global variable,
// except portable.
//
// The return value is freshly allocated and it should be freed with
// g_strfreev() when it is no longer needed.
func GetEnviron() []string {

	cret := xGetEnviron()
	return cret
}

var xGetenv func(string) string

// Returns the value of an environment variable.
//
// On UNIX, the name and value are byte strings which might or might not
// be in some consistent character set and encoding. On Windows, they are
// in UTF-8.
// On Windows, in case the environment variable's value contains
// references to other environment variables, they are expanded.
func Getenv(VariableVar string) string {

	cret := xGetenv(VariableVar)
	return cret
}

var xListenv func() []string

// Gets the names of all variables set in the environment.
//
// Programs that want to be portable to Windows should typically use
// this function and g_getenv() instead of using the environ array
// from the C library directly. On Windows, the strings in the environ
// array are in system codepage encoding, while in most of the typical
// use cases for environment variables in GLib-using programs you want
// the UTF-8 encoding that this function and g_getenv() provide.
func Listenv() []string {

	cret := xListenv()
	return cret
}

var xSetenv func(string, string, bool) bool

// Sets an environment variable. On UNIX, both the variable's name and
// value can be arbitrary byte strings, except that the variable's name
// cannot contain '='. On Windows, they should be in UTF-8.
//
// Note that on some systems, when variables are overwritten, the memory
// used for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling
// in UNIX is not thread-safe, and your program may crash if one thread
// calls g_setenv() while another thread is calling getenv(). (And note
// that many functions, such as gettext(), call getenv() internally.)
// This function is only safe to use at the very start of your program,
// before creating any other threads (or creating objects that create
// worker threads of their own).
//
// If you need to set up the environment for a child process, you can
// use g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that
// array directly to execvpe(), g_spawn_async(), or the like.
func Setenv(VariableVar string, ValueVar string, OverwriteVar bool) bool {

	cret := xSetenv(VariableVar, ValueVar, OverwriteVar)
	return cret
}

var xUnsetenv func(string)

// Removes an environment variable from the environment.
//
// Note that on some systems, when variables are overwritten, the
// memory used for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling
// in UNIX is not thread-safe, and your program may crash if one thread
// calls g_unsetenv() while another thread is calling getenv(). (And note
// that many functions, such as gettext(), call getenv() internally.) This
// function is only safe to use at the very start of your program, before
// creating any other threads (or creating objects that create worker
// threads of their own).
//
// If you need to set up the environment for a child process, you can
// use g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that
// array directly to execvpe(), g_spawn_async(), or the like.
func Unsetenv(VariableVar string) {

	xUnsetenv(VariableVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xEnvironGetenv, lib, "g_environ_getenv")
	core.PuregoSafeRegister(&xEnvironSetenv, lib, "g_environ_setenv")
	core.PuregoSafeRegister(&xEnvironUnsetenv, lib, "g_environ_unsetenv")
	core.PuregoSafeRegister(&xGetEnviron, lib, "g_get_environ")
	core.PuregoSafeRegister(&xGetenv, lib, "g_getenv")
	core.PuregoSafeRegister(&xListenv, lib, "g_listenv")
	core.PuregoSafeRegister(&xSetenv, lib, "g_setenv")
	core.PuregoSafeRegister(&xUnsetenv, lib, "g_unsetenv")

}
