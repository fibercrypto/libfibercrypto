package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_PEX_BroadcastTxn
func FC_mocks_PEX_BroadcastTxn(__m *C.PEXMocks__Handle, _txn *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupPEXMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__txn, oktxn := lookupTransactionHandle(*_txn)
	if !oktxn {
		____error_code = FC_BAD_HANDLE
		return
	}
	txn := *__txn
	____return_err := _m.BroadcastTxn(txn)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_PEX_GetConnections
func FC_mocks_PEX_GetConnections(__m *C.PEXMocks__Handle, _arg0 *C.PexNodeSet__Handle) (____error_code uint32) {
	_m, ok_m := lookupPEXMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetConnections()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = registerPexNodeSetHandle(&__arg0)
	}
	return
}

//export FC_mocks_PEX_GetTxnPool
func FC_mocks_PEX_GetTxnPool(__m *C.PEXMocks__Handle, _arg0 *C.TransactionIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupPEXMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetTxnPool()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = registerTransactionIteratorHandle(&__arg0)
	}
	return
}
