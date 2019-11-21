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

// // Error codes defined in cipher package
// //nolint megacheck
// const (
// 	// SKY_ErrAddressInvalidLength Unexpected size of address bytes buffer
// 	SKY_ErrAddressInvalidLength = SKY_PKG_CIPHER + iota
// 	// SKY_ErrAddressInvalidChecksum Computed checksum did not match expected value
// 	SKY_ErrAddressInvalidChecksum
// 	// SKY_ErrAddressInvalidVersion Unsupported address version value
// 	SKY_ErrAddressInvalidVersion
// 	// SKY_ErrAddressInvalidPubKey Public key invalid for address
// 	SKY_ErrAddressInvalidPubKey
// 	// SKY_ErrAddressInvalidFirstByte Invalid first byte in wallet import format string
// 	SKY_ErrAddressInvalidFirstByte
// 	// SKY_ErrAddressInvalidLastByte 33rd byte in wallet import format string is invalid
// 	SKY_ErrAddressInvalidLastByte
// 	// SKY_ErrBufferUnderflow bytes in input buffer not enough to deserialize expected type
// 	SKY_ErrBufferUnderflow
// 	// SKY_ErrInvalidOmitEmpty field tagged with omitempty and it's not last one in struct
// 	SKY_ErrInvalidOmitEmpty
// 	// SKY_ErrInvalidLengthPubKey  Invalid public key length
// 	SKY_ErrInvalidLengthPubKey
// 	// SKY_ErrPubKeyFromNullSecKey PubKeyFromSecKey, attempt to load null seckey, unsafe
// 	SKY_ErrPubKeyFromNullSecKey
// 	// SKY_ErrPubKeyFromBadSecKey  PubKeyFromSecKey, pubkey recovery failed. Function
// 	SKY_ErrPubKeyFromBadSecKey
// 	// SKY_ErrInvalidLengthSecKey Invalid secret key length
// 	SKY_ErrInvalidLengthSecKey
// 	// SKY_ErrECHDInvalidPubKey   ECDH invalid pubkey input
// 	SKY_ErrECHDInvalidPubKey
// 	// SKY_ErrECHDInvalidSecKey   ECDH invalid seckey input
// 	SKY_ErrECHDInvalidSecKey
// 	// SKY_ErrInvalidLengthSig    Invalid signature length
// 	SKY_ErrInvalidLengthSig
// 	// SKY_ErrInvalidLengthRipemd160 Invalid ripemd160 length
// 	SKY_ErrInvalidLengthRipemd160
// 	// SKY_ErrInvalidLengthSHA256 Invalid sha256 length
// 	SKY_ErrInvalidLengthSHA256
// 	// SKY_ErrInvalidBase58Char   Invalid base58 character
// 	SKY_ErrInvalidBase58Char
// 	// SKY_ErrInvalidBase58String Invalid base58 string
// 	SKY_ErrInvalidBase58String
// 	// SKY_ErrInvalidBase58Length Invalid base58 length
// 	SKY_ErrInvalidBase58Length
// 	// SKY_ErrInvalidHexLength       Invalid hex length
// 	SKY_ErrInvalidHexLength
// 	// SKY_ErrInvalidBytesLength     Invalid bytes length
// 	SKY_ErrInvalidBytesLength
// 	// SKY_ErrInvalidPubKey       Invalid public key
// 	SKY_ErrInvalidPubKey
// 	// SKY_ErrInvalidSecKey       Invalid public key
// 	SKY_ErrInvalidSecKey
// 	// SKY_ErrInvalidSigPubKeyRecovery Invalig sig: PubKey recovery failed
// 	SKY_ErrInvalidSigPubKeyRecovery
// 	// SKY_ErrInvalidSecKeyHex    Invalid SecKey: not valid hex
// 	SKY_ErrInvalidSecKeyHex // nolint megacheck
// 	// SKY_ErrInvalidAddressForSig Invalid sig: address does not match output address
// 	SKY_ErrInvalidAddressForSig
// 	// SKY_ErrInvalidHashForSig   Signature invalid for hash
// 	SKY_ErrInvalidHashForSig
// 	// SKY_ErrPubKeyRecoverMismatch Recovered pubkey does not match pubkey
// 	SKY_ErrPubKeyRecoverMismatch
// 	// SKY_ErrInvalidSigInvalidPubKey VerifySignedHash, secp256k1.VerifyPubkey failed
// 	SKY_ErrInvalidSigInvalidPubKey
// 	// SKY_ErrInvalidSigValidity  VerifySignedHash, VerifySignatureValidity failed
// 	SKY_ErrInvalidSigValidity
// 	// SKY_ErrInvalidSigForMessage Invalid signature for this message
// 	SKY_ErrInvalidSigForMessage
// 	// SKY_ErrInvalidSecKyVerification Seckey secp256k1 verification failed
// 	SKY_ErrInvalidSecKyVerification
// 	// SKY_ErrNullPubKeyFromSecKey Impossible error, CheckSecKey, nil pubkey recovered
// 	SKY_ErrNullPubKeyFromSecKey
// 	// SKY_ErrInvalidDerivedPubKeyFromSecKey impossible error, CheckSecKey, Derived Pubkey verification failed
// 	SKY_ErrInvalidDerivedPubKeyFromSecKey
// 	// SKY_ErrInvalidPubKeyFromHash Recovered pubkey does not match signed hash
// 	SKY_ErrInvalidPubKeyFromHash
// 	// SKY_ErrPubKeyFromSecKeyMismatch impossible error CheckSecKey, pubkey does not match recovered pubkey
// 	SKY_ErrPubKeyFromSecKeyMismatch
// 	// SKY_ErrInvalidLength Unexpected size of string or bytes buffer
// 	SKY_ErrInvalidLength
// 	// SKY_ErrBitcoinWIFInvalidFirstByte Unexpected value (!= 0x80) of first byte in Bitcoin Wallet Import Format
// 	SKY_ErrBitcoinWIFInvalidFirstByte
// 	// SKY_ErrBitcoinWIFInvalidSuffix Unexpected value (!= 0x01) of 33rd byte in Bitcoin Wallet Import Format
// 	SKY_ErrBitcoinWIFInvalidSuffix
// 	// SKY_ErrBitcoinWIFInvalidChecksum Invalid Checksum in Bitcoin WIF address
// 	SKY_ErrBitcoinWIFInvalidChecksum
// 	// SKY_ErrEmptySeed Seed input is empty
// 	SKY_ErrEmptySeed
// 	// SKY_ErrInvalidSig Invalid signature
// 	SKY_ErrInvalidSig
// 	// SKY_ErrMissingPassword missing password
// 	SKY_ErrMissingPassword
// 	// SKY_SKY_ErrDataTooLarge data length overflowed, it must <= math.MaxUint32(4294967295)
// 	SKY_ErrDataTooLarge
// 	// SKY_ErrInvalidChecksumLength invalid checksum length
// 	SKY_ErrInvalidChecksumLength
// 	// SKY_ErrInvalidChecksum invalid data, checksum is not matched
// 	SKY_ErrInvalidChecksum
// 	// SKY_ErrInvalidNonceLength invalid nonce length
// 	SKY_ErrInvalidNonceLength
// 	// SKY_ErrInvalidBlockSize invalid block size, must be multiple of 32 bytes
// 	SKY_ErrInvalidBlockSize
// 	// SKY_ErrReadDataHashFailed read data hash failed: read length != 32
// 	SKY_ErrReadDataHashFailed
// 	// SKY_ErrInvalidPassword invalid password SHA256or
// 	SKY_ErrInvalidPassword
// 	// SKY_ErrReadDataLengthFailed read data length failed
// 	SKY_ErrReadDataLengthFailed
// 	// SKY_ErrInvalidDataLength invalid data length
// 	SKY_ErrInvalidDataLength

// 	// bip32
// 	SKY_ErrSerializedKeyWrongSize
// 	SKY_ErrHardenedChildPublicKey
// 	SKY_bip32_ErrInvalidChecksum
// 	SKY_ErrDerivedInvalidPrivateKey
// 	SKY_ErrDerivedInvalidPublicKey
// 	SKY_ErrInvalidPrivateKeyVersion
// 	SKY_ErrInvalidPublicKeyVersion
// 	SKY_ErrInvalidSeedLength
// 	SKY_ErrDeserializePrivateFromPublic
// 	SKY_ErrInvalidKeyVersion
// 	SKY_ErrInvalidFingerprint
// 	SKY_ErrInvalidChildNumber
// 	SKY_ErrInvalidPrivateKey
// 	SKY_ErrInvalidPublicKey
// 	SKY_ErrMaxDepthReached
// 	// bip44
// 	// ErrInvalidCoinType coin_type is >= 0x80000000
// 	SKY_ErrInvalidCoinType
// 	// ErrInvalidAccount account is >= 0x80000000
// 	SKY_ErrInvalidAccount
// 	// bip32.path
// 	// SKY_ErrPathNoMaster HD wallet path does not start with m
// 	SKY_ErrPathNoMaster
// 	// SKY_ErrPathChildMaster HD wallet path contains m in a child node
// 	SKY_ErrPathChildMaster
// 	// SKY_ErrPathNodeNotNumber HD wallet path node is not a valid uint32 number
// 	SKY_ErrPathNodeNotNumber
// 	// SKY_ErrPathNodeNumberTooLarge HD wallet path node is >= 2^31
// 	SKY_ErrPathNodeNumberTooLarge
// )

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
		// 		// cipher
		// 		cipher.ErrAddressInvalidLength:    SKY_ErrAddressInvalidLength,
		// 		cipher.ErrAddressInvalidChecksum:  SKY_ErrAddressInvalidChecksum,
		// 		cipher.ErrAddressInvalidVersion:   SKY_ErrAddressInvalidVersion,
		// 		cipher.ErrAddressInvalidPubKey:    SKY_ErrAddressInvalidPubKey,
		// 		cipher.ErrAddressInvalidFirstByte: SKY_ErrAddressInvalidFirstByte,
		// 		cipher.ErrAddressInvalidLastByte:  SKY_ErrAddressInvalidLastByte,
		// 		encoder.ErrBufferUnderflow:        SKY_ErrBufferUnderflow,
		// 		encoder.ErrInvalidOmitEmpty:       SKY_ErrInvalidOmitEmpty,
		// 		cipher.ErrInvalidLengthPubKey:     SKY_ErrInvalidLengthPubKey,
		// 		cipher.ErrPubKeyFromNullSecKey:    SKY_ErrPubKeyFromNullSecKey,
		// 		cipher.ErrPubKeyFromBadSecKey:     SKY_ErrPubKeyFromBadSecKey,
		// 		cipher.ErrInvalidLengthSecKey:     SKY_ErrInvalidLengthSecKey,
		// 		cipher.ErrECHDInvalidPubKey:       SKY_ErrECHDInvalidPubKey,
		// 		cipher.ErrECHDInvalidSecKey:       SKY_ErrECHDInvalidSecKey,
		// 		cipher.ErrInvalidLengthSig:        SKY_ErrInvalidLengthSig,
		// 		cipher.ErrInvalidLengthRipemd160:  SKY_ErrInvalidLengthRipemd160,
		// 		cipher.ErrInvalidLengthSHA256:     SKY_ErrInvalidLengthSHA256,
		// 		// base58.ErrInvalidBase58Char:        SKY_ErrInvalidBase58Char,
		// 		// base58.ErrInvalidBase58String:      SKY_ErrInvalidBase58String,
		// 		// base58.ErrInvalidBase58Length:      SKY_ErrInvalidBase58Length,
		// 		cipher.ErrInvalidHexLength:         SKY_ErrInvalidHexLength,
		// 		cipher.ErrInvalidBytesLength:       SKY_ErrInvalidBytesLength,
		// 		cipher.ErrInvalidPubKey:            SKY_ErrInvalidPubKey,
		// 		cipher.ErrInvalidSecKey:            SKY_ErrInvalidSecKey,
		// 		cipher.ErrInvalidSigPubKeyRecovery: SKY_ErrInvalidSigPubKeyRecovery,
		// 		// Removed in ea0aafbffb76
		// 		// cipher.ErrInvalidSecKeyHex:               SKY_ErrInvalidSecKeyHex,
		// 		cipher.ErrInvalidAddressForSig:           SKY_ErrInvalidAddressForSig,
		// 		cipher.ErrInvalidHashForSig:              SKY_ErrInvalidHashForSig,
		// 		cipher.ErrPubKeyRecoverMismatch:          SKY_ErrPubKeyRecoverMismatch,
		// 		cipher.ErrInvalidSigInvalidPubKey:        SKY_ErrInvalidSigInvalidPubKey,
		// 		cipher.ErrInvalidSigValidity:             SKY_ErrInvalidSigValidity,
		// 		cipher.ErrInvalidSigForMessage:           SKY_ErrInvalidSigForMessage,
		// 		cipher.ErrInvalidSecKyVerification:       SKY_ErrInvalidSecKyVerification,
		// 		cipher.ErrNullPubKeyFromSecKey:           SKY_ErrNullPubKeyFromSecKey,
		// 		cipher.ErrInvalidDerivedPubKeyFromSecKey: SKY_ErrInvalidDerivedPubKeyFromSecKey,
		// 		cipher.ErrInvalidPubKeyFromHash:          SKY_ErrInvalidPubKeyFromHash,
		// 		cipher.ErrPubKeyFromSecKeyMismatch:       SKY_ErrPubKeyFromSecKeyMismatch,
		// 		cipher.ErrInvalidLength:                  SKY_ErrInvalidLength,
		// 		cipher.ErrBitcoinWIFInvalidFirstByte:     SKY_ErrBitcoinWIFInvalidFirstByte,
		// 		cipher.ErrBitcoinWIFInvalidSuffix:        SKY_ErrBitcoinWIFInvalidSuffix,
		// 		cipher.ErrBitcoinWIFInvalidChecksum:      SKY_ErrBitcoinWIFInvalidChecksum,
		// 		cipher.ErrEmptySeed:                      SKY_ErrEmptySeed,
		// 		cipher.ErrInvalidSig:                     SKY_ErrInvalidSig,
		// 		encrypt.ErrMissingPassword:               SKY_ErrMissingPassword,
		// 		encrypt.ErrDataTooLarge:                  SKY_ErrDataTooLarge,
		// 		encrypt.ErrInvalidChecksumLength:         SKY_ErrInvalidChecksumLength,
		// 		encrypt.ErrInvalidChecksum:               SKY_ErrInvalidChecksum,
		// 		encrypt.ErrInvalidNonceLength:            SKY_ErrInvalidNonceLength,
		// 		encrypt.ErrInvalidBlockSize:              SKY_ErrInvalidBlockSize,
		// 		encrypt.ErrReadDataHashFailed:            SKY_ErrReadDataHashFailed,
		// 		encrypt.ErrInvalidPassword:               SKY_ErrInvalidPassword,
		// 		encrypt.ErrReadDataLengthFailed:          SKY_ErrReadDataLengthFailed,
		// 		encrypt.ErrInvalidDataLength:             SKY_ErrInvalidDataLength,

		// 		// cli
		// 		cli.ErrTemporaryInsufficientBalance: SKY_ErrTemporaryInsufficientBalance,
		// 		cli.ErrAddress:                      SKY_ErrAddress,
		// 		cli.ErrWalletName:                   SKY_ErrWalletName,
		// 		cli.ErrJSONMarshal:                  SKY_ErrJSONMarshal,
		// 		// coin
		// 		coin.ErrAddEarnedCoinHoursAdditionOverflow: SKY_ErrAddEarnedCoinHoursAdditionOverflow,
		// 		mathutil.ErrUint64MultOverflow:             SKY_ErrUint64MultOverflow,
		// 		mathutil.ErrUint64AddOverflow:              SKY_ErrUint64AddOverflow,
		// 		mathutil.ErrUint32AddOverflow:              SKY_ErrUint32AddOverflow,
		// 		mathutil.ErrUint64OverflowsInt64:           SKY_ErrUint64OverflowsInt64,
		// 		mathutil.ErrInt64UnderflowsUint64:          SKY_ErrInt64UnderflowsUint64,
		// 		mathutil.ErrIntUnderflowsUint32:            SKY_ErrIntUnderflowsUint32,
		// 		mathutil.ErrIntOverflowsUint32:             SKY_ErrIntOverflowsUint32,
		// 		// daemon
		// 		// Removed in 34ad39ddb350
		// 		// gnet.ErrMaxDefaultConnectionsReached:           SKY_ErrMaxDefaultConnectionsReached,
		// 		pex.ErrPeerlistFull:       SKY_ErrPeerlistFull,
		// 		pex.ErrInvalidAddress:     SKY_ErrInvalidAddress,
		// 		pex.ErrNoLocalhost:        SKY_ErrNoLocalhost,
		// 		pex.ErrNotExternalIP:      SKY_ErrNotExternalIP,
		// 		pex.ErrPortTooLow:         SKY_ErrPortTooLow,
		// 		pex.ErrBlacklistedAddress: SKY_ErrBlacklistedAddress,
		// 		// gnet.ErrDisconnectReadFailed:              SKY_ErrDisconnectReadFailed,
		// 		// gnet.ErrDisconnectWriteFailed:             SKY_ErrDisconnectWriteFailed,
		// 		gnet.ErrDisconnectSetReadDeadlineFailed: SKY_ErrDisconnectSetReadDeadlineFailed,
		// 		gnet.ErrDisconnectInvalidMessageLength:  SKY_ErrDisconnectInvalidMessageLength,
		// 		gnet.ErrDisconnectMalformedMessage:      SKY_ErrDisconnectMalformedMessage,
		// 		gnet.ErrDisconnectUnknownMessage:        SKY_ErrDisconnectUnknownMessage,
		// 		gnet.ErrConnectionPoolClosed:            SKY_ErrConnectionPoolClosed,
		// 		gnet.ErrWriteQueueFull:                  SKY_ErrWriteQueueFull,
		// 		gnet.ErrNoReachableConnections:          SKY_ErrNoReachableConnections,
		// 		daemon.ErrDisconnectVersionNotSupported: SKY_ErrDisconnectVersionNotSupported,
		// 		daemon.ErrDisconnectIntroductionTimeout: SKY_ErrDisconnectIntroductionTimeout,
		// 		daemon.ErrDisconnectIsBlacklisted:       SKY_ErrDisconnectIsBlacklisted,
		// 		daemon.ErrDisconnectSelf:                SKY_ErrDisconnectSelf,
		// 		daemon.ErrDisconnectConnectedTwice:      SKY_ErrDisconnectConnectedTwice,
		// 		daemon.ErrDisconnectIdle:                SKY_ErrDisconnectIdle,
		// 		daemon.ErrDisconnectNoIntroduction:      SKY_ErrDisconnectNoIntroduction,
		// 		daemon.ErrDisconnectIPLimitReached:      SKY_ErrDisconnectIPLimitReached,
		// 		// Removed
		// 		//		daemon.ErrDisconnectMaxDefaultConnectionReached:   SKY_ErrDisconnectMaxDefaultConnectionReached,
		// 		daemon.ErrDisconnectMaxOutgoingConnectionsReached: SKY_ErrDisconnectMaxOutgoingConnectionsReached,
		// 		// util
		// 		fee.ErrTxnNoFee:                 SKY_ErrTxnNoFee,
		// 		fee.ErrTxnInsufficientFee:       SKY_ErrTxnInsufficientFee,
		// 		fee.ErrTxnInsufficientCoinHours: SKY_ErrTxnInsufficientCoinHours,
		// 		droplet.ErrNegativeValue:        SKY_ErrNegativeValue,
		// 		droplet.ErrTooManyDecimals:      SKY_ErrTooManyDecimals,
		// 		droplet.ErrTooLarge:             SKY_ErrTooLarge,
		// 		file.ErrEmptyDirectoryName:      SKY_ErrEmptyDirectoryName,
		// 		file.ErrDotDirectoryName:        SKY_ErrDotDirectoryName,
		// 		// visor
		// 		blockdb.ErrNoHeadBlock: SKY_ErrNoHeadBlock,
		// 		visor.ErrVerifyStopped: SKY_ErrVerifyStopped,
		// 		// wallet
		// 		transaction.ErrInsufficientBalance:       SKY_ErrInsufficientBalance,
		// 		transaction.ErrInsufficientHours:         SKY_ErrInsufficientHours,
		// 		transaction.ErrZeroSpend:                 SKY_ErrZeroSpend,
		// 		visor.ErrSpendingUnconfirmed:             SKY_ErrSpendingUnconfirmed,
		// 		wallet.ErrInvalidEncryptedField:          SKY_ErrInvalidEncryptedField,
		// 		wallet.ErrWalletEncrypted:                SKY_ErrWalletEncrypted,
		// 		wallet.ErrWalletNotEncrypted:             SKY_ErrWalletNotEncrypted,
		// 		wallet.ErrMissingPassword:                SKY_ErrWalletMissingPassword,
		// 		wallet.ErrMissingEncrypt:                 SKY_ErrMissingEncrypt,
		// 		wallet.ErrInvalidPassword:                SKY_ErrWalletInvalidPassword,
		// 		wallet.ErrMissingSeed:                    SKY_ErrMissingSeed,
		// 		wallet.ErrMissingAuthenticated:           SKY_ErrMissingAuthenticated,
		// 		wallet.ErrWrongCryptoType:                SKY_ErrWrongCryptoType,
		// 		wallet.ErrWalletNotExist:                 SKY_ErrWalletNotExist,
		// 		wallet.ErrSeedUsed:                       SKY_ErrSeedUsed,
		// 		wallet.ErrWalletAPIDisabled:              SKY_ErrWalletAPIDisabled,
		// 		wallet.ErrSeedAPIDisabled:                SKY_ErrSeedAPIDisabled,
		// 		wallet.ErrWalletNameConflict:             SKY_ErrWalletNameConflict,
		// 		transaction.ErrInvalidHoursSelectionType: SKY_ErrInvalidHoursSelectionType,
		// 		wallet.ErrUnknownAddress:                 SKY_ErrUnknownAddress,
		// 		wallet.ErrUnknownUxOut:                   SKY_ErrUnknownUxOut,
		// 		transaction.ErrNoUnspents:                SKY_ErrNoUnspents,
		// 		transaction.ErrNullChangeAddress:         SKY_ErrNullChangeAddress,
		// 		visor.ErrIncludesNullAddress:             SKY_ErrIncludesNullAddress,
		// 		visor.ErrDuplicateAddresses:              SKY_ErrDuplicateAddresses,
		// 		transaction.ErrMissingShareFactor:        SKY_ErrMissingShareFactor,
		// 		transaction.ErrInvalidShareFactor:        SKY_ErrInvalidShareFactor,
		// 		transaction.ErrShareFactorOutOfRange:     SKY_ErrShareFactorOutOfRange,
		// 		visor.ErrDuplicateUxOuts:                 SKY_ErrDuplicateUxOuts,
		// 		// params
		// 		params.ErrInvalidDecimals: SKY_ErrInvalidDecimals,
		// 		// bip32
		// 		bip32.ErrSerializedKeyWrongSize:       SKY_ErrSerializedKeyWrongSize,
		// 		bip32.ErrHardenedChildPublicKey:       SKY_ErrHardenedChildPublicKey,
		// 		bip32.ErrInvalidChecksum:              SKY_bip32_ErrInvalidChecksum,
		// 		bip32.ErrDerivedInvalidPrivateKey:     SKY_ErrDerivedInvalidPrivateKey,
		// 		bip32.ErrDerivedInvalidPublicKey:      SKY_ErrDerivedInvalidPublicKey,
		// 		bip32.ErrInvalidPrivateKeyVersion:     SKY_ErrInvalidPrivateKeyVersion,
		// 		bip32.ErrInvalidPublicKeyVersion:      SKY_ErrInvalidPublicKeyVersion,
		// 		bip32.ErrInvalidSeedLength:            SKY_ErrInvalidSeedLength,
		// 		bip32.ErrDeserializePrivateFromPublic: SKY_ErrDeserializePrivateFromPublic,
		// 		bip32.ErrInvalidKeyVersion:            SKY_ErrInvalidKeyVersion,
		// 		bip32.ErrInvalidFingerprint:           SKY_ErrInvalidFingerprint,
		// 		bip32.ErrInvalidChildNumber:           SKY_ErrInvalidChildNumber,
		// 		bip32.ErrInvalidPrivateKey:            SKY_ErrInvalidPrivateKey,
		// 		bip32.ErrInvalidPublicKey:             SKY_ErrInvalidPublicKey,
		// 		bip32.ErrMaxDepthReached:              SKY_ErrMaxDepthReached,
		// 		bip44.ErrInvalidCoinType:              SKY_ErrInvalidCoinType,
		// 		bip44.ErrInvalidAccount:               SKY_ErrInvalidAccount,
		// 		bip32.ErrPathNoMaster:                 SKY_ErrPathNoMaster,
		// 		bip32.ErrPathChildMaster:              SKY_ErrPathChildMaster,
		// 		bip32.ErrPathNodeNotNumber:            SKY_ErrPathNodeNotNumber,
		// 		bip32.ErrPathNodeNumberTooLarge:       SKY_ErrPathNodeNumberTooLarge,
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
