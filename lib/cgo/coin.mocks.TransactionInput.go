package main

import "reflect"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_TransactionInput_GetCoins
func FC_mocks_TransactionInput_GetCoins(__m *C.TransactionInputMocks__Handle, _ticker string, _arg1 *uint64) (____error_code uint32) {
	_m, ok_m := lookupTransactionInputMocksHandle(*__m)
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

//export FC_mocks_TransactionInput_GetId
func FC_mocks_TransactionInput_GetId(__m *C.TransactionInputMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupTransactionInputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetId()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_TransactionInput_GetSpentOutput
func FC_mocks_TransactionInput_GetSpentOutput(__m *C.TransactionInputMocks__Handle, _arg0 *C.TransactionOutput__Handle) (____error_code uint32) {
	_m, ok_m := lookupTransactionInputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetSpentOutput()
	*_arg0 = registerTransactionOutputHandle(&__arg0)
	return
}

//export FC_mocks_TransactionInput_SupportedAssets
func FC_mocks_TransactionInput_SupportedAssets(__m *C.TransactionInputMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupTransactionInputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.SupportedAssets()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}
