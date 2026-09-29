package p0007

func SolveArray() int64 {
	primes := make([]int64, 0)
	primes = append(primes, 3, 5, 7, 11, 13, 17, 19)

	for n := int64(21); len(primes) < LIMIT-1; n += 2 {
		isPrime := true
		for _, p := range primes {
			if (n % p) == 0 {
				isPrime = false
				break
			}

			if (p * p) > n {
				break
			}
		}

		if isPrime {
			primes = append(primes, n)
		}
	}

	return primes[len(primes)-1]
}

func SolveArrayWithCapacity() int64 {
	primes := make([]int64, 0, LIMIT)
	primes = append(primes, 3, 5, 7, 11, 13, 17, 19)

	for n := int64(21); len(primes) < LIMIT-1; n += 2 {
		isPrime := true
		for _, p := range primes {
			if (n % p) == 0 {
				isPrime = false
				break
			}

			if (p * p) > n {
				break
			}
		}

		if isPrime {
			primes = append(primes, n)
		}
	}

	return primes[len(primes)-1]
}

func SolveArrayWithPreAllocated() int64 {
	primes := make([]int64, LIMIT)
	primes[0] = 3
	primes[1] = 5
	primes[2] = 7
	primes[3] = 11
	primes[4] = 13
	primes[5] = 17
	primes[6] = 19

	i := 7
	for n := int64(21); i < LIMIT-1; n += 2 {
		isPrime := true
		for j, p := range primes {
			if j >= i {
				break
			}

			if (n % p) == 0 {
				isPrime = false
				break
			}

			if (p * p) > n {
				break
			}
		}

		if isPrime {
			primes[i] = n
			i += 1
		}
	}

	return primes[i-1]
}
