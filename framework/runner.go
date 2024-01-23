package framework

import (
	"context"
	"fmt"
	"time"
)

func emptyCancel() {}

func NewTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	background := context.Background()
	if timeout <= 0 {
		return background, emptyCancel
	}

	return context.WithTimeout(background, timeout)
}

type resultPackage struct {
	Result *Result
	Err    error
}

type Runner struct {
	Problems []Problem
	Index    map[int]Problem
	Pipe     chan Result
}

func NewRunner() *Runner {
	r := &Runner{
		Index: make(map[int]Problem),
		Pipe:  make(chan Result),
	}
	return r
}

func (r *Runner) Add(p Problem) {
	r.Problems = append(r.Problems, p)
	r.Index[p.Id] = p
}

func (r *Runner) Import(problems []Problem) {
	for _, p := range problems {
		r.Add(p)
	}
}

func (r *Runner) RunProblem(id int, method string) (*Result, error) {
	problem, found := r.Index[id]
	if !found {
		return nil, fmt.Errorf("no problem %d", id)
	}

	if method == "" {
		return problem.RunAll(), nil
	} else {
		return problem.RunMethod(method), nil
	}
}

func (r *Runner) runProblemWrap(ch chan<- resultPackage, id int, method string) {
	result, err := r.RunProblem(id, method)
	ch <- resultPackage{
		Result: result,
		Err:    err,
	}
}

func (r *Runner) RunProblemWithTimeout(ctx context.Context, id int, method string) (*Result, error) {
	ch := make(chan resultPackage)
	go r.runProblemWrap(ch, id, method)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case result := <-ch:
		return result.Result, result.Err
	}
}
