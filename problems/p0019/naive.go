package p0019

var MONTH_DAYS = [13]int{
	0,  // Placeholder for 0th month
	31, // January
	28, // February
	31, // March
	30, // April
	31, // May
	30, // June
	31, // July
	31, // August
	30, // September
	31, // October
	30, // November
	31, // December
}

func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	if (year % 100) == 0 {
		return (year % 400) == 0
	}

	return true
}

func SolveNaive() int64 {
	result := 0
	days := 365 // count from 1901 Jan 1

	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			// 1900-Jan-01 is Monday, so Sunday is 6 days after
			if days%7 == 6 {
				result += 1
			}

			days += MONTH_DAYS[month]
			if month == 2 && isLeapYear(year) {
				days += 1
			}
		}
	}

	return int64(result)
}
