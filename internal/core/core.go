// package core implements core functionality for the generated files
// this core lib is imported by the generated code
package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"

	"github.com/jwijenbergh/purego"
)

func PuregoSafeRegister(fptr interface{}, handle uintptr, name string) error {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		return err
	}
	purego.RegisterFunc(fptr, sym)
	return nil
}

// paths to where the shared object files should be located
// this is unique per architecture
// Debian/Ubuntu has it split into specific arch folder, Fedora is just /usr/lib64
// see:
// https://fedora.pkgs.org/38/fedora-x86_64/gtk4-4.10.1-1.fc38.x86_64.rpm.html
// https://fedora.pkgs.org/38/fedora-aarch64/gtk4-4.10.1-1.fc38.aarch64.rpm.html
// https://ubuntu.pkgs.org/23.04/ubuntu-main-amd64/libgtk-4-1_4.10.1+ds-2ubuntu1_amd64.deb.html
// https://ubuntu.pkgs.org/23.04/ubuntu-main-arm64/libgtk-4-1_4.10.1+ds-2ubuntu1_arm64.deb.html
var paths = map[string][]string{
	"amd64": {"/usr/lib/x86_64-linux-gnu/", "/usr/lib64/", "/usr/lib/"},
	"arm64": {"/usr/lib/aarch64-linux-gnu/", "/usr/lib64/", "/usr/lib/"},
}

// names is a lookup from library names to shared object filenames
var names = map[string][]string{
	"ADW":        {"libadwaita-1.so"},
	"CAIRO":      {"libcairo.so"},
	"GDKPIXBUF":  {"libgdk_pixbuf-2.0.so"},
	"GIO":        {"libgio-2.0.so"},
	"GLIB":       {"libglib-2.0.so"},
	"GMODULE":    {"libgmodule-2.0.so"},
	"GOBJECT":    {"libgobject-2.0.so"},
	"GRAPHENE":   {"libgraphene-1.0.so"},
	"GTK":        {"libgtk-4.so"},
	"PANGO":      {"libpango-1.0.so"},
	"PANGOCAIRO": {"libpangocairo-1.0.so"},
}

// aliases are lib aliases
// for example when we load the GSK lib, we should get the functions from the GTK shared library
var aliases = map[string]string{
	"GSK": "GTK",
	"GDK": "GTK",
}

// pkgConfNames is a lookup from library names to pkg-config library names
// TODO: Get this from the package name in the gir files with a fallback to gtk's package name
var pkgConfNames = map[string]string{
	"ADW":        "libadwaita-1",
	"CAIRO":      "cairo",
	"GDKPIXBUF":  "gdk-pixbuf-2.0",
	"GIO":        "gio-2.0",
	"GLIB":       "glib-2.0",
	"GMODULE":    "gmodule-2.0",
	"GOBJECT":    "gobject-2.0",
	"GRAPHENE":   "graphene-gobject-1.0",
	"GTK":        "gtk4",
	"PANGO":      "pango",
	"PANGOCAIRO": "pangocairo",
}

// findSo tries to find a shared object from a path and a library name
// It does this by mapping the library name to a suitable shared object filename and then trying some suffixes
func findSo(path string, name string) string {
	for _, n := range names[name] {
		suffixes := []string{"", ".0", ".1", ".2"}
		fn := filepath.Join(path, n)
		for _, s := range suffixes {
			if _, err := os.Stat(fn + s); err == nil {
				return fn + s
			}
		}
	}
	return ""
}

// findPkgConf finds a shared object file with pkg-config
// it does this by running pkg-config --libs-only-L libname
// and then it loops over the directories returned and finds a suitable one
func findPkgConf(name string) string {
	cmd := exec.Command("pkg-config", "--libs-only-L", pkgConfNames[name])
	var out, outerr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outerr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pkg-config, failed with: %v and stderr: %s\n", err, outerr.String())
		return ""
	}
	outs := strings.Split(out.String(), "-L")
	for _, v := range outs {
		c := strings.TrimSpace(v)
		if c == "" {
			continue
		}
		g := findSo(c, name)
		if g != "" {
			return g
		}
	}
	return ""
}

// GetPath gets a shared object file from a library name
// it does it in the following order
// see if PUREGOTK_LIBNAME_PATH is set (full path to the lib)
// - e.g. PUREGOTK_GTK_PATH
// see if PUREGOTK_LIB_FOLDER is set (root folder where to look for libs)
// go over the hardcoded paths
// find a library name with pkg-config
// panic if failed
// TODO: Hardcore a library shared object with linker -X flag
// This is useful for packaging
func GetPath(name string) string {
	// resolve alias
	if v, ok := aliases[name]; ok {
		name = v
	}
	// try to get from env var
	ev := fmt.Sprintf("PUREGOTK_%s_PATH", name)
	if v := os.Getenv(ev); v != "" {
		return v
	}

	// Or if a general folder is set where everywhere is located, return that
	ep := os.Getenv("PUREGOTK_LIB_FOLDER")
	if ep != "" {
		g := findSo(ep, name)
		if g == "" {
			panic(fmt.Sprintf("Could not find lib: %s, at path: %s with env: %s", name, ep, "PUREGOTK_FOLDER"))
		}
		return g
	}

	// fallback to lookup a path if no env var is found
	gp, ok := paths[runtime.GOARCH]
	if ok {
		// try to loop over paths
		for _, p := range gp {
			g := findSo(p, name)
			if g != "" {
				return g
			}

		}
	}
	// last effort: pkg-config
	g := findPkgConf(name)
	if g != "" {
		return g
	}

	panic(fmt.Sprintf("Path for library: %s not found. Please set the path to this library shared object file manually with env variable: %s or PUREGOTK_LIB_FOLDER. Or make sure pkg-config is setup correctly", strings.ToLower(name), ev))
}

// GoString copies a char* to a Go string.
// This function was copied from purego
func GoString(c uintptr) string {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	if ptr == nil {
		return ""
	}
	var length int
	for {
		if *(*byte)(unsafe.Add(ptr, uintptr(length))) == '\x00' {
			break
		}
		length++
	}
	return string(unsafe.Slice((*byte)(ptr), length))
}
