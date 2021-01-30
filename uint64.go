package easyrand

import (
	"crypto/rand"
)

func UInt64() (uint64, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56, nil
}

func UInt64Range(min, max uint64) (uint64, error) {
	return randomUInt64(min, max)
}

func UInt64Checked() uint64 {
	n, err := UInt64()
	if err != nil {
		panic(err)
	}
	return n
}

func UInt64RangeChecked(min, max uint64) uint64 {
	n, err := UInt64Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UInt64Unchecked() uint64 {
	n, _ := UInt64()
	return n
}

func UInt64RangeUnchecked(min, max uint64) uint64 {
	n, _ := UInt64Range(min, max)
	return n
}
