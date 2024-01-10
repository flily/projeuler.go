package p0001

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.Problem{
	Id:    1,
	Title: "Multiples of 3 and 5",
	Description: []string{
		`If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6`,
		`and 9. The sum of these multiples is 23.`,
		``,
		`Find the sum of all the multiples of 3 or 5 below 1000.`,
	},
	Answer: 233168,
	Methods: map[string]framework.Solution{
		"naive": SolveNaive,
	},
}
