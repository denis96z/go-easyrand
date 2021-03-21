package easyrand

func Bool() (bool, error) {
	b, err := randomBytes(1)
	if err != nil {
		return false, err
	}
	return b[0]%2 == 0, nil
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
