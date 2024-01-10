package p0001

import (
	"testing"
)

func TestNaive(t *testing.T) {
	Problem.Check(t).On(SolveNaive, "naive")
}
