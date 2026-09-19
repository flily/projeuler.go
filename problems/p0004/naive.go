package p0004

import (
	"fmt"
)

func reverseString(s string) string {
	b := []byte(s)
	l := len(b)
	for i := 0; i < l/2; i++ {
		b[i], b[l-1-i] = b[l-1-i], b[i]
	}
	return string(b)
}

func isPalindromeString(n int64) bool {
	s := fmt.Sprintf("%d", n)
	return s == reverseString(s)
}

func isPalindromeInt(n int64) bool {
	digits := [20]int64{}
	index := 0

	for n > 0 {
		digits[index] = n % 10
		n /= 10
		index++
	}

	for i := 0; i < index/2; i++ {
		if digits[i] != digits[index-1-i] {
			return false
		}
	}

	return true
}

func isPalindromeInt6(n int64) bool {
	if (n / 100_000) != (n % 10) {
		return false
	}

	if (n / 10_000 % 10) != (n / 10 % 10) {
		return false
	}

	if (n / 1_000 % 10) != (n / 100 % 10) {
		return false
	}

	return true
}

func isPalindromeInt6And(n int64) bool {
	return (n/100_000) == (n%10) &&
		(n/10_000%10) == (n/10%10) &&
		(n/1_000%10) == (n/100%10)
}

func SolveNaiveString() int64 {
	result := int64(0)
	for i := int64(100); i < 1000; i++ {
		for j := int64(100); j < 1000; j++ {
			p := i * j
			if p > result && isPalindromeString(p) {
				result = p
			}
		}
	}

	return result
}

func SolveNaiveInt() int64 {
	result := int64(0)
	for i := int64(100); i < 1000; i++ {
		for j := int64(100); j < 1000; j++ {
			p := i * j
			if p > result && isPalindromeInt(p) {
				result = p
			}
		}
	}

	return result
}

func SolveNaiveInt6() int64 {
	result := int64(0)
	for i := int64(100); i < 1000; i++ {
		for j := int64(100); j < 1000; j++ {
			p := i * j
			if p > result && isPalindromeInt6(p) {
				result = p
			}
		}
	}

	return result
}

func SolveNaiveInt6And() int64 {
	result := int64(0)
	for i := int64(100); i < 1000; i++ {
		for j := int64(100); j < 1000; j++ {
			p := i * j
			if p > result && isPalindromeInt6And(p) {
				result = p
			}
		}
	}

	return result
}
