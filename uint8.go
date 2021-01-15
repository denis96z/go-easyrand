package easyrand

func UInt8() (uint8, error) {
	return UInt8Range(0, (1<<8)-1)
}

func UInt8Range(min, max uint8) (uint8, error) {
	n, err := randomUInt64(uint64(min), uint64(max))
	return uint8(n), err
}

func UInt8Checked() uint8 {
	n, err := UInt8()
	if err != nil {
		panic(err)
	}
	return n
}

func UInt8RangeChecked(min, max uint8) uint8 {
	n, err := UInt8Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UInt8Unchecked() uint8 {
	n, _ := UInt8()
	return n
}

func UInt8RangeUnchecked(min, max uint8) uint8 {
	n, _ := UInt8Range(min, max)
	return n
}
