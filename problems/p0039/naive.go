package p0039

func canBeRightTriangle(a, b, c int) bool {
	return a*a+b*b == c*c ||
		a*a+c*c == b*b ||
		b*b+c*c == a*a
}

func findRightTriangleSolutions(p int) int {
	result := 0
	for a := 0; a < p; a++ {
		for b := 0; b < p; b++ {
			c := p - a - b
			if c <= 0 {
				break
			}

			if canBeRightTriangle(a, b, c) {
				result += 1
			}
		}
	}

	return result
}

func SolveNaive() int64 {
	result, count := 0, 0
	for n := 1; n <= 1000; n++ {
		c := findRightTriangleSolutions(n)
		if c > count {
			result = n
			count = c
		}
	}

	return int64(result)
}
