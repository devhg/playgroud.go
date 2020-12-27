package algorithm

/**
快速幂
*/
func Pow(base float64, exponent int) (ret float64) {
	if equal(base, 0.0) {
		return
	}
	var absExponent uint = uint(exponent)
	if exponent < 0 {
		absExponent = uint(-exponent)
	}

	ret = PoWithUnsignedExponent(base, absExponent)
	if exponent < 0 {
		return 1.0 / ret
	}
	return
}

func PoWithUnsignedExponent(base float64, absExponent uint) (ret float64) {
	if absExponent == 0 {
		return 1
	}
	if absExponent == 1 {
		return base
	}

	ret = PoWithUnsignedExponent(base, absExponent>>1)
	ret *= ret
	if absExponent&1 == 1 {
		ret *= base
	}
	return
}

func equal(a, b float64) bool {
	if (a-b > -0.000001) || (a-b < 0.000001) {
		return true
	}
	return false
}
