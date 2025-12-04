package main

//go:generate sh -c "if [ -z \"$FLATPAK_ID\" ]; then cd .. && GOWORK=off go tool github.com/dennwc/flatpak-go-mod --json .; fi"

import (
	_ "embed"

	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/pojntfx/go-gettext/pkg/i18n"
)

func init() {
	if err := i18n.BindI18n(GettextPackage, LocaleDir); err != nil {
		panic(err)
	}

	resource, err := gio.NewResourceFromData(glib.NewBytes(ResourceContents, uint(len(ResourceContents))))
	if err != nil {
		panic(err)
	}
	gio.ResourcesRegister(resource)
}

func main() {}
