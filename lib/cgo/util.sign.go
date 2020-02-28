package main

import (
	"errors"
	"unsafe"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
  #include "fccallback.h"
*/
import "C"

//export FC_util_AttachSignService
func FC_util_AttachSignService(_signer *C.TxnSigner__Handle) (____error_code uint32) {
	__signer, oksigner := lookupTxnSignerHandle(*_signer)
	if !oksigner {
		____error_code = FC_BAD_HANDLE
		return
	}
	signer := *__signer
	____return_err := util.AttachSignService(signer)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_util_LookupSignService
func FC_util_LookupSignService(_signerID *C.core__UID, _arg1 *C.TxnSigner__Handle) (____error_code uint32) {
	signerID := *(*core.UID)(unsafe.Pointer(_signerID))
	__arg1 := util.LookupSignService(signerID)
	*_arg1 = registerTxnSignerHandle(&__arg1)
	return
}

//export FC_util_RemoveSignService
func FC_util_RemoveSignService(_signerID *C.core__UID) (____error_code uint32) {
	signerID := *(*core.UID)(unsafe.Pointer(_signerID))
	____return_err := util.RemoveSignService(signerID)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_util_EnumerateSignServices
func FC_util_EnumerateSignServices(_arg0 *C.TxnSignerIterator__Handle) (____error_code uint32) {
	__arg0 := util.EnumerateSignServices()
	*_arg0 = registerTxnSignerIteratorHandle(&__arg0)
	return
}

//export FC_util_SignServicesForTxn
func FC_util_SignServicesForTxn(_wlt *C.Wallet__Handle, _txn *C.Transaction__Handle, _arg2 *C.TxnSignerIterator__Handle) (____error_code uint32) {
	__wlt, okwlt := lookupWalletHandle(*_wlt)
	if !okwlt {
		____error_code = FC_BAD_HANDLE
		return
	}
	wlt := *__wlt
	__txn, oktxn := lookupTransactionHandle(*_txn)
	if !oktxn {
		____error_code = FC_BAD_HANDLE
		return
	}
	txn := *__txn
	__arg2 := util.SignServicesForTxn(wlt, txn)
	*_arg2 = registerTxnSignerIteratorHandle(&__arg2)
	return
}

//export FC_util_ReadyForTxn
func FC_util_ReadyForTxn(_signerID *C.core__UID, _wallet *C.Wallet__Handle, _txn *C.Transaction__Handle, _arg3 *bool) (____error_code uint32) {
	signerID := *(*core.UID)(unsafe.Pointer(_signerID))
	__wallet, okwallet := lookupWalletHandle(*_wallet)
	if !okwallet {
		____error_code = FC_BAD_HANDLE
		return
	}
	wallet := *__wallet
	__txn, oktxn := lookupTransactionHandle(*_txn)
	if !oktxn {
		____error_code = FC_BAD_HANDLE
		return
	}
	txn := *__txn
	__arg3, ____return_err := util.ReadyForTxn(signerID, wallet, txn)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg3 = __arg3
	}
	return
}

//export FC_util_SignTransaction
func FC_util_SignTransaction(_signerID *C.core__UID, _txn *C.Transaction__Handle, _pwd *C.PasswordReader, _indices []string, _arg4 *C.Transaction__Handle) (____error_code uint32) {
	signerID := *(*core.UID)(unsafe.Pointer(_signerID))
	__txn, oktxn := lookupTransactionHandle(*_txn)
	if !oktxn {
		____error_code = FC_BAD_HANDLE
		return
	}
	txn := *__txn
	pwd := func(pString string, pKVS core.KeyValueStore) (string, error) {
		var pStringOut C.GoString_
		var pStringIn C.GoString_
		copyString(pString, &pStringIn)
		pKVSHandle := registerKeyValueStoreHandle(&pKVS)
		result := C.callPasswordReader(_pwd, pStringIn, pKVSHandle, &pStringOut)
		closeHandle(Handle(pKVSHandle))
		if result == FC_OK {
			pStringOut_ := *(*string)(unsafe.Pointer(&pStringOut))
			return string(pStringOut_), nil
		}
		return "", errors.New("Error in PasswdReader")
	}
	indices := *(*[]string)(unsafe.Pointer(&_indices))
	__arg4, ____return_err := util.SignTransaction(signerID, txn, pwd, indices)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg4 = registerTransactionHandle(&__arg4)
	}
	return
}

//export FC_util_GetSignerDescription
func FC_util_GetSignerDescription(_signerID *C.core__UID, _arg1 *C.GoString_) (____error_code uint32) {
	signerID := *(*core.UID)(unsafe.Pointer(_signerID))
	__arg1, ____return_err := util.GetSignerDescription(signerID)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyString(__arg1, _arg1)
	}
	return
}

//export FC_util_LookupSignServiceForWallet
func FC_util_LookupSignServiceForWallet(_wlt *C.Wallet__Handle, _signerID *C.core__UID, _arg2 *C.TxnSigner__Handle) (____error_code uint32) {
	__wlt, okwlt := lookupWalletHandle(*_wlt)
	if !okwlt {
		____error_code = FC_BAD_HANDLE
		return
	}
	wlt := *__wlt
	signerID := *(*core.UID)(unsafe.Pointer(_signerID))
	__arg2, ____return_err := util.LookupSignServiceForWallet(wlt, signerID)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg2 = registerTxnSignerHandle(&__arg2)
	}
	return
}

//export FC_util_GenericMultiWalletSign
func FC_util_GenericMultiWalletSign(_txn *C.Transaction__Handle, _signSpec *C.InputSignDescriptor__Handle, _pwd *C.PasswordReader, _arg3 *C.Transaction__Handle) (____error_code uint32) {
	__txn, oktxn := lookupTransactionHandle(*_txn)
	if !oktxn {
		____error_code = FC_BAD_HANDLE
		return
	}
	txn := *__txn
	signSpec := *(*[]core.InputSignDescriptor)(unsafe.Pointer(&_signSpec))
	// type PasswordReader func(string, KeyValueStore) (string, error)
	pwd := func(pString string, pKVS core.KeyValueStore) (string, error) {
		var pStringOut C.GoString_
		var pStringIn C.GoString_
		copyString(pString, &pStringIn)
		pKVSHandle := registerKeyValueStoreHandle(&pKVS)
		result := C.callPasswordReader(_pwd, pStringIn, pKVSHandle, &pStringOut)
		closeHandle(Handle(pKVSHandle))
		if result == FC_OK {
			pStringOut_ := *(*string)(unsafe.Pointer(&pStringOut))
			return string(pStringOut_), nil
		}
		return "", errors.New("Error in PasswdReader")
	}
	__arg3, ____return_err := util.GenericMultiWalletSign(txn, signSpec, pwd)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg3 = registerTransactionHandle(&__arg3)
	}
	return
}
