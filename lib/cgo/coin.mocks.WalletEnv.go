package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_WalletEnv_GetStorage
func FC_mocks_WalletEnv_GetStorage(__m *C.WalletEnvMocks__Handle, _arg0 *C.WalletStorage__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletEnvMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetStorage()
	*_arg0 = registerWalletStorageHandle(&__arg0)
	return
}

//export FC_mocks_WalletEnv_GetWalletSet
func FC_mocks_WalletEnv_GetWalletSet(__m *C.WalletEnvMocks__Handle, _arg0 *C.WalletSet__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletEnvMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetWalletSet()
	*_arg0 = registerWalletSetHandle(&__arg0)
	return
}

//export FC_mocks_WalletEnv_LookupWallet
func FC_mocks_WalletEnv_LookupWallet(__m *C.WalletEnvMocks__Handle, _firstAddr string, _arg1 *C.Wallet__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletEnvMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	firstAddr := _firstAddr
	__arg1, ____return_err := _m.LookupWallet(firstAddr)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerWalletHandle(&__arg1)
	}
	return
}
