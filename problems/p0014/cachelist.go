package p0014

type CacheList []int64

func NewCacheList(size int) CacheList {
	c := make(CacheList, size)
	c[0] = 1
	c[1] = 1
	c[2] = 2
	c[3] = 9
	c[4] = 4

	return c
}

func (c CacheList) CalcLength(n int64) int64 {
	m := n
	result := int64(1)

	for m > 1 {
		if m < int64(len(c)) {
			if v := c[m]; v != 0 {
				result += v
				break
			}
		}

		result += 1
		if m%2 == 0 {
			m = m / 2
		} else {
			m = 3*m + 1
		}
	}

	c[n] = result
	return result
}

func SolveCacheList() int64 {
	cache := NewCacheList(LIMIT + 1)
	maxSize := int64(0)
	result := int64(0)

	for i := int64(1); i < LIMIT; i++ {
		size := cache.CalcLength(i)
		if size > maxSize {
			maxSize = size
			result = i
		}
	}

	return result
}
