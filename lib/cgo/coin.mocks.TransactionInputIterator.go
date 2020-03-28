package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_TransactionInputIterator_HasNext
func FC_mocks_TransactionInputIterator_HasNext(__m *C.TransactionInputIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionInputIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.HasNext()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TransactionInputIterator_Next
func FC_mocks_TransactionInputIterator_Next(__m *C.TransactionInputIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionInputIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Next()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TransactionInputIterator_Value
func FC_mocks_TransactionInputIterator_Value(__m *C.TransactionInputIteratorMocks__Handle, _arg0 *C.TransactionInput__Handle) (____error_code uint32) {
	_m, ok_m := lookupTransactionInputIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Value()
	*_arg0 = registerTransactionInputHandle(&__arg0)
	return
}
