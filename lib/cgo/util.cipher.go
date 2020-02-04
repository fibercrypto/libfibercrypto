package main

import (
	"reflect"
	"unsafe"

	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_NewGenericAddress
func FC_util_NewGenericAddress(_addr string, _arg1 *C.util__GenericAddress) (____error_code uint32) {
	addr := _addr
	__arg1 := util.NewGenericAddress(addr)
	*_arg1 = *(*C.util__GenericAddress)(unsafe.Pointer(&__arg1))
	return
}

//export FC_util_GenericAddress_Bytes
func FC_util_GenericAddress_Bytes(_ga *C.util__GenericAddress, _arg0 *C.GoSlice_) (____error_code uint32) {
	ga := (*util.GenericAddress)(unsafe.Pointer(_ga))
	__arg0 := ga.Bytes()
	copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	return
}

//export FC_util_GenericAddress_Checksum
func FC_util_GenericAddress_Checksum(_ga *C.util__GenericAddress, _arg0 *C.core__Checksum) (____error_code uint32) {
	ga := (*util.GenericAddress)(unsafe.Pointer(_ga))
	__arg0 := ga.Checksum()
	*_arg0 = *(*C.core__Checksum)(unsafe.Pointer(&__arg0))
	return
}

//export FC_util_GenericAddress_Verify
func FC_util_GenericAddress_Verify(_ga *C.util__GenericAddress) (____error_code uint32) {
	ga := (*util.GenericAddress)(unsafe.Pointer(_ga))
	____return_err := ga.Verify(nil)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_util_GenericAddress_Null
func FC_util_GenericAddress_Null(_ga *C.util__GenericAddress, _arg0 *bool) (____error_code uint32) {
	ga := (*util.GenericAddress)(unsafe.Pointer(_ga))
	__arg0 := ga.Null()
	*_arg0 = __arg0
	return
}

//export FC_util_GenericAddress_IsBip32
func FC_util_GenericAddress_IsBip32(_ga *C.util__GenericAddress, _arg0 *bool) (____error_code uint32) {
	ga := (*util.GenericAddress)(unsafe.Pointer(_ga))
	__arg0 := ga.IsBip32()
	*_arg0 = __arg0
	return
}

//export FC_util_GenericAddress_String
func FC_util_GenericAddress_String(_ga *C.util__GenericAddress, _arg0 *C.GoString_) (____error_code uint32) {
	ga := (*util.GenericAddress)(unsafe.Pointer(_ga))
	__arg0 := ga.String()
	copyString(__arg0, _arg0)
	return
}

//export FC_util_GenericAddress_GetCryptoAccount
func FC_util_GenericAddress_GetCryptoAccount(_ga *C.util__GenericAddress, _arg0 *C.core__CryptoAccount) (____error_code uint32) {
	ga := (*util.GenericAddress)(unsafe.Pointer(_ga))
	__arg0 := ga.GetCryptoAccount()
	*_arg0 = *(*C.core__CryptoAccount)(unsafe.Pointer(&__arg0))
	return
}
