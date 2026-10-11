package p0015

func callLatticePaths(width int64, height int64) int64 {
	row := make([]int64, width)
	for i := range row {
		row[i] = 1
	}

	for range height - 1 {
		for i := int64(1); i < width; i++ {
			row[i] += row[i-1]
		}
	}

	return row[width-1]
}

func SolveNaive() int64 {
	return callLatticePaths(20+1, 20+1)
}
