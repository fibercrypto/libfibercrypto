package main

import (
	"unsafe"

	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_Min
func FC_util_Min(_a, _b int, _arg1 *int) (____error_code uint32) {
	a := _a
	b := _b
	__arg1 := util.Min(a, b)
	*_arg1 = __arg1
	return
}

//export FC_util_GetCoinValue
func FC_util_GetCoinValue(_value string, _ticker string, _arg2 *uint64) (____error_code uint32) {
	value := _value
	ticker := _ticker
	__arg2, ____return_err := util.GetCoinValue(value, ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg2 = __arg2
	}
	return
}

//export FC_util_FormatUint64
func FC_util_FormatUint64(_n uint64, _arg1 *C.GoString_) (____error_code uint32) {
	n := _n
	__arg1 := util.FormatUint64(n)
	copyString(__arg1, _arg1)
	return
}

//export FC_util_FormatCoins
func FC_util_FormatCoins(_n uint64, _quotient uint64, _arg2 *C.GoString_) (____error_code uint32) {
	n := _n
	quotient := _quotient
	__arg2 := util.FormatCoins(n, quotient)
	copyString(__arg2, _arg2)
	return
}

//export FC_util_RemoveZeros
func FC_util_RemoveZeros(_s string, _arg1 *C.GoString_) (____error_code uint32) {
	s := _s
	__arg1 := util.RemoveZeros(s)
	copyString(__arg1, _arg1)
	return
}

//export FC_util_StringInList
func FC_util_StringInList(_s string, _list []string, _arg2 *bool) (____error_code uint32) {
	s := _s
	list := *(*[]string)(unsafe.Pointer(&_list))
	__arg2 := util.StringInList(s, list)
	*_arg2 = __arg2
	return
}

//export FC_util_AddressFromString
func FC_util_AddressFromString(_addrs, _coinTicket string, _arg1 *C.Address__Handle) (____error_code uint32) {
	addrs := _addrs
	coinTicket := _coinTicket
	__arg1, ____return_err := util.AddressFromString(addrs, coinTicket)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerAddressHandle(&__arg1)
	}
	return
}
