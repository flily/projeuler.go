package p0092

func checkChainsVector(n int64, set1 []bool, set89 []bool) bool {
	ok1, ok89 := set1[n], set89[n]

	m := n
	for !ok1 && !ok89 {
		m = DigitSquareSum(m)
		ok1, ok89 = set1[m], set89[m]
	}

	if ok89 {
		set89[n] = true
	} else {
		set1[n] = true
	}

	return ok89
}

func SolveVector() int64 {
	result := int64(0)

	set1 := make([]bool, Limit)
	set89 := make([]bool, Limit)
	set1[1] = true
	set89[89] = true

	for x := int64(1); x < Limit; x++ {
		if checkChainsVector(x, set1, set89) {
			result++
		}
	}

	return result
}

func checkChain(n int64) int64 {
	m := n
	for m != 1 && m != 89 {
		m = DigitSquareSum(m)
	}
	return m
}

func SolveVectorReduced() int64 {
	vectorSize := (7 * 81) + 1
	set89 := make([]bool, vectorSize)
	set89[89] = true

	for x := int64(1); x < int64(vectorSize); x++ {
		m := checkChain(x)
		if m == 89 {
			set89[x] = true
		}
	}

	count := int64(0)
	for x := int64(1); x < Limit; x++ {
		s := DigitSquareSum(x)
		if set89[s] {
			count++
		}
	}

	return count
}
