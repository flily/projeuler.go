package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/flily/projeuler.go/framework"
	"github.com/flily/projeuler.go/framework/problems"
)

func runWorker(conf *framework.Configure) {
	worker, err := framework.NewWorker("127.0.0.1", conf.ServePort)
	if err != nil {
		fmt.Printf("start worker failed: %s\n", err)
		return
	}

	worker.Import(problems.Problems)
	go worker.Serve()
	worker.Process()
}

func doClient(conf *framework.Configure) {
	client, err := framework.NewClient("127.0.0.1", conf.ServePort)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	for _, problem := range conf.Problems {
		info, err := framework.ParseProblemId(problem)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}

		methods := make([]string, 0, 1)
		problem, found := problems.GetProblem(info.ProblemId)
		if found && info.Method == "" {
			for method := range problem.Methods {
				methods = append(methods, method)
			}

		} else {
			methods = append(methods, info.Method)
		}

		fmt.Printf("run problem %d\n", info.ProblemId)
		for _, method := range methods {
			result, err := client.Run(info.ProblemId, method)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
				continue
			}

			for _, item := range result.Results {
				fmt.Printf("  %d %s: %s\n", item.ProblemId, item.Method, item.TimeCost)
			}
		}
	}
}

func doRunRaw(conf *framework.Configure) {
	ctx, cancel := framework.NewTimeoutContext(conf.TotalTimeout)
	defer cancel()

	runner := framework.NewRunner()
	runner.Import(problems.Problems)

	infoList, err := framework.ParseProblemIdList(conf.Problems)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var results []*framework.Result

	if len(infoList) > 0 {
		results, err = runner.RunProblemsWithTimeout(ctx, infoList)

	} else {
		results, err = runner.RunAllProblemsWithTimeout(ctx)
	}

	if err != nil {
		log.Printf("run problem solution error: %s\n", err)
		return
	}

	for _, result := range results {
		for _, item := range result.Results {
			fmt.Printf("  %d %s: %s\n", item.ProblemId, item.Method, item.TimeCost)
		}
	}
}

func main() {
	conf := &framework.Configure{}

	flag.BoolVar(&conf.RunnerMode, "runner", true, "run in runner mode")
	flag.DurationVar(&conf.TotalTimeout, "total-timeout", 0, "total timeout, 0 means no timeout")
	flag.DurationVar(&conf.ProblemTimeout, "problem-timeout", 5*time.Second, "problem timeout")
	flag.DurationVar(&conf.MethodTimeout, "method-timeout", 1*time.Second, "method timeout")

	flag.BoolVar(&conf.WorkerMode, "worker", false, "run in worker mode")
	flag.BoolVar(&conf.ClientMode, "client", false, "run in client mode")
	flag.BoolVar(&conf.RawMode, "raw", false, "run in raw mode")
	flag.IntVar(&conf.ServePort, "port", 1707, "server port")
	flag.BoolVar(&conf.DebugMode, "debug", false, "debug mode")

	flag.Parse()

	conf.Problems = flag.Args()

	if conf.WorkerMode {
		runWorker(conf)

	} else if conf.ClientMode {
		doClient(conf)

	} else if conf.RawMode {
		doRunRaw(conf)

	} else if conf.RunnerMode {
		runProblems(conf, problems.Problems)

	} else {
		flag.Usage()
	}
}
