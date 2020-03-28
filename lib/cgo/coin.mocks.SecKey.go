package main

import "reflect"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_SecKey_Bytes
func FC_mocks_SecKey_Bytes(__m *C.SecKeyMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupSecKeyMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Bytes()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_SecKey_Null
func FC_mocks_SecKey_Null(__m *C.SecKeyMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupSecKeyMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Null()
	*_arg0 = __arg0
	return
}

//export FC_mocks_SecKey_Verify
func FC_mocks_SecKey_Verify(__m *C.SecKeyMocks__Handle) (____error_code uint32) {
	_m, ok_m := lookupSecKeyMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	____return_err := _m.Verify()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}
