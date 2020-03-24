package main

import "reflect"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_TransactionOutput_GetAddress
func FC_mocks_TransactionOutput_GetAddress(__m *C.TransactionOutputMocks__Handle, _arg0 *C.Address__Handle) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetAddress()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = registerAddressHandle(&__arg0)
	}
	return
}

//export FC_mocks_TransactionOutput_GetCoins
func FC_mocks_TransactionOutput_GetCoins(__m *C.TransactionOutputMocks__Handle, _ticker string, _arg1 *uint64) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, ____return_err := _m.GetCoins(ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_mocks_TransactionOutput_GetId
func FC_mocks_TransactionOutput_GetId(__m *C.TransactionOutputMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetId()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_TransactionOutput_IsSpent
func FC_mocks_TransactionOutput_IsSpent(__m *C.TransactionOutputMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.IsSpent()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TransactionOutput_SupportedAssets
func FC_mocks_TransactionOutput_SupportedAssets(__m *C.TransactionOutputMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.SupportedAssets()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}
