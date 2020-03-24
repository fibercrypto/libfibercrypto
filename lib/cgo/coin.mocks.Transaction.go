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

//export FC_mocks_Transaction_ComputeFee
func FC_mocks_Transaction_ComputeFee(__m *C.TransactionMocks__Handle, _ticker string, _arg1 *uint64) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, ____return_err := _m.ComputeFee(ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_mocks_Transaction_GetId
func FC_mocks_Transaction_GetId(__m *C.TransactionMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetId()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_Transaction_GetInputs
func FC_mocks_Transaction_GetInputs(__m *C.TransactionMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetInputs()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_Transaction_GetOutputs
func FC_mocks_Transaction_GetOutputs(__m *C.TransactionMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetOutputs()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_Transaction_GetStatus
func FC_mocks_Transaction_GetStatus(__m *C.TransactionMocks__Handle, _arg0 *C.core__TransactionStatus) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetStatus()
	*_arg0 = *(*C.core__TransactionStatus)(unsafe.Pointer(&__arg0))
	return
}

//export FC_mocks_Transaction_GetTimestamp
func FC_mocks_Transaction_GetTimestamp(__m *C.TransactionMocks__Handle, _arg0 *C.core__Timestamp) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetTimestamp()
	*_arg0 = *(*C.core__Timestamp)(unsafe.Pointer(&__arg0))
	return
}

//export FC_mocks_Transaction_IsFullySigned
func FC_mocks_Transaction_IsFullySigned(__m *C.TransactionMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.IsFullySigned()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = __arg0
	}
	return
}

//export FC_mocks_Transaction_SupportedAssets
func FC_mocks_Transaction_SupportedAssets(__m *C.TransactionMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.SupportedAssets()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_Transaction_VerifySigned
func FC_mocks_Transaction_VerifySigned(__m *C.TransactionMocks__Handle) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	____return_err := _m.VerifySigned()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_Transaction_VerifyUnsigned
func FC_mocks_Transaction_VerifyUnsigned(__m *C.TransactionMocks__Handle) (____error_code uint32) {
	_m, ok_m := lookupTransactionMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	____return_err := _m.VerifyUnsigned()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}
