package easyrand

import (
	"math"
)

func Int64() (int64, error) {
	return Int64Range(math.MinInt64, math.MaxInt64)
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
