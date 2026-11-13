package p0026

import (
	"math"
	"slices"
)

func power10(n int64) int64 {
	result := int64(1)
	for range n {
		result *= 10
	}
	return result
}

type keyDigit struct {
	r int64
	d int64
}

func getCycleLength(n int64) int64 {
	base := power10(int64(math.Ceil(math.Log10(float64(n)))))
	m := base
	digits := make([]keyDigit, 0)
	loopStart := -1
	for m != 0 {
		r := m / n
		d := m % n
		if d == 0 {
			break
		}

		key := keyDigit{r: r, d: d}
		if slices.Contains(digits, key) {
			loopStart = slices.Index(digits, key)
			break
		}

		digits = append(digits, key)
		m = d
		for m < n {
			m *= 10
		}
	}

	if loopStart < 0 {
		return 0
	}

	return int64(len(digits) - loopStart)
}

func SolveNaive() int64 {
	maxCycleLength := int64(0)
	maxCycleNumber := int64(0)

	for i := int64(2); i < 1000; i++ {
		l := getCycleLength(i)
		if l > maxCycleLength {
			maxCycleLength = l
			maxCycleNumber = i
		}
	}

	return maxCycleNumber
}

func isPrime(n int64) bool {
	// assume n > 3 and n is odd
	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func SolveNaivePrime() int64 {
	maxCycleLength := int64(0)
	maxCycleNumber := int64(0)

	for i := int64(3); i < 1000; i += 2 {
		if !isPrime(i) {
			continue
		}

		l := getCycleLength(i)
		if l > maxCycleLength {
			maxCycleLength = l
			maxCycleNumber = i
		}
	}

	return maxCycleNumber
}
