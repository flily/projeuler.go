package p0023

type FactorSumCache struct {
	cache map[int]int
}

func NewFactorSumCache() *FactorSumCache {
	cache := &FactorSumCache{
		cache: make(map[int]int),
	}
	return cache
}

func (c *FactorSumCache) CheckNumberType(n int) int {
	sum, ok := c.cache[n]
	if !ok {
		sum = sumOfFactors(n)
		c.cache[n] = sum
	}

	if sum < n {
		return Deficient

	} else if sum == n {
		return Perfect

	} else {
		return Abundant
	}
}

func SolveWithFactorSumCache() int64 {
	result := int64(0)

	cache := NewFactorSumCache()
	for i := 1; i <= Limit; i++ {
		canBeSumOfTwoAbundant := false
		for j := 1; j < i-1; j++ {
			k := i - j
			if cache.CheckNumberType(j) == Abundant && cache.CheckNumberType(k) == Abundant {
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
