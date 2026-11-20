package p0030

import (
	"fmt"
)

func digitPowerSumString(n int64, pow int64) int64 {
	sum := int64(0)
	s := fmt.Sprintf("%d", n)
	for _, c := range s {
		digit := int64(c - '0')
		sum += powInt64(digit, pow)
	}
	return sum
}

func SolveString() int64 {
	sum := int64(0)

	for i := int64(2); i < MAX_NUMBER; i++ {
		if digitPowerSumString(i, 5) == i {
			sum += i
		}
	}

	return sum
}

func digitPowerSumStringInline(n int64, pow int64) int64 {
	sum := int64(0)
	s := fmt.Sprintf("%d", n)
	for _, c := range s {
		digit := int64(c - '0')
		r := int64(1)
		for range pow {
			r *= digit
		}
		sum += r
	}
	return sum
}

func SolveStringInline() int64 {
	sum := int64(0)

	for i := int64(2); i < MAX_NUMBER; i++ {
		if digitPowerSumStringInline(i, 5) == i {
			sum += i
		}
	}

	return sum
}
