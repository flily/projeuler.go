package p0018

func maxPathSumRecursive(triangle [][]int64, row int, col int, sum int64) int64 {
	current := triangle[row][col]

	if row >= len(triangle)-1 {
		return sum + current
	}

	next := sum + current
	left := maxPathSumRecursive(triangle, row+1, col, next)
	right := maxPathSumRecursive(triangle, row+1, col+1, next)

	return max(left, right)
}

func SolveBruteforceRecursive() int64 {
	return maxPathSumRecursive(triangle, 0, 0, 0)
}
