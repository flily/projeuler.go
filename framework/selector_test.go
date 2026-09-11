package framework

import (
	"slices"
	"testing"
)

func TestParseSelectorProblemIdOnly(t *testing.T) {
	input := "42"
	expected := &Selector{
		Problem: SelectorInfo{
			String: "42",
			Int:    new(42),
		},
		Solutions: nil,
	}

	got, err := ParseSelector(input)
	if err != nil {
		t.Fatalf("unexpected error\n%s", err)
	}

	if !got.Equal(expected) {
		t.Errorf("got wrong result")
		t.Errorf("expect: %v", expected)
		t.Errorf("got   : %v", got)
	}

	p1 := NewProblem(42, "title").
		Solution("lorem", nil).
		Solution("ipsum", nil)
	if !got.Match(p1) {
		t.Fatalf("failed to match problem ID")
	}

	expS1 := []SolutionSelector{
		NewToRunSolution("lorem"),
		NewToRunSolution("ipsum"),
	}
	gotS1, _ := got.MatchedSolutions(p1)
	if !slices.Equal(gotS1, expS1) {
		t.Errorf("got wrong matched solutions")
		t.Errorf("expect: %v", expS1)
		t.Fatalf("got   : %v", gotS1)
	}

	p2 := NewProblem(1, "the answer of life, the universe, everything is 42")
	if !got.Match(p2) {
		t.Fatalf("failed to match problem title pattern")
	}
}

func TestParseSelectorProblemTitlePatternOnly(t *testing.T) {
	input := "lorem"
	expected := &Selector{
		Problem: SelectorInfo{
			String: "lorem",
			Int:    nil,
		},
		Solutions: nil,
	}

	got, err := ParseSelector(input)
	if err != nil {
		t.Fatalf("unexpected error\n%s", err)
	}

	if !got.Equal(expected) {
		t.Errorf("got wrong result")
		t.Errorf("expect: %v", expected)
		t.Errorf("got   : %v", got)
	}

	p1 := NewProblem(1, "lorem ipsum")
	if !got.Match(p1) {
		t.Fatalf("failed to match problem title pattern")
	}
}

func TestParseSelectorProblemIdWithDot(t *testing.T) {
	input := "42."
	expected := &Selector{
		Problem: SelectorInfo{
			String: "42",
			Int:    new(42),
		},
		Solutions: nil,
	}

	got, err := ParseSelector(input)
	if err != nil {
		t.Fatalf("unexpected error\n%s", err)
	}

	if !got.Equal(expected) {
		t.Errorf("got wrong result")
		t.Errorf("expect: %v", expected)
		t.Errorf("got   : %v", got)
	}

	p1 := NewProblem(42, "title").
		Solution("lorem", nil).
		Solution("ipsum", nil)
	if !got.Match(p1) {
		t.Fatalf("failed to match problem ID")
	}

	expS1 := []SolutionSelector{
		NewToRunSolution("lorem"),
		NewToRunSolution("ipsum"),
	}
	gotS1, _ := got.MatchedSolutions(p1)
	if !slices.Equal(gotS1, expS1) {
		t.Errorf("got wrong matched solutions")
		t.Errorf("expect: %v", expS1)
		t.Fatalf("got   : %v", gotS1)
	}

	p2 := NewProblem(1, "the answer of life, the universe, everything is 42")
	if !got.Match(p2) {
		t.Fatalf("failed to match problem title pattern")
	}
}

func TestParseSelectorWithOnlySolutionNoProblemId(t *testing.T) {
	input := ".m"
	expected := &Selector{
		Problem: SelectorInfo{
			String: "",
			Int:    nil,
		},
		Solutions: []*SelectorInfo{
			NewSelectorInfo("m"),
		},
	}

	got, err := ParseSelector(input)
	if err != nil {
		t.Fatalf("unexpected error\n%s", err)
	}

	if !got.Equal(expected) {
		t.Errorf("got wrong result")
		t.Errorf("expect: %v", expected)
		t.Errorf("got   : %v", got)
	}

	p1 := NewProblem(42, "example").
		Solution("lorem", nil).
		Solution("ipsum", nil).
		Solution("dolor", nil).
		Solution("sit", nil).
		Solution("amet", nil)

	gotS1, _ := got.MatchedSolutions(p1)
	expS1 := []SolutionSelector{
		NewToRunSolution("lorem"),
		NewToRunSolution("ipsum"),
		NewToSkipSolution("dolor"),
		NewToSkipSolution("sit"),
		NewToRunSolution("amet"),
	}
	if !slices.Equal(gotS1, expS1) {
		t.Errorf("got wrong matched solutions")
		t.Errorf("expect: %v", expS1)
		t.Fatalf("got   : %v", gotS1)
	}

	if !got.Match(p1) {
		t.Fatalf("failed to match problem with only solution selector")
	}

	p2 := NewProblem(1, "another example").
		Solution("aaaa", nil).
		Solution("bbbb", nil).
		Solution("cccc", nil)

	gotS2, _ := got.MatchedSolutions(p2)
	expS2 := []SolutionSelector{
		NewToSkipSolution("aaaa"),
		NewToSkipSolution("bbbb"),
		NewToSkipSolution("cccc"),
	}
	if !slices.Equal(gotS2, expS2) {
		t.Errorf("got wrong matched solutions")
		t.Errorf("expect: %v", expS2)
		t.Fatalf("got   : %v", gotS2)
	}

	if got.Match(p2) {
		t.Fatalf("failed to match problem with only solution selector")
	}
}
