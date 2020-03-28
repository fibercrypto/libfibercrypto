package main

import (
	"reflect"
	"unsafe"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_Block_GetFee
func FC_mocks_Block_GetFee(__m *C.BlockMocks__Handle, _ticker string, _arg1 *uint64) (____error_code uint32) {
	_m, ok_m := lookupBlockMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, ____return_err := _m.GetFee(ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_mocks_Block_GetHash
func FC_mocks_Block_GetHash(__m *C.BlockMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupBlockMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetHash()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	}
	return
}

//export FC_mocks_Block_GetHeight
func FC_mocks_Block_GetHeight(__m *C.BlockMocks__Handle, _arg0 *uint64) (____error_code uint32) {
	_m, ok_m := lookupBlockMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetHeight()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = __arg0
	}
	return
}

//export FC_mocks_Block_GetPrevHash
func FC_mocks_Block_GetPrevHash(__m *C.BlockMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupBlockMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetPrevHash()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	}
	return
}

//export FC_mocks_Block_GetTime
func FC_mocks_Block_GetTime(__m *C.BlockMocks__Handle, _arg0 *C.core__Timestamp) (____error_code uint32) {
	_m, ok_m := lookupBlockMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetTime()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = *(*C.core__Timestamp)(unsafe.Pointer(&__arg0))
	}
	return
}

//export FC_mocks_Block_GetVersion
func FC_mocks_Block_GetVersion(__m *C.BlockMocks__Handle, _arg0 *uint32) (____error_code uint32) {
	_m, ok_m := lookupBlockMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetVersion()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = __arg0
	}
	return
}

//export FC_mocks_Block_IsGenesisBlock
func FC_mocks_Block_IsGenesisBlock(__m *C.BlockMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupBlockMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.IsGenesisBlock()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = __arg0
	}
	return
}
