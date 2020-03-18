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

//export FC_mocks_Address_Bytes
func FC_mocks_Address_Bytes(__m *C.AddressMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Bytes()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_Address_Checksum
func FC_mocks_Address_Checksum(__m *C.AddressMocks__Handle, _arg0 *C.core__Checksum) (____error_code uint32) {
	_m, ok_m := lookupAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Checksum()
	*_arg0 = *(*C.core__Checksum)(unsafe.Pointer(&__arg0))
	return
}

//export FC_mocks_Address_GetCryptoAccount
func FC_mocks_Address_GetCryptoAccount(__m *C.AddressMocks__Handle, _arg0 *C.CryptoAccount__Handle) (____error_code uint32) {
	_m, ok_m := lookupAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetCryptoAccount()
	*_arg0 = registerCryptoAccountHandle(&__arg0)
	return
}

//export FC_mocks_Address_IsBip32
func FC_mocks_Address_IsBip32(__m *C.AddressMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.IsBip32()
	*_arg0 = __arg0
	return
}

//export FC_mocks_Address_Null
func FC_mocks_Address_Null(__m *C.AddressMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Null()
	*_arg0 = __arg0
	return
}

//export FC_mocks_Address_String
func FC_mocks_Address_String(__m *C.AddressMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.String()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_Address_Verify
func FC_mocks_Address_Verify(__m *C.AddressMocks__Handle, __a0 *C.PubKey__Handle) (____error_code uint32) {
	_m, ok_m := lookupAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	___a0, ok_a0 := lookupPubKeyHandle(*__a0)
	if !ok_a0 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *___a0
	____return_err := _m.Verify(_a0)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}
