package p0092

func DigitSquareSum(n int64) int64 {
	sum := int64(0)
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return sum
}
