package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_TxnSignerIterator_Count
func FC_mocks_TxnSignerIterator_Count(__m *C.TxnSignerIteratorMocks__Handle, _arg0 *int) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Count()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TxnSignerIterator_HasNext
func FC_mocks_TxnSignerIterator_HasNext(__m *C.TxnSignerIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.HasNext()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TxnSignerIterator_Next
func FC_mocks_TxnSignerIterator_Next(__m *C.TxnSignerIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Next()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TxnSignerIterator_Value
func FC_mocks_TxnSignerIterator_Value(__m *C.TxnSignerIteratorMocks__Handle, _arg0 *C.TxnSigner__Handle) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Value()
	*_arg0 = registerTxnSignerHandle(&__arg0)
	return
}
