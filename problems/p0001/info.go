package p0001

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(1, "Multiples of 3 and 5").
	WithAnswer(233168).
	Solution("naive", SolveNaive).
	Solution("formula", SolveFormula).
	WithDescription(
		`If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6`,
		`and 9. The sum of these multiples is 23.`,
		``,
		`Find the sum of all the multiples of 3 or 5 below 1000.`,
	)
