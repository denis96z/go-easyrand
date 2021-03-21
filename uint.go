package easyrand

import (
	"unsafe"
)

func UInt() (uint, error) {
	var x uint

	b, err := randomBytes(int(unsafe.Sizeof(x)))
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(b); i++ {
		x |= uint(b[i]) << (8 * i)
	}

	return x, nil
}

func UIntRange(min, max uint) (uint, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := UInt()
	if err != nil {
		return 0, err
	}
	return min + (x % d), nil
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
