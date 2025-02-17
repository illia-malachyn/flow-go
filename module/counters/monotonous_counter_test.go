package counters

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/utils/unittest"
)

func TestSet(t *testing.T) {
	counter := NewMonotonicCounter(3)
	require.True(t, counter.Set(4))
	require.Equal(t, uint64(4), counter.Value())
	require.False(t, counter.Set(2))
	require.Equal(t, uint64(4), counter.Value())
}

func TestIncrement(t *testing.T) {
	counter := NewMonotonicCounter(1)
	require.Equal(t, uint64(2), counter.Increment())
	require.Equal(t, uint64(3), counter.Increment())
}

// TestIncrementConcurrently tests that the MonotonicCounter's Increment method
// works correctly when called concurrently from multiple goroutines
func TestIncrementConcurrently(t *testing.T) {
	counter := NewMonotonicCounter(0)

	unittest.Concurrently(100, func(i int) {
		counter.Increment()
	})

	require.Equal(t, uint64(100), counter.Value())
}

func TestFuzzy(t *testing.T) {
	counter := NewMonotonicCounter(3)
	require.True(t, counter.Set(4))
	require.False(t, counter.Set(2))
	require.True(t, counter.Set(7))
	require.True(t, counter.Set(9))
	require.True(t, counter.Set(12))
	require.False(t, counter.Set(10))
	require.True(t, counter.Set(18))

	for i := 20; i < 100; i++ {
		require.True(t, counter.Set(uint64(i)))
	}

	for i := 20; i < 100; i++ {
		require.False(t, counter.Set(uint64(i)))
	}
}

func TestConcurrent(t *testing.T) {
	counter := NewMonotonicCounter(3)

	unittest.Concurrently(100, func(i int) {
		counter.Set(uint64(i))
	})

	require.Equal(t, uint64(99), counter.Value())
}
