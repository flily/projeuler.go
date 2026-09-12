package p0010

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(10, "Summation of primes").
	WithAnswer(142913828922).
	Solution("naive", SolveNaive).
	WithDescription(
		`The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.`,
		``,
		`Find the sum of all the primes below two million.`,
	)
