package main

import (
	"unsafe"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_AltcoinCaption
func FC_util_AltcoinCaption(_ticker string, _arg1 *C.GoString_) (____error_code uint32) {
	ticker := _ticker
	__arg1 := util.AltcoinCaption(ticker)
	copyString(__arg1, _arg1)
	return
}

//export FC_util_AltcoinQuotient
func FC_util_AltcoinQuotient(_ticker string, _arg1 *uint64) (____error_code uint32) {
	ticker := _ticker
	__arg1, ____return_err := util.AltcoinQuotient(ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_util_RegisterAltcoin
func FC_util_RegisterAltcoin(_p *C.AltcoinPlugin__Handle) (____error_code uint32) {
	__p, okp := lookupAltcoinPluginHandle(*_p)
	if !okp {
		____error_code = FC_BAD_HANDLE
		return
	}
	p := *__p
	util.RegisterAltcoin(p)
	return
}

//export FC_util_LookupSignerByUID
func FC_util_LookupSignerByUID(_wlt *C.Wallet__Handle, _id *C.core__UID, _arg2 *C.TxnSigner__Handle) (____error_code uint32) {
	__wlt, okwlt := lookupWalletHandle(*_wlt)
	if !okwlt {
		____error_code = FC_BAD_HANDLE
		return
	}
	wlt := *__wlt
	id := *(*core.UID)(unsafe.Pointer(_id))
	__arg2 := util.LookupSignerByUID(wlt, id)
	*_arg2 = registerTxnSignerHandle(&__arg2)
	return
}
