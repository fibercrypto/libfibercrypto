package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_WalletAddress_GetAddress
func FC_mocks_WalletAddress_GetAddress(__m *C.WalletAddressMocks__Handle, _arg0 *C.Address__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetAddress()
	*_arg0 = registerAddressHandle(&__arg0)
	return
}

//export FC_mocks_WalletAddress_GetWallet
func FC_mocks_WalletAddress_GetWallet(__m *C.WalletAddressMocks__Handle, _arg0 *C.Wallet__Handle) (____error_code uint32) {
	_m, ok_m := lookupWalletAddressMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetWallet()
	*_arg0 = registerWalletHandle(&__arg0)
	return
}
