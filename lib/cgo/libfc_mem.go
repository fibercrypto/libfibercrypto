package main

import (
	"reflect"
	"unsafe"
)

/*
	#include <string.h>

  #include "fctypes.h"

	void eos(char *str, int len) {
		str[len] = 0;
	}

*/
import "C"

const ()

/**
 * Copy helpers
 */

func copyString(src string, dest *C.GoString_) {
	srcLen := len(src)
	dest.p = (*C.char)(C.malloc(C.size_t(srcLen + 1)))
	strAddr := (*C.GoString_)(unsafe.Pointer(&src))
	result := C.memcpy(unsafe.Pointer(dest.p), unsafe.Pointer(strAddr.p), C.size_t(srcLen))
	if result != nil {
		C.eos(dest.p, C.int(srcLen))
		dest.n = C.GoInt_(srcLen)
	}
}

// Determine the memory address of a slice buffer and the
// size of its underlaying element type
func getBufferData(src reflect.Value) (bufferAddr unsafe.Pointer, elemSize C.size_t) {
	firstElem := src.Index(0)
	elemSize = C.size_t(firstElem.Type().Size())
	bufferAddr = unsafe.Pointer(src.Pointer())
	return
}

// Copy n items in source slice/array/string onto C-managed memory buffer
//
// This function takes for granted that all values in src
// will be instances of the same type, and that src and dest
// element types will be aligned exactly the same
// in memory of the same size
func copyToBuffer(src reflect.Value, dest unsafe.Pointer, n uint) {
	srcLen := src.Len()
	if srcLen == 0 {
		return
	}
	srcAddr, elemSize := getBufferData(src)
	if C.memcpy(dest, srcAddr, C.size_t(n)*elemSize) != nil {
	}
}

// Copy source slice/array/string onto instance of C.GSlice struct
//
// This function takes for granted that all values in src
// will be instances of the same type, and that src and dest
// element types will be aligned exactly the same
// in memory of the same size
func copyToGoSlice(src reflect.Value, dest *C.GoSlice_) {
	srcLen := src.Len()
	if srcLen == 0 {
		dest.len = 0
		return
	}
	srcAddr, elemSize := getBufferData(src)
	if dest.cap == 0 {
		dest.data = C.malloc(C.size_t(srcLen) * elemSize)
		dest.cap = C.GoInt_(srcLen)
	}
	n, overflow := srcLen, srcLen > int(dest.cap)
	if overflow {
		n = int(dest.cap)
	}
	result := C.memcpy(dest.data, srcAddr, C.size_t(n)*elemSize)
	if result != nil {
		// Do not modify slice metadata until memory is actually copied
		if overflow {
			dest.len = dest.cap - C.GoInt_(srcLen)
		} else {
			dest.len = C.GoInt_(srcLen)
		}
	}
}
