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
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := UInt16()
	if err != nil {
		return 0, err
	}

	return min + (x % d), nil
}

func UInt16Checked() uint16 {
	x, err := UInt16()
	if err != nil {
		panic(err)
	}
	return x
}

func UInt16RangeChecked(min, max uint16) uint16 {
	x, err := UInt16Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func UInt16Unchecked() uint16 {
	x, _ := UInt16()
	return x
}

func UInt16RangeUnchecked(min, max uint16) uint16 {
	x, _ := UInt16Range(min, max)
	return x
}
