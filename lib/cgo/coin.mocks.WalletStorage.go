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

//export FC_mocks_WalletStorage_Decrypt
func FC_mocks_WalletStorage_Decrypt(__m *C.WalletStorageMocks__Handle, _walletName string, _password *C.PasswordReader) (____error_code uint32) {
	_m, ok_m := lookupWalletStorageMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	walletName := _walletName
	password := func(pString string, pKVS core.KeyValueStore) (string, error) {
		var pStringOut C.GoString_
		var pStringIn C.GoString_
		copyString(pString, &pStringIn)
		pKVSHandle := registerKeyValueStoreHandle(&pKVS)
		result := C.callPasswordReader(_password, pStringIn, pKVSHandle, &pStringOut)
		closeHandle(Handle(pKVSHandle))
		if result == FC_OK {
			pStringOut_ := *(*string)(unsafe.Pointer(&pStringOut))
			return string(pStringOut_), nil
		}
		return "", errors.New("Error in PasswdReader")
	}
	_m.Decrypt(walletName, password)
	return
}

//export FC_mocks_WalletStorage_Encrypt
func FC_mocks_WalletStorage_Encrypt(__m *C.WalletStorageMocks__Handle, _walletName string, _password *C.PasswordReader) (____error_code uint32) {
	_m, ok_m := lookupWalletStorageMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	walletName := _walletName
	password := func(pString string, pKVS core.KeyValueStore) (string, error) {
		var pStringOut C.GoString_
		var pStringIn C.GoString_
		copyString(pString, &pStringIn)
		pKVSHandle := registerKeyValueStoreHandle(&pKVS)
		result := C.callPasswordReader(_password, pStringIn, pKVSHandle, &pStringOut)
		closeHandle(Handle(pKVSHandle))
		if result == FC_OK {
			pStringOut_ := *(*string)(unsafe.Pointer(&pStringOut))
			return string(pStringOut_), nil
		}
		return "", errors.New("Error in PasswdReader")
	}
	_m.Encrypt(walletName, password)
	return
}

//export FC_mocks_WalletStorage_IsEncrypted
func FC_mocks_WalletStorage_IsEncrypted(__m *C.WalletStorageMocks__Handle, _walletName string, _arg1 *bool) (____error_code uint32) {
	_m, ok_m := lookupWalletStorageMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	walletName := _walletName
	__arg1, ____return_err := _m.IsEncrypted(walletName)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}
