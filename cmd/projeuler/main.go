package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/flily/projeuler.go/framework"
	"github.com/flily/projeuler.go/framework/problems"
)

func doRunner(conf *framework.Configure) {
	ctx, cancel := framework.NewTimeoutContext(conf.TotalTimeout)
	defer cancel()

	runner := framework.NewRunner()
	runner.Import(problems.Problems)

	for _, problemId := range conf.Problems {
		id, method, err := framework.ParseProblemId(problemId)
		if err != nil {
			fmt.Printf("invalid problem id: '%s'\n", problemId)
			continue
		}

		result, err := runner.RunProblemWithTimeout(ctx, id, method)
		if err != nil {
			fmt.Printf("run problem solution error: %s\n", err)
			continue
		}

		fmt.Printf("problem %d\n", id)
		for _, item := range result.Items {
			fmt.Printf("  %s: %s\n", item.Method, item.TimeCost)
		}
	}
}

func main() {
	conf := &framework.Configure{}

	flag.BoolVar(&conf.RunnerMode, "runner", true, "run in runner mode")
	flag.DurationVar(&conf.TotalTimeout, "total-timeout", 0, "total timeout, 0 means no timeout")
	flag.DurationVar(&conf.ProblemTimeout, "problem-timeout", 5*time.Second, "problem timeout")
	flag.DurationVar(&conf.MethodTimeout, "method-timeout", 1*time.Second, "method timeout")

	flag.BoolVar(&conf.ServerMode, "server", false, "run in server mode")
	flag.IntVar(&conf.ServerPort, "port", 1707, "server port")

	flag.BoolVar(&conf.WorkerMode, "worker", false, "run in worker mode")
	flag.Parse()

	conf.Problems = flag.Args()

	if conf.RunnerMode {
		doRunner(conf)

	} else {
		flag.Usage()
	}
}
