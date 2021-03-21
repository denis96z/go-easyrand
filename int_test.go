package easyrand

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		_, err := Int()
		require.NoError(t, err)
	}
}

func TestIntRange(t *testing.T) {
	t.Parallel()

	intMin := func(x int, n uintptr) int {
		return x << n
	}(
		1, bits.UintSize-1,
	)

	intMax := (1 << (bits.UintSize - 1)) - 1

	tCases := []struct {
		Min int
		Max int
	}{
		{
			Min: intMin,
			Max: intMin,
		},
		{
			Min: intMin,
			Max: intMin + 1,
		},
		{
			Min: -10,
			Max: 0,
		},
		{
			Min: -5,
			Max: 5,
		},
		{
			Min: 0,
			Max: 10,
		},
		{
			Min: intMax - 1,
			Max: intMax,
		},
		{
			Min: intMax,
			Max: intMax,
		},
	}

	for _, tCase := range tCases {
		x, err := IntRange(tCase.Min, tCase.Max)
		require.NoError(t, err)
		require.LessOrEqual(t, tCase.Min, x)
		require.GreaterOrEqual(t, tCase.Max, x)
	}
}
