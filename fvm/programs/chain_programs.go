package programs

import (
	"fmt"
	"sync"

	"github.com/hashicorp/golang-lru/simplelru"

	"github.com/onflow/flow-go/model/flow"
)

const DefaultProgramsCacheSize = 1000

// ChainPrograms is a cache of BlockPrograms databases used for speeding up
// cadence execution.
//
// Since programs are derived from external source, the BlockPrograms databases
// need not be durable and can be recreated on the fly.
type ChainPrograms struct {
	// NOTE: It's unsafe to use RWMutex since lru updates the data structure
	// on Get.
	mutex sync.Mutex

	lru *simplelru.LRU
}

func NewChainPrograms(chainCacheSize uint) (*ChainPrograms, error) {
	lru, err := simplelru.NewLRU(int(chainCacheSize), nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create LRU cache: %w", err)
	}

	return &ChainPrograms{
		lru: lru,
	}, nil
}

func (chain *ChainPrograms) unsafeGet(
	currentBlockId flow.Identifier,
) *Programs {
	currentEntry, ok := chain.lru.Get(currentBlockId)
	if ok {
		return currentEntry.(*Programs)
	}

	return nil
}

func (chain *ChainPrograms) Get(
	currentBlockId flow.Identifier,
) *Programs {
	chain.mutex.Lock()
	defer chain.mutex.Unlock()

	return chain.unsafeGet(currentBlockId)
}

func (chain *ChainPrograms) GetOrCreateBlockPrograms(
	currentBlockId flow.Identifier,
	parentBlockId flow.Identifier,
) *Programs {
	chain.mutex.Lock()
	defer chain.mutex.Unlock()

	currentEntry := chain.unsafeGet(currentBlockId)
	if currentEntry != nil {
		return currentEntry
	}

	var current *Programs
	parentEntry, ok := chain.lru.Get(parentBlockId)
	if ok {
		current = parentEntry.(*Programs).ChildPrograms()
	} else {
		current = NewEmptyPrograms()
	}

	chain.lru.Add(currentBlockId, current)
	return current
}

func (chain *ChainPrograms) NewBlockProgramsForScript(
	currentBlockId flow.Identifier,
) *Programs {
	block := chain.Get(currentBlockId)
	if block != nil {
		return block.ChildPrograms()
	}

	return NewEmptyPrograms()
}
