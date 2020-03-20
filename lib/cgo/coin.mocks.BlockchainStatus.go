package main

import (
	"unsafe"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_BlockchainStatus_GetCoinValue
func FC_mocks_BlockchainStatus_GetCoinValue(__m *C.BlockchainStatusMocks__Handle, _coinvalue *C.core__CoinValueMetric, _ticker string, _arg2 *uint64) (____error_code uint32) {
	_m, ok_m := lookupBlockchainStatusMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	coinvalue := *(*core.CoinValueMetric)(unsafe.Pointer(_coinvalue))
	ticker := _ticker
	__arg2, ____return_err := _m.GetCoinValue(coinvalue, ticker)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg2 = __arg2
	}
	return
}

//export FC_mocks_BlockchainStatus_GetLastBlock
func FC_mocks_BlockchainStatus_GetLastBlock(__m *C.BlockchainStatusMocks__Handle, _arg0 *C.Block__Handle) (____error_code uint32) {
	_m, ok_m := lookupBlockchainStatusMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetLastBlock()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = registerBlockHandle(&__arg0)
	}
	return
}

//export FC_mocks_BlockchainStatus_GetNumberOfBlocks
func FC_mocks_BlockchainStatus_GetNumberOfBlocks(__m *C.BlockchainStatusMocks__Handle, _arg0 *uint64) (____error_code uint32) {
	_m, ok_m := lookupBlockchainStatusMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetNumberOfBlocks()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = __arg0
	}
	return
}
