package p0005

func SolveNaive() int64 {
	factors := make([]int64, 0, 20)
	for i := int64(2); i <= 20; i++ {
		factors = append(factors, i)
	}

	for i := range factors {
		n := factors[i]
		if n == 1 {
			continue
		}

		for j := i + 1; j < len(factors); j++ {
			m := factors[j]
			if (m % n) == 0 {
				m /= n
			}
			factors[j] = m
		}
	}

	result := int64(1)
	for _, f := range factors {
		result *= f
	}

	return result
}

func maxMultiple2(n int64) int64 {
	if n == 0 {
		return 0
	}

	result := int64(1)
	for result <= n {
		result <<= 1
	}

	return result >> 1
}

func SolveNaiveOdds() int64 {
	factors := make([]int64, 0, 10)
	for i := int64(1); i <= 20; i += 2 {
		factors = append(factors, i)
	}

	for i := range factors {
		n := factors[i]
		if n == 1 {
			continue
		}

		for j := i + 1; j < len(factors); j++ {
			m := factors[j]
			if (m % n) == 0 {
				m /= n
			}
			factors[j] = m
		}
	}

	result := int64(1)
	for _, f := range factors {
		result *= f
	}

	m := maxMultiple2(20)
	return m * result
}
