package p0012

func getDivisors(n int64) int {
	result := 0
	i := int64(2)

	for ; i*i < n; i++ {
		if n%i == 0 {
			result += 2
		}
	}

	if i*i == n {
		result += 1
	}

	return result
}

func SolveNaive() int64 {
	n, i := int64(1), int64(2)

	for getDivisors(n) <= 500 {
		n += i
		i += 1
	}

	return n
}
