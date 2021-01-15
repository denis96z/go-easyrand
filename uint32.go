package easyrand

func UInt32() (uint32, error) {
	return UInt32Range(0, (1<<32)-1)
}

func UInt32Range(min, max uint32) (uint32, error) {
	n, err := randomUInt64(uint64(min), uint64(max))
	return uint32(n), err
}

func UInt32Checked() uint32 {
	n, err := UInt32()
	if err != nil {
		panic(err)
	}
	return n
}

func UInt32RangeChecked(min, max uint32) uint32 {
	n, err := UInt32Range(min, max)
	if err != nil {
		panic(err)
	}
	return n
}

func UInt32Unchecked() uint32 {
	n, _ := UInt32()
	return n
}

func UInt32RangeUnchecked(min, max uint32) uint32 {
	n, _ := UInt32Range(min, max)
	return n
}
