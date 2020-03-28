package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_PexNode_GetBlockHeight
func FC_mocks_PexNode_GetBlockHeight(__m *C.PexNodeMocks__Handle, _arg0 *uint64) (____error_code uint32) {
	_m, ok_m := lookupPexNodeMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetBlockHeight()
	*_arg0 = __arg0
	return
}

//export FC_mocks_PexNode_GetIp
func FC_mocks_PexNode_GetIp(__m *C.PexNodeMocks__Handle, _arg0 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupPexNodeMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetIp()
	copyString(__arg0, _arg0)
	return
}

//export FC_mocks_PexNode_GetLastSeenIn
func FC_mocks_PexNode_GetLastSeenIn(__m *C.PexNodeMocks__Handle, _arg0 *int64) (____error_code uint32) {
	_m, ok_m := lookupPexNodeMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetLastSeenIn()
	*_arg0 = __arg0
	return
}

//export FC_mocks_PexNode_GetLastSeenOut
func FC_mocks_PexNode_GetLastSeenOut(__m *C.PexNodeMocks__Handle, _arg0 *int64) (____error_code uint32) {
	_m, ok_m := lookupPexNodeMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetLastSeenOut()
	*_arg0 = __arg0
	return
}

//export FC_mocks_PexNode_GetPort
func FC_mocks_PexNode_GetPort(__m *C.PexNodeMocks__Handle, _arg0 *uint16) (____error_code uint32) {
	_m, ok_m := lookupPexNodeMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetPort()
	*_arg0 = __arg0
	return
}

//export FC_mocks_PexNode_IsTrusted
func FC_mocks_PexNode_IsTrusted(__m *C.PexNodeMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupPexNodeMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.IsTrusted()
	*_arg0 = __arg0
	return
}
