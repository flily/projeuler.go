package framework

import (
	"strings"
	"testing"
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

func (p Problem) Check(t *testing.T) TestContext {
	ctx := TestContext{
		t:        t,
		answer:   p.Answer,
		noAnswer: p.NoAnswer,
	}

	return ctx
}
