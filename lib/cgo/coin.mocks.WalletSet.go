package main

import (
	"errors"
	"reflect"
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

//export FC_mocks_WalletSet_CreateWallet
func FC_mocks_WalletSet_CreateWallet(__m *C.WalletSetMocks__Handle, _name string, _seed string, _walletType string, _isEncryptrd bool, _pwd *C.PasswordReader, _scanAddressesN int, _arg6 *C.Wallet__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletSetMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	name := _name
	seed := _seed
	walletType := _walletType
	isEncryptrd := _isEncryptrd
	pwd := func(pString string, pKVS core.KeyValueStore) (string, error) {
		var pStringOut C.GoString_
		var pStringIn C.GoString_
		copyString(pString, &pStringIn)
		pKVSHandle := registerKeyValueStoreHandle(&pKVS)
		result := C.callPasswordReader(_pwd, pStringIn, pKVSHandle, &pStringOut)
		closeHandle(Handle(pKVSHandle))
		if result == FC_OK {
			pStringOut_ := *(*string)(unsafe.Pointer(&pStringOut))
			return pStringOut_, nil
		}
		return "", errors.New("Error in PasswdReader")
	}
	scanAddressesN := _scanAddressesN
	__arg6, ____return_err := _m.CreateWallet(name, seed, walletType, isEncryptrd, pwd, scanAddressesN)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg6 = registerWalletHandle(&__arg6)
	}
	return
}

//export FC_mocks_WalletSet_DefaultWalletType
func FC_mocks_WalletSet_DefaultWalletType(__m *C.WalletSetMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupWalletSetMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.DefaultWalletType()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_WalletSet_GetWallet
func FC_mocks_WalletSet_GetWallet(__m *C.WalletSetMocks__Handle, _id string, _arg1 *C.Wallet__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletSetMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	id := _id
	__arg1 := _m.GetWallet(id)
	*_arg1 = registerWalletHandle(&__arg1)
	return
}

//export FC_mocks_WalletSet_ListWallets
func FC_mocks_WalletSet_ListWallets(__m *C.WalletSetMocks__Handle, _arg0 *C.WalletIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletSetMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ListWallets()
	*_arg0 = registerWalletIteratorHandle(&__arg0)
	return
}

//export FC_mocks_WalletSet_SupportedWalletTypes
func FC_mocks_WalletSet_SupportedWalletTypes(__m *C.WalletSetMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupWalletSetMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.SupportedWalletTypes()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}
