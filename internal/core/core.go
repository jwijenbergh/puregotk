// package core implements core functionality for the generated files
// this core lib is imported by the generated code
package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
// TODO: This needs to be split in architecture and more paths should be added
var paths = []string{
	"/usr/lib/x86_64-linux-gnu/",
}

// names is a lookup from library names to shared object filenames
var names = map[string][]string{
	"ADW":        {"libadwaita-1.so"},
	"CAIRO":      {"libcairo.so"},
	"GDK":        {"libgtk-4.so"},
	"GDKPIXBUF":  {"libgdk_pixbuf-2.0.so"},
	"GIO":        {"libgio-2.0.so"},
	"GLIB":       {"libglib-2.0.so"},
	"GMODULE":    {"libgmodule-2.0.so"},
	"GOBJECT":    {"libgobject-2.0.so"},
	"GRAPHENE":   {"libgraphene-1.0.so"},
	"GSK":        {"libgtk-4.so"},
	"GTK":        {"libgtk-4.so"},
	"PANGO":      {"libpango-1.0.so"},
	"PANGOCAIRO": {"libpangocairo-1.0.so"},
}

// pkgConfNames is a lookup from library names to pkg-config library names
// TODO: Get this from the package name in the gir files with a fallback to gtk's package name
var pkgConfNames = map[string]string{
	"ADW":        "libadwaita-1",
	"CAIRO":      "cairo",
	"GDK":        "gtk4",
	"GDKPIXBUF":  "gdk-pixbuf-2.0",
	"GIO":        "gio-2.0",
	"GLIB":       "glib-2.0",
	"GMODULE":    "gmodule-2.0",
	"GOBJECT":    "gobject-2.0",
	"GRAPHENE":   "graphene-gobject-1.0",
	"GSK":        "gtk4",
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
				return fn+s
			}
		}
	}
	return ""
}

// findPkgConf finds a shared object file with pkg-config
// it does this by running pkg-config --libs-only-L libname
// and then it loops over the directories returned and finds a suitable one
func findPkgConf(name string) string {
	// try to get with pkg-config
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
// see if PUREGOTK_LIBNAME_PATH is set
// go over the hardcoded paths
// find a library name with pkg-config
// panic if failed
// TODO: Hardcore a library shared object with linker -X flag
// This is useful for packaging
func GetPath(name string) string {
	// try to get from env var
	ev := fmt.Sprintf("PUREGOTK_%s_PATH", name)
	if v := os.Getenv(ev); v != "" {
		return v
	}

	// try to loop over paths
	for _, p := range paths {
		g := findSo(p, name)
		if g != "" {
			return g
		}

	}
	g := findPkgConf(name)
	if g != "" {
		return g
	}

	panic(fmt.Sprintf("Path for library: %s not found. Please set the path to this library shared object file manually with env variable: %s. Or make sure pkg-config is setup correctly", strings.ToLower(name), ev))
}
