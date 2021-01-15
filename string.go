package easyrand

func String(alphabet string, n int) (string, error) {
	b, err := randomBytes([]byte(alphabet), n)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

const (
	num = "0123456789"

	alphaLower = "abcdefghijklmnopqrstuvwxyz"
	alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	alpha = alphaLower + alphaUpper

	alphaNum = alphaLower + alphaUpper + num
)

func NumString(n int) (string, error) {
	return String(num, n)
}

func NumStringChecked(n int) string {
	s, err := NumString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func NumStringUnchecked(n int) string {
	s, _ := NumString(n)
	return s
}

func NumVarLenString(nmin, nmax int) (string, error) {
	n, err := IntRange(nmin, nmax)
	if err != nil {
		return "", err
	}
	return NumString(n)
}

func NumVarLenStringChecked(nmin, nmax int) string {
	s, err := NumVarLenString(nmin, nmax)
	if err != nil {
		panic(err)
	}
	return s
}

func NumVarLenStringUnchecked(nmin, nmax int) string {
	s, _ := NumVarLenString(nmin, nmax)
	return s
}

func AlphaString(n int) (string, error) {
	return String(alpha, n)
}

func AlphaStringChecked(n int) string {
	s, err := AlphaString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaStringUnchecked(n int) string {
	s, _ := AlphaString(n)
	return s
}

func AlphaVarLenString(nmin, nmax int) (string, error) {
	n, err := IntRange(nmin, nmax)
	if err != nil {
		return "", err
	}
	return AlphaString(n)
}

func AlphaVarLenStringChecked(nmin, nmax int) string {
	s, err := AlphaVarLenString(nmin, nmax)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaVarLenStringUnchecked(nmin, nmax int) string {
	s, _ := AlphaVarLenString(nmin, nmax)
	return s
}

func AlphaNumString(n int) (string, error) {
	return String(alphaNum, n)
}

func AlphaNumStringChecked(n int) string {
	s, err := AlphaString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaNumStringUnchecked(n int) string {
	s, _ := AlphaString(n)
	return s
}

func AlphaNumVarLenString(nmin, nmax int) (string, error) {
	n, err := IntRange(nmin, nmax)
	if err != nil {
		return "", err
	}
	return AlphaNumString(n)
}

func AlphaNumVarLenStringChecked(nmin, nmax int) string {
	s, err := AlphaNumVarLenString(nmin, nmax)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaNumVarLenStringUnchecked(nmin, nmax int) string {
	s, _ := AlphaNumVarLenString(nmin, nmax)
	return s
}
