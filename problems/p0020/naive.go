package p0020

import (
	"math/big"
)

func SolveNaive() int64 {
	one := big.NewInt(1)
	i := big.NewInt(1)
	result := big.NewInt(1)

	limit := big.NewInt(100)
	for i.Cmp(limit) <= 0 {
		result.Mul(result, i)
		i.Add(i, one)
	}

	sum := int64(0)
	for _, digit := range result.String() {
		sum += int64(digit - '0')
	}
	return sum
}
