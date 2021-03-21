package easyrand

func Int16() (int16, error) {
	b, err := randomBytes(2)
	if err != nil {
		return 0, err
	}
	return int16(b[0]) | int16(b[1])<<8, nil
}

func Int16Range(min, max int16) (int16, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := Int16()
	if err != nil {
		return 0, err
	}

	// XXX: x should always be positive
	x &= int16(^(uint16(1) << 15))

	return min + (x % d), nil
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
