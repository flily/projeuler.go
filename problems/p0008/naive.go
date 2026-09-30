package p0008

func product(nums ...int64) int64 {
	result := int64(1)
	for _, n := range nums {
		result *= n
	}
	return result
}

func SolveNaive() int64 {
	largest := int64(0)

	for i := 0; i <= len(NUM)-Length; i++ {
		digits := make([]int64, Length)
		for j := 0; j < Length; j++ {
			digits[j] = getDigitList()[i+j]
		}

		p := product(digits...)
		if p > largest {
			largest = p
		}
	}

	return largest
}

func SolveNaiveStaticBuffer() int64 {
	largest := int64(0)

	digits := make([]int64, Length)
	for i := 0; i <= len(NUM)-Length; i++ {
		for j := 0; j < Length; j++ {
			digits[j] = getDigitList()[i+j]
		}

		p := product(digits...)
		if p > largest {
			largest = p
		}
	}

	return largest
}
