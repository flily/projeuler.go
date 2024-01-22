package p0010

func SolveNaive() int64 {
	sum := int64(2 + 3 + 5 + 7 + 11 + 13 + 17 + 19)
	basic_primes := []int64{3, 5, 7, 11, 13, 17, 19}
	primes_list := make([]int64, 0, 230)
	primes_set := make(map[int64]bool)
	for _, p := range basic_primes {
		primes_list = append(primes_list, p)
		primes_set[p] = true
	}

	for i := int64(21); i < 2000000; i += 2 {
		isPrime := true
		for _, p := range primes_list {
			if i%p == 0 {
				isPrime = false
				break
			}

			if p*p >= i {
				break
			}
		}

		if isPrime {
			primes_list = append(primes_list, i)
			primes_set[i] = true
			sum += i
		}
	}

	return sum
}
