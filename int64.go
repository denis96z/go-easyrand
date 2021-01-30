package easyrand

import (
	"crypto/rand"
)

func Int64() (int64, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
		int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56, nil
}

func Int64Range(min, max int64) (int64, error) {
	return randomInt64(min, max)
}

func Int64Checked() int64 {
	n, err := Int64()
	if err != nil {
		panic(err)
	}
	return n
}

func Int64RangeChecked(min, max int64) int64 {
	n, err := Int64Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func Int64Unchecked() int64 {
	n, _ := Int64()
	return n
}

func Int64RangeUnchecked(min, max int64) int64 {
	n, _ := Int64Range(min, max)
	return n
}
