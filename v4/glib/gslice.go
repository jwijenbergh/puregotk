// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type SliceConfig int

const (
	GSliceConfigAlwaysMallocValue SliceConfig = 1

	GSliceConfigBypassMagazinesValue SliceConfig = 2

	GSliceConfigWorkingSetMsecsValue SliceConfig = 3

	GSliceConfigColorIncrementValue SliceConfig = 4

	GSliceConfigChunkSizesValue SliceConfig = 5

	GSliceConfigContentionCounterValue SliceConfig = 6
)

var xSliceAlloc func(uint) uintptr

// Allocates a block of memory from the slice allocator.
//
// The block address handed out can be expected to be aligned
// to at least `1 * sizeof (void*)`, though in general slices
// are `2 * sizeof (void*)` bytes aligned; if a `malloc()`
// fallback implementation is used instead, the alignment may
// be reduced in a libc dependent fashion.
//
// Note that the underlying slice allocation mechanism can
// be changed with the [`G_SLICE=always-malloc`][G_SLICE]
// environment variable.
func SliceAlloc(BlockSizeVar uint) uintptr {

	cret := xSliceAlloc(BlockSizeVar)
	return cret
}

var xSliceAlloc0 func(uint) uintptr

// Allocates a block of memory via g_slice_alloc() and initializes
// the returned memory to 0. Note that the underlying slice allocation
// mechanism can be changed with the [`G_SLICE=always-malloc`][G_SLICE]
// environment variable.
func SliceAlloc0(BlockSizeVar uint) uintptr {

	cret := xSliceAlloc0(BlockSizeVar)
	return cret
}

var xSliceCopy func(uint, uintptr) uintptr

// Allocates a block of memory from the slice allocator
// and copies @block_size bytes into it from @mem_block.
//
// @mem_block must be non-%NULL if @block_size is non-zero.
func SliceCopy(BlockSizeVar uint, MemBlockVar uintptr) uintptr {

	cret := xSliceCopy(BlockSizeVar, MemBlockVar)
	return cret
}

var xSliceFree1 func(uint, uintptr)

// Frees a block of memory.
//
// The memory must have been allocated via g_slice_alloc() or
// g_slice_alloc0() and the @block_size has to match the size
// specified upon allocation. Note that the exact release behaviour
// can be changed with the [`G_DEBUG=gc-friendly`][G_DEBUG] environment
// variable, also see [`G_SLICE`][G_SLICE] for related debugging options.
//
// If @mem_block is %NULL, this function does nothing.
func SliceFree1(BlockSizeVar uint, MemBlockVar uintptr) {

	xSliceFree1(BlockSizeVar, MemBlockVar)

}

var xSliceFreeChainWithOffset func(uint, uintptr, uint)

// Frees a linked list of memory blocks of structure type @type.
//
// The memory blocks must be equal-sized, allocated via
// g_slice_alloc() or g_slice_alloc0() and linked together by a
// @next pointer (similar to #GSList). The offset of the @next
// field in each block is passed as third argument.
// Note that the exact release behaviour can be changed with the
// [`G_DEBUG=gc-friendly`][G_DEBUG] environment variable, also see
// [`G_SLICE`][G_SLICE] for related debugging options.
//
// If @mem_chain is %NULL, this function does nothing.
func SliceFreeChainWithOffset(BlockSizeVar uint, MemChainVar uintptr, NextOffsetVar uint) {

	xSliceFreeChainWithOffset(BlockSizeVar, MemChainVar, NextOffsetVar)

}

var xSliceGetConfig func(SliceConfig) int64

func SliceGetConfig(CkeyVar SliceConfig) int64 {

	cret := xSliceGetConfig(CkeyVar)
	return cret
}

var xSliceGetConfigState func(SliceConfig, int64, uint) int64

func SliceGetConfigState(CkeyVar SliceConfig, AddressVar int64, NValuesVar uint) int64 {

	cret := xSliceGetConfigState(CkeyVar, AddressVar, NValuesVar)
	return cret
}

var xSliceSetConfig func(SliceConfig, int64)

func SliceSetConfig(CkeyVar SliceConfig, ValueVar int64) {

	xSliceSetConfig(CkeyVar, ValueVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xSliceAlloc, lib, "g_slice_alloc")
	core.PuregoSafeRegister(&xSliceAlloc0, lib, "g_slice_alloc0")
	core.PuregoSafeRegister(&xSliceCopy, lib, "g_slice_copy")
	core.PuregoSafeRegister(&xSliceFree1, lib, "g_slice_free1")
	core.PuregoSafeRegister(&xSliceFreeChainWithOffset, lib, "g_slice_free_chain_with_offset")
	core.PuregoSafeRegister(&xSliceGetConfig, lib, "g_slice_get_config")
	core.PuregoSafeRegister(&xSliceGetConfigState, lib, "g_slice_get_config_state")
	core.PuregoSafeRegister(&xSliceSetConfig, lib, "g_slice_set_config")

}
