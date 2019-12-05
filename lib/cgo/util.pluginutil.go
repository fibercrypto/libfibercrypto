package main

import "github.com/fibercrypto/fibercryptowallet/src/util"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_AltcoinCaption
func FC_util_AltcoinCaption(_arg0 string, _arg1 *C.GoString_) (____error_code uint32) {
	callNode()
	arg0 := string(_arg0)
	arg1 := util.AltcoinCaption(arg0)
	copyString(arg1, _arg1)
	return
}
