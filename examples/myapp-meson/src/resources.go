package main

import (
	_ "embed"
	"path"
)

const (
	AppID      = "io.github.puregotk.MyAppMeson"
	AppVersion = "0.1.0"
)

//go:embed myapp-meson.gresource
var ResourceContents []byte

var (
	AppPath = path.Join("/io", "github", "puregotk", "MyAppMeson")

	ResourceWindowUIPath = path.Join(AppPath, "window.ui")
)
