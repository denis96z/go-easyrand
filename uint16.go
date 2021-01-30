package easyrand

import (
	"crypto/rand"
)

func UInt16() (uint16, error) {
	b := make([]byte, 2)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return uint16(b[0]) | uint16(b[1])<<8, nil
}

func UInt16Range(min, max uint16) (uint16, error) {
	n, err := randomUInt64(uint64(min), uint64(max))
	return uint16(n), err
}

func UInt16Checked() uint16 {
	n, err := UInt16()
	if err != nil {
		panic(err)
	}
	return n
}

func UInt16RangeChecked(min, max uint16) uint16 {
	n, err := UInt16Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UInt16Unchecked() uint16 {
	n, _ := UInt16()
	return n
}

func UInt16RangeUnchecked(min, max uint16) uint16 {
	n, _ := UInt16Range(min, max)
	return n
}
