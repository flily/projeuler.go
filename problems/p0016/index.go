package p0016

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(16, "Power Digit Sum").
	WithAnswer(1366).
	Solution("bigint (as integer)", SolveBigIntInteger).
	Solution("bigint (as string)", SolveBigintString).
	WithDescription(
		`2 ^ 15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.`,
		``,
		`What is the sum of the digits of the number 2 ^ 1000?`,
	)
