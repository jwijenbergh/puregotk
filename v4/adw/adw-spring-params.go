// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// Physical parameters of a spring for [class@SpringAnimation].
//
// Any spring can be described by three parameters: mass, stiffness and damping.
//
// An undamped spring will produce an oscillatory motion which will go on
// forever.
//
// The frequency and amplitude of the oscillations will be determined by the
// stiffness (how "strong" the spring is) and its mass (how much "inertia" it
// has).
//
// If damping is larger than 0, the amplitude of that oscillating motion will
// exponientally decrease over time. If that damping is strong enough that the
// spring can't complete a full oscillation, it's called an overdamped spring.
//
// If we the spring can oscillate, it's called an underdamped spring.
//
// The value between these two behaviors is called critical damping; a
// critically damped spring will comes to rest in the minimum possible time
// without producing oscillations.
//
// The damping can be replaced by damping ratio, which produces the following
// springs:
//
// * 0: an undamped spring.
// * Between 0 and 1: an underdamped spring.
// * 1: a critically damped spring.
// * Larger than 1: an overdamped spring.
//
// As such
type SpringParams struct {
	_ structs.HostLayout
}

var xSpringParamsGLibType func() types.GType

func SpringParamsGLibType() types.GType {
	return xSpringParamsGLibType()
}

func (x *SpringParams) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewSpringParams func(float64, float64, float64) *SpringParams

// Creates a new `AdwSpringParams` from @mass, @stiffness and @damping_ratio.
//
// The damping value is calculated from @damping_ratio and the other two
// parameters.
//
//   - If @damping_ratio is 0, the spring will not be damped and will oscillate
//     endlessly.
//   - If @damping_ratio is between 0 and 1, the spring is underdamped and will
//     always overshoot.
//   - If @damping_ratio is 1, the spring is critically damped and will reach its
//     resting position the quickest way possible.
//   - If @damping_ratio is larger than 1, the spring is overdamped and will reach
//     its resting position faster than it can complete an oscillation.
//
// [ctor@SpringParams.new_full] allows to pass a raw damping value instead.
func NewSpringParams(DampingRatioVar float64, MassVar float64, StiffnessVar float64) *SpringParams {

	cret := xNewSpringParams(DampingRatioVar, MassVar, StiffnessVar)
	return cret
}

var xNewSpringParamsFull func(float64, float64, float64) *SpringParams

// Creates a new `AdwSpringParams` from @mass, @stiffness and @damping.
//
// See [ctor@SpringParams.new] for a simplified constructor using damping ratio
// instead of @damping.
func NewSpringParamsFull(DampingVar float64, MassVar float64, StiffnessVar float64) *SpringParams {

	cret := xNewSpringParamsFull(DampingVar, MassVar, StiffnessVar)
	return cret
}

var xSpringParamsGetDamping func(uintptr) float64

// Gets the damping of @self.
func (x *SpringParams) GetDamping() float64 {

	cret := xSpringParamsGetDamping(x.GoPointer())
	return cret
}

var xSpringParamsGetDampingRatio func(uintptr) float64

// Gets the damping ratio of @self.
func (x *SpringParams) GetDampingRatio() float64 {

	cret := xSpringParamsGetDampingRatio(x.GoPointer())
	return cret
}

var xSpringParamsGetMass func(uintptr) float64

// Gets the mass of @self.
func (x *SpringParams) GetMass() float64 {

	cret := xSpringParamsGetMass(x.GoPointer())
	return cret
}

var xSpringParamsGetStiffness func(uintptr) float64

// Gets the stiffness of @self.
func (x *SpringParams) GetStiffness() float64 {

	cret := xSpringParamsGetStiffness(x.GoPointer())
	return cret
}

var xSpringParamsRef func(uintptr) *SpringParams

// Increases the reference count of @self.
func (x *SpringParams) Ref() *SpringParams {

	cret := xSpringParamsRef(x.GoPointer())
	return cret
}

var xSpringParamsUnref func(uintptr)

// Decreases the reference count of @self.
//
// If the last reference is dropped, the structure is freed.
func (x *SpringParams) Unref() {

	xSpringParamsUnref(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSpringParamsGLibType, lib, "adw_spring_params_get_type")

	core.PuregoSafeRegister(&xNewSpringParams, lib, "adw_spring_params_new")
	core.PuregoSafeRegister(&xNewSpringParamsFull, lib, "adw_spring_params_new_full")

	core.PuregoSafeRegister(&xSpringParamsGetDamping, lib, "adw_spring_params_get_damping")
	core.PuregoSafeRegister(&xSpringParamsGetDampingRatio, lib, "adw_spring_params_get_damping_ratio")
	core.PuregoSafeRegister(&xSpringParamsGetMass, lib, "adw_spring_params_get_mass")
	core.PuregoSafeRegister(&xSpringParamsGetStiffness, lib, "adw_spring_params_get_stiffness")
	core.PuregoSafeRegister(&xSpringParamsRef, lib, "adw_spring_params_ref")
	core.PuregoSafeRegister(&xSpringParamsUnref, lib, "adw_spring_params_unref")

}
