package main

import (
	"errors"

	err "github.com/fibercrypto/fibercryptowallet/src/errors"
)

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
	FC_PKG = (1 + iota) << 24 // nolint megacheck
)

// Define all error
const (
	// FC_ErrInvalidPoolSection no section for name found in the pool
	FC_ErrInvalidPoolSection = FC_PKG + iota
	// FC_ErrObjectPoolUndeflow no objects can be allocated from the pool
	FC_ErrObjectPoolUndeflow
	// ErrInvalidLogLevelName invalid log level
	FC_ErrInvalidLogLevelName
	// ErrInvalidAltcoinTicker ticker string not bound to any altcoin plugin
	FC_ErrInvalidAltcoinTicker
	// ErrUnsupportedSigner signer not supported
	FC_ErrUnsupportedSigner
	// ErrInvalidTxn invalid transaction
	FC_ErrInvalidTxn
	// ErrInvalidOptions invalid options
	FC_ErrInvalidOptions
	// ErrIntegerInputsRequired Input IDs must be integer values
	FC_ErrIntegerInputsRequired
	// ErrBlockNotSet could not find reference to block
	FC_ErrBlockNotSet
	// ErrInvalidTransactionNoBlocks unknown number of blocks for unconfirmed transaction
	FC_ErrInvalidTransactionNoBlocks
	// ErrInvalidUnconfirmedTxn invalid unconfirmed transaction
	FC_ErrInvalidUnconfirmedTxn
	// ErrInvalidNetworkType invalid network type
	FC_ErrInvalidNetworkType
	// ErrInvalidID invalid ID
	FC_ErrInvalidID
	// ErrNotFound target item not found in collection
	FC_ErrNotFound
	// ErrParseTxID invalid string value for transaction hash ID
	FC_ErrParseTxID
	// ErrParseSHA256 invalid SHA256 hash
	FC_ErrParseSHA256
	// ErrParseTxnFee invalid string value for transaction fee
	FC_ErrParseTxnFee
	// ErrParseTxnCoins transaction coins can not be parsed
	FC_ErrParseTxnCoins
	// ErrInvalidAddressString could not decode base58 address
	FC_ErrInvalidAddressString
	// ErrTxnSignFailure signing strategy reported an error whie signing transaction
	FC_ErrTxnSignFailure
	// ErrUnexpectedUxOutAddress unexpected address
	FC_ErrUnexpectedUxOutAddress
	// ErrInvalidPoolObjectType clients in the pool do not match expected type
	FC_ErrInvalidPoolObjectType
	// ErrInvalidWalletEntropy entropy complexity does not match supported values
	FC_ErrInvalidWalletEntropy
	// ErrInvalidValue invalid value was supplied in to function
	FC_ErrInvalidValue
	// ErrWalletCantSign wallet can not sign transactions
	FC_ErrWalletCantSign
	// ErrNotImplemented feature not implemented
	FC_ErrNotImplemented
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
		ErrorBadHandle:         FC_BAD_HANDLE,
		ErrorUnknown:           FC_ERROR,
		ErrorInvalidTimeString: FC_INVALID_TIMESTRING,
		//	Error
		err.ErrInvalidPoolSection:         FC_ErrInvalidPoolSection,
		err.ErrObjectPoolUndeflow:         FC_ErrObjectPoolUndeflow,
		err.ErrInvalidLogLevelName:        FC_ErrInvalidLogLevelName,
		err.ErrInvalidAltcoinTicker:       FC_ErrInvalidAltcoinTicker,
		err.ErrUnsupportedSigner:          FC_ErrUnsupportedSigner,
		err.ErrInvalidTxn:                 FC_ErrInvalidTxn,
		err.ErrInvalidOptions:             FC_ErrInvalidOptions,
		err.ErrIntegerInputsRequired:      FC_ErrIntegerInputsRequired,
		err.ErrBlockNotSet:                FC_ErrBlockNotSet,
		err.ErrInvalidTransactionNoBlocks: FC_ErrInvalidTransactionNoBlocks,
		err.ErrInvalidUnconfirmedTxn:      FC_ErrInvalidUnconfirmedTxn,
		err.ErrInvalidNetworkType:         FC_ErrInvalidNetworkType,
		err.ErrInvalidID:                  FC_ErrInvalidID,
		err.ErrNotFound:                   FC_ErrNotFound,
		err.ErrParseTxID:                  FC_ErrParseTxID,
		err.ErrParseSHA256:                FC_ErrParseSHA256,
		err.ErrParseTxnFee:                FC_ErrParseTxnFee,
		err.ErrParseTxnCoins:              FC_ErrParseTxnCoins,
		err.ErrInvalidAddressString:       FC_ErrInvalidAddressString,
		err.ErrTxnSignFailure:             FC_ErrTxnSignFailure,
		err.ErrUnexpectedUxOutAddress:     FC_ErrUnexpectedUxOutAddress,
		err.ErrInvalidPoolObjectType:      FC_ErrInvalidPoolObjectType,
		err.ErrInvalidWalletEntropy:       FC_ErrInvalidWalletEntropy,
		err.ErrInvalidValue:               FC_ErrInvalidValue,
		err.ErrWalletCantSign:             FC_ErrWalletCantSign,
		err.ErrNotImplemented:             FC_ErrNotImplemented,
	}
)

func libErrorCode(err error) uint32 {
	if err == nil {
		return FC_OK
	}
	if errCode, isKnownError := errorToCodeMap[err]; isKnownError {
		return errCode
	}
	return FC_ERROR
}

func init() {
	// Init reverse error code map
	for _err := range errorToCodeMap {
		codeToErrorMap[errorToCodeMap[_err]] = _err
	}
}
