package p0027

type Cache struct {
	cache    map[int64]bool
	maxPrime int64
}

func NewCache() *Cache {
	c := &Cache{
		cache:    make(map[int64]bool),
		maxPrime: 0,
	}

	return c
}

func (c *Cache) IsPrime(n int64) bool {
	if n < c.maxPrime {
		if v, ok := c.cache[n]; ok {
			return v
		}

		return false
	}

	result := IsPrime(n)
	c.cache[n] = result
	if result && n > c.maxPrime {
		c.maxPrime = n
	}

	return result
}

func consecutivePrimeSizeCache(cache *Cache, a int64, b int64) int64 {
	x := int64(0)
	for cache.IsPrime(Func(a, b, x)) {
		x += 1
	}

	return x
}

func SolveCache() int64 {
	cache := NewCache()
	maxPrimeSize := int64(0)
	max_a, max_b := int64(0), int64(0)

	for a := int64(-999); a < 1000; a++ {
		for b := int64(-1000); b <= 1000; b++ {
			size := consecutivePrimeSizeCache(cache, a, b)
			if size > maxPrimeSize {
				maxPrimeSize = size
				max_a, max_b = a, b
			}
		}
	}

	return max_a * max_b
}
