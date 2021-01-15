package easyrand

func Bool() (bool, error) {
	n, err := randomInt64(0, 2)
	return n%2 == 0, err
}

func BoolChecked() bool {
	b, err := Bool()
	if err != nil {
		panic(err)
	}
	return b
}

func BoolUnchecked() bool {
	b, _ := Bool()
	return b
}
