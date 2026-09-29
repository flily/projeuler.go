package p0009

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(9, "Special Pythagorean Triplet").
	WithAnswer(31875000).
	Solution("naive", SolveNaive).
	WithDescription(
		`Description of the problem.`,
	)
