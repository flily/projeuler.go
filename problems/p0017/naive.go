package p0017

import "strings"

var NUMBERS = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
}

func numberInEnglish1000(n int) string {
	if n == 1000 {
		return "one thousand"
	}

	parts := make([]string, 0, 8)
	if n >= 100 {
		hundreds := n / 100
		n = n % 100
		parts = append(parts, NUMBERS[hundreds], NUMBERS[100])
	}

	if n == 0 {
		return strings.Join(parts, " ")
	}

	if len(parts) > 0 {
		parts = append(parts, "and")
	}

	if n >= 20 {
		tens := n / 10 * 10
		n = n % 10
		if n == 0 {
			parts = append(parts, NUMBERS[tens])
		} else {
			parts = append(parts, NUMBERS[tens], NUMBERS[n])
		}

	} else if n > 0 {
		parts = append(parts, NUMBERS[n])
	}

	return strings.Join(parts, " ")
}

func SolveNaive() int64 {
	result := int64(0)
	for i := 1; i <= 1000; i++ {
		eng := numberInEnglish1000(i)
		result += int64(len(strings.ReplaceAll(eng, " ", "")))
	}

	return result
}
