package p0003

func IsPrime(n int64) bool {
	if n <= 2 {
		return true
	}

	i := int64(3)
	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i += 2
	}
	return true
}

func SolveNaive() int64 {
	n := int64(NUMBER)

	i := int64(3)
	largest := i
	for 2*i <= n {
		if n%i == 0 {
			if IsPrime(i) {
				largest = i
			}
		}
		i += 2
	}

	return largest
}
