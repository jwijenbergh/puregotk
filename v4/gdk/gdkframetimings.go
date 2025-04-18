// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"structs"
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject/types"
)

// A `GdkFrameTimings` object holds timing information for a single frame
// of the application’s displays.
//
// To retrieve `GdkFrameTimings` objects, use [method@Gdk.FrameClock.get_timings]
// or [method@Gdk.FrameClock.get_current_timings]. The information in
// `GdkFrameTimings` is useful for precise synchronization of video with
// the event or audio streams, and for measuring quality metrics for the
// application’s display, such as latency and jitter.
type FrameTimings struct {
	_ structs.HostLayout
}

var xFrameTimingsGLibType func() types.GType

func FrameTimingsGLibType() types.GType {
	return xFrameTimingsGLibType()
}

func (x *FrameTimings) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xFrameTimingsGetComplete func(uintptr) bool

// Returns whether @timings are complete.
//
// The timing information in a `GdkFrameTimings` is filled in
// incrementally as the frame as drawn and passed off to the
// window system for processing and display to the user. The
// accessor functions for `GdkFrameTimings` can return 0 to
// indicate an unavailable value for two reasons: either because
// the information is not yet available, or because it isn't
// available at all.
//
// Once this function returns %TRUE for a frame, you can be
// certain that no further values will become available and be
// stored in the `GdkFrameTimings`.
func (x *FrameTimings) GetComplete() bool {

	cret := xFrameTimingsGetComplete(x.GoPointer())
	return cret
}

var xFrameTimingsGetFrameCounter func(uintptr) int64

// Gets the frame counter value of the `GdkFrameClock` when
// this frame was drawn.
func (x *FrameTimings) GetFrameCounter() int64 {

	cret := xFrameTimingsGetFrameCounter(x.GoPointer())
	return cret
}

var xFrameTimingsGetFrameTime func(uintptr) int64

// Returns the frame time for the frame.
//
// This is the time value that is typically used to time
// animations for the frame. See [method@Gdk.FrameClock.get_frame_time].
func (x *FrameTimings) GetFrameTime() int64 {

	cret := xFrameTimingsGetFrameTime(x.GoPointer())
	return cret
}

var xFrameTimingsGetPredictedPresentationTime func(uintptr) int64

// Gets the predicted time at which this frame will be displayed.
//
// Although no predicted time may be available, if one is available,
// it will be available while the frame is being generated, in contrast
// to [method@Gdk.FrameTimings.get_presentation_time], which is only
// available after the frame has been presented.
//
// In general, if you are simply animating, you should use
// [method@Gdk.FrameClock.get_frame_time] rather than this function,
// but this function is useful for applications that want exact control
// over latency. For example, a movie player may want this information
// for Audio/Video synchronization.
func (x *FrameTimings) GetPredictedPresentationTime() int64 {

	cret := xFrameTimingsGetPredictedPresentationTime(x.GoPointer())
	return cret
}

var xFrameTimingsGetPresentationTime func(uintptr) int64

// Reurns the presentation time.
//
// This is the time at which the frame became visible to the user.
func (x *FrameTimings) GetPresentationTime() int64 {

	cret := xFrameTimingsGetPresentationTime(x.GoPointer())
	return cret
}

var xFrameTimingsGetRefreshInterval func(uintptr) int64

// Gets the natural interval between presentation times for
// the display that this frame was displayed on.
//
// Frame presentation usually happens during the “vertical
// blanking interval”.
func (x *FrameTimings) GetRefreshInterval() int64 {

	cret := xFrameTimingsGetRefreshInterval(x.GoPointer())
	return cret
}

var xFrameTimingsRef func(uintptr) *FrameTimings

// Increases the reference count of @timings.
func (x *FrameTimings) Ref() *FrameTimings {

	cret := xFrameTimingsRef(x.GoPointer())
	return cret
}

var xFrameTimingsUnref func(uintptr)

// Decreases the reference count of @timings.
//
// If @timings is no longer referenced, it will be freed.
func (x *FrameTimings) Unref() {

	xFrameTimingsUnref(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xFrameTimingsGLibType, lib, "gdk_frame_timings_get_type")

	core.PuregoSafeRegister(&xFrameTimingsGetComplete, lib, "gdk_frame_timings_get_complete")
	core.PuregoSafeRegister(&xFrameTimingsGetFrameCounter, lib, "gdk_frame_timings_get_frame_counter")
	core.PuregoSafeRegister(&xFrameTimingsGetFrameTime, lib, "gdk_frame_timings_get_frame_time")
	core.PuregoSafeRegister(&xFrameTimingsGetPredictedPresentationTime, lib, "gdk_frame_timings_get_predicted_presentation_time")
	core.PuregoSafeRegister(&xFrameTimingsGetPresentationTime, lib, "gdk_frame_timings_get_presentation_time")
	core.PuregoSafeRegister(&xFrameTimingsGetRefreshInterval, lib, "gdk_frame_timings_get_refresh_interval")
	core.PuregoSafeRegister(&xFrameTimingsRef, lib, "gdk_frame_timings_ref")
	core.PuregoSafeRegister(&xFrameTimingsUnref, lib, "gdk_frame_timings_unref")

}
