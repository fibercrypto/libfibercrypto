package main

import (
	"unsafe"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_BlockchainSignService_Sign
func FC_mocks_BlockchainSignService_Sign(__m *C.BlockchainSignServiceMocks__Handle, _txn *C.Transaction__Handle, _signSpec *C.InputSignDescriptor__Handle, _pwd *C.PasswordReader, _arg3 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupBlockchainSignServiceMocksHandle(*__m)
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
	signSpec := *(*[]core.InputSignDescriptor)(unsafe.Pointer(&_signSpec))
	pwd := *(*core.PasswordReader)(unsafe.Pointer(_pwd))
	__arg3, ____return_err := _m.Sign(txn, signSpec, pwd)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg3 = registerTransactionHandle(&__arg3)
	}
	return
}
