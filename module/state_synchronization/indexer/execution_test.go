package indexer

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/cmd/util/ledger/migrations"
	"github.com/onflow/flow-go/engine/execution/state"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/storage/snapshot"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/pathfinder"
	"github.com/onflow/flow-go/ledger/complete"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/executiondatasync/execution_data"
	synctest "github.com/onflow/flow-go/module/state_synchronization/requester/unittest"
	"github.com/onflow/flow-go/storage"
	storagemock "github.com/onflow/flow-go/storage/mock"
	"github.com/onflow/flow-go/utils/unittest"
)

type indexTest struct {
	t                *testing.T
	indexer          *ExecutionState
	registers        *storagemock.Registers
	events           *storagemock.Events
	ctx              context.Context
	blocks           []*flow.Block
	data             *execution_data.BlockExecutionDataEntity
	lastHeightStore  func(t *testing.T) (uint64, error)
	firstHeightStore func(t *testing.T) (uint64, error)
	registersStore   func(t *testing.T, entries flow.RegisterEntries, height uint64) error
	eventsStore      func(t *testing.T, ID flow.Identifier, events []flow.EventsList) error
	registersGet     func(t *testing.T, IDs flow.RegisterID, height uint64) (flow.RegisterValue, error)
}

func newIndexTest(
	t *testing.T,
	blocks []*flow.Block,
	exeData *execution_data.BlockExecutionDataEntity,
) *indexTest {
	registers := storagemock.NewRegisters(t)
	events := storagemock.NewEvents(t)

	return &indexTest{
		t:         t,
		registers: registers,
		events:    events,
		blocks:    blocks,
		ctx:       context.Background(),
		data:      exeData,
	}
}

func (i *indexTest) setLastHeight(f func(t *testing.T) (uint64, error)) *indexTest {
	i.registers.
		On("LatestHeight").
		Return(func() (uint64, error) {
			return f(i.t)
		})
	return i
}

func (i *indexTest) useDefaultLastHeight() *indexTest {
	i.registers.
		On("LatestHeight").
		Return(func() (uint64, error) {
			return i.blocks[len(i.blocks)-1].Header.Height, nil
		})
	return i
}

func (i *indexTest) useDefaultFirstHeight() *indexTest {
	i.registers.
		On("FirstHeight").
		Return(func() (uint64, error) {
			return i.blocks[0].Header.Height, nil
		})
	return i
}

func (i *indexTest) useDefaultHeights() *indexTest {
	i.useDefaultFirstHeight()
	return i.useDefaultLastHeight()
}

func (i *indexTest) setFirstHeight(f func(t *testing.T) (uint64, error)) *indexTest {
	i.registers.
		On("FirstHeight").
		Return(func() (uint64, error) {
			return f(i.t)
		})
	return i
}

func (i *indexTest) setStoreRegisters(f func(t *testing.T, entries flow.RegisterEntries, height uint64) error) *indexTest {
	i.registers.
		On("Store", mock.AnythingOfType("flow.RegisterEntries"), mock.AnythingOfType("uint64")).
		Return(func(entries flow.RegisterEntries, height uint64) error {
			return f(i.t, entries, height)
		})
	return i
}

func (i *indexTest) setStoreEvents(f func(t *testing.T, ID flow.Identifier, events []flow.EventsList) error) *indexTest {
	i.events.
		On("Store", mock.AnythingOfType("flow.Identifier"), mock.AnythingOfType("[]flow.EventsList")).
		Return(func(ID flow.Identifier, events []flow.EventsList) error {
			return f(i.t, ID, events)
		})
	return i
}

func (i *indexTest) setGetRegisters(f func(t *testing.T, ID flow.RegisterID, height uint64) (flow.RegisterValue, error)) *indexTest {
	i.registers.
		On("Get", mock.AnythingOfType("flow.RegisterID"), mock.AnythingOfType("uint64")).
		Return(func(IDs flow.RegisterID, height uint64) (flow.RegisterValue, error) {
			return f(i.t, IDs, height)
		})
	return i
}

func (i *indexTest) initIndexer() {
	headers := newBlockHeadersStorage(i.blocks)
	i.useDefaultHeights()
	indexer, err := New(i.registers, headers, i.blocks[0].Header.Height)
	require.NoError(i.t, err)
	i.indexer = indexer
}

func (i *indexTest) runIndexBlockData() error {
	i.initIndexer()
	return i.indexer.IndexBlockData(i.ctx, i.data)
}

func (i *indexTest) runGetRegisters(IDs flow.RegisterIDs, height uint64) ([]flow.RegisterValue, error) {
	i.initIndexer()
	return i.indexer.RegisterValues(IDs, height)
}

func TestExecutionState_HeightByBlockID(t *testing.T) {
	blocks := blocksFixture(5)
	indexer := ExecutionState{headers: newBlockHeadersStorage(blocks)}

	for _, b := range blocks {
		ret, err := indexer.HeightByBlockID(b.ID())
		require.NoError(t, err)
		require.Equal(t, b.Header.Height, ret)
	}
}

// test cases:
// - no chunk data
// - no registers data
// - multiple inserts, same height
// - smaller invalid height
// - bigger invalid height
// - error on register updates
// - error on events
// - full register data, events, collections...

func TestExecutionState_IndexBlockData(t *testing.T) {
	blocks := blocksFixture(5)
	block := blocks[len(blocks)-1]

	t.Run("Index Single Chunk and Single Register", func(t *testing.T) {
		trie := trieUpdateFixture()
		ed := &execution_data.BlockExecutionData{
			BlockID: block.ID(),
			ChunkExecutionDatas: []*execution_data.ChunkExecutionData{
				{TrieUpdate: trie},
			},
		}
		execData := execution_data.NewBlockExecutionDataEntity(block.ID(), ed)
		// crate a lookup map that matches flow register ID to index in the payloads slice
		payloadRegID := make(map[flow.RegisterID]int)
		for i, p := range trie.Payloads {
			k, _ := p.Key()
			regKey, _ := migrations.KeyToRegisterID(k)
			payloadRegID[regKey] = i
		}

		err := newIndexTest(t, blocks, execData).
			// make sure update registers match in length and are same as block data ledger payloads
			setStoreRegisters(func(t *testing.T, entries flow.RegisterEntries, height uint64) error {
				assert.Equal(t, height, block.Header.Height)
				assert.Len(t, trie.Payloads, entries.Len())

				// make sure all the registers from the execution data have been stored as well the value matches
				for _, entry := range entries {
					index, ok := payloadRegID[entry.Key]
					assert.True(t, ok)
					trie.Payloads[index].Value().Equals(entry.Value)
				}
				return nil
			}).
			runIndexBlockData()

		assert.NoError(t, err)
	})

	t.Run("Index Multiple Chunks and Merge Same Register Updates", func(t *testing.T) {
		tries := []*ledger.TrieUpdate{trieUpdateFixture(), trieUpdateFixture()}
		// make sure we have two register updates that are updating the same value, so we can check
		// if the value from the second update is being persisted instead of first
		tries[1].Paths[0] = tries[0].Paths[0]
		testValue := tries[1].Payloads[0]
		key, err := testValue.Key()
		require.NoError(t, err)
		testRegisterID, err := migrations.KeyToRegisterID(key)
		require.NoError(t, err)

		ed := &execution_data.BlockExecutionData{
			BlockID: block.ID(),
			ChunkExecutionDatas: []*execution_data.ChunkExecutionData{
				{TrieUpdate: tries[0]},
				{TrieUpdate: tries[1]},
			},
		}
		execData := execution_data.NewBlockExecutionDataEntity(block.ID(), ed)

		testRegisterFound := false
		err = newIndexTest(t, blocks, execData).
			// make sure update registers match in length and are same as block data ledger payloads
			setStoreRegisters(func(t *testing.T, entries flow.RegisterEntries, height uint64) error {
				for _, entry := range entries {
					if entry.Key.String() == testRegisterID.String() {
						testRegisterFound = true
						assert.True(t, testValue.Value().Equals(entry.Value))
					}
				}
				// we should make sure the register updates are equal to both payloads' length -1 since we don't
				// duplicate the same register
				assert.Equal(t, len(tries[0].Payloads)+len(tries[1].Payloads)-1, len(entries))
				return nil
			}).
			runIndexBlockData()

		assert.NoError(t, err)
		assert.True(t, testRegisterFound)
	})

	t.Run("Invalid Heights", func(t *testing.T) {
		last := blocks[len(blocks)-1]
		ed := &execution_data.BlockExecutionData{
			BlockID: last.Header.ID(),
		}
		execData := execution_data.NewBlockExecutionDataEntity(last.ID(), ed)

		err := newIndexTest(t, blocks, execData).
			// return a height one smaller than the latest block in storage
			setLastHeight(func(t *testing.T) (uint64, error) {
				return blocks[len(blocks)-3].Header.Height, nil
			}).
			runIndexBlockData()

		assert.True(t, errors.Is(err, ErrIndexValue))
	})

	t.Run("Unknown block ID", func(t *testing.T) {
		unknownBlock := blocksFixture(1)[0]
		ed := &execution_data.BlockExecutionData{
			BlockID: unknownBlock.Header.ID(),
		}
		execData := execution_data.NewBlockExecutionDataEntity(unknownBlock.Header.ID(), ed)

		err := newIndexTest(t, blocks, execData).runIndexBlockData()

		assert.True(t, errors.Is(err, storage.ErrNotFound))
	})

}

func TestExecutionState_RegisterValues(t *testing.T) {
	t.Run("Get value for single register", func(t *testing.T) {
		blocks := blocksFixture(5)
		height := blocks[1].Header.Height
		ids := []flow.RegisterID{{
			Owner: "1",
			Key:   "2",
		}}
		val := flow.RegisterValue("0x1")

		values, err := newIndexTest(t, blocks, nil).
			setGetRegisters(func(t *testing.T, ID flow.RegisterID, height uint64) (flow.RegisterValue, error) {
				return val, nil
			}).
			runGetRegisters(ids, height)

		assert.NoError(t, err)
		assert.Equal(t, values, []flow.RegisterValue{val})
	})
}

func newBlockHeadersStorage(blocks []*flow.Block) storage.Headers {
	blocksByID := make(map[flow.Identifier]*flow.Block, 0)
	for _, b := range blocks {
		blocksByID[b.ID()] = b
	}

	return synctest.MockBlockHeaderStorage(synctest.WithByID(blocksByID))
}

func blocksFixture(n int) []*flow.Block {
	blocks := make([]*flow.Block, n)

	genesis := unittest.BlockFixture()
	blocks[0] = &genesis
	for i := 1; i < n; i++ {
		blocks[i] = unittest.BlockWithParentFixture(blocks[i-1].Header)
	}

	return blocks
}

func bootstrapTrieUpdates() *ledger.TrieUpdate {
	opts := []fvm.Option{
		fvm.WithChain(flow.Testnet.Chain()),
	}
	ctx := fvm.NewContext(opts...)
	vm := fvm.NewVirtualMachine()

	snapshotTree := snapshot.NewSnapshotTree(nil)

	bootstrapOpts := []fvm.BootstrapProcedureOption{
		fvm.WithInitialTokenSupply(unittest.GenesisTokenSupply),
	}

	executionSnapshot, _, _ := vm.Run(
		ctx,
		fvm.Bootstrap(unittest.ServiceAccountPublicKey, bootstrapOpts...),
		snapshotTree)

	payloads := make([]*ledger.Payload, 0)
	for regID, regVal := range executionSnapshot.WriteSet {
		key := ledger.Key{
			KeyParts: []ledger.KeyPart{
				{
					Type:  state.KeyPartOwner,
					Value: []byte(regID.Owner),
				},
				{
					Type:  state.KeyPartKey,
					Value: []byte(regID.Key),
				},
			},
		}

		payloads = append(payloads, ledger.NewPayload(key, regVal))
	}

	return trieUpdateWithPayloadsFixture(payloads)
}

func trieUpdateWithPayloadsFixture(payloads []*ledger.Payload) *ledger.TrieUpdate {
	keys := make([]ledger.Key, 0)
	values := make([]ledger.Value, 0)
	for _, payload := range payloads {
		key, _ := payload.Key()
		keys = append(keys, key)
		values = append(values, payload.Value())
	}

	update, _ := ledger.NewUpdate(ledger.DummyState, keys, values)
	trie, _ := pathfinder.UpdateToTrieUpdate(update, complete.DefaultPathFinderVersion)
	return trie
}

func trieUpdateFixture() *ledger.TrieUpdate {
	return trieUpdateWithPayloadsFixture(
		[]*ledger.Payload{
			ledgerPayloadFixture(),
			ledgerPayloadFixture(),
			ledgerPayloadFixture(),
			ledgerPayloadFixture(),
		})
}

func ledgerPayloadFixture() *ledger.Payload {
	owner := unittest.RandomAddressFixture()
	key := make([]byte, 8)
	rand.Read(key)
	val := make([]byte, 8)
	rand.Read(val)
	return ledgerPayloadWithValuesFixture(owner.String(), fmt.Sprintf("%x", key), val)
}

func ledgerPayloadWithValuesFixture(owner string, key string, value []byte) *ledger.Payload {
	k := ledger.Key{
		KeyParts: []ledger.KeyPart{
			{
				Type:  state.KeyPartOwner,
				Value: []byte(owner),
			},
			{
				Type:  state.KeyPartKey,
				Value: []byte(key),
			},
		},
	}

	return ledger.NewPayload(k, value)
}
