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
	n, err := randomUInt64(uint64(min), uint64(max))
	return uint(n), err
}

func UIntChecked() uint {
	n, err := UInt()
	if err != nil {
		panic(err)
	}
	return n
}

func UIntRangeChecked(min, max uint) uint {
	n, err := UIntRange(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UIntUnchecked() uint {
	n, _ := UInt()
	return n
}

func UIntRangeUnchecked(min, max uint) uint {
	n, _ := UIntRange(min, max)
	return n
}
