package problems

import (
	"github.com/flily/projeuler.go/framework"

	"github.com/flily/projeuler.go/problems/p0001"
	"github.com/flily/projeuler.go/problems/p0010"
	"github.com/flily/projeuler.go/problems/p0023"
)

var Problems = []framework.Problem{
	p0001.Problem,
	p0010.Problem,
	p0023.Problem,
}

func init() {
	for _, problem := range Problems {
		_, foundEmpty := problem.Methods[""]
		if foundEmpty {
			panic("empty method name")
		}
	}
}

func GetProblem(id int) (framework.Problem, bool) {
	for _, problem := range Problems {
		if problem.Id == id {
			return problem, true
		}
	}

	return framework.Problem{}, false
}
