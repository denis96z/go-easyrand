package easyrand

func UInt64() (uint64, error) {
	b, err := randomBytes(8)
	if err != nil {
		return 0, err
	}
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56, nil
}

func UInt64Range(min, max uint64) (uint64, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := UInt64()
	if err != nil {
		return 0, err
	}

	return min + (x % d), nil
}

func UInt64Checked() uint64 {
	x, err := UInt64()
	if err != nil {
		panic(err)
	}
	return x
}

func UInt64RangeChecked(min, max uint64) uint64 {
	x, err := UInt64Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func UInt64Unchecked() uint64 {
	x, _ := UInt64()
	return x
}

func UInt64RangeUnchecked(min, max uint64) uint64 {
	x, _ := UInt64Range(min, max)
	return x
}
