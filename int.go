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
	n, err := randomInt64(int64(min), int64(max))
	return int(n), err
}

func IntChecked() int {
	n, err := Int()
	if err != nil {
		panic(err)
	}
	return n
}

func IntRangeChecked(min, max int) int {
	n, err := IntRange(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func IntUnchecked() int {
	n, _ := Int()
	return n
}

func IntRangeUnchecked(min, max int) int {
	n, _ := IntRange(min, max)
	return n
}
