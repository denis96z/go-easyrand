package easyrand

func Bytes(n int) ([]byte, error) {
	return randomBytes(n)
}

func BytesChecked(n int) []byte {
	b, err := Bytes(n)
	if err != nil {
		panic(err)
	}
	return b
}

func BytesUnchecked(n int) []byte {
	b, _ := Bytes(n)
	return b
}

func BytesAlphabet(alphabet []byte, n int) ([]byte, error) {
	return randomBytesAlphabet(alphabet, n)
}

func BytesAlphabetChecked(alphabet []byte, n int) []byte {
	b, err := BytesAlphabet(alphabet, n)
	if err != nil {
		panic(err)
	}
	return b
}

func BytesAlphabetUnchecked(alphabet []byte, n int) []byte {
	b, _ := BytesAlphabet(alphabet, n)
	return b
}
