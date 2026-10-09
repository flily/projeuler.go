package p0013

import (
	"math/big"
	"strconv"
)

func SolveNaive() int64 {
	sum := big.NewInt(0)

	for _, num := range NUMS {
		n := new(big.Int)
		n.SetString(num, 10)
		sum.Add(sum, n)
	}

	s := sum.String()[:10]
	result, _ := strconv.ParseInt(s, 10, 64)
	return result
}
