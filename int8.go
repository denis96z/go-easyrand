package easyrand

func Int8() (int8, error) {
	b, err := randomBytes(1)
	if err != nil {
		return 0, err
	}
	return int8(b[0]), nil
}

func Int8Range(min, max int8) (int8, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := Int8()
	if err != nil {
		return 0, err
	}

	//XXX: x should always be positive
	x &= int8(^(uint8(1) << 7))

	return min + (x % d), nil
}

func Int8Checked() int8 {
	x, err := Int8()
	if err != nil {
		panic(err)
	}
	return x
}

func Int8RangeChecked(min, max int8) int8 {
	x, err := Int8Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func Int8Unchecked() int8 {
	x, _ := Int8()
	return x
}

func Int8RangeUnchecked(min, max int8) int8 {
	x, _ := Int8Range(min, max)
	return x
}
