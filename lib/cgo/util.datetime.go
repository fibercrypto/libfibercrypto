package main

import util "github.com/fibercrypto/fibercryptowallet/src/util"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_ParseDate
func FC_util_ParseDate(_timeStamp int64, _arg1 *int, _arg2 *int, _arg3 *int, _arg4 *int, _arg5 *int, _arg6 *int) (____error_code uint32) {
	timeStamp := _timeStamp
	__arg1, __arg2, __arg3, __arg4, __arg5, __arg6 := util.ParseDate(timeStamp)
	*_arg1 = __arg1
	*_arg2 = __arg2
	*_arg3 = __arg3
	*_arg4 = __arg4
	*_arg5 = __arg5
	*_arg6 = __arg6
	return
}
