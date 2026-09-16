package p0003

func RemoveFactor(n int64, factor int64) int64 {
	for n%factor == 0 {
		n /= factor
	}

	return n
}

func SolveRemoveFactor() int64 {
	n := int64(NUMBER)

	i := int64(3)
	largest := i
	for n > 0 && i <= n {
		if n%i == 0 && IsPrime(i) {
			largest = i
			n = RemoveFactor(n, i)
		}

		i += 2
	}

	return largest
}
