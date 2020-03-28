package main

import (
	"reflect"
)

/*
  #include <string.h>
  #include "fctypes.h"



*/
import "C"

//export FC_StringtoByte
func FC_StringtoByte(_src string, _dest *C.GoSlice_) {

	dest := []byte(_src)
	copyToGoSlice(reflect.ValueOf(dest), _dest)
}
