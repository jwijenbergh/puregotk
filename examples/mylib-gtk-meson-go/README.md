# mylib-gtk-meson-go

This example shows how to generate Go bindings from the custom GTK library created in [../mylib-gtk-meson](../mylib-gtk-meson/). It uses puregotk's code generation pipeline to create type-safe Go bindings from the library's GIR (GObject Introspection Repository) file.

First, make sure the `mylib-gtk-meson` library is built (see [../mylib-gtk-meson](../mylib-gtk-meson/) for instructions). Then copy the resulting GIR file from [../mylib-gtk-meson/\_build/src/MyLibGtkMeson-0.1.gir](../mylib-gtk-meson/_build/src/MyLibGtkMeson-0.1.gir) to [internal/gir/spec/MyLibGtkMeson-0.1.gir](./internal/gir/spec/MyLibGtkMeson-0.1.gir); this GIR file is what the bindings are generated from.

To generate the Go bindings from this GIR file, run:

```shell
./gen.sh
```

The generated Go bindings can then be imported from [mylibgtkmeson](./mylibgtkmeson/) and used from other Go applications, as demonstrated in [../myapp-gnome-meson](../myapp-gnome-meson/) and [../myapp-gnome-gomod](../myapp-gnome-gomod/).
