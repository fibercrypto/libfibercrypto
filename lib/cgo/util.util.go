package main

import (
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_Min
func FC_util_Min(_arg0 int, _arg1 int, _arg2 *int) (____error_code uint32) {
	arg2 := util.Min(_arg0, _arg1)
	*_arg2 = arg2
	return
}

//export FC_util_GetCoinValue
func FC_util_GetCoinValue(_arg0 string, _arg1 string, _arg2 *uint64) (____error_code uint32) {
	arg0 := _arg0
	arg1 := _arg1

	arg2, err := util.GetCoinValue(arg0, arg1)
	____error_code = libErrorCode(err)
	if err == nil {
		*_arg2 = arg2
		return
	}
	return
}

//export FC_util_FormatUint64
func FC_util_FormatUint64(_arg0 uint64, _arg1 *C.GoString_) (____error_code uint32) {
	arg0 := uint64(_arg0)
	arg1 := util.FormatUint64(arg0)
	copyString(arg1, _arg1)
	return
}

//export FC_util_FormatCoins
func FC_util_FormatCoins(_arg0 uint64, _arg1 uint64, _arg2 *C.GoString_) (____error_code uint32) {
	arg0 := uint64(_arg0)
	arg1 := uint64(_arg1)
	arg2 := util.FormatCoins(arg0, arg1)
	copyString(arg2, _arg2)
	return
}

//export FC_util_RemoveZeros
func FC_util_RemoveZeros(_arg0 string, _arg1 *C.GoString_) (____error_code uint32) {
	arg0 := string(_arg0)
	arg1 := util.RemoveZeros(arg0)
	copyString(arg1, _arg1)
	return
}

//export FC_util_StringInList
func FC_util_StringInList(_arg0 string, _arg1 C.Strings__Handle, _arg2 *bool) (____error_code uint32) {
	arg0 := string(_arg0)
	arg1, ok := lookupStringsHandle(_arg1)
	if ok != true {
		____error_code = FC_BAD_HANDLE
		return
	}

	arg2 := util.StringInList(arg0, arg1)
	*_arg2 = arg2

	return
}
