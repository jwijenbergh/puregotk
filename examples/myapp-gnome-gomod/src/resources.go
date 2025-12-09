package main

import (
	_ "embed"
	"path"
)

const (
	AppID      = "io.github.puregotk.MyAppGnomeGomod"
	AppVersion = "0.1.0"
)

//go:generate sh -c "blueprint-compiler batch-compile . . *.blp && glib-compile-resources *.gresource.xml"
//go:embed myapp-gnome-gomod.gresource
var ResourceContents []byte

var (
	AppPath = path.Join("/io", "github", "puregotk", "MyAppGnomeGomod")

	ResourceWindowUIPath = path.Join(AppPath, "window.ui")
)
