package main

import (
	"reflect"
	"unsafe"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_AltcoinManager_AttachSignService
func FC_mocks_AltcoinManager_AttachSignService(__m *C.AltcoinManagerMocks__Handle, __a0 *C.TxnSigner__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	___a0, ok_a0 := lookupTxnSignerHandle(*__a0)
	if !ok_a0 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *___a0
	____return_err := _m.AttachSignService(_a0)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_AltcoinManager_DescribeAltcoin
func FC_mocks_AltcoinManager_DescribeAltcoin(__m *C.AltcoinManagerMocks__Handle, _ticker string, _arg1 *C.core__AltcoinMetadata, _arg2 *bool) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, __arg2 := _m.DescribeAltcoin(ticker)
	*_arg1 = *(*C.core__AltcoinMetadata)(unsafe.Pointer(&__arg1))
	*_arg2 = __arg2
	return
}

//export FC_mocks_AltcoinManager_EnumerateSignServices
func FC_mocks_AltcoinManager_EnumerateSignServices(__m *C.AltcoinManagerMocks__Handle, _arg0 *C.TxnSignerIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.EnumerateSignServices()
	*_arg0 = registerTxnSignerIteratorHandle(&__arg0)
	return
}

//export FC_mocks_AltcoinManager_ListRegisteredPlugins
func FC_mocks_AltcoinManager_ListRegisteredPlugins(__m *C.AltcoinManagerMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ListRegisteredPlugins()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_AltcoinManager_LookupAltcoinPlugin
func FC_mocks_AltcoinManager_LookupAltcoinPlugin(__m *C.AltcoinManagerMocks__Handle, _ticker string, _arg1 *C.AltcoinPlugin__Handle, _arg2 *bool) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	ticker := _ticker
	__arg1, __arg2 := _m.LookupAltcoinPlugin(ticker)
	*_arg1 = registerAltcoinPluginHandle(&__arg1)
	*_arg2 = __arg2
	return
}

//export FC_mocks_AltcoinManager_LookupSignService
func FC_mocks_AltcoinManager_LookupSignService(__m *C.AltcoinManagerMocks__Handle, __a0 *C.core__UID, _arg1 *C.TxnSigner__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *(*core.UID)(unsafe.Pointer(__a0))
	__arg1 := _m.LookupSignService(_a0)
	*_arg1 = registerTxnSignerHandle(&__arg1)
	return
}

//export FC_mocks_AltcoinManager_RegisterAltcoin
func FC_mocks_AltcoinManager_RegisterAltcoin(__m *C.AltcoinManagerMocks__Handle, _info *C.core__AltcoinMetadata, _plugin *C.AltcoinPlugin__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	info := *(*core.AltcoinMetadata)(unsafe.Pointer(_info))
	__plugin, okplugin := lookupAltcoinPluginHandle(*_plugin)
	if !okplugin {
		____error_code = FC_BAD_HANDLE
		return
	}
	plugin := *__plugin
	_m.RegisterAltcoin(info, plugin)
	return
}

//export FC_mocks_AltcoinManager_RegisterPlugin
func FC_mocks_AltcoinManager_RegisterPlugin(__m *C.AltcoinManagerMocks__Handle, _p *C.AltcoinPlugin__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__p, okp := lookupAltcoinPluginHandle(*_p)
	if !okp {
		____error_code = FC_BAD_HANDLE
		return
	}
	p := *__p
	_m.RegisterPlugin(p)
	return
}

//export FC_mocks_AltcoinManager_RemoveSignService
func FC_mocks_AltcoinManager_RemoveSignService(__m *C.AltcoinManagerMocks__Handle, __a0 *C.TxnSigner__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	___a0, ok_a0 := lookupTxnSignerHandle(*__a0)
	if !ok_a0 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *___a0
	____return_err := _m.RemoveSignService(_a0)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_AltcoinManager_SignServicesForTxn
func FC_mocks_AltcoinManager_SignServicesForTxn(__m *C.AltcoinManagerMocks__Handle, __a0 *C.Wallet__Handle, __a1 *C.Transaction__Handle, _arg2 *C.TxnSignerIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinManagerMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	___a0, ok_a0 := lookupWalletHandle(*__a0)
	if !ok_a0 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *___a0
	___a1, ok_a1 := lookupTransactionHandle(*__a1)
	if !ok_a1 {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a1 := *___a1
	__arg2 := _m.SignServicesForTxn(_a0, _a1)
	*_arg2 = registerTxnSignerIteratorHandle(&__arg2)
	return
}
