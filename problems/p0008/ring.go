package p0008

func SolveRingBuffer() int64 {
	digits := make([]int64, Length)
	for i := range Length {
		digits[i] = int64(NUM[i] - '0')
	}

	largest := product(digits...)

	for i := Length; i < len(NUM); i++ {
		digits[i%Length] = int64(NUM[i] - '0')
		p := product(digits...)
		if p > largest {
			largest = p
		}
	}

	return largest
}

func SolveRingBufferDigits() int64 {
	digits := getDigitList()
	buffer := make([]int64, Length)
	copy(buffer, digits[:Length])
	largest := product(buffer...)

	for i := Length; i < len(digits); i++ {
		buffer[i%Length] = digits[i]
		p := product(buffer...)
		if p > largest {
			largest = p
		}
	}

	return largest
}
