package p0013

func toDecimalInt(s string, target []int) {
	sSize := len(s)
	for i, c := range s {
		target[sSize-1-i] = int(c - '0')
	}
}

func decimalAdd(a []int, b []int) {
	carry := 0
	l := min(len(a), len(b))

	for i := range l {
		s := a[i] + b[i] + carry
		a[i] = s % 10
		carry = s / 10
	}
}

func SolveStringInt() int64 {
	resultSize := 100
	sum := make([]int, resultSize)
	n := make([]int, resultSize)

	for _, num := range NUMS {
		toDecimalInt(num, n)
		decimalAdd(sum, n)
	}

	result := int64(0)
	c, found := 10, false
	for i := resultSize - 1; i > 0 && c > 0; i-- {
		if !found {
			if sum[i] == 0 {
				continue
			} else {
				found = true
			}
		}

		result = result*10 + int64(sum[i])
		c--
	}

	return result
}
