package main

import (
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"reflect"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_NewGenericOutput
func FC_util_NewGenericOutput(_addr *C.Address__Handle, _id string, _arg2 *C.GenericOutput__Handle) (____error_code uint32) {
	__addr, okaddr := lookupAddressHandle(*_addr)
	if !okaddr {
		____error_code = FC_BAD_HANDLE
		return
	}
	addr := *__addr
	id := _id
	__arg2 := util.NewGenericOutput(addr, id)
	*_arg2 = registerGenericOutputHandle(&__arg2)
	return
}

//export FC_util_GenericOutput_GetId
func FC_util_GenericOutput_GetId(_gOut *C.GenericOutput__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	gOut, okgOut := lookupGenericOutputHandle(*_gOut)
	if !okgOut {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := gOut.GetId()
	copyString(__arg0, _arg0)
	return
}

//export FC_util_GenericOutput_IsSpent
func FC_util_GenericOutput_IsSpent(_gOut *C.GenericOutput__Handle, _arg0 *bool) (____error_code uint32) {
	gOut, okgOut := lookupGenericOutputHandle(*_gOut)
	if !okgOut {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := gOut.IsSpent()
	*_arg0 = __arg0
	return
}

//export FC_util_GenericOutput_GetAddress
func FC_util_GenericOutput_GetAddress(_gOut *C.GenericOutput__Handle, _arg0 *C.Address__Handle) (____error_code uint32) {
	gOut, okgOut := lookupGenericOutputHandle(*_gOut)
	if !okgOut {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := gOut.GetAddress()
	*_arg0 = registerAddressHandle(&__arg0)
	return
}

//export FC_util_GenericOutput_SetCoins
func FC_util_GenericOutput_SetCoins(_gOut *C.GenericOutput__Handle, _ticker string, _coins uint64) (____error_code uint32) {
	gOut, okgOut := lookupGenericOutputHandle(*_gOut)
	if !okgOut {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	coins := _coins
	gOut.SetCoins(ticker, coins)
	return
}

//export FC_util_GenericOutput_PushCoins
func FC_util_GenericOutput_PushCoins(_gOut *C.GenericOutput__Handle, _ticker string, _coinStr string) (____error_code uint32) {
	gOut, okgOut := lookupGenericOutputHandle(*_gOut)
	if !okgOut {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	coinStr := _coinStr
	____return_err := gOut.PushCoins(ticker, coinStr)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_util_GenericOutput_GetCoins
func FC_util_GenericOutput_GetCoins(_gOut *C.GenericOutput__Handle, _ticker string, _arg1 *uint64) (____error_code uint32) {
	gOut, okgOut := lookupGenericOutputHandle(*_gOut)
	if !okgOut {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, ____return_err := gOut.GetCoins(ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_util_GenericOutput_SupportedAssets
func FC_util_GenericOutput_SupportedAssets(_gOut *C.GenericOutput__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	gOut, okgOut := lookupGenericOutputHandle(*_gOut)
	if !okgOut {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := gOut.SupportedAssets()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}
