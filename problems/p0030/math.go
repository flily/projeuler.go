package p0030

func powInt64(base int64, exp int64) int64 {
	result := int64(1)
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func digitPowerSumMath(n int64, pow int64) int64 {
	sum := int64(0)

	for n > 0 {
		digit := n % 10
		sum += powInt64(digit, pow)
		n /= 10
	}

	return sum
}

func SolveMath() int64 {
	sum := int64(0)

	for i := int64(2); i < MAX_NUMBER; i++ {
		if digitPowerSumMath(i, 5) == i {
			sum += i
		}
	}

	return sum
}

func digitPowerSumMathInline(n int64, pow int64) int64 {
	sum := int64(0)

	for n > 0 {
		digit := n % 10
		r := int64(1)
		for range pow {
			r *= digit
		}
		sum += r
		n /= 10
	}

	return sum
}

func SolveMathInline() int64 {
	sum := int64(0)

	for i := int64(2); i < MAX_NUMBER; i++ {
		if digitPowerSumMathInline(i, 5) == i {
			sum += i
		}
	}

	return sum
}
