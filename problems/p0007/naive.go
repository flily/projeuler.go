package p0007

func isPrime(n int64) bool {
	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func SolveNaive() int64 {
	count := int64(1) // counting the prime number 2

	n := int64(3)
	for ; count < LIMIT; n += 2 {
		if isPrime(n) {
			count++
		}
	}

	return n - 2
}
