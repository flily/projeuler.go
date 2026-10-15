package p0018

import (
	"slices"
)

func maxSumPathFlood(triangle [][]int64) int64 {
	result := triangle[0]

	for _, row := range triangle[1:] {
		new_row := make([]int64, len(row))

		for i, v := range row {
			left, right := int64(0), int64(0)

			if i > 0 {
				left = result[i-1]
			}

			if i < len(result) {
				right = result[i]
			}

			new_row[i] = v + max(left, right)
		}

		result = new_row
	}

	return slices.Max(result)
}

func SolveFlood() int64 {
	return maxSumPathFlood(triangle)
}
