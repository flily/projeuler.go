package p0016

import (
	"math/big"
)

func SolveBigIntInteger() int64 {
	two := big.NewInt(2)
	thousand := big.NewInt(1000)
	result := big.NewInt(1)
	result.Exp(two, thousand, nil)

	zero, ten := big.NewInt(0), big.NewInt(10)
	sum := int64(0)
	for result.Cmp(zero) > 0 {
		digit := big.NewInt(0)
		digit.Mod(result, ten)
		result.Div(result, ten)
		sum += digit.Int64()
	}

	return sum
}

func SolveBigintString() int64 {
	two := big.NewInt(2)
	thousand := big.NewInt(1000)
	result := big.NewInt(1)
	result.Exp(two, thousand, nil)

	s := result.String()
	sum := int64(0)
	for _, digit := range s {
		sum += int64(digit - '0')
	}
	return sum
}
