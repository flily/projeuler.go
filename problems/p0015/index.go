package p0015

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(15, "Lattice Paths").
	WithAnswer(137846528820).
	Solution("naive", SolveNaive).
	WithDescription(
		`Starting in the top left corner of a 2 x 2 grid, and only being able to move to the right`,
		`and down, there are exactly 6 routes to the bottom right corner.`,
		``,
		`    * == * == *    * == * -- *    * == * -- *`,
		`    |    |    ||   |    ||   |    |    ||   |`,
		`    * -- * -- *    * -- * == *    * -- * -- *`,
		`    |    |    ||   |    |    ||   |    ||   |`,
		`    * -- * -- *    * -- * -- *    * -- * == *`,
		``,
		`    * -- * -- *    * -- * -- *    * -- * -- *`,
		`    ||   |    |    ||   |    |    ||   |    |`,
		`    * == * == *    * == * -- *    * -- * -- *`,
		`    |    |    ||   |    ||   |    ||   |    |`,
		`    * -- * -- *    * -- * == *    * == * == *`,
		``,
		`How many such routes are there through a 20 x 20 grid?`,
	)
