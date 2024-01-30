package p0022

import (
	"testing"
)

func TestNaive(t *testing.T) {
	Problem.Check(t).On(SolveNaive, "naive")
}
