package main

/*

  #include <string.h>
  #include <stdlib.h>


  #include "fctypes.h"
*/
import "C"

import (
	"encoding/json"
)

//export FC_JsonEncode_Handle
func FC_JsonEncode_Handle(handle C.Handle, json_string *C.GoString_) uint32 {
	obj, ok := lookupHandle(handle)
	if ok {
		jsonBytes, err := json.Marshal(obj)
		if err == nil {
			copyString(string(jsonBytes), json_string)
			return FC_OK
		}
	}
	return FC_BAD_HANDLE
}
