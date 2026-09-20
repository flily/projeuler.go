package p0004

func power10(n int64) int64 {
	result := int64(1)
	for i := int64(0); i < n; i++ {
		result *= 10
	}
	return result
}

func palindromeGenerateGeneric(n int, ch chan<- int64) {
	defer close(ch)

	if n <= 0 {
		ch <- 1
		return
	}

	if n == 1 {
		for i := int64(2); i < 10; i++ {
			ch <- i
		}
		return
	}

	half := int64(n / 2)
	size := int64((n + 1) / 2) // ceiling of n/2
	x := power10(size) - 1
	lower := power10(size - 1)
	for ; x >= lower; x-- {
		y := int64(0)
		for i := int64(0); i < half; i++ {
			di := (x / power10(size-i-1)) % 10
			y += di * power10(i)
		}

		m := x*power10(half) + y
		ch <- m
	}
}

func palindromeGenerate6(ch chan<- int64) {
	defer close(ch)

	for x := int64(999); x >= 100; x-- {
		y := 100_000 * (x % 10)
		y += 10_000 * ((x / 10) % 10)
		y += 1_000 * ((x / 100) % 10)
		n := x + y
		if n > 100_000 {
			ch <- n
		}
	}
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

func SolveGeneratorChannel() int64 {
	result := int64(0)
	ch := make(chan int64)
	go palindromeGenerateGeneric(6, ch)
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

func SolveGenerator6Channel() int64 {
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
