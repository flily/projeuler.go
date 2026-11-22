package p0029

import (
	"math/big"
)

func SolveNaiveStringMap() int64 {
	terms := make(map[string]struct{})
	flag := struct{}{}

	for a := int64(2); a <= LIMIT; a++ {
		for b := int64(2); b <= LIMIT; b++ {
			bigA := big.NewInt(a)
			bigB := big.NewInt(b)
			bigA.Exp(bigA, bigB, nil)

			terms[bigA.String()] = flag
		}
	}

	return int64(len(terms))
}
