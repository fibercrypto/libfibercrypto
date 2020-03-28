package main

import util "github.com/fibercrypto/fibercryptowallet/src/util"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_NewBalanceSnapshot
func FC_util_NewBalanceSnapshot(_updInterval int64, _arg1 *C.BalanceSnapshot__Handle) (____error_code uint32) {
	updInterval := _updInterval
	__arg1 := util.NewBalanceSnapshot(updInterval)
	*_arg1 = registerBalanceSnapshotHandle(__arg1)
	return
}

//export FC_util_BalanceSnapshot_SetCoins
func FC_util_BalanceSnapshot_SetCoins(_bs *C.BalanceSnapshot__Handle, _ticker string, _coins uint64) (____error_code uint32) {
	bs, okbs := lookupBalanceSnapshotHandle(*_bs)
	if !okbs {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	coins := _coins
	bs.SetCoins(ticker, coins)
	return
}

//export FC_util_BalanceSnapshot_GetCoins
func FC_util_BalanceSnapshot_GetCoins(_bs *C.BalanceSnapshot__Handle, _ticker string, _arg1 *uint64) (____error_code uint32) {
	bs, okbs := lookupBalanceSnapshotHandle(*_bs)
	if !okbs {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, ____return_err := bs.GetCoins(ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_util_BalanceSnapshot_IsUpdated
func FC_util_BalanceSnapshot_IsUpdated(_bs *C.BalanceSnapshot__Handle, _arg0 *bool) (____error_code uint32) {
	bs, okbs := lookupBalanceSnapshotHandle(*_bs)
	if !okbs {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := bs.IsUpdated()
	*_arg0 = __arg0
	return
}
