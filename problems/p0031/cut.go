package p0031

func findCoinWaysCut(target int64, coins []int64, index int, coinMax []int) int64 {
	if index >= len(coins) {
		if target == 0 {
			return 1
		}
		return 0
	}

	result := int64(0)
	for i := 0; i <= coinMax[index]; i++ {
		change := coins[index] * int64(i)
		if change > target {
			break
		}

		result += findCoinWaysCut(target-change, coins, index+1, coinMax)
	}

	return result
}

func SolveBranchCut() int64 {
	coinMax := make([]int, len(coins))
	for i := range coinMax {
		coinMax[i] = int(200 / coins[i])
	}

	return findCoinWaysCut(200, coins, 0, coinMax)
}
