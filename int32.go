package easyrand

func Int32() (int32, error) {
	b, err := randomBytes(4)
	if err != nil {
		return 0, err
	}
	return int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24, nil
}

func Int32Range(min, max int32) (int32, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := Int32()
	if err != nil {
		return 0, err
	}

	//XXX: x should always be positive
	x &= int32(^(uint32(1) << 31))

	return min + (x % d), nil
}

func Int32Checked() int32 {
	x, err := Int32()
	if err != nil {
		panic(err)
	}
	return x
}

func Int32RangeChecked(min, max int32) int32 {
	x, err := Int32Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func Int32Unchecked() int32 {
	x, _ := Int32()
	return x
}

func Int32RangeUnchecked(min, max int32) int32 {
	x, _ := Int32Range(min, max)
	return x
}
