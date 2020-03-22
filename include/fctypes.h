
#ifndef FCTYPES_H
#define FCTYPES_H

#include <inttypes.h>
#ifndef __SIZE_TYPE__
#define __SIZE_TYPE__ uintptr_t
#endif

/**
 * Go 8-bit signed integer values.
 */
typedef signed char GoInt8_;
/**
 * Go 8-bit unsigned integer values.
 */
typedef unsigned char GoUint8_;
/**
 * Go 16-bit signed integer values.
 */
typedef short GoInt16_;
/**
 * Go 16-bit unsigned integer values.
 */
typedef unsigned short GoUint16_;
/**
 * Go 32-bit signed integer values.
 */
typedef int GoInt32_;
/**
 * Go 32-bit unsigned integer values.
 */
typedef unsigned int GoUint32_;
/**
 * Go 64-bit signed integer values.
 */
typedef long long GoInt64_;
/**
 * Go 64-bit unsigned integer values.
 */
typedef unsigned long long GoUint64_;
/**
 * Go integer values aligned to the word size of the underlying architecture.
 */
#if __x86_64__ || __ppc64__
typedef GoInt64_ GoInt_;
#else
typedef GoInt32_ GoInt_;
#endif
/**
 * Go unsigned integer values aligned to the word size of the underlying
 * architecture.
 */
#if __x86_64__ || __ppc64__
typedef GoUint64_ GoUint_;
#else
typedef GoUint32_ GoUint_;
#endif
/**
 * Architecture-dependent type representing instances Go `uintptr` type.
 * Used as a generic representation of pointer types.
 */
typedef __SIZE_TYPE__ GoUintptr_;
/**
 * Go single precision 32-bits floating point values.
 */
typedef float GoFloat32_;
/**
 * Go double precision 64-bits floating point values.
 */
typedef double GoFloat64_;
/**
 * Instances of Go `complex` type.
 */
typedef struct {
  float real;
  float imaginary;
} GoComplex64_;
/**
 * Instances of Go `complex` type.
 */
typedef struct {
  double real;
  double imaginary;
} GoComplex128_;
typedef unsigned int BOOL;
typedef unsigned int error;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt._
*/
#if __x86_64__ || __ppc64__
typedef char
    _check_for_64_bit_pointer_matchingGoInt[sizeof(void *) == 64 / 8 ? 1 : -1];
#endif

/**
 * Instances of Go `string` type.
 */
typedef struct {
  const char *p; ///< Pointer to string characters buffer.
  GoInt_ n;      ///< String size not counting trailing `\0` char
                 ///< if at all included.
} GoString_;
/**
 * Instances of Go `map` type.
 */
typedef void *GoMap_;

/**
 * Instances of Go `chan` channel types.
 */
typedef void *GoChan_;

/**
 * Instances of Go interface types.
 */
typedef struct {
  void *t; ///< Pointer to the information of the concrete Go type
           ///< bound to this interface reference.
  void *v; ///< Pointer to the data corresponding to the value
           ///< bound to this interface type.
} GoInterface_;
/**
 * Instances of Go slices
 */
typedef struct {
  void *data; ///< Pointer to buffer containing slice data.
  GoInt_ len; ///< Number of items stored in slice buffer
  GoInt_ cap; ///< Maximum number of items that fits in this slice
              ///< considering allocated memory and item type's
              ///< size.
} GoSlice_;

#include "fctypes.gen.h"

/**
 * Memory handles returned back to the caller and manipulated
 * internally by API functions. Usually used to avoid type dependencies
 * with internal implementation types.
 */
typedef GoInt64_ Handle;

/**
 * String Slice Handle
 */
typedef Handle Strings__Handle;

/**
 * AltcoinPlugin__Handle Handle, interface core.AltcoinPlugin
 */
typedef Handle AltcoinPlugin__Handle;

/**
 * TxnSigner__Handle Handle, interface core.TxnSigner
 */
typedef Handle TxnSigner__Handle;

/**
 * Wallet__Handle Handle, interface core.Wallet
 */
typedef Handle Wallet__Handle;

/**
 * CryptoAccount__Handle Handle, interface core.CryptoAccount
 */
typedef Handle CryptoAccount__Handle;

/**
 * Address__Handle Handle, interface core.Address
 */
typedef Handle Address__Handle;

/**
 * GenericOutput__Handle Handle, interface util.GenericOutput
 */
typedef Handle GenericOutput__Handle;

/**
 * TxnSignerIterator__Handle Handle, interface core.TxnSignerIterator
 */
typedef Handle TxnSignerIterator__Handle;

/**
 * Transaction__Handle Handle, interface core.Transaction
 */
typedef Handle Transaction__Handle;

/**
 * KeyValueStore__Handle Handle, interface core.KeyValueStore
 */
typedef Handle KeyValueStore__Handle;

/**
 * InputSignDescriptor__Handle Handle, interface core.InputSignDescriptor
 */
typedef Handle InputSignDescriptor__Handle;

/**
 * SimpleWalletOutput__Handle Handle, struct util.SimpleWalletOutput
 */
typedef Handle SimpleWalletOutput__Handle;

/**
 * TransactionOutput__Handle Handle, interface core.TransactionOutput
 */
typedef Handle TransactionOutput__Handle;

/**
 * SimpleWalletAddress__Handle Handle, struct util.SimpleWalletAddress
 */
typedef Handle SimpleWalletAddress__Handle;

/**
 * BalanceSnapshot__Handle Handle, struct util.BalanceSnapshot
 */
typedef Handle BalanceSnapshot__Handle;

/**
 * AddressBook__Handle Handle, interface core.AddressBook
 */
typedef Handle AddressBook__Handle;

/**
 * Contact__Handle Handle, interface core.Contact
 */
typedef Handle Contact__Handle;

/**
 * Storage__Handle Handle, interface core.Storage
 */
typedef Handle Storage__Handle;

/**
 * CryptoAccountMocks__Handle Handle, interface mocks.CryptoAccount
 */
typedef Handle CryptoAccountMocks__Handle;

/**
 * TransactionIterator__Handle Handle, interface core.TransactionIterator
 */
typedef Handle TransactionIterator__Handle;

/**
 * TransactionOutputIterator__Handle Handle, interface
 * core.TransactionOutputIterator
 */
typedef Handle TransactionOutputIterator__Handle;

/**
 * AddressMocks__Handle Handle, interface mocks.Address
 */
typedef Handle AddressMocks__Handle;

/**
 * AddressBookMocks__Handle Handle, interface mocks.AddressBook
 */
typedef Handle AddressBookMocks__Handle;

/**
 * PubKey__Handle Handle, interface core.PubKey
 */
typedef Handle PubKey__Handle;

/**
 * AddressIteratorMocks__Handle Handle, interface mocks.AddressIterator
 */
typedef Handle AddressIteratorMocks__Handle;

/**
 * AltcoinManagerMocks__Handle Handle, interface mocks.AltcoinManager
 */
typedef Handle AltcoinManagerMocks__Handle;

/**
 * AltcoinPluginMocks__Handle Handle, interface mocks.AltcoinPlugin
 */
typedef Handle AltcoinPluginMocks__Handle;

/**
 * SecKey__Handle Handle, interface core.SecKey
 */
typedef Handle SecKey__Handle;

/**
 * AltcoinManager__Handle Handle, interface core.AltcoinManager
 */
typedef Handle AltcoinManager__Handle;

/**
 * BlockchainSignService__Handle Handle, interface core.BlockchainSignService
 */
typedef Handle BlockchainSignService__Handle;

/**
 * BlockchainTransactionAPI__Handle Handle, interface
 * core.BlockchainTransactionAPI
 */
typedef Handle BlockchainTransactionAPI__Handle;

/**
 * PEX__Handle Handle, interface core.PEX
 */
typedef Handle PEX__Handle;

/**
 * BlockMocks__Handle Handle, interface mocks.Block
 */
typedef Handle BlockMocks__Handle;

/**
 * BlockchainSignServiceMocks__Handle Handle, struct mocks.BlockchainSignService
 */
typedef Handle BlockchainSignServiceMocks__Handle;

/**
 * BlockchainStatusMocks__Handle Handle, struct mocks.BlockchainStatus
 */
typedef Handle BlockchainStatusMocks__Handle;

/**
 * Block__Handle Handle, interface core.Block
 */
typedef Handle Block__Handle;

/**
 * BlockchainTransactionAPIMocks__Handle Handle, struct
 * mocks.BlockchainTransactionAPI
 */
typedef Handle BlockchainTransactionAPIMocks__Handle;

/**
 * WalletAddress__Handle Handle, interface core.WalletAddress
 */
typedef Handle WalletAddress__Handle;

/**
 * WalletOutput__Handle Handle, interface core.WalletOutput
 */
typedef Handle WalletOutput__Handle;

/**
 * ContactMocks__Handle Handle, struct mocks.Contact
 */
typedef Handle ContactMocks__Handle;

/**
 * StringAddress__Handle Handle, interface core.StringAddress
 */
typedef Handle StringAddress__Handle;

/**
 * PexNodeIterator__Handle Handle, interface core.PexNodeIterator
 */
typedef Handle PexNodeIterator__Handle;

/**
 * PexNodeSet__Handle Handle, interface core.PexNodeSet
 */
typedef Handle PexNodeSet__Handle;

/**
 * PexNode__Handle Handle, interface core.PexNode
 */
typedef Handle PexNode__Handle;

/**
 * PooledObjectFactory__Handle Handle, interface core.PooledObjectFactory
 */
typedef Handle PooledObjectFactory__Handle;

/**
 * MultiPool__Handle Handle, interface core.MultiPool
 */
typedef Handle MultiPool__Handle;

/**
 * MultiPoolSection__Handle Handle, interface core.MultiPoolSection
 */
typedef Handle MultiPoolSection__Handle;

/**
 * MultiConnectionsPool__Handle Handle, struct core.MultiConnectionsPool
 */
typedef Handle MultiConnectionsPool__Handle;

/**
 * PoolSection__Handle Handle, struct core.PoolSection
 */
typedef Handle PoolSection__Handle;

// Callbacks

/**
 *  PasswordReaderFunc callback , func(string, KeyValueStore) (string, error)
 */

typedef GoUint32_ (*PasswordReaderFunc)(GoString_ pString,
                                        KeyValueStore__Handle pKeyValue,
                                        GoString_ *pStringOut, void *context);

typedef struct {
  PasswordReaderFunc callback;
  void *context;
} PasswordReader;

#endif
