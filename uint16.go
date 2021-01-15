package easyrand

func UInt16() (uint16, error) {
	return UInt16Range(0, (1<<16)-1)
}

func UInt16Range(min, max uint16) (uint16, error) {
	n, err := randomUInt64(uint64(min), uint64(max))
	return uint16(n), err
}

func UInt16Checked() uint16 {
	n, err := UInt16()
	if err != nil {
		panic(err)
	}
	return n
}

func UInt16RangeChecked(min, max uint16) uint16 {
	n, err := UInt16Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UInt16Unchecked() uint16 {
	n, _ := UInt16()
	return n
}

func UInt16RangeUnchecked(min, max uint16) uint16 {
	n, _ := UInt16Range(min, max)
	return n
}
