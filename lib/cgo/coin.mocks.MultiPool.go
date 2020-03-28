package main

import "reflect"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_MultiPool_CreateSection
func FC_mocks_MultiPool_CreateSection(__m *C.MultiPoolMocks__Handle, __a0 string, __a1 *C.PooledObjectFactory__Handle) (____error_code uint32) {
	_m, ok_m := lookupMultiPoolMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := __a0
	___a1, ok_a1 := lookupPooledObjectFactoryHandle(*__a1)
	if !ok_a1 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a1 := *___a1
	____return_err := _m.CreateSection(_a0, _a1)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_MultiPool_GetSection
func FC_mocks_MultiPool_GetSection(__m *C.MultiPoolMocks__Handle, __a0 string, _arg1 *C.MultiPoolSection__Handle) (____error_code uint32) {
	_m, ok_m := lookupMultiPoolMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := __a0
	__arg1, ____return_err := _m.GetSection(_a0)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerMultiPoolSectionHandle(&__arg1)
	}
	return
}

//export FC_mocks_MultiPool_ListSections
func FC_mocks_MultiPool_ListSections(__m *C.MultiPoolMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupMultiPoolMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.ListSections()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	}
	return
}
