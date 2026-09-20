package p0004

func SolveOrderInString() int64 {
	i, j := int64(0), int64(0)
	for i < 999 && j < 999 {
		x, y := 999-i, 999-j
		for x < 1000 {
			n := x * y
			if isPalindromeString(n) {
				return n
			}

			x += 1
			y -= 1
		}

		if i == j {
			j += 1
		} else {
			i += 1
		}
	}

	// impossible to here
	return 0
}

func SolveOrderInInt6() int64 {
	i, j := int64(0), int64(0)
	for i < 999 && j < 999 {
		x, y := 999-i, 999-j
		for x < 1000 {
			n := x * y
			if isPalindromeInt6(n) {
				return n
			}

			x += 1
			y -= 1
		}

		if i == j {
			j += 1
		} else {
			i += 1
		}
	}

	// impossible to here
	return 0
}
