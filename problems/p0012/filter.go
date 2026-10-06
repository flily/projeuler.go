package p0012

func SolveFilter() int64 {
	n, i := int64(1), int64(2)
	commonFactors := int64(2 * 3 * 5 * 7 * 11 * 13)

	for {
		if n%commonFactors == 0 {
			if getDivisors(n) >= LIMIT {
				break
			}
		}

		n += i
		i += 1
	}

	return n
}
