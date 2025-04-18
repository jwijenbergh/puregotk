// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

var xBuilderErrorQuark func() glib.Quark

func BuilderErrorQuark() glib.Quark {

	cret := xBuilderErrorQuark()
	return cret
}

var xConstraintVflParserErrorQuark func() glib.Quark

func ConstraintVflParserErrorQuark() glib.Quark {

	cret := xConstraintVflParserErrorQuark()
	return cret
}

var xCssParserErrorQuark func() glib.Quark

func CssParserErrorQuark() glib.Quark {

	cret := xCssParserErrorQuark()
	return cret
}

var xIconThemeErrorQuark func() glib.Quark

func IconThemeErrorQuark() glib.Quark {

	cret := xIconThemeErrorQuark()
	return cret
}

var xRecentManagerErrorQuark func() glib.Quark

func RecentManagerErrorQuark() glib.Quark {

	cret := xRecentManagerErrorQuark()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xBuilderErrorQuark, lib, "gtk_builder_error_quark")
	core.PuregoSafeRegister(&xConstraintVflParserErrorQuark, lib, "gtk_constraint_vfl_parser_error_quark")
	core.PuregoSafeRegister(&xCssParserErrorQuark, lib, "gtk_css_parser_error_quark")
	core.PuregoSafeRegister(&xIconThemeErrorQuark, lib, "gtk_icon_theme_error_quark")
	core.PuregoSafeRegister(&xRecentManagerErrorQuark, lib, "gtk_recent_manager_error_quark")

}
