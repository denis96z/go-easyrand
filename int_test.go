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
