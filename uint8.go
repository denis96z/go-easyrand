package easyrand

func UInt8() (uint8, error) {
	b, err := randomBytes(1)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func UInt8Range(min, max uint8) (uint8, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := UInt8()
	if err != nil {
		return 0, err
	}

	return min + (x % d), nil
}

func UInt8Checked() uint8 {
	x, err := UInt8()
	if err != nil {
		panic(err)
	}
	return x
}

func UInt8RangeChecked(min, max uint8) uint8 {
	x, err := UInt8Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func UInt8Unchecked() uint8 {
	x, _ := UInt8()
	return x
}

func UInt8RangeUnchecked(min, max uint8) uint8 {
	x, _ := UInt8Range(min, max)
	return x
}
