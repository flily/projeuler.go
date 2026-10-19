package p0021

func findDivisors(n int64) []int64 {
	result := make([]int64, 0, LIMIT/2)
	result = append(result, 1)

	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	return result
}

func sum(nums []int64) int64 {
	result := int64(0)
	for _, n := range nums {
		result += n
	}
	return result
}

func SolveNaive() int64 {
	result := int64(0)
	divisorSums := map[int64]int64{}

	n := int64(1)
	for n < LIMIT {
		divisors := findDivisors(n)
		s := sum(divisors)
		divisorSums[n] = s

		if s != n && divisorSums[s] == n {
			result += n + s
		}

		n += 1
	}

	return result
}
