package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/flily/projeuler.go/framework"
)

func rightPadding(s string, width int, padding string) string {
	if len(s) >= width {
		return s
	}

	paddingLength := width - len(s)
	paddingString := strings.Repeat(padding, paddingLength/len(padding))
	return s + paddingString
}

func toMsString(d time.Duration) string {
	ms := d.Milliseconds()
	ns := d.Nanoseconds() - (ms * 1_000_000)

	msf := float64(ms) + (float64(ns) / 1_000_000.0)
	return fmt.Sprintf("%10.3fms", msf)
}

func makeRunProblemEntryMap(problems []string) (map[int][]string, error) {
	m := make(map[int][]string)
	for _, problem := range problems {
		info, err := framework.ParseProblemId(problem)
		if err != nil {
			return nil, err
		}

		m[info.ProblemId] = append(m[info.ProblemId], info.Method)
	}

	return m, nil
}

func startWorker(conf *framework.Configure) *framework.WorkerProc {
	args := []string{os.Args[0], "-worker", "-port", fmt.Sprintf("%d", conf.ServePort)}
	files := []*os.File{nil, os.Stdout, nil}
	if conf.DebugMode {
		files[2] = os.Stderr
	}

	attrs := &os.ProcAttr{
		Files: files,
	}

	proc, err := os.StartProcess(os.Args[0], args, attrs)
	if err != nil {
		panic(err)
	}

	time.Sleep(100 * time.Millisecond) // wait for worker to start
	log.Printf("start background worker pid=%d", proc.Pid)
	return framework.NewWorkerProc(proc)
}

func initConnection(conf *framework.Configure) (*framework.WorkerProc, *framework.Client) {
	worker := startWorker(conf)

	time.Sleep(100 * time.Millisecond) // wait for worker to start
	client, err := conf.NewClient("127.0.0.1")
	if err != nil {
		log.Printf("create client failed: %s\n", err)
		panic(err)
	}

	client.SetTimeout(conf.ProblemTimeout, conf.MethodTimeout)
	return worker, client
}

func runProblems(conf *framework.Configure, allProblems []framework.Problem) {
	problemEntry, err := makeRunProblemEntryMap(conf.Problems)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	worker, client := initConnection(conf)
	defer func() {
		worker.Kill()
	}()

	for _, problem := range allProblems {
		methods, found := problemEntry[problem.Id]
		if len(problemEntry) > 0 && !found {
			continue
		}

		if methods == nil {
			methods = problem.MethodList()
		}

		finalResult := framework.NewResult()
		for _, method := range methods {
			resultSet, err := client.Run(problem.Id, method)
			if err != nil {
				fmt.Printf("Run problem %d %s error: %s\n", problem.Id, method, err)
				return
			}

			if resultSet.HasTimeoutedResult() {
				client.Close()
				worker.Kill()
				time.Sleep(100 * time.Millisecond)
				worker, client = initConnection(conf)
			}

			finalResult.Append(resultSet)
		}

		printResult(conf, problem, finalResult)
	}
}

func printResult(conf *framework.Configure, problem framework.Problem, result *framework.Result) {
	if result.Length() == 1 {
		item := result.Results[0]
		fmt.Printf("%-5d %-40s %-15d %-15s\n",
			problem.Id, rightPadding(problem.Title, 40, "."), item.Result, toMsString(item.TimeCost))

	} else {
		fmt.Printf("%-5d %-40s\n", problem.Id, rightPadding(problem.Title, 40, "."))
		for _, item := range result.Results {
			fmt.Printf("      + %-38s %-15d %-15s\n", rightPadding(item.Method, 38, "."), item.Result, toMsString(item.TimeCost))
		}
	}
}
