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

func (r *Runner) RunProblem(info ProblemRunInfo) (*Result, error) {
	problem, found := r.Index[info.ProblemId]
	if !found {
		return nil, fmt.Errorf("no problem %d", info.ProblemId)
	}

	if info.IsAllMethods() {
		return problem.RunAll(), nil
	} else {
		return problem.RunMethod(info.Method), nil
	}
}

func (r *Runner) runProblemWrap(ch chan<- resultPackage, info ProblemRunInfo) {
	result, err := r.RunProblem(info)
	ch <- resultPackage{
		Result: result,
		Err:    err,
	}
}

func (r *Runner) RunProblemWithTimeout(ctx context.Context, info ProblemRunInfo) (*Result, error) {
	ch := make(chan resultPackage)
	go r.runProblemWrap(ch, info)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case result := <-ch:
		return result.Result, result.Err
	}
}

func (r *Runner) RunProblemsWithTimeout(ctx context.Context, problems []ProblemRunInfo) ([]*Result, error) {
	results := make([]*Result, 0, len(problems))
	for _, info := range problems {
		result, err := r.RunProblemWithTimeout(ctx, info)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *Runner) RunAllProblemsWithTimeout(ctx context.Context) ([]*Result, error) {
	results := make([]*Result, 0, len(r.Problems))
	for _, p := range r.Problems {
		info := NewProblemRunInfo(p.Id, "")
		result, err := r.RunProblemWithTimeout(ctx, info)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}
