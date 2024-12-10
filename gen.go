package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jwijenbergh/puregotk/internal/gir/pass"
	"github.com/jwijenbergh/puregotk/internal/gir/util"
)

//go:generate go run gen.go

func main() {
	dir := "v4"
	os.RemoveAll(dir)
	var girs []string
	filepath.Walk("internal/gir/spec", func(path string, f os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".gir") {
			return nil
		}
		girs = append(girs, path)
		return nil
	})
	p, err := pass.New(girs)
	if err != nil {
		panic(err)
	}
	// collect basic type info
	p.First()

	// Create the template
	gotemp, err := template.New("go").Funcs(template.FuncMap{"conv": util.ConvertArgs, "convc": util.ConvertArgsComma}).ParseFiles("templates/go")
	if err != nil {
		panic(err)
	}

	// Write go files by making the second pass
	p.Second(dir, gotemp)

	// Finally copy some extra code that we want in the API
	data, err := os.ReadFile("templates/gobject")
	if err == nil {
		os.WriteFile("v4/gobject/more.go", data, 0o644)
	}
	data, err = os.ReadFile("templates/gtype")
	if err == nil {
		mkerr := os.MkdirAll("v4/gobject/types", 0o755)
		if mkerr != nil {
			panic(mkerr)
		}
		os.WriteFile("v4/gobject/types/types.go", data, 0o644)
	}
	data, err = os.ReadFile("templates/glib")
	if err == nil {
		os.WriteFile("v4/glib/more.go", data, 0o644)
	}
	data, err = os.ReadFile("templates/glib_sysv")
	if err == nil {
		os.WriteFile("v4/glib/more_sysv.go", data, 0o644)
	}
	data, err = os.ReadFile("templates/glib_windows")
	if err == nil {
		os.WriteFile("v4/glib/more_windows.go", data, 0o644)
	}
	data, err = os.ReadFile("templates/glib_other")
	if err == nil {
		os.WriteFile("v4/glib/more_other.go", data, 0o644)
	}
}
