package p0007

import (
	"github.com/flily/projeuler.go/framework"
)

const LIMIT = 10001

var Problem = framework.InitProblem(7, "10001st Prime").
	WithAnswer(104743).
	Solution("count", SolveNaive).
	Solution("array", SolveArray).
	Solution("array-capacity", SolveArrayWithCapacity).
	Solution("array-preallocated", SolveArrayWithPreAllocated).
	WithDescription(
		`By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we`,
		`can see that the 6th prime is 13.`,
		``,
		`What is the 10001st prime number?`,
	)
