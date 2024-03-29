package problems

import (
	"fmt"
	"strings"

	"github.com/flily/projeuler.go/framework"

	"github.com/flily/projeuler.go/problems/p0001"
	"github.com/flily/projeuler.go/problems/p0010"
	"github.com/flily/projeuler.go/problems/p0022"
	"github.com/flily/projeuler.go/problems/p0023"
)

var Problems = []framework.Problem{
	p0001.Problem,
	p0010.Problem,
	p0022.Problem,
	p0023.Problem,
}

func init() {
	for _, problem := range Problems {
		_, foundEmpty := problem.Methods[""]
		if foundEmpty {
			err := fmt.Sprintf("empty method name MUST NOT be used, found in problem %d",
				problem.Id)
			panic(err)
		}

		for name := range problem.Methods {
			if strings.Contains(name, " ") {
				err := fmt.Sprintf(
					"method name MUST NOT contain space, found in problem %d, method '%s'",
					problem.Id, name)
				panic(err)
			}
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
