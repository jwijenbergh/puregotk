package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jwijenbergh/puregotk/pkg/gir/pass"
	"github.com/jwijenbergh/puregotk/pkg/gir/util"
)

//go:generate go run gen.go

func main() {
	dir := "."
	os.RemoveAll(dir)
	var girs []string
	var localNamespaces []string
	filepath.Walk("internal/gir/spec", func(path string, f os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".gir") {
			return nil
		}
		girs = append(girs, path)

		// Extract namespace from filename (e.g., "MyLibMeson-1.0.gir" -> "MyLibMeson")
		base := filepath.Base(path)
		if idx := strings.Index(base, "-"); idx != -1 {
			localNamespaces = append(localNamespaces, base[:idx])
		}

		return nil
	})

	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", "github.com/jwijenbergh/puregotk")
	output, err := cmd.Output()
	if err != nil {
		panic("puregotk dependency not found: " + err.Error())
	}
	puregotk := strings.TrimSpace(string(output))

	filepath.Walk(filepath.Join(puregotk, "internal/gir/spec"), func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".gir") {
			girs = append(girs, path)
		}
		return nil
	})

	p, err := pass.New(girs)
	if err != nil {
		panic(err)
	}
	// collect basic type info
	p.First()

	// Create the template
	gotemp, err := template.New("go").Funcs(template.FuncMap{
		"conv":     util.ConvertArgs,
		"convc":    util.ConvertArgsComma,
		"convcb":   util.ConvertCallbackArgs,
		"convcd":   util.ConvertArgsCommaDeref,
		"convd":    util.ConvertArgsDeref,
		"convcbne": util.ConvertCallbackArgsNoErr,
		"propsset": util.PropertyScalarSet,
		"propsget": util.PropertyScalarGet,
		"propvset": util.PropertyVectorSet,
		"propvget": util.PropertyVectorGet,
	}).ParseFiles(filepath.Join(puregotk, "templates/go"))
	if err != nil {
		panic(err)
	}

	// Only generate code for local namespaces
	original := p.Parsed
	p.Parsed = nil
	for _, repo := range original {
		for _, ns := range repo.Namespaces {
			for _, localNs := range localNamespaces {
				if ns.Name == localNs {
					p.Parsed = append(p.Parsed, repo)
					break
				}
			}
		}
	}

	// Write go files by making the second pass
	p.Second(dir, gotemp)
}
