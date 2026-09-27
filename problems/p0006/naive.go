package p0006

func SolveNaive() int64 {
	sumOfSquares := int64(0)
	for i := int64(1); i <= 100; i++ {
		sumOfSquares += i * i
	}

	squareOfSum := int64((1 + 100) * 100 / 2)
	squareOfSum *= squareOfSum
	return squareOfSum - sumOfSquares
}
