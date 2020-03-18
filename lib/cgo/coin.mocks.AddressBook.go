package main

import "reflect"

/*

  #include <string.h>
  #include <stdlib.h>

  #include "fctypes.h"
*/
import "C"

//export FC_mocks_AddressBook_Authenticate
func FC_mocks_AddressBook_Authenticate(__m *C.AddressBookMocks__Handle, _password string) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	password := _password
	____return_err := _m.Authenticate(password)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_AddressBook_ChangeSecurity
func FC_mocks_AddressBook_ChangeSecurity(__m *C.AddressBookMocks__Handle, _NewSecType int, _oldPassword string, _newPassword string) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	NewSecType := _NewSecType
	oldPassword := _oldPassword
	newPassword := _newPassword
	____return_err := _m.ChangeSecurity(NewSecType, oldPassword, newPassword)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_AddressBook_Close
func FC_mocks_AddressBook_Close(__m *C.AddressBookMocks__Handle) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	____return_err := _m.Close()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_AddressBook_DeleteContact
func FC_mocks_AddressBook_DeleteContact(__m *C.AddressBookMocks__Handle, _id uint64) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	id := _id
	____return_err := _m.DeleteContact(id)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_AddressBook_GetContact
func FC_mocks_AddressBook_GetContact(__m *C.AddressBookMocks__Handle, _id uint64, _arg1 *C.Contact__Handle) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	id := _id
	__arg1, ____return_err := _m.GetContact(id)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = registerContactHandle(&__arg1)
	}
	return
}

//export FC_mocks_AddressBook_GetSecType
func FC_mocks_AddressBook_GetSecType(__m *C.AddressBookMocks__Handle, _arg0 *int) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.GetSecType()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg0 = __arg0
	}
	return
}

//export FC_mocks_AddressBook_GetStorage
func FC_mocks_AddressBook_GetStorage(__m *C.AddressBookMocks__Handle, _arg0 *C.Storage__Handle) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.GetStorage()
	*_arg0 = registerStorageHandle(&__arg0)
	return
}

//export FC_mocks_AddressBook_HasInit
func FC_mocks_AddressBook_HasInit(__m *C.AddressBookMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.HasInit()
	*_arg0 = __arg0
	return
}

//export FC_mocks_AddressBook_Init
func FC_mocks_AddressBook_Init(__m *C.AddressBookMocks__Handle, _secType int, _password string) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	secType := _secType
	password := _password
	____return_err := _m.Init(secType, password)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}

//export FC_mocks_AddressBook_InsertContact
func FC_mocks_AddressBook_InsertContact(__m *C.AddressBookMocks__Handle, _contact *C.Contact__Handle, _arg1 *uint64) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__contact, okcontact := lookupContactHandle(*_contact)
	if !okcontact {
		____error_code = FC_BAD_HANDLE
		return
	}
	contact := *__contact
	__arg1, ____return_err := _m.InsertContact(contact)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		*_arg1 = __arg1
	}
	return
}

//export FC_mocks_AddressBook_IsOpen
func FC_mocks_AddressBook_IsOpen(__m *C.AddressBookMocks__Handle, _arg0 *bool) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0 := _m.IsOpen()
	*_arg0 = __arg0
	return
}

//export FC_mocks_AddressBook_ListContact
func FC_mocks_AddressBook_ListContact(__m *C.AddressBookMocks__Handle, _arg0 *C.GoSlice_) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	__arg0, ____return_err := _m.ListContact()
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
		copyToGoSlice(reflect.ValueOf(__arg0), _arg0)
	}
	return
}

//export FC_mocks_AddressBook_UpdateContact
func FC_mocks_AddressBook_UpdateContact(__m *C.AddressBookMocks__Handle, _id uint64, _contact *C.Contact__Handle) (____error_code uint32) {
	_m, ok_m := lookupAddressBookMocksHandle(*__m)
	if !ok_m {
		____error_code = FC_BAD_HANDLE
		return
	}
	id := _id
	__contact, okcontact := lookupContactHandle(*_contact)
	if !okcontact {
		____error_code = FC_BAD_HANDLE
		return
	}
	contact := *__contact
	____return_err := _m.UpdateContact(id, contact)
	____error_code = libErrorCode(____return_err)
	if ____return_err == nil {
	}
	return
}
