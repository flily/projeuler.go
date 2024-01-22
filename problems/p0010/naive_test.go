package p0010

import (
	"testing"
)

func TestNaive(t *testing.T) {
	Problem.Check(t).On(SolveNaive, "naive")
}
