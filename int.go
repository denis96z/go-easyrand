package easyrand

import (
	"crypto/rand"
	"math/bits"
	"unsafe"
)

func Int() (int, error) {
	var x int

	b := make([]byte, unsafe.Sizeof(x))
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}

	for i := 0; i < len(b); i++ {
		x |= int(b[i]) << (8 * i)
	}

	return x, nil
}

func IntRange(min, max int) (int, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := Int()
	if err != nil {
		return 0, err
	}

	//XXX: x should always be positive
	x &= int(^(uint(1) << (bits.UintSize - 1)))

	return min + (x % d), nil
}

func IntChecked() int {
	x, err := Int()
	if err != nil {
		panic(err)
	}
	return x
}

func IntRangeChecked(min, max int) int {
	x, err := IntRange(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func IntUnchecked() int {
	x, _ := Int()
	return x
}

func IntRangeUnchecked(min, max int) int {
	x, _ := IntRange(min, max)
	return x
}
