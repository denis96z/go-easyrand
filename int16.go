package easyrand

import (
	"crypto/rand"
)

func Int16() (int16, error) {
	b := make([]byte, 2)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return int16(b[0]) | int16(b[1])<<8, nil
}

func Int16Range(min, max int16) (int16, error) {
	x, err := randomInt64(int64(min), int64(max))
	return int16(x), err
}

func Int16Checked() int16 {
	x, err := Int16()
	if err != nil {
		panic(err)
	}
	return x
}

func Int16RangeChecked(min, max int16) int16 {
	x, err := Int16Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func Int16Unchecked() int16 {
	x, _ := Int16()
	return x
}

func Int16RangeUnchecked(min, max int16) int16 {
	x, _ := Int16Range(min, max)
	return x
}
