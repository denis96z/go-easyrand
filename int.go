package easyrand

import (
	"crypto/rand"
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
	x, err := randomInt64(int64(min), int64(max))
	return int(x), err
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
