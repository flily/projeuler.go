package p0003

import (
	"github.com/flily/projeuler.go/framework"
)

const NUMBER = 600851475143

var Problem = framework.InitProblem(3, "Largest Prime Factor").
	WithAnswer(6857).
	Solution("naive", SolveNaive).
	Solution("remove-factor", SolveRemoveFactor).
	WithDescription(
		`The prime factors of 13195 are 5, 7, 13 and 29.`,
		``,
		`What is the largest prime factor of the number 600851475143?`,
	)
