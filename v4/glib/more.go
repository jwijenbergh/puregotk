package glib

import (
	"fmt"

	"github.com/jwijenbergh/puregotk/internal/core"
)

func (e *Error) Error() string {
	return fmt.Sprintf("Gtk reported an error with message: '%s', Domain: '%v' and Code: '%v'", e.MessageGo(), e.Domain, e.Code)
}

func (e *Error) MessageGo() string {
	return core.GoString(e.Message)
}
