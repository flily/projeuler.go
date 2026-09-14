package p0002

func SolveNaive() int64 {
	sum := int64(2)

	a, b := int64(1), int64(2)
	for b < LIMIT {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}

	return sum
}

func SolveNaiveNoBranch() int64 {
	sum := int64(0)

	a, b := int64(1), int64(2)
	for b < LIMIT {
		sum += b
		a, b = b, a+b
		a, b = b, a+b
		a, b = b, a+b
	}

	return sum
}
