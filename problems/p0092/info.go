package p0092

import (
	"github.com/flily/projeuler.go/framework"
)

const (
	Limit = 10_000_000
)

var Problem = framework.NewProblem(92, "Square Digit Chains").
	WithAnswer(8581146).
	Solution("set", SolveSet).
	Solution("set-optimized", SolveSetOptimized).
	Solution("vector", SolveVector).
	Solution("vector-reduced", SolveVectorReduced).
	WithDescription(
		`A number chain is created by continuously adding the square of the digits in a number to form a new`,
		`number until it has been seen before.`,

		`For example,`,
		`    44 -> 32 -> 13 -> 10 -> [1] -> [1]`,
		`    85 -> [89] -> 145 -> 42 -> 20 -> 4 -> 16 -> 37 -> 58 -> [89]`,

		`Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. What is most`,
		`amazing is that EVERY starting number will eventually arrive at 1 or 89.`,

		`How many starting numbers below ten million will arrive at 89?`,
	)
