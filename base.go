package easyrand

import (
	"crypto/rand"
	"math/big"
)

func randomInt64(mn, mx int64) (int64, error) {
	x := big.NewInt(mx - mn)

	n, err := rand.Int(rand.Reader, x)
	if err != nil {
		return 0, err
	}

	return mn + n.Int64(), nil
}

func randomUInt64(mn, mx uint64) (uint64, error) {
	x := new(big.Int).SetUint64(mx - mn)

	n, err := rand.Int(rand.Reader, x)
	if err != nil {
		return 0, err
	}

	return mn + n.Uint64(), nil
}

func randomBytes(alphabet []byte, n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	res := make([]byte, n)
	for i, c := range b {
		res[i] = alphabet[int(c)%len(alphabet)]
	}

	return res, nil
}
