package easyrand

import (
	"crypto/rand"
)

func randomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func randomBytesAlphabet(alphabet []byte, n int) ([]byte, error) {
	b, err := randomBytes(n)
	if err != nil {
		return nil, err
	}

	for i := 0; i < n; i++ {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}

	return b, nil
}
