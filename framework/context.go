package framework

import (
	"context"
	"time"
)

type Context struct {
	TotalTimeout          time.Duration
	TotalTimeoutContext   context.Context
	ProblemTimeout        time.Duration
	ProblemTimeoutContext context.Context
	MethodTimeout         time.Duration
	MethodTimeoutContext  context.Context
}
