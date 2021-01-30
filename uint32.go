package easyrand

import "crypto/rand"

func UInt32() (uint32, error) {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24, nil
}

func UInt32Range(min, max uint32) (uint32, error) {
	n, err := randomUInt64(uint64(min), uint64(max))
	return uint32(n), err
}

func UInt32Checked() uint32 {
	n, err := UInt32()
	if err != nil {
		panic(err)
	}
	return n
}

func UInt32RangeChecked(min, max uint32) uint32 {
	n, err := UInt32Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UInt32Unchecked() uint32 {
	n, _ := UInt32()
	return n
}

func UInt32RangeUnchecked(min, max uint32) uint32 {
	n, _ := UInt32Range(min, max)
	return n
}
