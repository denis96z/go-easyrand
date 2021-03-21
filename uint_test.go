package easyrand

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUInt(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		_, err := UInt()
		require.NoError(t, err)
	}
}

func TestUIntRange(t *testing.T) {
	t.Parallel()
	tCases := []struct {
		Min uint
		Max uint
	}{
		{
			Min: 0,
			Max: 0,
		},
		{
			Min: 0,
			Max: 1,
		},
		{
			Min: 5,
			Max: 10,
		},
		{
			Min: (uint(1) << (bits.UintSize - 1)) - 1,
			Max: uint(1) << (bits.UintSize - 1),
		},
		{
			Min: uint(1) << (bits.UintSize - 1),
			Max: uint(1) << (bits.UintSize - 1),
		},
	}

	for _, tCase := range tCases {
		x, err := UIntRange(tCase.Min, tCase.Max)
		require.NoError(t, err)
		require.LessOrEqual(t, tCase.Min, x)
		require.GreaterOrEqual(t, tCase.Max, x)
	}
}
