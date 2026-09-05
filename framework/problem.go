package framework

import (
	"fmt"
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
	problem  *Problem
	t        *testing.T
	answer   Answer
	noAnswer bool
}

func (c TestContext) on(t *testing.T, solution Solution, name string) {
	got := solution()
	if c.noAnswer {
		t.Logf("method '%s': %d", name, got)

	} else if !c.answer.Equals(got) {
		t.Errorf("Got wrong answer '%d' of method '%s', expect %d", got, name, c.answer)
	}
}

func (c TestContext) On(solution Solution, name string) {
	c.on(c.t, solution, name)
}

func (c TestContext) All() {
	for name, entry := range c.problem.Methods {
		testName := fmt.Sprintf("TestSolutionP%04d.%s", c.problem.Id, name)
		c.t.Run(testName, func(tt *testing.T) {
			c.on(tt, entry, name)
		})
	}
}

type FinalResult int

const (
	FinalResultNone    FinalResult = iota // not run yet
	FinalResultUnknown                    // run but not checked
	FinalResultSkipped                    // skipped by user
	FinalResultCorrect
	FinalResultWrong
	FinalResultTimeout
	FinalResultCrash
)

var finalResultStrings = map[FinalResult]string{
	FinalResultNone:    "-",
	FinalResultUnknown: "?",
	FinalResultSkipped: "skipped",
	FinalResultCorrect: "correct",
	FinalResultWrong:   "wrong",
	FinalResultTimeout: "timeout",
	FinalResultCrash:   "crash",
}

func (r FinalResult) String() string {
	if s, ok := finalResultStrings[r]; ok {
		return s
	}

	return "unknown"
}

func (r FinalResult) Style() DisplayStyle {
	s := DefaultDisplayStyle()
	switch r {
	case FinalResultCorrect:
		s = s.Colour(ColourGreen)

	case FinalResultUnknown, FinalResultSkipped, FinalResultTimeout:
		s = s.Colour(ColourYellow)

	case FinalResultWrong, FinalResultCrash:
		s = s.Colour(ColourRed)
	}

	return s
}

type ResultItem struct {
	ProblemId int
	Method    string
	Result    FinalResult
	Answer    int64
	TimeCost  time.Duration
}

func (i *ResultItem) ToMessage() *message.MessageResultItem {
	item := message.NewResultItem(i.ProblemId, i.Method, i.Answer, i.TimeCost)
	item.IsTimeout = i.Result == FinalResultTimeout
	return item
}

func (i *ResultItem) FromMessage(message *message.MessageResultItem) {
	i.ProblemId = message.ProblemId
	i.Method = message.Method
	i.Answer = message.Result
	i.TimeCost = message.Duration

	if message.IsTimeout {
		i.Result = FinalResultTimeout

	} else {
		i.Result = FinalResultUnknown
	}
}

func (i *ResultItem) Check(answer Answer) FinalResult {
	if i.Result == FinalResultUnknown {
		if i.Answer == int64(answer) {
			i.Result = FinalResultCorrect
		} else {
			i.Result = FinalResultWrong
		}
	}

	return i.Result
}

type Result struct {
	Message string
	Results []*ResultItem
}

func NewResult() *Result {
	r := &Result{
		Results: make([]*ResultItem, 0),
	}

	return r
}

func (r *Result) Add(item *ResultItem) {
	copy := *item
	r.Results = append(r.Results, &copy)
}

func (r *Result) AddTimeoutResult(problemId int, method string, cost time.Duration) {
	item := &ResultItem{
		ProblemId: problemId,
		Method:    method,
		Result:    FinalResultTimeout,
		TimeCost:  cost,
	}

	r.Add(item)
}

func (r *Result) CheckResult(answer Answer) (int, int) {
	countCorrect, countTotal := 0, 0
	for _, item := range r.Results {
		if item.Check(answer) == FinalResultCorrect {
			countCorrect++
		}
		countTotal++
	}

	return countCorrect, countTotal
}

func (r *Result) GetProblemResult() (FinalResult, int) {
	result := FinalResultUnknown
	bestIndex := -1
	bestTime := time.Duration(0)

	for index, item := range r.Results {
		stop := false

		switch item.Result {
		case FinalResultCorrect:
			result = FinalResultCorrect
			if bestIndex < 0 || item.TimeCost < bestTime {
				bestIndex = index
				bestTime = item.TimeCost
			}

		case FinalResultWrong, FinalResultCrash:
			result = FinalResultWrong
			stop = true
		}

		if stop {
			break
		}
	}

	return result, bestIndex
}

func (r *Result) IsCorrect(answer Answer) bool {
	result := false
	for _, item := range r.Results {
		if item.Result == FinalResultUnknown {
			if answer.Equals(item.Answer) {
				result = true
				break
			}
		}
	}

	return result
}

func (r *Result) TotalCost() time.Duration {
	var total time.Duration
	for _, item := range r.Results {
		total += item.TimeCost
	}

	return total
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
		if item.Result == FinalResultTimeout {
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
		r.Add(&item)
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
		Result:    FinalResultNone,
	}

	start := time.Now()
	answer := solution()
	finished := time.Now()
	item.Answer = answer
	item.TimeCost = finished.Sub(start)
	item.Result = FinalResultUnknown

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
	result.Add(item)
	return result
}

func (p Problem) RunAll() *Result {
	result := NewResult()
	for method := range p.Methods {
		item := p.runMethod(method)
		if item != nil {
			result.Add(item)
		}
	}

	return result
}

func (p Problem) Check(t *testing.T) TestContext {
	ctx := TestContext{
		problem:  &p,
		t:        t,
		answer:   p.Answer,
		noAnswer: p.NoAnswer,
	}

	return ctx
}
