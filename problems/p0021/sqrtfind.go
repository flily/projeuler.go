package p0021

func findDivisorSqrt(n int64) []int64 {
	result := make([]int64, 0, LIMIT/2)
	result = append(result, 1)

	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
			j := n / i
			if i != j {
				result = append(result, j)
			}
		}
	}

	return result
}

func SolveSqrtFind() int64 {
	result := int64(0)
	divisorSums := map[int64]int64{}

	n := int64(1)
	for n < LIMIT {
		divisors := findDivisorSqrt(n)
		s := sum(divisors)
		divisorSums[n] = s

		if s != n && divisorSums[s] == n {
			result += n + s
		}

		n += 1
	}

	return result
}

func findDivisorSqrtNoArray(n int64) int64 {
	result := int64(1)

	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			if i != n/i {
				result += i + n/i
			} else {
				result += i
			}
		}
	}

	return result
}

func SolveSqrtFindNoArray() int64 {
	result := int64(0)
	divisorSums := map[int64]int64{}

	n := int64(1)
	for n < LIMIT {
		s := findDivisorSqrtNoArray(n)
		divisorSums[n] = s

		if s != n && divisorSums[s] == n {
			result += n + s
		}

		n += 1
	}

	return result
}
