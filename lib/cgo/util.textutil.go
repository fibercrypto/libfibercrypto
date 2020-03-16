package main

import (
	"unsafe"

	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_util_EmptyPassword
func FC_util_EmptyPassword(_arg2 *C.GoString_) (____error_code uint32) {
	__arg2, ____return_err := util.EmptyPassword("", nil)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyString(__arg2, _arg2)
	}
	return
}

//export FC_util_ConstantPassword
func FC_util_ConstantPassword(_pwdText string, _arg1 *C.PasswordReader) (____error_code uint32) {
	pwdText := _pwdText
	__arg1 := util.ConstantPassword(pwdText)
	*_arg1 = *(*C.PasswordReader)(unsafe.Pointer(&__arg1))
	return
}

//export FC_util_MessageFromMsgAndArgs
func FC_util_MessageFromMsgAndArgs(_msgAndArgs, _arg1 *C.GoString_) (____error_code uint32) {
	msgAndArgs := _msgAndArgs
	__arg1 := util.MessageFromMsgAndArgs(msgAndArgs)
	copyString(__arg1, _arg1)
	return
}

//export FC_util_IndentMessageLines
func FC_util_IndentMessageLines(_message string, _longestLabelLen int, _arg2 *C.GoString_) (____error_code uint32) {
	message := _message
	longestLabelLen := _longestLabelLen
	__arg2 := util.IndentMessageLines(message, longestLabelLen)
	copyString(__arg2, _arg2)
	return
}
