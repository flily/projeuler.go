package p0023

/*
def solve_with_substraction() -> int:
    abundant_numbers = []
    abundant_set = set()

    n = 1
    result = 0
    while n <= LIMIT:
        factor_sum = sum_of_factors(n)
        if factor_sum > n:
            abundant_numbers.append(n)
            abundant_set.add(n)

        is_sum_of_abundant = False
        for x in abundant_numbers:
            if x >= n:
                break

            if n - x in abundant_set:
                is_sum_of_abundant = True
                break

        if not is_sum_of_abundant:
            result += n

        n += 1

    return result
*/

func SolveWithSubstraction() int64 {
	abundantNumbers := make([]int, 0, Limit)
	abundantSet := make(map[int]bool)

	n := 1
	result := 0
	for n <= Limit {
		factorSum := sumOfFactors(n)
		if factorSum > n {
			abundantNumbers = append(abundantNumbers, n)
			abundantSet[n] = true
		}

		isSumOfAbundant := false
		for _, x := range abundantNumbers {
			if x >= n {
				break
			}

			if _, ok := abundantSet[n-x]; ok {
				isSumOfAbundant = true
				break
			}
		}

		if !isSumOfAbundant {
			result += n
		}

		n += 1
	}

	return int64(result)
}
