package glib

import (
	"fmt"
	"sync"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var callbacks = struct {
	sync.RWMutex
	refs map[uintptr]uintptr
}{
	refs: make(map[uintptr]uintptr),
}

// GetCallback retrives a callback reference by value.
// Users should not need to call this.
func GetCallback(cbPtr uintptr) (uintptr, bool) {
	callbacks.RLock()
	defer callbacks.RUnlock()
	refPtr, ok := callbacks.refs[cbPtr]
	return refPtr, ok
}

// SaveCallback saves a reference to the callback value.
// Users should not need to call this.
func SaveCallback(cbPtr uintptr, refPtr uintptr) {
	callbacks.Lock()
	callbacks.refs[cbPtr] = refPtr
	callbacks.Unlock()
}

// UnrefCallbackValue unreferences the provided callback by reflect.value to free a purego slot
//
// NOTE: Windows does not support unreferencing callbacks, so on that platform this operation is
// a NOOP, callback memory is never freed, and there is a limit on maximum total callbacks.
// See the purego documentation for further details.
func UnrefCallback(fnPtr interface{}) error {
	return unrefCallback(fnPtr)
}

// NewCallback is an alias to purego.NewCallback
func NewCallback(fnPtr interface{}) uintptr {
	return purego.NewCallbackFnPtr(fnPtr)
}

func (e *Error) Error() string {
	return fmt.Sprintf("Gtk reported an error with message: '%s', domain: '%v' and code: '%v'", e.MessageGo(), e.Domain, e.Code)
}

func (e *Error) MessageGo() string {
	return core.GoString(e.Message)
}
