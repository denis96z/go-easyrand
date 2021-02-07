package easyrand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomBytesAlphabet(t *testing.T) {
	m := map[byte]struct{}{
		1: {}, 2: {}, 3: {},
	}

	a := make([]byte, 0, len(m))
	for k := range m {
		a = append(a, k)
	}

	for i := 0; i < 1000; i++ {
		v, err := randomBytesAlphabet(a, i)

		assert.NoError(t, err)
		assert.Equal(t, i, len(v))

		for j := 0; j < i; j++ {
			_, ok := m[v[j]]
			assert.True(t, ok)
		}
	}
}
