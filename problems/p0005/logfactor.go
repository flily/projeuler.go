package p0005

import "math"

func SolveLogFactor() int64 {
	primes := []int64{2, 3, 5, 7, 11, 13, 17, 19}
	result := int64(1)

	log20 := math.Log10(20)
	for _, p := range primes {
		b := math.Log10(float64(p))
		exp := int64(log20 / b)
		result *= int64(math.Pow(float64(p), float64(exp)))
	}

	return result
}
