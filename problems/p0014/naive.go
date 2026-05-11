package p0014

const LIMIT = 1_000_000

func collatzSeqSize(n int64) int64 {
	m := n
	result := int64(1)

	for m > 1 {
		if m%2 == 0 {
			m = m / 2
		} else {
			m = 3*m + 1
		}
		result += 1
	}

	return result
}

func SolveNaive() int64 {
	maxSize := int64(0)
	result := int64(0)

	for i := int64(1); i < LIMIT; i++ {
		size := collatzSeqSize(i)
		if size > maxSize {
			maxSize = size
			result = i
		}
	}

	return result
}
