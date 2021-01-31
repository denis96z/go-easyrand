package easyrand

import (
	"crypto/rand"
)

func UInt8() (uint8, error) {
	b := make([]byte, 1)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return b[0], nil
}

func UInt8Range(min, max uint8) (uint8, error) {
	x, err := randomUInt64(uint64(min), uint64(max))
	return uint8(x), err
}

func UInt8Checked() uint8 {
	x, err := UInt8()
	if err != nil {
		panic(err)
	}
	return x
}

func UInt8RangeChecked(min, max uint8) uint8 {
	x, err := UInt8Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func UInt8Unchecked() uint8 {
	x, _ := UInt8()
	return x
}

func UInt8RangeUnchecked(min, max uint8) uint8 {
	x, _ := UInt8Range(min, max)
	return x
}
