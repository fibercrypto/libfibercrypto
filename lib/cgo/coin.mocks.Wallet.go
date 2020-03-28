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

//export FC_mocks_Wallet_GenAddresses
func FC_mocks_Wallet_GenAddresses(__m *C.WalletMocks__Handle, _addrType *C.core__AddressType, _startIndex uint32, _count uint32, _pwd *C.PasswordReader, _arg4 *C.AddressIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	addrType := *(*core.AddressType)(unsafe.Pointer(_addrType))
	startIndex := _startIndex
	count := _count
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
	__arg4 := _m.GenAddresses(addrType, startIndex, count, pwd)
	*_arg4 = registerAddressIteratorHandle(&__arg4)
	return
}

//export FC_mocks_Wallet_GetCryptoAccount
func FC_mocks_Wallet_GetCryptoAccount(__m *C.WalletMocks__Handle, _arg0 *C.CryptoAccount__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetCryptoAccount()
	*_arg0 = registerCryptoAccountHandle(&__arg0)
	return
}

//export FC_mocks_Wallet_GetId
func FC_mocks_Wallet_GetId(__m *C.WalletMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetId()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_Wallet_GetLabel
func FC_mocks_Wallet_GetLabel(__m *C.WalletMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetLabel()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_Wallet_GetLoadedAddresses
func FC_mocks_Wallet_GetLoadedAddresses(__m *C.WalletMocks__Handle, _arg0 *C.AddressIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetLoadedAddresses()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = registerAddressIteratorHandle(&__arg0)
	}
	return
}

//export FC_mocks_Wallet_SendFromAddress
func FC_mocks_Wallet_SendFromAddress(__m *C.WalletMocks__Handle, _from *C.Address__Handle, _to *C.TransactionOutput__Handle, _change *C.Address__Handle, _options *C.KeyValueStore__Handle, _arg4 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	from := *(*[]core.Address)(unsafe.Pointer(&_from))
	to := *(*[]core.TransactionOutput)(unsafe.Pointer(&_to))
	__change, okchange := lookupAddressHandle(*_change)
	if !okchange {
		____error_code = FC_BAD_HANDLE
		return
	}
	change := *__change
	__options, okoptions := lookupKeyValueStoreHandle(*_options)
	if !okoptions {
		____error_code = FC_BAD_HANDLE
		return
	}
	options := *__options
	__arg4, ____return_err := _m.SendFromAddress(from, to, change, options)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg4 = registerTransactionHandle(&__arg4)
	}
	return
}

//export FC_mocks_Wallet_SetLabel
func FC_mocks_Wallet_SetLabel(__m *C.WalletMocks__Handle, _wltName string) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	wltName := _wltName
	_m.SetLabel(wltName)
	return
}

//export FC_mocks_Wallet_Sign
func FC_mocks_Wallet_Sign(__m *C.WalletMocks__Handle, _txn *C.Transaction__Handle, _signer *C.TxnSigner__Handle, _pwd *C.PasswordReader, _index []string, _arg4 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
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
	__signer, oksigner := lookupTxnSignerHandle(*_signer)
	if !oksigner {
		____error_code = FC_BAD_HANDLE
		return
	}
	signer := *__signer
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
	index := *(*[]string)(unsafe.Pointer(&_index))
	__arg4, ____return_err := _m.Sign(txn, signer, pwd, index)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg4 = registerTransactionHandle(&__arg4)
	}
	return
}

//export FC_mocks_Wallet_Spend
func FC_mocks_Wallet_Spend(__m *C.WalletMocks__Handle, _unspent *C.TransactionOutput__Handle, _new *C.TransactionOutput__Handle, _change *C.Address__Handle, _options *C.KeyValueStore__Handle, _arg4 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	unspent := *(*[]core.TransactionOutput)(unsafe.Pointer(&_unspent))
	new := *(*[]core.TransactionOutput)(unsafe.Pointer(&_new))
	__change, okchange := lookupAddressHandle(*_change)
	if !okchange {
		____error_code = FC_BAD_HANDLE
		return
	}
	change := *__change
	__options, okoptions := lookupKeyValueStoreHandle(*_options)
	if !okoptions {
		____error_code = FC_BAD_HANDLE
		return
	}
	options := *__options
	__arg4, ____return_err := _m.Spend(unspent, new, change, options)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg4 = registerTransactionHandle(&__arg4)
	}
	return
}

//export FC_mocks_Wallet_Transfer
func FC_mocks_Wallet_Transfer(__m *C.WalletMocks__Handle, _to *C.TransactionOutput__Handle, _options *C.KeyValueStore__Handle, _arg2 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__to, okto := lookupTransactionOutputHandle(*_to)
	if !okto {
		____error_code = FC_BAD_HANDLE
		return
	}
	to := *__to
	__options, okoptions := lookupKeyValueStoreHandle(*_options)
	if !okoptions {
		____error_code = FC_BAD_HANDLE
		return
	}
	options := *__options
	__arg2, ____return_err := _m.Transfer(to, options)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg2 = registerTransactionHandle(&__arg2)
	}
	return
}
