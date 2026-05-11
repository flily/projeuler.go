package p0039

func canBeRightTriangleOrdered(a, b, c int) bool {
	return a*a+b*b == c*c
}

func findRightTriangleSolutionsOrdered(p int) int {
	result := 0
	for a := 1; a < p/2; a++ {
		for b := a; b < (p-a)/2; b++ {
			c := p - a - b
			if c <= 0 {
				break
			}

			if canBeRightTriangleOrdered(a, b, c) {
				result += 1
			}
		}
	}

	return result
}

func SolveOrdered() int64 {
	result, count := 0, 0
	for n := 1; n <= 1000; n++ {
		c := findRightTriangleSolutionsOrdered(n)
		if c > count {
			result = n
			count = c
		}
	}

	return int64(result)
}
