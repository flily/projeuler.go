package p0001

func SolveNaive() int64 {
	sum := int64(0)
	for i := int64(1); i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum
}
