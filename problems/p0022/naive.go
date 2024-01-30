package p0022

import (
	"sort"
)

func NameScore(name string) int64 {
	score := 0
	for _, c := range name {
		score += int(c - 'A' + 1)
	}

	return int64(score)
}

func SolveNaive() int64 {
	names := Load()
	sort.Strings(names)

	result := int64(0)
	for i, name := range names {
		result += int64(i+1) * NameScore(name)
	}

	return result
}
