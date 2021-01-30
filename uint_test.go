package easyrand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := UInt()
		assert.NoError(t, err)
	}
}
