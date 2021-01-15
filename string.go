package easyrand

func String(alphabet string, n int) (string, error) {
	b, err := randomBytesAlphabet([]byte(alphabet), n)
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

	alphaNumLower = alphaLower + num
	alphaNumUpper = alphaUpper + num

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

func AlphaLowerString(n int) (string, error) {
	return String(alphaLower, n)
}

func AlphaLowerStringChecked(n int) string {
	s, err := AlphaLowerString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaLowerStringUnchecked(n int) string {
	s, _ := AlphaLowerString(n)
	return s
}

func AlphaLowerVarLenString(nmin, nmax int) (string, error) {
	n, err := IntRange(nmin, nmax)
	if err != nil {
		return "", err
	}
	return AlphaLowerString(n)
}

func AlphaLowerVarLenStringChecked(nmin, nmax int) string {
	s, err := AlphaLowerVarLenString(nmin, nmax)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaLowerVarLenStringUnchecked(nmin, nmax int) string {
	s, _ := AlphaLowerVarLenString(nmin, nmax)
	return s
}

func AlphaUpperString(n int) (string, error) {
	return String(alphaUpper, n)
}

func AlphaUpperStringChecked(n int) string {
	s, err := AlphaUpperString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaUpperStringUnchecked(n int) string {
	s, _ := AlphaUpperString(n)
	return s
}

func AlphaUpperVarLenString(nmin, nmax int) (string, error) {
	n, err := IntRange(nmin, nmax)
	if err != nil {
		return "", err
	}
	return AlphaUpperString(n)
}

func AlphaUpperVarLenStringChecked(nmin, nmax int) string {
	s, err := AlphaLowerVarLenString(nmin, nmax)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaUpperVarLenStringUnchecked(nmin, nmax int) string {
	s, _ := AlphaLowerVarLenString(nmin, nmax)
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

func AlphaNumLowerString(n int) (string, error) {
	return String(alphaNumLower, n)
}

func AlphaNumLowerStringChecked(n int) string {
	s, err := AlphaLowerString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaNumLowerStringUnchecked(n int) string {
	s, _ := AlphaLowerString(n)
	return s
}

func AlphaNumLowerVarLenString(nmin, nmax int) (string, error) {
	n, err := IntRange(nmin, nmax)
	if err != nil {
		return "", err
	}
	return AlphaNumLowerString(n)
}

func AlphaNumLowerVarLenStringChecked(nmin, nmax int) string {
	s, err := AlphaNumLowerVarLenString(nmin, nmax)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaNumLowerVarLenStringUnchecked(nmin, nmax int) string {
	s, _ := AlphaNumLowerVarLenString(nmin, nmax)
	return s
}

func AlphaNumUpperString(n int) (string, error) {
	return String(alphaNumUpper, n)
}

func AlphaNumUpperStringChecked(n int) string {
	s, err := AlphaUpperString(n)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaNumUpperStringUnchecked(n int) string {
	s, _ := AlphaLowerString(n)
	return s
}

func AlphaNumUpperVarLenString(nmin, nmax int) (string, error) {
	n, err := IntRange(nmin, nmax)
	if err != nil {
		return "", err
	}
	return AlphaNumUpperString(n)
}

func AlphaNumUpperVarLenStringChecked(nmin, nmax int) string {
	s, err := AlphaNumUpperVarLenString(nmin, nmax)
	if err != nil {
		panic(err)
	}
	return s
}

func AlphaNumUpperVarLenStringUnchecked(nmin, nmax int) string {
	s, _ := AlphaNumUpperVarLenString(nmin, nmax)
	return s
}
