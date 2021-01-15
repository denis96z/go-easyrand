package easyrand

import (
	"math"
)

func Int32() (int32, error) {
	return Int32Range(math.MinInt32, math.MaxInt32)
}

func Int32Range(min, max int32) (int32, error) {
	n, err := randomInt64(int64(min), int64(max))
	return int32(n), err
}

func Int32Checked() int32 {
	n, err := Int32()
	if err != nil {
		panic(err)
	}
	return n
}

func Int32RangeChecked(min, max int32) int32 {
	n, err := Int32Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func Int32Unchecked() int32 {
	n, _ := Int32()
	return n
}

func Int32RangeUnchecked(min, max int32) int32 {
	n, _ := Int32Range(min, max)
	return n
}
