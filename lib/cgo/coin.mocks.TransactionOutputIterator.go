package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_TransactionOutputIterator_HasNext
func FC_mocks_TransactionOutputIterator_HasNext(__m *C.TransactionOutputIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.HasNext()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TransactionOutputIterator_Next
func FC_mocks_TransactionOutputIterator_Next(__m *C.TransactionOutputIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Next()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TransactionOutputIterator_Value
func FC_mocks_TransactionOutputIterator_Value(__m *C.TransactionOutputIteratorMocks__Handle, _arg0 *C.TransactionOutput__Handle) (____error_code uint32) {
	_m, ok_m := lookupTransactionOutputIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Value()
	*_arg0 = registerTransactionOutputHandle(&__arg0)
	return
}
