// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A type corresponding to the appropriate struct type for the stat()
// system call, depending on the platform and/or compiler being used.
//
// See g_stat() for more information.
type StatBuf struct {
	_ structs.HostLayout
}

func (x *StatBuf) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xAccess func(string, int) int

// A wrapper for the POSIX access() function. This function is used to
// test a pathname for one or several of read, write or execute
// permissions, or just existence.
//
// On Windows, the file protection mechanism is not at all POSIX-like,
// and the underlying function in the C library only checks the
// FAT-style READONLY attribute, and does not look at the ACL of a
// file at all. This function is this in practise almost useless on
// Windows. Software that needs to handle file permissions on Windows
// more exactly should use the Win32 API.
//
// See your C library manual for more details about access().
func Access(FilenameVar string, ModeVar int) int {

	cret := xAccess(FilenameVar, ModeVar)
	return cret
}

var xChdir func(string) int

// A wrapper for the POSIX chdir() function. The function changes the
// current directory of the process to @path.
//
// See your C library manual for more details about chdir().
func Chdir(PathVar string) int {

	cret := xChdir(PathVar)
	return cret
}

var xClose func(int, **Error) bool

// This wraps the close() call; in case of error, %errno will be
// preserved, but the error will also be stored as a #GError in @error.
//
// Besides using #GError, there is another major reason to prefer this
// function over the call provided by the system; on Unix, it will
// attempt to correctly handle %EINTR, which has platform-specific
// semantics.
func Close(FdVar int) (bool, error) {
	var cerr *Error

	cret := xClose(FdVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xRmdir func(string) int

// A wrapper for the POSIX rmdir() function. The rmdir() function
// deletes a directory from the filesystem.
//
// See your C library manual for more details about how rmdir() works
// on your system.
func Rmdir(FilenameVar string) int {

	cret := xRmdir(FilenameVar)
	return cret
}

var xUnlink func(string) int

// A wrapper for the POSIX unlink() function. The unlink() function
// deletes a name from the filesystem. If this was the last link to the
// file and no processes have it opened, the diskspace occupied by the
// file is freed.
//
// See your C library manual for more details about unlink(). Note
// that on Windows, it is in general not possible to delete files that
// are open to some process, or mapped into memory.
func Unlink(FilenameVar string) int {

	cret := xUnlink(FilenameVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xAccess, lib, "g_access")
	core.PuregoSafeRegister(&xChdir, lib, "g_chdir")
	core.PuregoSafeRegister(&xClose, lib, "g_close")
	core.PuregoSafeRegister(&xRmdir, lib, "g_rmdir")
	core.PuregoSafeRegister(&xUnlink, lib, "g_unlink")

}
