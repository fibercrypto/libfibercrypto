package main

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_SeedGenerator_GenerateMnemonic
func FC_mocks_SeedGenerator_GenerateMnemonic(__m *C.SeedGeneratorMocks__Handle, _wordCount int, _arg1 *C.GoString_) (____error_code uint32) {
	_m, ok_m := lookupSeedGeneratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	wordCount := _wordCount
	__arg1, ____return_err := _m.GenerateMnemonic(wordCount)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyString(__arg1, _arg1)
	}
	return
}

//export FC_mocks_SeedGenerator_VerifyMnemonic
func FC_mocks_SeedGenerator_VerifyMnemonic(__m *C.SeedGeneratorMocks__Handle, _seed string, _arg1 *bool) (____error_code uint32) {
	_m, ok_m := lookupSeedGeneratorMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	seed := _seed
	__arg1, ____return_err := _m.VerifyMnemonic(seed)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}
