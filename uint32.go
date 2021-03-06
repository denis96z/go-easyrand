package easyrand

func UInt32() (uint32, error) {
	b, err := randomBytes(4)
	if err != nil {
		return 0, err
	}
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24, nil
}

func UInt32Range(min, max uint32) (uint32, error) {
	d := max - min
	if d == 0 {
		return min, nil
	}

	x, err := UInt32()
	if err != nil {
		return 0, err
	}

	return min + (x % d), nil
}

func UInt32Checked() uint32 {
	x, err := UInt32()
	if err != nil {
		panic(err)
	}
	return x
}

func UInt32RangeChecked(min, max uint32) uint32 {
	x, err := UInt32Range(min, max)
	if err != nil {
		panic(err)
	}
	return x
}

func UInt32Unchecked() uint32 {
	x, _ := UInt32()
	return x
}

func UInt32RangeUnchecked(min, max uint32) uint32 {
	x, _ := UInt32Range(min, max)
	return x
}
