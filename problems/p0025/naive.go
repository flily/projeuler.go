package p0025

import (
	"math/big"
)

func SolveNaiveArray() int64 {
	a := big.NewInt(1)
	b := big.NewInt(1)

	fibos := make([]*big.Int, 0, 10000)
	fibos = append(fibos, a, b)

	for len(b.String()) < 1000 {
		c := big.NewInt(1)
		c.Add(a, b)
		fibos = append(fibos, c)
		a, b = b, c
	}

	return int64(len(fibos))
}

func SolveNaiveSwap() int64 {
	a := big.NewInt(1)
	b := big.NewInt(1)
	count := int64(2)

	for len(b.String()) < 1000 {
		c := big.NewInt(1)
		c.Add(a, b)

		a, b = b, c
		count++
	}

	return count
}

func SolveNaiveCompare() int64 {
	a := big.NewInt(1)
	b := big.NewInt(1)
	limit := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(999), nil)

	count := int64(2)
	for b.Cmp(limit) < 0 {
		c := big.NewInt(1)
		c.Add(a, b)
		a, b = b, c
		count++
	}

	return count
}
