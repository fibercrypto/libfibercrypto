package main

import "reflect"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_CryptoAccount_GetBalance
func FC_mocks_CryptoAccount_GetBalance(__m *C.CryptoAccountMocks__Handle, _ticker string, _arg1 *uint64) (____error_code uint32) {
	_m, ok_m := lookupCryptoAccountMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, ____return_err := _m.GetBalance(ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_mocks_CryptoAccount_ListAssets
func FC_mocks_CryptoAccount_ListAssets(__m *C.CryptoAccountMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupCryptoAccountMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ListAssets()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_CryptoAccount_ListPendingTransactions
func FC_mocks_CryptoAccount_ListPendingTransactions(__m *C.CryptoAccountMocks__Handle, _arg0 *C.TransactionIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupCryptoAccountMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.ListPendingTransactions()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = registerTransactionIteratorHandle(&__arg0)
	}
	return
}

//export FC_mocks_CryptoAccount_ListTransactions
func FC_mocks_CryptoAccount_ListTransactions(__m *C.CryptoAccountMocks__Handle, _arg0 *C.TransactionIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupCryptoAccountMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ListTransactions()
	*_arg0 = registerTransactionIteratorHandle(&__arg0)
	return
}

//export FC_mocks_CryptoAccount_ScanUnspentOutputs
func FC_mocks_CryptoAccount_ScanUnspentOutputs(__m *C.CryptoAccountMocks__Handle, _arg0 *C.TransactionOutputIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupCryptoAccountMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ScanUnspentOutputs()
	*_arg0 = registerTransactionOutputIteratorHandle(&__arg0)
	return
}
