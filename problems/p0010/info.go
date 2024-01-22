package p0010

import "github.com/flily/projeuler.go/framework"

var Problem = framework.Problem{
	Id:    10,
	Title: "Summation of primes",
	Description: []string{
		`The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.`,
		``,
		`Find the sum of all the primes below two million.`,
	},
	Answer: 142913828922,
	Methods: map[string]framework.Solution{
		"naive": SolveNaive,
	},
}
