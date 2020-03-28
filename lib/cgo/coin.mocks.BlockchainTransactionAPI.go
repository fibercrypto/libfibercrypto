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

//export FC_mocks_BlockchainTransactionAPI_SendFromAddress
func FC_mocks_BlockchainTransactionAPI_SendFromAddress(__m *C.BlockchainTransactionAPIMocks__Handle, _from *C.WalletAddress__Handle, _to *C.TransactionOutput__Handle, _change *C.Address__Handle, _options *C.KeyValueStore__Handle, _arg4 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupBlockchainTransactionAPIMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	from := *(*[]core.WalletAddress)(unsafe.Pointer(&_from))
	to := *(*[]core.TransactionOutput)(unsafe.Pointer(&_to))
	__change, okchange := lookupAddressHandle(*_change)
	if !okchange {
		____error_code = FC_BAD_HANDLE
		return
	}
	change := *__change
	__options, okoptions := lookupKeyValueStoreHandle(*_options)
	if !okoptions {
		____error_code = FC_BAD_HANDLE
		return
	}
	options := *__options
	__arg4, ____return_err := _m.SendFromAddress(from, to, change, options)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg4 = registerTransactionHandle(&__arg4)
	}
	return
}

//export FC_mocks_BlockchainTransactionAPI_Spend
func FC_mocks_BlockchainTransactionAPI_Spend(__m *C.BlockchainTransactionAPIMocks__Handle, _unspent *C.WalletOutput__Handle, _new *C.TransactionOutput__Handle, _change *C.Address__Handle, _options *C.KeyValueStore__Handle, _arg4 *C.Transaction__Handle) (____error_code uint32) {
	_m, ok_m := lookupBlockchainTransactionAPIMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	unspent := *(*[]core.WalletOutput)(unsafe.Pointer(&_unspent))
	new := *(*[]core.TransactionOutput)(unsafe.Pointer(&_new))
	__change, okchange := lookupAddressHandle(*_change)
	if !okchange {
		____error_code = FC_BAD_HANDLE
		return
	}
	change := *__change
	__options, okoptions := lookupKeyValueStoreHandle(*_options)
	if !okoptions {
		____error_code = FC_BAD_HANDLE
		return
	}
	options := *__options
	__arg4, ____return_err := _m.Spend(unspent, new, change, options)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg4 = registerTransactionHandle(&__arg4)
	}
	return
}
