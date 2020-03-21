package main

import (
	"reflect"
	"unsafe"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_Contact_GetAddresses
func FC_mocks_Contact_GetAddresses(__m *C.ContactMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupContactMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetAddresses()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_Contact_GetID
func FC_mocks_Contact_GetID(__m *C.ContactMocks__Handle, _arg0 *uint64) (____error_code uint32) {
	_m, ok_m := lookupContactMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetID()
	*_arg0 = __arg0
	return
}

//export FC_mocks_Contact_GetName
func FC_mocks_Contact_GetName(__m *C.ContactMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupContactMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetName()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_Contact_IsValid
func FC_mocks_Contact_IsValid(__m *C.ContactMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupContactMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.IsValid()
	*_arg0 = __arg0
	return
}

//export FC_mocks_Contact_SetAddresses
func FC_mocks_Contact_SetAddresses(__m *C.ContactMocks__Handle, __a0 *C.StringAddress__Handle) (____error_code uint32) {
	_m, ok_m := lookupContactMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *(*[]core.StringAddress)(unsafe.Pointer(&__a0))
	_m.SetAddresses(_a0)
	return
}

//export FC_mocks_Contact_SetID
func FC_mocks_Contact_SetID(__m *C.ContactMocks__Handle, _id uint64) (____error_code uint32) {
	_m, ok_m := lookupContactMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	id := _id
	_m.SetID(id)
	return
}

//export FC_mocks_Contact_SetName
func FC_mocks_Contact_SetName(__m *C.ContactMocks__Handle, __a0 string) (____error_code uint32) {
	_m, ok_m := lookupContactMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := __a0
	_m.SetName(_a0)
	return
}
