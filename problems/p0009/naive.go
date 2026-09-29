package p0009

func SolveNaive() int64 {
	for a := int64(1); a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			c := 1000 - a - b
			if (a*a)+(b*b) == c*c {
				return a * b * c
			}
		}
	}

	return 0
}
