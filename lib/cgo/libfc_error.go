package main

import "errors"

const (
	// FC_PKG_LIBCGO package prefix for internal API errors
	FC_PKG_LIBCGO = 0x7F000000 // nolint megacheck
	// FC_OK error code is used to report success
	FC_OK = 0
	// FC_ERROR generic error condition
	FC_ERROR = 0x7FFFFFFF
)

const (
	// FC_BAD_HANDLE invalid handle argument
	FC_BAD_HANDLE = FC_PKG_LIBCGO + iota + 1
	// FC_INVALID_TIMESTRING invalid time value
	FC_INVALID_TIMESTRING
)

// Package prefixes for error codes
//nolint megacheck
const (
	// Error code prefix for api package
	FC_PKG_API = (1 + iota) << 24 // nolint megacheck
	FC_PKG_COIN
)

// Error codes defined in params package
// nolint megacheck

var (
	// ErrorBadHandle invalid handle value
	ErrorBadHandle = errors.New("Invalid or unknown handle value")
	// ErrorUnknown unexpected error
	ErrorUnknown = errors.New("Unexpected error")
	// ErrorInvalidTimeString time string does not match expected time format
	// More precise errors conditions can be found in the logs
	ErrorInvalidTimeString = errors.New("Invalid time value")

	codeToErrorMap = make(map[uint32]error)
	errorToCodeMap = map[error]uint32{
		// libcgo
		ErrorBadHandle:         FC_BAD_HANDLE,
		ErrorUnknown:           FC_ERROR,
		ErrorInvalidTimeString: FC_INVALID_TIMESTRING,
	}
)

func libErrorCode(err error) uint32 {
	if err == nil {
		return FC_OK
	}
	if errcode, isKnownError := errorToCodeMap[err]; isKnownError {
		return errcode
	}
	return FC_ERROR
}

func init() {
	// Init reverse error code map
	for _err := range errorToCodeMap {
		codeToErrorMap[errorToCodeMap[_err]] = _err
	}
}
