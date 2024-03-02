//go:build darwin || freebsd || (linux && (amd64 || arm64))

package glib

import (
	"fmt"
	"reflect"

	"github.com/jwijenbergh/purego"
)

func unrefCallback(fnPtr interface{}) error {
	val := reflect.ValueOf(fnPtr)
	if val.IsNil() {
		return fmt.Errorf("function pointer must not be nil")
	}
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Func {
		return fmt.Errorf("type must be a function pointer")
	}
	cbPtr := reflect.ValueOf(fnPtr).Pointer()
	refPtr, ok := GetCallback(cbPtr)
	if !ok {
		return purego.UnrefCallbackFnPtr(fnPtr)
	}
	defer func() {
		callbacks.Lock()
		delete(callbacks.refs, cbPtr)
		callbacks.Unlock()
	}()
	return purego.UnrefCallback(refPtr)
}
