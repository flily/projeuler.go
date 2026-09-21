package main

import (
	"flag"
	"fmt"

	"github.com/flily/projeuler.go/framework"
)

func runClient(port int, problems []string) {
	client, err := framework.NewClient("127.0.0.1", port)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	for _, problem := range problems {
		info, err := framework.ParseProblemId(problem)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			continue
		}

		methods := make([]string, 0, 1)
		problem, found := framework.GetProblem(info.ProblemId)
		if found && info.Method == "" {
			for _, method := range problem.Methods {
				methods = append(methods, method.Name)
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

func doClient(args []string, _ []*framework.Problem) {
	set := flag.NewFlagSet("client", flag.ExitOnError)
	port := set.Int("port", 1707, "server port")
	_ = set.Parse(args)

	problems := set.Args()
	initLogger(false)
	runClient(*port, problems)
}
