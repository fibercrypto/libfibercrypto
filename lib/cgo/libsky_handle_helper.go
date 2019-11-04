package main

/*

  #include <string.h>
  #include <stdlib.h>


  #include "skytypes.h"
*/
import "C"

import (
	"encoding/json"
)

//export SKY_JsonEncode_Handle
func SKY_JsonEncode_Handle(handle C.Handle, json_string *C.GoString_) uint32 {
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

// //export SKY_Handle_Progress_GetCurrent
// func SKY_Handle_Progress_GetCurrent(handle C.Handle, current *uint64) uint32 {
// 	obj, ok := lookupHandle(C.Handle(handle))
// 	if ok {
// 		if obj, isOK := (obj).(*daemon.BlockchainProgress); isOK {
// 			*current = obj.Current
// 			return SKY_OK
// 		}
// 	}
// 	return SKY_BAD_HANDLE
// }
