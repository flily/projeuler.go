package p0003

import "slices"

type PrimeTable struct {
	Primes []int64
	Set    map[int64]struct{}
}

var initPrimes = []int64{
	3, 5, 7, 11, 13, 17, 19,
}

func NewPrimeTable(cap int) PrimeTable {
	size := max(cap, len(initPrimes))
	t := PrimeTable{
		Primes: make([]int64, 0, size),
		Set:    make(map[int64]struct{}, size),
	}

	for _, p := range initPrimes {
		t.Primes = append(t.Primes, p)
		t.Set[p] = struct{}{}
	}

	return t
}

func (t *PrimeTable) Last() int64 {
	return t.Primes[len(t.Primes)-1]
}

func (t *PrimeTable) Check(n int64) bool {
	largest := t.Last()
	if n < largest {
		_, ok := t.Set[n]
		return ok
	}

	for _, p := range t.Primes {
		if n%p == 0 {
			return false
		}
	}

	t.Primes = append(t.Primes, n)
	t.Set[n] = struct{}{}
	return true
}

func SolvePrimeTable() int64 {
	t := NewPrimeTable(1000)
	factors := make([]int64, 0, 100)

	n := int64(NUMBER)
	i := int64(3)
	for ; n > 0 && i <= n; i += 2 {
		if t.Check(i) && n%i == 0 {
			factors = append(factors, i)
			n = RemoveFactor(n, i)
		}
	}

	return factors[len(factors)-1]
}

type PrimeList []int64

func NewPrimeList(cap int) PrimeList {
	size := max(cap, len(initPrimes))
	l := make([]int64, 0, size)
	l = append(l, initPrimes...)

	return l
}

func (l *PrimeList) Last() int64 {
	return (*l)[len(*l)-1]
}

func (l *PrimeList) Check(n int64) bool {
	largest := l.Last()
	if n < largest {
		_, ok := slices.BinarySearch(*l, n)
		return ok
	}

	for _, p := range *l {
		if n%p == 0 {
			return false
		}
	}

	*l = append(*l, n)
	return true
}

func SolvePrimeList() int64 {
	t := NewPrimeList(1000)
	factors := make([]int64, 0, 100)

	n := int64(NUMBER)
	i := int64(3)
	for ; n > 0 && i <= n; i += 2 {
		if t.Check(i) && n%i == 0 {
			factors = append(factors, i)
			n = RemoveFactor(n, i)
		}
	}

	return factors[len(factors)-1]
}
