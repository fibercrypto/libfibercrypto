package main

import (
	"reflect"

	core "github.com/fibercrypto/fibercryptowallet/src/core"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_core_MultiConnectionsPool_GetSection
func FC_core_MultiConnectionsPool_GetSection(_mp *C.MultiConnectionsPool__Handle, _poolSection string, _arg1 *C.MultiPoolSection__Handle) (____error_code uint32) {
	mp, okmp := lookupMultiConnectionsPoolHandle(*_mp)
	if !okmp {
		____error_code = FC_BAD_HANDLE
		return
	}
	poolSection := _poolSection
	__arg1, ____return_err := mp.GetSection(poolSection)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerMultiPoolSectionHandle(&__arg1)
	}
	return
}

//export FC_core_MultiConnectionsPool_CreateSection
func FC_core_MultiConnectionsPool_CreateSection(_mp *C.MultiConnectionsPool__Handle, _name string, _factory *C.PooledObjectFactory__Handle) (____error_code uint32) {
	mp, okmp := lookupMultiConnectionsPoolHandle(*_mp)
	if !okmp {
		____error_code = FC_BAD_HANDLE
		return
	}
	name := _name
	__factory, okfactory := lookupPooledObjectFactoryHandle(*_factory)
	if !okfactory {
		____error_code = FC_BAD_HANDLE
		return
	}
	factory := *__factory
	____return_err := mp.CreateSection(name, factory)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_core_MultiConnectionsPool_ListSections
func FC_core_MultiConnectionsPool_ListSections(_mp *C.MultiConnectionsPool__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	mp, okmp := lookupMultiConnectionsPoolHandle(*_mp)
	if !okmp {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := mp.ListSections()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	}
	return
}

//export FC_core_GetMultiPool
func FC_core_GetMultiPool(_arg0 *C.MultiPool__Handle) (____error_code uint32) {
	__arg0 := core.GetMultiPool()
	*_arg0 = registerMultiPoolHandle(&__arg0)
	return
}
