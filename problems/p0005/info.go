package p0005

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(5, "Smallest Multiple").
	WithAnswer(232792560).
	Solution("naive", SolveNaive).
	Solution("naive odds", SolveNaiveOdds).
	Solution("log factor", SolveLogFactor).
	WithDescription(
		`2520 is the smallest number that can be divided by each of the numbers`,
		`from 1 to 10 without any`,
		`remainder.`,
		``,
		`What is the smallest positive number that is evenly divisible by all`,
		`of the numbers from 1 to 20?`,
	)
