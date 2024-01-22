package p0023

const (
	Perfect   = 0
	Deficient = 1
	Abundant  = 2

	Limit = 28123
)

func sumOfFactors(n int) int {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i*i != n {
				sum += n / i
			}
		}
	}

	return sum
}

func checkNumberType(n int) int {
	sum := sumOfFactors(n)
	if sum < n {
		return Deficient

	} else if sum == n {
		return Perfect

	} else {
		return Abundant
	}
}

func SolveNaive() int64 {
	result := int64(0)
	for i := 1; i <= Limit; i++ {
		canBeSumOfTwoAbundant := false
		for j := 1; j < i-1; j++ {
			k := i - j
			if checkNumberType(j) == Abundant && checkNumberType(k) == Abundant {
				canBeSumOfTwoAbundant = true
				break
			}
		}

		if !canBeSumOfTwoAbundant {
			result += int64(i)
		}
	}

	return result
}
