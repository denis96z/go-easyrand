package easyrand

func UInt64() (uint64, error) {
	return UInt64Range(0, (1<<64)-1)
}

func UInt64Range(min, max uint64) (uint64, error) {
	return randomUInt64(min, max)
}

func UInt64Checked() uint64 {
	n, err := UInt64()
	if err != nil {
		panic(err)
	}
	return n
}

func UInt64RangeChecked(min, max uint64) uint64 {
	n, err := UInt64Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UInt64Unchecked() uint64 {
	n, _ := UInt64()
	return n
}

func UInt64RangeUnchecked(min, max uint64) uint64 {
	n, _ := UInt64Range(min, max)
	return n
}
