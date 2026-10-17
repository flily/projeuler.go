package p0020

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(20, "Factorial Digit Sum").
	WithAnswer(648).
	Solution("naive (bigint)", SolveNaive).
	WithDescription(
		`n! means n × (n − 1) × ... × 3 × 2 × 1.`,
		``,
		`For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800, and the sum`,
		`of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.`,
		``,
		`Find the sum of the digits in 100!.`,
	)
