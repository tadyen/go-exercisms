package rotationalcipher

type AlphaCases struct {
	lower, upper, invalid string
}

var alphaCases = AlphaCases{"lower", "upper", "invalid"}

func testLatinAlphabet(c rune) (isAlpha bool, alphaCase string) {
	// a~z
	if c >= 'a' && c <= 'z' {
		return true, alphaCases.lower
	}
	// A~Z
	if c >= 'A' && c <= 'Z' {
		return true, alphaCases.upper
	}
	return false, alphaCases.invalid
}

func RotationalCipher(plain string, shiftKey int) string {
	var outStr string = ""
	for _, c := range plain {
		isAlpha, alphaCase := testLatinAlphabet(c)
		if isAlpha {
			var a rune
			switch alphaCase {
			case alphaCases.lower:
				a = 'a'
			case alphaCases.upper:
				a = 'A'
			}
			atoi := (int(c-a)+shiftKey)%26 + int(a)
			outStr += string(rune(atoi))
		} else {
			outStr += string(c)
		}
	}
	return outStr
}
