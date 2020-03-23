package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_PexNodeSet_ListPeers
func FC_mocks_PexNodeSet_ListPeers(__m *C.PexNodeSetMocks__Handle, _arg0 *C.PexNodeIterator__Handle) (____error_code uint32) {
	_m, ok_m := lookupPexNodeSetMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.ListPeers()
	*_arg0 = registerPexNodeIteratorHandle(&__arg0)
	return
}
