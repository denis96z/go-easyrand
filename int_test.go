package easyrand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := Int()
		assert.NoError(t, err)
	}
}

func TestIntRange(t *testing.T) {
	for i := -100; i < 100; i += 10 {
		x1, x2 := i, i + 20
		x, err := IntRange(x1, x2)
		assert.NoError(t, err)
		assert.LessOrEqual(t, x1, x)
		assert.GreaterOrEqual(t, x2, x)
	}
}
