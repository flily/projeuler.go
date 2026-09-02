package p0001

func sumOfMultiples(n int64, k int64) int64 {
	m := (n - 1) / k
	return k * m * (m + 1) / 2
}

func SolveFormula() int64 {
	n := int64(1000)
	return sumOfMultiples(n, 3) + sumOfMultiples(n, 5) - sumOfMultiples(n, 15)
}
