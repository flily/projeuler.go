package p0027

func Func(a int64, b int64, n int64) int64 {
	return n*n + a*n + b
}

func IsPrime(n int64) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func consecutivePrimeSize(a int64, b int64) int64 {
	x := int64(0)
	for IsPrime(Func(a, b, x)) {
		x += 1
	}

	return x
}

func SolveNaive() int64 {
	maxPrimeSize := int64(0)
	max_a, max_b := int64(0), int64(0)

	for a := int64(-999); a < 1000; a++ {
		for b := int64(-1000); b <= 1000; b++ {
			size := consecutivePrimeSize(a, b)
			if size > maxPrimeSize {
				maxPrimeSize = size
				max_a, max_b = a, b
			}
		}
	}

	return max_a * max_b
}
