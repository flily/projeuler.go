package p0024

func toInt64(digits []int) int64 {
	result := int64(0)
	for _, d := range digits {
		result = result*10 + int64(d)
	}
	return result
}

func generatePermutations(digits []bool, got []int, i int64, count int64, stop_at int64) (int64, int64) {
	if int(i) >= len(digits) {
		c := count + 1
		v := int64(0)
		if c >= stop_at {
			v = toInt64(got)
		}

		return c, v
	}

	foundCount := count
	for j := 0; j < len(digits); j++ {
		if !digits[j] {
			digits[j] = true
			got[i] = j
			c, v := generatePermutations(digits, got, i+1, foundCount, stop_at)
			got[i] = -1
			digits[j] = false
			foundCount = c
			if v > 0 {
				return c, v
			}
		}
	}

	return foundCount, 0
}

func SolveNaive() int64 {
	digits := make([]bool, 10)
	got := make([]int, 10)
	_, result := generatePermutations(digits, got, 0, 0, LIMIT)

	return result
}
