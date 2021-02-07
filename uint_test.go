package easyrand

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := UInt()
		assert.NoError(t, err)
	}
}

func TestUIntRange(t *testing.T) {
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
		assert.NoError(t, err)
		assert.LessOrEqual(t, tCase.Min, x)
		assert.GreaterOrEqual(t, tCase.Max, x)
	}
}
