//go:build !windows && !darwin && !freebsd && (!linux || (!amd64 && !arm64))

package glib

// unrefCallback on unsupported on other platforms
func unrefCallback(_ interface{}) error {
	panic("glib.UnrefCallback is unsupported on this platform")
}
