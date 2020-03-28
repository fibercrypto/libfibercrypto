package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_TransactionIterator_HasNext
func FC_mocks_TransactionIterator_HasNext(__m *C.TransactionIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.HasNext()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TransactionIterator_Next
func FC_mocks_TransactionIterator_Next(__m *C.TransactionIteratorMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupTransactionIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Next()
	*_arg0 = __arg0
	return
}

//export FC_mocks_TransactionIterator_Value
func FC_mocks_TransactionIterator_Value(__m *C.TransactionIteratorMocks__Handle, _arg0 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupTransactionIteratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.Value()
	*_arg0 = registerTransactionHandle(&__arg0)
	return
}
