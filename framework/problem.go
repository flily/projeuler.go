package framework

import (
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/flily/projeuler.go/framework/message"
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
	ProblemId int
	Method    string
	Result    int64
	IsTimeout bool
	TimeCost  time.Duration
}

func (i *ResultItem) ToMessage() *message.MessageResultItem {
	item := message.NewResultItem(i.ProblemId, i.Method, i.Result, i.TimeCost)
	item.IsTimeout = i.IsTimeout
	return item
}

func (i *ResultItem) FromMessage(message *message.MessageResultItem) {
	i.ProblemId = message.ProblemId
	i.Method = message.Method
	i.Result = message.Result
	i.TimeCost = message.Duration
	i.IsTimeout = message.IsTimeout
}

type Result struct {
	Message string
	Results []ResultItem
}

func NewResult() *Result {
	r := &Result{
		Results: make([]ResultItem, 0),
	}

	return r
}

func (r *Result) Add(item ResultItem) {
	r.Results = append(r.Results, item)
}

func (r *Result) AddTimeoutResult(problemId int, method string, cost time.Duration) {
	item := ResultItem{
		ProblemId: problemId,
		Method:    method,
		IsTimeout: true,
		TimeCost:  cost,
	}

	r.Add(item)
}

func (r *Result) Append(other *Result) {
	if other != nil {
		r.Results = append(r.Results, other.Results...)
	}
}

func (r *Result) Length() int {
	return len(r.Results)
}

func (r *Result) HasTimeoutedResult() bool {
	for _, item := range r.Results {
		if item.IsTimeout {
			return true
		}
	}

	return false
}

func (r *Result) ToMessage() *message.MessageResult {
	result := message.NewResult()

	for _, item := range r.Results {
		itemMessage := item.ToMessage()
		result.AddResult(itemMessage)
	}

	result.Message = r.Message
	return result
}

func (r *Result) FromMessage(message *message.MessageResult) {
	for _, itemMessage := range message.Results {
		item := ResultItem{}
		item.FromMessage(&itemMessage)
		r.Add(item)
	}

	r.Message = message.Message
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
		ProblemId: p.Id,
		Method:    method,
	}

	start := time.Now()
	answer := solution()
	finished := time.Now()
	item.Result = answer
	item.TimeCost = finished.Sub(start)
	return item
}

func (p Problem) MethodList() []string {
	result := make([]string, 0, len(p.Methods))
	for method := range p.Methods {
		result = append(result, method)
	}

	sort.Strings(result)
	return result
}

func (p Problem) RunMethod(method string) *Result {
	item := p.runMethod(method)
	if item == nil {
		return nil
	}

	result := NewResult()
	result.Add(*item)
	return result
}

func (p Problem) RunAll() *Result {
	result := NewResult()
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
