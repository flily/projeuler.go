package framework

import (
	"time"
)

type Configure struct {
	RunnerMode     bool
	TotalTimeout   time.Duration
	ServerMode     bool
	ServerPort     int
	ProblemTimeout time.Duration
	MethodTimeout  time.Duration
	WorkerMode     bool
	Problems       []string
}
