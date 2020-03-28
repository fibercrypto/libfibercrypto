package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_WalletOutput_GetOutput
func FC_mocks_WalletOutput_GetOutput(__m *C.WalletOutputMocks__Handle, _arg0 *C.TransactionOutput__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletOutputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetOutput()
	*_arg0 = registerTransactionOutputHandle(&__arg0)
	return
}

//export FC_mocks_WalletOutput_GetWallet
func FC_mocks_WalletOutput_GetWallet(__m *C.WalletOutputMocks__Handle, _arg0 *C.Wallet__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletOutputMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetWallet()
	*_arg0 = registerWalletHandle(&__arg0)
	return
}
