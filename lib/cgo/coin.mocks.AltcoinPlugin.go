package main

import (
	"reflect"
	"unsafe"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_AltcoinPlugin_AddressFromString
func FC_mocks_AltcoinPlugin_AddressFromString(__m *C.AltcoinPluginMocks__Handle, __a0 string, _arg1 *C.Address__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := __a0
	__arg1, ____return_err := _m.AddressFromString(_a0)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerAddressHandle(&__arg1)
	}
	return
}

//export FC_mocks_AltcoinPlugin_GetDescription
func FC_mocks_AltcoinPlugin_GetDescription(__m *C.AltcoinPluginMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetDescription()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_AltcoinPlugin_GetName
func FC_mocks_AltcoinPlugin_GetName(__m *C.AltcoinPluginMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetName()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_AltcoinPlugin_ListSupportedAltcoins
func FC_mocks_AltcoinPlugin_ListSupportedAltcoins(__m *C.AltcoinPluginMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ListSupportedAltcoins()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_AltcoinPlugin_ListSupportedFamilies
func FC_mocks_AltcoinPlugin_ListSupportedFamilies(__m *C.AltcoinPluginMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ListSupportedFamilies()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_AltcoinPlugin_LoadPEX
func FC_mocks_AltcoinPlugin_LoadPEX(__m *C.AltcoinPluginMocks__Handle, _netType string, _arg1 *C.PEX__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	netType := _netType
	__arg1, ____return_err := _m.LoadPEX(netType)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerPEXHandle(&__arg1)
	}
	return
}

//export FC_mocks_AltcoinPlugin_LoadSignService
func FC_mocks_AltcoinPlugin_LoadSignService(__m *C.AltcoinPluginMocks__Handle, _arg0 *C.BlockchainSignService__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.LoadSignService()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = registerBlockchainSignServiceHandle(&__arg0)
	}
	return
}

//export FC_mocks_AltcoinPlugin_LoadTransactionAPI
func FC_mocks_AltcoinPlugin_LoadTransactionAPI(__m *C.AltcoinPluginMocks__Handle, _netType string, _arg1 *C.BlockchainTransactionAPI__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	netType := _netType
	__arg1, ____return_err := _m.LoadTransactionAPI(netType)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerBlockchainTransactionAPIHandle(&__arg1)
	}
	return
}

//export FC_mocks_AltcoinPlugin_LoadWalletEnvs
func FC_mocks_AltcoinPlugin_LoadWalletEnvs(__m *C.AltcoinPluginMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.LoadWalletEnvs()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_mocks_AltcoinPlugin_PubKeyFromBytes
func FC_mocks_AltcoinPlugin_PubKeyFromBytes(__m *C.AltcoinPluginMocks__Handle, __a0 []byte, _arg1 *C.PubKey__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *(*[]byte)(unsafe.Pointer(&__a0))
	__arg1, ____return_err := _m.PubKeyFromBytes(_a0)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerPubKeyHandle(&__arg1)
	}
	return
}

//export FC_mocks_AltcoinPlugin_RegisterTo
func FC_mocks_AltcoinPlugin_RegisterTo(__m *C.AltcoinPluginMocks__Handle, _manager *C.AltcoinManager__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__manager, okmanager := lookupAltcoinManagerHandle(*_manager)
	if !okmanager {
		____error_code = FC_BAD_HANDLE
		return
	}
	manager := *__manager
	_m.RegisterTo(manager)
	return
}

//export FC_mocks_AltcoinPlugin_SecKeyFromBytes
func FC_mocks_AltcoinPlugin_SecKeyFromBytes(__m *C.AltcoinPluginMocks__Handle, __a0 []byte, _arg1 *C.SecKey__Handle) (____error_code uint32) {
	_m, ok_m := lookupAltcoinPluginMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	_a0 := *(*[]byte)(unsafe.Pointer(&__a0))
	__arg1, ____return_err := _m.SecKeyFromBytes(_a0)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerSecKeyHandle(&__arg1)
	}
	return
}
