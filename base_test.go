package easyrand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		const min, max int64 = -1, 1
		v, err := randomInt64(min, max)

		assert.NoError(t, err)
		assert.LessOrEqual(t, min, v)
		assert.GreaterOrEqual(t, max, v)
	}
}

func TestRandomUInt64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		const min, max uint64 = 0, 1
		v, err := randomUInt64(min, max)

		assert.NoError(t, err)
		assert.LessOrEqual(t, min, v)
		assert.GreaterOrEqual(t, max, v)
	}
}

func TestRandomBytes(t *testing.T) {
	m := map[byte]struct{}{
		1: {}, 2: {}, 3: {},
	}

	a := make([]byte, 0, len(m))
	for k := range m {
		a = append(a, k)
	}

	for i := 0; i < 1000; i++ {
		v, err := randomBytes(a, i)

		assert.NoError(t, err)
		assert.Equal(t, i, len(v))

		for j := 0; j < i; j++ {
			_, ok := m[v[j]]
			assert.True(t, ok)
		}
	}
}
