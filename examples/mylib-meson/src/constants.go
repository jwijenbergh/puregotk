package main

import (
	_ "embed"

	"path"
)

const (
	dataKeyGoInstance = "go_instance"

	propertyIdTestButtonSensitive = 1
)

var (
	appPath = path.Join("/io", "github", "puregotk", "MyLibMeson")

	resourceWindowUIPath = path.Join(appPath, "window.ui")
)

//go:embed mylib-meson.gresource
var ResourceContents []byte
