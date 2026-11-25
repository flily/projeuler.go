package p0031

var coins = []int64{
	200, 100, 50, 20, 10, 5, 2, 1,
}

func findCoinWays(target int64, coins []int64, index int, coinMax []int) int64 {
	if index >= len(coins) {
		if target == 0 {
			return 1
		}
		return 0
	}

	result := int64(0)
	for i := 0; i <= coinMax[index]; i++ {
		change := coins[index] * int64(i)
		result += findCoinWays(target-change, coins, index+1, coinMax)
	}

	return result
}

// SolveNaive is checked correct, use about 8s
func SolveNaive() int64 {
	coinMax := make([]int, len(coins))
	for i := range coinMax {
		coinMax[i] = int(200 / coins[i])
	}

	return findCoinWays(200, coins, 0, coinMax)
}
