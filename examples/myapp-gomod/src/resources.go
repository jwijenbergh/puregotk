package main

import (
	_ "embed"
	"path"
)

const (
	AppID      = "io.github.puregotk.MyAppGomod"
	AppVersion = "0.1.0"
)

//go:generate sh -c "blueprint-compiler batch-compile . . *.blp && glib-compile-resources *.gresource.xml"
//go:embed myapp-gomod.gresource
var ResourceContents []byte

var (
	AppPath = path.Join("/io", "github", "puregotk", "MyAppGomod")

	ResourceWindowUIPath = path.Join(AppPath, "window.ui")
)
