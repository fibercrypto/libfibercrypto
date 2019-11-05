package main

/*

  #include <string.h>
  #include <stdlib.h>


  #include "fctypes.h"
*/
import "C"

type Handle uint64

var (
	handlesCounter uint64 = 0
	handleMap             = make(map[Handle]interface{})
)

func registerHandle(obj interface{}) C.Handle {
	handlesCounter++
	handle := handlesCounter
	//handle := *(*Handle)(unsafe.Pointer(&obj))
	handleMap[Handle(handle)] = obj
	return (C.Handle)(handle)
}

func lookupHandle(handle C.Handle) (interface{}, bool) {
	obj, ok := handleMap[Handle(handle)]
	return obj, ok
}

func overwriteHandle(handle C.Handle, obj interface{}) bool {
	_, ok := handleMap[Handle(handle)]
	if ok {
		handleMap[Handle(handle)] = obj
		return true
	}
	return false
}

func closeHandle(handle Handle) {
	delete(handleMap, handle)
}

//export FC_handle_close
func FC_handle_close(handle C.Handle) {
	closeHandle(Handle(handle))
}
func registerStringsHandle(obj []string) C.Strings__Handle {
	return (C.Strings__Handle)(registerHandle(obj))
}

func lookupStringsHandle(handle C.Strings__Handle) ([]string, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).([]string); isOK {
			return obj, true
		}
	}
	return nil, false
}
