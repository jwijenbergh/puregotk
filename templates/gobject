package gobject

import (
	"reflect"

	"github.com/jwijenbergh/puregotk/v4/glib"
)

type Ptr interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
}

func ConvertPtr(a interface{}) *uintptr {
	if a == nil || (reflect.ValueOf(a).Kind() == reflect.Ptr && reflect.ValueOf(a).IsNil()) {
		return nil
	}
	ptr := reflect.ValueOf(a).Elem()
	v, ok := ptr.Interface().(Ptr)
	if !ok {
		panic("not valid")
	}
	g := v.GoPointer()
	return &g
}

func IncreaseRef(a uintptr) {
	xObjectRefSink(a)
}

func SignalConnect(a uintptr, b string, c uintptr) uint32 {
	return xSignalConnectData(a, b, c, 0, 0, 0)
}

func (o Object) Cast(v Ptr) {
	v.SetGoPointer(o.GoPointer())
}

func (o Object) ConnectSignal(signal string, cb *func()) uint32 {
	return SignalConnect(o.GoPointer(), signal, glib.NewCallback(cb))
}

func (o Object) DisconnectSignal(handler uint32) {
	SignalHandlerDisconnect(&o, handler)
}

// types
const (
	TypeInvalidVal           Type = 0
	TypeNoneVal                   = 1 << 2
	TypeInterfaceVal              = 2 << 2
	TypeCharVal                   = 3 << 2
	TypeUcharVal                  = 4 << 2
	TypeBooleanVal                = 5 << 2
	TypeIntVal                    = 6 << 2
	TypeUintVal                   = 7 << 2
	TypeLongVal                   = 8 << 2
	TypeUlongVal                  = 9 << 2
	TypeInt64Val                  = 10 << 2
	TypeUint64Val                 = 11 << 2
	TypeEnumVal                   = 12 << 2
	TypeFlagsVal                  = 13 << 2
	TypeFloatVal                  = 14 << 2
	TypeDoubleVal                 = 15 << 2
	TypeStringVal                 = 16 << 2
	TypePointerVal                = 17 << 2
	TypeBoxedVal                  = 18 << 2
	TypeParamVal                  = 19 << 2
	TypeObjectVal                 = 20 << 2
	TypeReservedGLibLastVal       = 31 << 2
	TypeReservedBseFirstVal       = 32 << 2
	TypeReservedBseLastVal        = 48 << 2
	TypeReservedUserFirstVal      = 49 << 2
)
