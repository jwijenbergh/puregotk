module github.com/jwijenbergh/puregotk/examples/mylib-gtk-meson-go

go 1.25.0

require (
	github.com/jwijenbergh/purego v0.0.0-20251017112123-b71757b9ba42
	github.com/jwijenbergh/puregotk v0.0.0-00010101000000-000000000000
)

replace github.com/jwijenbergh/puregotk v0.0.0-00010101000000-000000000000 => ../..
