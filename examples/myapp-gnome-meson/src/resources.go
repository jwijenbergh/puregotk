package main

import (
	_ "embed"
	"path"
)

const (
	AppID      = "io.github.puregotk.MyAppGnomeMeson"
	AppVersion = "0.1.0"
)

//go:embed myapp-gnome-meson.gresource
var ResourceContents []byte

var (
	AppPath = path.Join("/io", "github", "puregotk", "MyAppGnomeMeson")

	ResourceWindowUIPath = path.Join(AppPath, "window.ui")
)
