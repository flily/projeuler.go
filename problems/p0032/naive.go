package p0032

func checkPandigital(nums ...int) bool {
	var digits [10]bool

	for _, num := range nums {
		if num == 0 {
			return false
		}

		for num > 0 {
			digit := num % 10
			if digit == 0 {
				return false
			}

			if digits[digit] {
				return false
			}

			digits[digit] = true
			num /= 10
		}
	}

	for i := 1; i < 10; i++ {
		if !digits[i] {
			return false
		}
	}

	return true
}

func SolveNaive() int64 {
	result := make(map[int64]struct{})
	for a := range 1000 {
		for b := a; b < 10000; b++ {
			c := a * b
			if checkPandigital(a, b, c) {
				result[int64(c)] = struct{}{}
			}
		}
	}

	sum := int64(0)
	for c := range result {
		sum += c
	}

	return sum
}
