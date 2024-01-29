package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
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

func toMsString(d time.Duration) (float64, string) {
	ms := d.Milliseconds()
	ns := d.Nanoseconds() - (ms * 1_000_000)

	msf := float64(ms) + (float64(ns) / 1_000_000.0)
	return msf, fmt.Sprintf("%10.3fms", msf)
}

func toMsColour(d time.Duration, isTimeout bool) string {
	msf, mss := toMsString(d)

	var result string
	switch {
	case isTimeout:
		result = mss

	case msf < 100.0:
		result = color.GreenString(mss)

	case msf < 200.0:
		result = color.CyanString(mss)

	case msf < 500.0:
		result = color.YellowString(mss)

	default:
		result = color.RedString(mss)
	}

	return result
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
	args := []string{os.Args[0], "-worker", "-port", fmt.Sprintf("%d", conf.RunPort)}
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

	exitSignal := make(chan struct{})
	go func() {
		_, _ = proc.Wait()
		exitSignal <- struct{}{}
	}()

	context, cancel := framework.NewTimeoutContext(100 * time.Millisecond)
	defer cancel()

	select {
	case <-context.Done():
		log.Printf("start background worker port=%d pid=%d", conf.RunPort, proc.Pid)
		// subprocess started and not exited

	case <-exitSignal:
		// subprocess exited
		log.Printf("failed to start background worker port=%d pid=%d", conf.RunPort, proc.Pid)
		proc = nil
	}

	return framework.NewWorkerProc(proc)
}

func initConnection(conf *framework.Configure) (*framework.WorkerProc, *framework.Client) {
	var worker *framework.WorkerProc
	for worker == nil {
		worker = startWorker(conf)
		if worker == nil {
			conf.RunPort += 1
			if conf.RunPort > 1783 {
				conf.RunPort = 1707
			}
		}
	}

	client, err := conf.NewClient("127.0.0.1")
	if err != nil {
		log.Printf("create client failed: %s\n", err)
		panic(err)
	}

	client.SetTimeout(conf.ProblemTimeout, conf.MethodTimeout)
	return worker, client
}

func runProblems(conf *framework.Configure, allProblems []framework.Problem) {
	conf.RunPort = conf.ServePort
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

func printResultItem(conf *framework.Configure, problem framework.Problem,
	result framework.ResultItem, isBest bool) string {
	parts := make([]string, 0, 3)
	if result.IsTimeout {

		parts = append(parts,
			//               1   5   10   15
			color.RedString("NO RESULT      "))
	} else {
		parts = append(parts, fmt.Sprintf("%-15d", result.Result))
	}

	if conf.CheckMode {
		if result.IsTimeout {
			parts = append(parts, color.YellowString("timeout   "))

		} else if problem.NoAnswer {
			parts = append(parts, color.YellowString("unknown   "))

		} else if problem.Answer == framework.Answer(result.Result) {
			parts = append(parts, color.GreenString("correct   "))

		} else {
			parts = append(parts, color.RedString("wrong     "))
		}
	}

	parts = append(parts, toMsColour(result.TimeCost, result.IsTimeout))

	if isBest {
		parts = append(parts, "*BEST")
	}

	return strings.Join(parts, " ")
}

func printResultTitleWithMultipleResults(conf *framework.Configure, problem framework.Problem, result *framework.Result) {
	var correct string
	switch {
	case problem.NoAnswer:
		correct = color.YellowString("unknown   ")

	case result.IsCorrect(problem.Answer):
		correct = color.GreenString("correct   ")

	default:
		correct = color.RedString("wrong     ")
	}

	args := make([]interface{}, 0, 5)
	args = append(args, problem.Id, rightPadding(problem.Title, 40, "."), "")
	format := "%-5d %-40s %15s %s\n"
	if conf.CheckMode {
		args = append(args, correct)
		format = "%-5d %-40s %15s %s %s\n"
	}

	_, timeCost := toMsString(result.TotalCost())
	args = append(args, timeCost)

	fmt.Printf(format, args...)
}

func printResult(conf *framework.Configure, problem framework.Problem, result *framework.Result) {
	if result.Length() == 1 {
		item := result.Results[0]
		resultColumn := printResultItem(conf, problem, item, false)
		fmt.Printf("%-5d %-40s %s\n",
			problem.Id, rightPadding(problem.Title, 40, "."), resultColumn)

	} else {
		printResultTitleWithMultipleResults(conf, problem, result)
		best := result.FindBest()
		for i, item := range result.Results {
			resultColumn := printResultItem(conf, problem, item, best == i)
			fmt.Printf("      + %-38s %s\n",
				rightPadding(item.Method, 38, "."), resultColumn)
		}
	}
}
