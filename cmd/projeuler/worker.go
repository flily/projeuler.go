package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/flily/projeuler.go/framework"
)

func runWorker(port int, allProblems []*framework.Problem) {
	worker, err := framework.NewWorker("127.0.0.1", port)
	if err != nil {
		fmt.Printf("start worker failed: %s\n", err)
		os.Exit(1)
		return
	}

	worker.Import(allProblems)
	go worker.Serve()
	worker.Process()
}

func doWorker(args []string, allProblems []*framework.Problem) {
	flag := flag.NewFlagSet("worker", flag.ExitOnError)
	port := flag.Int("port", 1707, "server port")

	_ = flag.Parse(args)

	initLogger(false)
	runWorker(*port, allProblems)
}
