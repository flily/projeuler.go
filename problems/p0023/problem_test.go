package p0023

import (
	"testing"
)

func TestNaive(t *testing.T) {
	Problem.Check(t).On(SolveNaive, "naive")
}

func TestSolveWithFactorSumCache(t *testing.T) {
	Problem.Check(t).On(SolveWithFactorSumCache, "with factor sum cache")
}

func TestSolveWithSubstraction(t *testing.T) {
	Problem.Check(t).On(SolveWithSubstraction, "with substraction")
}
