package framework

import (
	"fmt"
)

var (
	ErrNoSuchProblem  = fmt.Errorf("no such problem")
	ErrNoSuchSolution = fmt.Errorf("no such solution")
)
