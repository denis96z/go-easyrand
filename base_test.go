package easyrand

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomBytesAlphabet(t *testing.T) {
	t.Parallel()
	m := map[byte]struct{}{
		1: {}, 2: {}, 3: {},
	}

	a := make([]byte, 0, len(m))
	for k := range m {
		a = append(a, k)
	}

	for i := 0; i < 1000; i++ {
		v, err := randomBytesAlphabet(a, i)

		require.NoError(t, err)
		require.Equal(t, i, len(v))

		for j := 0; j < i; j++ {
			_, ok := m[v[j]]
			require.True(t, ok)
		}
	}
}
