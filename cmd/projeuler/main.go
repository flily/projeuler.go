package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/flily/projeuler.go/framework"
	"github.com/flily/projeuler.go/framework/problems"
)

func doServer(conf *framework.Configure) {
	server, err := framework.NewWorker("127.0.0.1", conf.ServerPort)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return
	}

	server.Import(problems.Problems)
	go server.Serve()
	server.Process()
}

func doClient(conf *framework.Configure) {
	client, err := framework.NewClient("127.0.0.1", conf.ServerPort)
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

		result, err := client.Run(info.ProblemId, info.Method)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}

		for _, item := range result.Results {
			fmt.Printf("  %d %s: %s\n", item.ProblemId, item.Method, item.TimeCost)
		}
	}
}

func doRunner(conf *framework.Configure) {
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

	flag.BoolVar(&conf.ServerMode, "server", false, "run in server mode")
	flag.BoolVar(&conf.ClientMode, "client", false, "run in client mode")
	flag.IntVar(&conf.ServerPort, "port", 1707, "server port")

	flag.BoolVar(&conf.WorkerMode, "worker", false, "run in worker mode")
	flag.Parse()

	conf.Problems = flag.Args()

	if conf.ServerMode {
		doServer(conf)

	} else if conf.ClientMode {
		doClient(conf)
		return
	}

	if conf.RunnerMode {
		doRunner(conf)

	} else {
		flag.Usage()
	}
}
