package easyrand

import (
	"crypto/rand"
)

func Int8() (int8, error) {
	b := make([]byte, 1)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return int8(b[0]), nil
}

func Int8Range(min, max int8) (int8, error) {
	n, err := randomInt64(int64(min), int64(max))
	return int8(n), err
}

func Int8Checked() int8 {
	n, err := Int8()
	if err != nil {
		panic(err)
	}
	return n
}

func Int8RangeChecked(min, max int8) int8 {
	n, err := Int8Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func Int8Unchecked() int8 {
	n, _ := Int8()
	return n
}

func Int8RangeUnchecked(min, max int8) int8 {
	n, _ := Int8Range(min, max)
	return n
}
