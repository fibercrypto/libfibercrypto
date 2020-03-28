package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_SimpleWalletOutput_GetWallet
func FC_util_SimpleWalletOutput_GetWallet(_wo *C.SimpleWalletOutput__Handle, _arg0 *C.Wallet__Handle) (____error_code uint32) {
	wo, okwo := lookupSimpleWalletOutputHandle(*_wo)
	if !okwo {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := wo.GetWallet()
	*_arg0 = registerWalletHandle(&__arg0)
	return
}

//export FC_util_SimpleWalletOutput_GetOutput
func FC_util_SimpleWalletOutput_GetOutput(_wo *C.SimpleWalletOutput__Handle, _arg0 *C.TransactionOutput__Handle) (____error_code uint32) {
	wo, okwo := lookupSimpleWalletOutputHandle(*_wo)
	if !okwo {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := wo.GetOutput()
	*_arg0 = registerTransactionOutputHandle(&__arg0)
	return
}

//export FC_util_SimpleWalletAddress_GetWallet
func FC_util_SimpleWalletAddress_GetWallet(_wa *C.SimpleWalletAddress__Handle, _arg0 *C.Wallet__Handle) (____error_code uint32) {
	wa, okwa := lookupSimpleWalletAddressHandle(*_wa)
	if !okwa {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := wa.GetWallet()
	*_arg0 = registerWalletHandle(&__arg0)
	return
}

//export FC_util_SimpleWalletAddress_GetAddress
func FC_util_SimpleWalletAddress_GetAddress(_wa *C.SimpleWalletAddress__Handle, _arg0 *C.Address__Handle) (____error_code uint32) {
	wa, okwa := lookupSimpleWalletAddressHandle(*_wa)
	if !okwa {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := wa.GetAddress()
	*_arg0 = registerAddressHandle(&__arg0)
	return
}
