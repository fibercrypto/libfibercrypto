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

//export FC_mocks_StringAddress_GetCoinType
func FC_mocks_StringAddress_GetCoinType(__m *C.StringAddressMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupStringAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetCoinType()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_StringAddress_GetValue
func FC_mocks_StringAddress_GetValue(__m *C.StringAddressMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupStringAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetValue()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_StringAddress_IsValid
func FC_mocks_StringAddress_IsValid(__m *C.StringAddressMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupStringAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.IsValid()
	*_arg0 = __arg0
	return
}

//export FC_mocks_StringAddress_SetCoinType
func FC_mocks_StringAddress_SetCoinType(__m *C.StringAddressMocks__Handle, _val []byte) (____error_code uint32) {
	_m, ok_m := lookupStringAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	val := *(*[]byte)(unsafe.Pointer(&_val))
	_m.SetCoinType(val)
	return
}

//export FC_mocks_StringAddress_SetValue
func FC_mocks_StringAddress_SetValue(__m *C.StringAddressMocks__Handle, _val []byte) (____error_code uint32) {
	_m, ok_m := lookupStringAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	val := *(*[]byte)(unsafe.Pointer(&_val))
	_m.SetValue(val)
	return
}
