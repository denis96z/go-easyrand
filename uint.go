package easyrand

import (
	"crypto/rand"
	"unsafe"
)

func UInt() (uint, error) {
	var x uint

	b := make([]byte, unsafe.Sizeof(x))
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}

	for i := 0; i < len(b); i++ {
		x |= uint(b[i]) << (8 * i)
	}

	return x, nil
}

func UIntRange(min, max uint) (uint, error) {
	x, err := randomUInt64(uint64(min), uint64(max))
	return uint(x), err
}

func UIntChecked() uint {
	x, err := UInt()
	if err != nil {
		panic(err)
	}
	return x
}

func UIntRangeChecked(min, max uint) uint {
	x, err := UIntRange(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func UIntUnchecked() uint {
	x, _ := UInt()
	return x
}

func UIntRangeUnchecked(min, max uint) uint {
	x, _ := UIntRange(min, max)
	return x
}
