package main

import (
	"errors"
	"unsafe"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
  #include "fccallback.h"
*/
import "C"

//export FC_mocks_TxnSigner_GetSignerDescription
func FC_mocks_TxnSigner_GetSignerDescription(__m *C.TxnSignerMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetSignerDescription()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyString(__arg0, _arg0)
	}
	return
}

//export FC_mocks_TxnSigner_GetSignerUID
func FC_mocks_TxnSigner_GetSignerUID(__m *C.TxnSignerMocks__Handle, _arg0 *C.core__UID) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetSignerUID()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = *(*C.core__UID)(unsafe.Pointer(&__arg0))
	}
	return
}

//export FC_mocks_TxnSigner_ReadyForTxn
func FC_mocks_TxnSigner_ReadyForTxn(__m *C.TxnSignerMocks__Handle, __a0 *C.Wallet__Handle, __a1 *C.Transaction__Handle, _arg2 *bool) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	___a0, ok_a0 := lookupWalletHandle(*__a0)
	if !ok_a0 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *___a0
	___a1, ok_a1 := lookupTransactionHandle(*__a1)
	if !ok_a1 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a1 := *___a1
	__arg2, ____return_err := _m.ReadyForTxn(_a0, _a1)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg2 = __arg2
	}
	return
}

//export FC_mocks_TxnSigner_SignTransaction
func FC_mocks_TxnSigner_SignTransaction(__m *C.TxnSignerMocks__Handle, __a0 *C.Transaction__Handle, __a1 *C.PasswordReader, __a2 []string, _arg3 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupTxnSignerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	___a0, ok_a0 := lookupTransactionHandle(*__a0)
	if !ok_a0 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *___a0
	_a1 := func(pString string, pKVS core.KeyValueStore) (string, error) {
		var pStringOut C.GoString_
		var pStringIn C.GoString_
		copyString(pString, &pStringIn)
		pKVSHandle := registerKeyValueStoreHandle(&pKVS)
		result := C.callPasswordReader(__a1, pStringIn, pKVSHandle, &pStringOut)
		closeHandle(Handle(pKVSHandle))
		if result == FC_OK {
			pStringOut_ := *(*string)(unsafe.Pointer(&pStringOut))
			return pStringOut_, nil
		}
		return "", errors.New("Error in PasswdReader")
	}
	_a2 := *(*[]string)(unsafe.Pointer(&__a2))
	__arg3, ____return_err := _m.SignTransaction(_a0, _a1, _a2)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg3 = registerTransactionHandle(&__arg3)
	}
	return
}
