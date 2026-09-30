package p0008

func getDigitList() []int64 {
	digits := make([]int64, len(NUM))
	for i, c := range NUM {
		digits[i] = int64(c - '0')
	}

	return digits
}

func SolvePrecomputedDigits() int64 {
	largest := int64(0)

	digits := getDigitList()
	for i := 0; i <= len(NUM)-Length; i++ {
		p := product(digits[i : i+Length]...)
		if p > largest {
			largest = p
		}
	}

	return largest
}
