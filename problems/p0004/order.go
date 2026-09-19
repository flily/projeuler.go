package p0004

func power10(n int64) int64 {
	result := int64(1)
	for i := int64(0); i < n; i++ {
		result *= 10
	}
	return result
}

func palindromeGenerate6(ch chan<- int64) {
	for x := int64(999); x >= 100; x-- {
		y := 100_000 * (x % 10)
		y += 10_000 * ((x / 10) % 10)
		y += 1_000 * ((x / 100) % 10)
		n := x + y
		if n > 100_000 {
			ch <- n
		}
	}

	close(ch)
}

func checkPalindromeProduct(n int64) int64 {
	for i := int64(100); i < 1000; i++ {
		if (n % i) != 0 {
			continue
		}

		j := n / i
		if j >= 100 && j < 1000 {
			return n
		}
	}

	return 0
}

func SolveGenerator() int64 {
	result := int64(0)
	ch := make(chan int64)
	go palindromeGenerate6(ch)
	for {
		p, ok := <-ch
		if !ok {
			break
		}

		if p > result && checkPalindromeProduct(p) != 0 {
			result = p
		}
	}

	return result
}
