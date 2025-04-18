// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Renders a spin button in a cell
//
// `GtkCellRendererSpin` renders text in a cell like `GtkCellRendererText` from
// which it is derived. But while `GtkCellRendererText` offers a simple entry to
// edit the text, `GtkCellRendererSpin` offers a `GtkSpinButton` widget. Of course,
// that means that the text has to be parseable as a floating point number.
//
// The range of the spinbutton is taken from the adjustment property of the
// cell renderer, which can be set explicitly or mapped to a column in the
// tree model, like all properties of cell renders. `GtkCellRendererSpin`
// also has properties for the `GtkCellRendererSpin:climb-rate` and the number
// of `GtkCellRendererSpin:digits` to display. Other `GtkSpinButton` properties
// can be set in a handler for the `GtkCellRenderer::editing-started` signal.
//
// The `GtkCellRendererSpin` cell renderer was added in GTK 2.10.
type CellRendererSpin struct {
	CellRendererText
}

var xCellRendererSpinGLibType func() types.GType

func CellRendererSpinGLibType() types.GType {
	return xCellRendererSpinGLibType()
}

func CellRendererSpinNewFromInternalPtr(ptr uintptr) *CellRendererSpin {
	cls := &CellRendererSpin{}
	cls.Ptr = ptr
	return cls
}

var xNewCellRendererSpin func() uintptr

// Creates a new `GtkCellRendererSpin`.
func NewCellRendererSpin() *CellRendererSpin {
	var cls *CellRendererSpin

	cret := xNewCellRendererSpin()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellRendererSpin{}
	cls.Ptr = cret
	return cls
}

func (c *CellRendererSpin) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellRendererSpin) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xCellRendererSpinGLibType, lib, "gtk_cell_renderer_spin_get_type")

	core.PuregoSafeRegister(&xNewCellRendererSpin, lib, "gtk_cell_renderer_spin_new")

}
