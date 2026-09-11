package p0092

func inMaps(n int64, m1 map[int64]bool, m89 map[int64]bool) (bool, bool) {
	_, ok1 := m1[n]
	_, ok2 := m89[n]
	return ok1, ok2
}

func checkChainsSet(n int64, set1 map[int64]bool, set89 map[int64]bool) bool {
	result := false
	m := n
	ok1, ok89 := inMaps(m, set1, set89)

	for !ok1 && !ok89 {
		m = DigitSquareSum(m)
		ok1, ok89 = inMaps(m, set1, set89)
	}

	if ok89 {
		result = true
		set89[n] = true
	} else {
		set1[n] = true
	}

	return result
}

func SolveSet() int64 {
	count := int64(0)
	set1 := make(map[int64]bool, Limit)
	set89 := make(map[int64]bool, Limit)

	set1[1] = true
	set89[89] = true

	for x := int64(1); x < Limit; x++ {
		r := checkChainsSet(x, set1, set89)
		if r {
			count++
		}
	}

	return count
}

func checkChainsSetOptimized(n int64, set1 map[int64]bool, set89 map[int64]bool) bool {
	result := false
	m := n
	ok1, ok89 := inMaps(m, set1, set89)

	for !ok1 && !ok89 {
		m = DigitSquareSum(m)
		ok1, ok89 = inMaps(m, set1, set89)
	}

	if ok89 {
		result = true
		set89[n] = true
	} else {
		set1[n] = true
	}

	return result
}

func SolveSetOptimized() int64 {
	count := int64(0)
	set1 := make(map[int64]bool, Limit)
	set89 := make(map[int64]bool, Limit)

	set1[1] = true
	set89[89] = true

	for x := int64(1); x < Limit; x++ {
		r := checkChainsSetOptimized(x, set1, set89)
		if r {
			count++
		}
	}

	return count
}
