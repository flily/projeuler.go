package p0028

func SolveCount() int64 {
	sum := int64(1)

	i, d, c := int64(3), int64(2), int64(0)
	for i <= (LIMIT * LIMIT) {
		sum += i
		c += 1
		if c == 4 {
			c = 0
			d += 2
		}

		i += d
	}

	return sum
}
