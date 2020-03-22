package main

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
)

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

func lookupAltcoinPluginHandle(handle C.AltcoinPlugin__Handle) (*core.AltcoinPlugin, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.AltcoinPlugin); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerAltcoinPluginHandle(obj *core.AltcoinPlugin) C.AltcoinPlugin__Handle {
	return (C.AltcoinPlugin__Handle)(registerHandle(obj))
}

func lookupWalletHandle(handle C.Wallet__Handle) (*core.Wallet, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Wallet); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerWalletHandle(obj *core.Wallet) C.Wallet__Handle {
	return (C.Wallet__Handle)(registerHandle(obj))
}

func lookupTxnSignerHandle(handle C.TxnSigner__Handle) (*core.TxnSigner, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TxnSigner); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTxnSignerHandle(obj *core.TxnSigner) C.TxnSigner__Handle {
	return (C.TxnSigner__Handle)(registerHandle(obj))
}

func lookupCryptoAccountHandle(handle C.CryptoAccount__Handle) (*core.CryptoAccount, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.CryptoAccount); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerCryptoAccountHandle(obj *core.CryptoAccount) C.CryptoAccount__Handle {
	return (C.CryptoAccount__Handle)(registerHandle(obj))
}

func lookupAddressHandle(handle C.Address__Handle) (*core.Address, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Address); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAddressHandle(obj *core.Address) C.Address__Handle {
	return (C.Address__Handle)(registerHandle(obj))
}

func lookupGenericOutputHandle(handle C.GenericOutput__Handle) (*util.GenericOutput, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.GenericOutput); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerGenericOutputHandle(obj *util.GenericOutput) C.GenericOutput__Handle {
	return (C.GenericOutput__Handle)(registerHandle(obj))
}

func lookupTxnSignerIteratorHandle(handle C.TxnSignerIterator__Handle) (*core.TxnSignerIterator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TxnSignerIterator); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTxnSignerIteratorHandle(obj *core.TxnSignerIterator) C.TxnSignerIterator__Handle {
	return (C.TxnSignerIterator__Handle)(registerHandle(obj))
}

func lookupTransactionHandle(handle C.Transaction__Handle) (*core.Transaction, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Transaction); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTransactionHandle(obj *core.Transaction) C.Transaction__Handle {
	return (C.Transaction__Handle)(registerHandle(obj))
}

func lookupKeyValueStoreHandle(handle C.KeyValueStore__Handle) (*core.KeyValueStore, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.KeyValueStore); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerKeyValueStoreHandle(obj *core.KeyValueStore) C.KeyValueStore__Handle {
	return (C.KeyValueStore__Handle)(registerHandle(obj))
}

func lookupSimpleWalletOutputHandle(handle C.SimpleWalletOutput__Handle) (*util.SimpleWalletOutput, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.SimpleWalletOutput); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerSimpleWalletOutputHandle(obj *util.SimpleWalletOutput) C.SimpleWalletOutput__Handle {
	return (C.SimpleWalletOutput__Handle)(registerHandle(obj))
}

func lookupTransactionOutputHandle(handle C.TransactionOutput__Handle) (*core.TransactionOutput, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TransactionOutput); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTransactionOutputHandle(obj *core.TransactionOutput) C.TransactionOutput__Handle {
	return (C.TransactionOutput__Handle)(registerHandle(obj))
}

func lookupSimpleWalletAddressHandle(handle C.SimpleWalletAddress__Handle) (*util.SimpleWalletAddress, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.SimpleWalletAddress); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerSimpleWalletAddressHandle(obj *util.SimpleWalletAddress) C.SimpleWalletAddress__Handle {
	return (C.SimpleWalletAddress__Handle)(registerHandle(obj))
}

func lookupBalanceSnapshotHandle(handle C.BalanceSnapshot__Handle) (*util.BalanceSnapshot, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*util.BalanceSnapshot); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBalanceSnapshotHandle(obj *util.BalanceSnapshot) C.BalanceSnapshot__Handle {
	return (C.BalanceSnapshot__Handle)(registerHandle(obj))
}

func lookupCryptoAccountMocksHandle(handle C.CryptoAccountMocks__Handle) (*mocks.CryptoAccount, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.CryptoAccount); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerCryptoAccountMocksHandle(obj *mocks.CryptoAccount) C.CryptoAccountMocks__Handle {
	return (C.CryptoAccountMocks__Handle)(registerHandle(obj))
}

func lookupTransactionIteratorHandle(handle C.TransactionIterator__Handle) (*core.TransactionIterator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TransactionIterator); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTransactionIteratorHandle(obj *core.TransactionIterator) C.TransactionIterator__Handle {
	return (C.TransactionIterator__Handle)(registerHandle(obj))
}

func lookupTransactionOutputIteratorHandle(handle C.TransactionOutputIterator__Handle) (*core.TransactionOutputIterator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.TransactionOutputIterator); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerTransactionOutputIteratorHandle(obj *core.TransactionOutputIterator) C.TransactionOutputIterator__Handle {
	return (C.TransactionOutputIterator__Handle)(registerHandle(obj))
}

func lookupAddressMocksHandle(handle C.AddressMocks__Handle) (*mocks.Address, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.Address); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAddressMocksHandle(obj *mocks.Address) C.AddressMocks__Handle {
	return (C.AddressMocks__Handle)(registerHandle(obj))
}

func lookupAddressBookMocksHandle(handle C.AddressBookMocks__Handle) (*mocks.AddressBook, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.AddressBook); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAddressBookMocksHandle(obj *mocks.AddressBook) C.AddressBookMocks__Handle {
	return (C.AddressBookMocks__Handle)(registerHandle(obj))
}

func lookupContactHandle(handle C.Contact__Handle) (*core.Contact, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Contact); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerContactHandle(obj *core.Contact) C.Contact__Handle {
	return (C.Contact__Handle)(registerHandle(obj))
}

func lookupStorageHandle(handle C.Storage__Handle) (*core.Storage, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Storage); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerStorageHandle(obj *core.Storage) C.Storage__Handle {
	return (C.Storage__Handle)(registerHandle(obj))
}

func lookupPubKeyHandle(handle C.PubKey__Handle) (*core.PubKey, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.PubKey); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerPubKeyHandle(obj *core.PubKey) C.PubKey__Handle {
	return (C.PubKey__Handle)(registerHandle(obj))
}

func lookupAddressIteratorMocksHandle(handle C.AddressIteratorMocks__Handle) (*mocks.AddressIterator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.AddressIterator); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAddressIteratorMocksHandle(obj *mocks.AddressIterator) C.AddressIteratorMocks__Handle {
	return (C.AddressIteratorMocks__Handle)(registerHandle(obj))
}

func lookupAltcoinManagerMocksHandle(handle C.AltcoinManagerMocks__Handle) (*mocks.AltcoinManager, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.AltcoinManager); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAltcoinManagerMocksHandle(obj *mocks.AltcoinManager) C.AltcoinManagerMocks__Handle {
	return (C.AddressIteratorMocks__Handle)(registerHandle(obj))
}

func lookupAltcoinPluginMocksHandle(handle C.AltcoinPluginMocks__Handle) (*mocks.AltcoinPlugin, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.AltcoinPlugin); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAltcoinPluginMocksHandle(obj *mocks.AltcoinPlugin) C.AltcoinPluginMocks__Handle {
	return (C.AltcoinPluginMocks__Handle)(registerHandle(obj))
}

func lookupSecKeyHandle(handle C.SecKey__Handle) (*core.SecKey, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.SecKey); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerSecKeyHandle(obj *core.SecKey) C.SecKey__Handle {
	return (C.SecKey__Handle)(registerHandle(obj))
}

func lookupAltcoinManagerHandle(handle C.AltcoinManager__Handle) (*core.AltcoinManager, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.AltcoinManager); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerAltcoinManagerHandle(obj *core.AltcoinManager) C.AltcoinManager__Handle {
	return (C.AltcoinManager__Handle)(registerHandle(obj))
}
func lookupBlockchainSignServiceHandle(handle C.BlockchainSignService__Handle) (*core.BlockchainSignService, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.BlockchainSignService); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBlockchainSignServiceHandle(obj *core.BlockchainSignService) C.BlockchainSignService__Handle {
	return (C.BlockchainSignService__Handle)(registerHandle(obj))
}

func lookupBlockchainTransactionAPIHandle(handle C.BlockchainTransactionAPI__Handle) (*core.BlockchainTransactionAPI, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.BlockchainTransactionAPI); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBlockchainTransactionAPIHandle(obj *core.BlockchainTransactionAPI) C.BlockchainTransactionAPI__Handle {
	return (C.BlockchainTransactionAPI__Handle)(registerHandle(obj))
}
func lookupPEXHandle(handle C.PEX__Handle) (*core.PEX, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.PEX); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerPEXHandle(obj *core.PEX) C.PEX__Handle {
	return (C.PEX__Handle)(registerHandle(obj))
}
func lookupBlockMocksHandle(handle C.BlockMocks__Handle) (*mocks.Block, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.Block); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBlockMocksHandle(obj *mocks.Block) C.BlockMocks__Handle {
	return (C.BlockMocks__Handle)(registerHandle(obj))
}
func lookupBlockchainSignServiceMocksHandle(handle C.BlockchainSignServiceMocks__Handle) (*mocks.BlockchainSignService, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.BlockchainSignService); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBlockchainSignServiceMocksHandle(obj *mocks.BlockchainSignService) C.BlockchainSignServiceMocks__Handle {
	return (C.BlockchainSignServiceMocks__Handle)(registerHandle(obj))
}
func lookupBlockchainStatusMocksHandle(handle C.BlockchainStatusMocks__Handle) (*mocks.BlockchainStatus, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.BlockchainStatus); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBlockchainStatusMocksHandle(obj *mocks.BlockchainStatus) C.BlockchainStatusMocks__Handle {
	return (C.BlockchainStatusMocks__Handle)(registerHandle(obj))
}
func lookupBlockHandle(handle C.Block__Handle) (*core.Block, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.Block); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBlockHandle(obj *core.Block) C.Block__Handle {
	return (C.Block__Handle)(registerHandle(obj))
}

func lookupBlockchainTransactionAPIMocksHandle(handle C.BlockchainTransactionAPIMocks__Handle) (*mocks.BlockchainTransactionAPI, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.BlockchainTransactionAPI); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerBlockchainTransactionAPIMocksHandle(obj *mocks.BlockchainTransactionAPI) C.BlockchainTransactionAPIMocks__Handle {
	return (C.BlockchainTransactionAPIMocks__Handle)(registerHandle(obj))
}

func lookupWalletAddressHandle(handle C.WalletAddress__Handle) (*core.WalletAddress, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.WalletAddress); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerWalletAddressHandle(obj *core.WalletAddress) C.WalletAddress__Handle {
	return (C.WalletAddress__Handle)(registerHandle(obj))
}

func lookupWalletOutputHandle(handle C.WalletOutput__Handle) (*core.WalletOutput, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.WalletOutput); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerWalletOutputHandle(obj *core.WalletOutput) C.WalletOutput__Handle {
	return (C.WalletOutput__Handle)(registerHandle(obj))
}

func lookupContactMocksHandle(handle C.ContactMocks__Handle) (*mocks.Contact, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.Contact); isOK {
			return obj, true
		}
	}
	return nil, false
}
func registerContactMocksHandle(obj *mocks.Contact) C.ContactMocks__Handle {
	return (C.ContactMocks__Handle)(registerHandle(obj))
}

func lookupPexNodeIteratorHandle(handle C.PexNodeIterator__Handle) (*core.PexNodeIterator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.PexNodeIterator); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerPexNodeIteratorHandle(obj *core.PexNodeIterator) C.PexNodeIterator__Handle {
	return (C.PexNodeIterator__Handle)(registerHandle(obj))
}

func lookupPexNodeSetHandle(handle C.PexNodeSet__Handle) (*core.PexNodeSet, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.PexNodeSet); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerPexNodeSetHandle(obj *core.PexNodeSet) C.PexNodeSet__Handle {
	return (C.PexNodeSet__Handle)(registerHandle(obj))
}

func lookupPexNodeHandle(handle C.PexNode__Handle) (*core.PexNode, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.PexNode); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerPexNodeHandle(obj *core.PexNode) C.PexNode__Handle {
	return (C.PexNode__Handle)(registerHandle(obj))
}

func lookupPooledObjectFactoryHandle(handle C.PooledObjectFactory__Handle) (*core.PooledObjectFactory, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.PooledObjectFactory); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerPooledObjectFactoryHandle(obj *core.PooledObjectFactory) C.PooledObjectFactory__Handle {
	return (C.PooledObjectFactory__Handle)(registerHandle(obj))
}

func lookupMultiPoolHandle(handle C.MultiPool__Handle) (*core.MultiPool, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.MultiPool); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerMultiPoolHandle(obj *core.MultiPool) C.MultiPool__Handle {
	return (C.MultiPool__Handle)(registerHandle(obj))
}

func lookupMultiPoolSectionHandle(handle C.MultiPoolSection__Handle) (*core.MultiPoolSection, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.MultiPoolSection); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerMultiPoolSectionHandle(obj *core.MultiPoolSection) C.MultiPoolSection__Handle {
	return (C.MultiPoolSection__Handle)(registerHandle(obj))
}

func lookupMultiConnectionsPoolHandle(handle C.MultiConnectionsPool__Handle) (*core.MultiConnectionsPool, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.MultiConnectionsPool); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerMultiConnectionsPoolHandle(obj *core.MultiConnectionsPool) C.MultiConnectionsPool__Handle {
	return (C.MultiConnectionsPool__Handle)(registerHandle(obj))
}

func lookupPoolSectionHandle(handle C.PoolSection__Handle) (*core.PoolSection, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.PoolSection); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerPoolSectionHandle(obj *core.PoolSection) C.PoolSection__Handle {
	return (C.PoolSection__Handle)(registerHandle(obj))
}

func lookupWalletIteratorHandle(handle C.WalletIterator__Handle) (*core.WalletIterator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.WalletIterator); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerWalletIteratorHandle(obj *core.WalletIterator) C.WalletIterator__Handle {
	return (C.WalletIterator__Handle)(registerHandle(obj))
}

func lookupWalletSetHandle(handle C.WalletSet__Handle) (*core.WalletSet, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.WalletSet); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerWalletSetHandle(obj *core.WalletSet) C.WalletSet__Handle {
	return (C.WalletSet__Handle)(registerHandle(obj))
}

func lookupWalletStorageHandle(handle C.WalletStorage__Handle) (*core.WalletStorage, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.WalletStorage); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerWalletStorageHandle(obj *core.WalletStorage) C.WalletStorage__Handle {
	return (C.WalletStorage__Handle)(registerHandle(obj))
}

func lookupSeedGeneratorHandle(handle C.SeedGenerator__Handle) (*core.SeedGenerator, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.SeedGenerator); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerSeedGeneratorHandle(obj *core.SeedGenerator) C.SeedGenerator__Handle {
	return (C.SeedGenerator__Handle)(registerHandle(obj))
}

func lookupWalletEnvHandle(handle C.WalletEnv__Handle) (*core.WalletEnv, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*core.WalletEnv); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerWalletEnvHandle(obj *core.WalletEnv) C.WalletEnv__Handle {
	return (C.WalletEnv__Handle)(registerHandle(obj))
}

func lookupMultiPoolMocksHandle(handle C.MultiPoolMocks__Handle) (*mocks.MultiPool, bool) {
	obj, ok := lookupHandle(C.Handle(handle))
	if ok {
		if obj, isOK := (obj).(*mocks.MultiPool); isOK {
			return obj, true
		}
	}
	return nil, false
}

func registerMultiPoolMocksHandle(obj *mocks.MultiPool) C.MultiPoolMocks__Handle {
	return (C.MultiPoolMocks__Handle)(registerHandle(obj))
}
