package badger

import (
	"errors"
	"fmt"

	"github.com/dapperlabs/flow-go/crypto"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/sdk/emulator/storage"
	"github.com/dapperlabs/flow-go/sdk/emulator/types"

	"github.com/dgraph-io/badger"
)

type Config struct {
	Path string
}

// Store is an embedded storage implementation using Badger as the underlying
// persistent key-value store.
type Store struct {
	db *badger.DB
}

// New returns a new Badger Store.
func New(config *Config) (storage.Store, error) {
	db, err := badger.Open(badger.DefaultOptions(config.Path))
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}
	return Store{db}, nil
}

func (s Store) GetBlockByHash(blockHash crypto.Hash) (block types.Block, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		// get block number by block hash
		encBlockNumber, err := getTx(txn)(blockHashIndexKey(blockHash))
		if err != nil {
			return err
		}

		// decode block number
		var blockNumber uint64
		if err := decodeUint64(&blockNumber, encBlockNumber); err != nil {
			return err
		}

		// get block by block number and decode
		encBlock, err := getTx(txn)(blockKey(blockNumber))
		if err != nil {
			return err
		}
		return decodeBlock(&block, encBlock)
	})
	return
}

func (s Store) GetBlockByNumber(blockNumber uint64) (block types.Block, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		encBlock, err := getTx(txn)(blockKey(blockNumber))
		if err != nil {
			return err
		}
		return decodeBlock(&block, encBlock)
	})
	return
}

func (s Store) GetLatestBlock() (block types.Block, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		// get latest block number
		latestBlockNumber, err := getLatestBlockNumberTx(txn)
		if err != nil {
			return err
		}

		// get corresponding block
		encBlock, err := getTx(txn)(blockKey(latestBlockNumber))
		if err != nil {
			return err
		}
		return decodeBlock(&block, encBlock)
	})
	return
}

func (s Store) InsertBlock(block types.Block) error {
	encBlock, err := encodeBlock(block)
	if err != nil {
		return err
	}
	encBlockNumber, err := encodeUint64(block.Number)
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		// get latest block number
		latestBlockNumber, err := getLatestBlockNumberTx(txn)
		if err != nil {
			return err
		}

		// insert the block by block number
		if err := txn.Set(blockKey(block.Number), encBlock); err != nil {
			return err
		}
		// add block hash to hash->number lookup
		if err := txn.Set(blockHashIndexKey(block.Hash()), encBlockNumber); err != nil {
			return err
		}

		// if this is latest block, set latest block
		if block.Number > latestBlockNumber {
			return txn.Set(latestBlockKey(), encBlockNumber)
		}
		return nil
	})
}

func (s Store) GetTransaction(txHash crypto.Hash) (tx flow.Transaction, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		encTx, err := getTx(txn)(transactionKey(txHash))
		if err != nil {
			return err
		}
		return decodeTransaction(&tx, encTx)
	})
	return
}

func (s Store) InsertTransaction(tx flow.Transaction) error {
	encTx, err := encodeTransaction(tx)
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(transactionKey(tx.Hash()), encTx)
	})
}

func (s Store) GetRegistersView(blockNumber uint64) (view flow.RegistersView, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		encRegisters, err := getTx(txn)(registersKey(blockNumber))
		if err != nil {
			return err
		}

		var registers flow.Registers
		if err := decodeRegisters(&registers, encRegisters); err != nil {
			return err
		}
		view = *registers.NewView()
		return nil
	})
	return
}

func (s Store) SetRegisters(blockNumber uint64, registers flow.Registers) error {
	encRegisters, err := encodeRegisters(registers)
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(registersKey(blockNumber), encRegisters)
	})
}

// TODO
func (s Store) GetEvents(blockNumber uint64, eventType string, startBlock, endBlock uint64) ([]flow.Event, error) {
	iterOpts := badger.DefaultIteratorOptions
	iterOpts.Prefix = []byte(eventsKeyPrefix)

	s.db.View(func(txn *badger.Txn) error {
		iter := txn.NewIterator(iterOpts)
		_ = iter
		return nil
	})
	return nil, nil
}

func (s Store) InsertEvents(blockNumber uint64, events ...flow.Event) error {
	encEvents, err := encodeEventList(events)
	if err != nil {
		return err
	}

	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(eventsKey(blockNumber), encEvents)
	})
}

// getTx returns a getter function bound to the input transaction that can be
// used to get values from Badger. The getter function checks for key-not-found
// errors and wraps them in storage.ErrNotFound.
//
// This saves a few lines of converting a badger.Item to []byte.
func getTx(txn *badger.Txn) func([]byte) ([]byte, error) {
	return func(key []byte) ([]byte, error) {
		// Badger returns an "item" upon GETs
		item, err := txn.Get(key)
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return nil, storage.ErrNotFound{}
			}
			return nil, err
		}

		val := make([]byte, item.ValueSize())
		return item.ValueCopy(val)
	}
}

// getLatestBlockNumberTx retrieves the latest block number and returns it.
// Must be called from within a Badger transaction.
func getLatestBlockNumberTx(txn *badger.Txn) (uint64, error) {
	encBlockNumber, err := getTx(txn)(latestBlockKey())
	if err != nil {
		return 0, err
	}

	var blockNumber uint64
	if err := decodeUint64(&blockNumber, encBlockNumber); err != nil {
		return 0, err
	}

	return blockNumber, nil
}

// The following *Key functions return keys to use when reading/writing values
// to Badger. The key name includes how it is indexed.

const (
	blockKeyPrefix          = "block_by_number"
	blockHashIndexKeyPrefix = "block_hash_to_number"
	transactionKeyPrefix    = "transaction_by_hash"
	registersKeyPrefix      = "registers_by_block_number"
	eventsKeyPrefix         = "events_by_block_number"
)

func blockKey(blockNumber uint64) []byte {
	return []byte(fmt.Sprintf("%s-%d", blockKeyPrefix, blockNumber))
}

func blockHashIndexKey(blockHash crypto.Hash) []byte {
	return []byte(fmt.Sprintf("%s-%s", blockHashIndexKeyPrefix, blockHash.Hex()))
}

func latestBlockKey() []byte {
	return []byte("latest_block_number")
}

func transactionKey(txHash crypto.Hash) []byte {
	return []byte(fmt.Sprintf("%s-%s", transactionKeyPrefix, txHash.Hex()))
}

func registersKey(blockNumber uint64) []byte {
	return []byte(fmt.Sprintf("%s-%d", registersKeyPrefix, blockNumber))
}

func eventsKey(blockNumber uint64) []byte {
	return []byte(fmt.Sprintf("%s-%d", eventsKeyPrefix, blockNumber))
}
