package p0029

import (
	"math/big"
	"slices"
)

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

func sortSortedInsertInto[T any](arr *[]T, item T, cmp func(a, b T) int) {
	*arr = append(*arr, item)
	i := 0
	for ; cmp(item, (*arr)[i]) > 0; i++ {
	}

	for j := len(*arr) - 1; j > i; j-- {
		(*arr)[j], (*arr)[j-1] = (*arr)[j-1], (*arr)[j]
	}

	(*arr)[i] = item
}

func SolveNaiveArrayInsertSort() int64 {
	terms := make([]*big.Int, 0)

	for a := int64(2); a <= LIMIT; a++ {
		for b := int64(2); b <= LIMIT; b++ {
			bigA := big.NewInt(a)
			bigB := big.NewInt(b)
			bigA.Exp(bigA, bigB, nil)

			_, found := slices.BinarySearchFunc(terms, bigA, bigCompare)

			if !found {
				sortSortedInsertInto(&terms, bigA, bigCompare)
			}
		}
	}

	return int64(len(terms))
}
