package p0030

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(30, "Digit Fifth Powers").
	WithAnswer(443839).
	Solution("math", SolveMath).
	Solution("math inline", SolveMathInline).
	Solution("string", SolveString).
	Solution("string inline", SolveStringInline).
	WithDescription(
		`Surprisingly there are only three numbers that can be written as the`,
		`sum of fourth powers of their digits:`,
		`        1634 = 1^4 + 6^4 + 3^4 + 4^4`,
		`        8208 = 8^4 + 2^4 + 0^4 + 8^4`,
		`        9474 = 9^4 + 4^4 + 7^4 + 4^4`,
		``,
		`As 1 = 1^4 is not a sum it is not included.`,
		`The sum of these numbers is 1634 + 8208 + 9474 = 19316.`,
		``,
		`Find the sum of all the numbers that can be written as the sum of`,
		`fifth powers of their digits.`,
		`        1634 = 1^4 + 6^4 + 3^4 + 4^4`,
		`        8208 = 8^4 + 2^4 + 0^4 + 8^4`,
		`        9474 = 9^4 + 4^4 + 7^4 + 4^4`,
		``,
		`As 1 = 1^4 is not a sum it is not included.`,
		`The sum of these numbers is 1634 + 8208 + 9474 = 19316.`,
		``,
		`Find the sum of all the numbers that can be written as the sum of`,
		`fifth powers of their digits.`,
	)

// 6 * 9^5 is the upper bound for the sum of fifth powers of digits
const MAX_NUMBER = 6 * 9 * 9 * 9 * 9 * 9
