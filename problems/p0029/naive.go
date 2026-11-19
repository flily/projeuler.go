package p0029

import (
	"math/big"
	"slices"
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

func bigCompare(a *big.Int, b *big.Int) int {
	return a.Cmp(b)
}

func SolveNaiveArrayBisearch() int64 {
	terms := make([]*big.Int, 0)

	for a := int64(2); a <= LIMIT; a++ {
		for b := int64(2); b <= LIMIT; b++ {
			bigA := big.NewInt(a)
			bigB := big.NewInt(b)
			bigA.Exp(bigA, bigB, nil)

			_, found := slices.BinarySearchFunc(terms, bigA, bigCompare)

			if !found {
				terms = append(terms, bigA)
				slices.SortFunc(terms, bigCompare)
			}
		}
	}

	return int64(len(terms))
}
