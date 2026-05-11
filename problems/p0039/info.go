package p0039

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.Problem{
	Id:    39,
	Title: "Integer right triangles",
	Description: []string{
		`If p is the perimeter of a right angle triangle with integral length sides, {a,b,c},`,
		`there are exactly three solutions for p = 120. {20,48,52}, {24,45,51}, {30,40,50}`,
		``,
		`For which value of p ≤ 1000, is the number of solutions maximised?`,
	},
	Answer: 840,
	Methods: map[string]framework.Solution{
		"naive":   SolveNaive,
		"ordered": SolveOrdered,
	},
}
