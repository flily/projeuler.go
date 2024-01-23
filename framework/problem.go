package framework

import (
	"strings"
	"testing"
	"time"
)

type Solution func() int64

type Answer int64

func (a Answer) Test(t *testing.T) TestContext {
	ctx := TestContext{
		t:        t,
		answer:   a,
		noAnswer: false,
	}

	return ctx
}

func (a Answer) Equals(b int64) bool {
	return int64(a) == b
}

type TestContext struct {
	t        *testing.T
	answer   Answer
	noAnswer bool
}

func (c TestContext) On(solution Solution, name string) {
	got := solution()
	if c.noAnswer {
		c.t.Logf("method '%s': %d", name, got)

	} else if !c.answer.Equals(got) {
		c.t.Errorf("Got wrong answer '%d' of method '%s', expect %d", got, name, c.answer)
	}
}

type ResultItem struct {
	Method    string
	Result    int64
	IsTimeout bool
	TimeCost  time.Duration
}

type Result struct {
	ProblemId int
	Results   []ResultItem
}

func NewResult(problemId int) *Result {
	r := &Result{
		ProblemId: problemId,
		Results:   make([]ResultItem, 0),
	}

	return r
}

func (r *Result) Add(item ResultItem) {
	r.Results = append(r.Results, item)
}

type Problem struct {
	Id          int
	Title       string
	Description []string
	Answer      Answer
	Methods     map[string]Solution
	NoAnswer    bool
}

func (p Problem) GetDescription() string {
	return strings.Join(p.Description, "\n")
}

func (p Problem) runMethod(method string) *ResultItem {
	solution, found := p.Methods[method]
	if !found {
		return nil
	}

	item := &ResultItem{
		Method: method,
	}

	start := time.Now()
	answer := solution()
	finished := time.Now()
	item.Result = answer
	item.TimeCost = finished.Sub(start)
	return item
}

func (p Problem) RunMethod(method string) *Result {
	item := p.runMethod(method)
	if item == nil {
		return nil
	}

	result := NewResult(p.Id)
	result.Add(*item)
	return result
}

func (p Problem) RunAll() *Result {
	result := NewResult(p.Id)
	for method := range p.Methods {
		item := p.runMethod(method)
		if item != nil {
			result.Add(*item)
		}
	}

	return result
}

func (p Problem) Check(t *testing.T) TestContext {
	ctx := TestContext{
		t:        t,
		answer:   p.Answer,
		noAnswer: p.NoAnswer,
	}

	return ctx
}
