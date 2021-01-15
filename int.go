package easyrand

import (
	"math"
)

func Int() (int, error) {
	//NOTE: sizeof(int32) <= sizeof(int) <= sizeof(int64)
	return IntRange(math.MinInt32, math.MaxInt32)
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
