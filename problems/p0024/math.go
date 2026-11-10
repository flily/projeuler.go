package p0024

var factorials = []int64{
	1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880,
}

func SolveMath() int64 {
	digitsLeft := int64(9)
	numbersLeft := int64(LIMIT - 1)
	digitsUse := make([]bool, 10)
	digitsPick := make([]int, 0, 10)

	for digitsLeft >= 0 {
		blockCount := factorials[digitsLeft]
		numbersIndex := numbersLeft / blockCount

		j := int64(0)
		for i := range digitsUse {
			if !digitsUse[i] {
				if j == numbersIndex {
					digitsUse[i] = true
					digitsPick = append(digitsPick, int(i))
					break
				}

				j += 1
			}
		}

		numbersLeft -= numbersIndex * blockCount
		digitsLeft -= 1
	}

	return toInt64(digitsPick)
}
