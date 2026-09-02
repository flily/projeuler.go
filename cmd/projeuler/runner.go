package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/flily/projeuler.go/framework"
)

type Style = framework.Style

const (
	ColumnWidthPID    = 4
	ColumnWidthTitle  = 40
	ColumnWidthAnswer = 30
	ColumnWidthResult = 9
	ColumnWidthTime   = 12
)

var resultTable = []framework.Column{
	{
		Name:  "PID",
		Style: framework.NewIntegerStyle(ColumnWidthPID),
	},
	{
		Name:  "Title / Solution",
		Style: framework.NewGenericStyle(ColumnWidthTitle),
	},
	{
		Name:  "Answer",
		Style: framework.NewGenericStyle(ColumnWidthAnswer).Center(),
	},
	{
		Name:  "Result",
		Style: framework.NewGenericStyle(ColumnWidthResult).Center(),
	},
	{
		Name:  "Time",
		Style: framework.NewGenericStyle(ColumnWidthTime).Right(),
	},
}

func rightPadding(s string, width int, padding string) string {
	if len(s) >= width {
		return s
	}

	paddingLength := width - len(s)
	paddingString := strings.Repeat(padding, paddingLength/len(padding))
	return s + paddingString
}

func costColour(d time.Duration, timeout time.Duration) framework.Colour {
	costMs := float64(d.Nanoseconds()) / 1_000_000.0
	totalTimeout := float64(timeout.Nanoseconds()) / 1_000_000.0
	prop := costMs / totalTimeout

	if prop < 0.1 {
		return framework.ColourGreen

	} else if prop < 0.2 {
		return framework.ColourCyan

	} else if prop < 0.5 {
		return framework.ColourYellow

	} else if prop < 0.8 {
		return framework.ColourMagenta

	} else {
		return framework.ColourRed
	}
}

func makeColourCost(d time.Duration, colour framework.Colour, isBest bool) framework.DisplayStyle {
	nsI := d.Nanoseconds()
	nsF := float64(nsI)

	text := ""
	if d == 0 {
		text = ">>> 0       "

	} else if d < 100*time.Nanosecond {
		text = fmt.Sprintf(">> %2d     ns", nsI)

	} else if d < 10*time.Microsecond {
		text = fmt.Sprintf(">> %6.3f µs", nsF/1_000.0)

	} else {
		text = fmt.Sprintf("%8.3f ms", nsF/1_000_000.0)
	}

	style := framework.DefaultDisplayStyle()

	if isBest {
		return style.BackgroundColour(colour).Bold().With(text)
	} else {
		return style.Colour(colour).With(text)
	}
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

	output := framework.NewOutputTableWith(resultTable)
	output.PrintHeader()

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

		printResult(output, conf, problem, finalResult)
	}

	output.PrintSeparator()
}

func printSolutionResult(out *framework.OutputTable, conf *framework.Configure, problem framework.Problem,
	pid *int, title string, result *framework.ResultItem, isBest bool) {
	parts := make([]framework.DisplayStyle, 0, 6)

	resultStyle := result.Result.Style()

	if pid != nil {
		parts = append(parts, resultStyle.Bold().With(*pid))
	} else {
		parts = append(parts, resultStyle.Bold().With(""))
	}

	if isBest {
		parts = append(parts, resultStyle.ToBackgroundColour().Bold().With("* "+title))
	} else {
		parts = append(parts, resultStyle.With("+ "+title))
	}

	switch result.Result {
	case framework.FinalResultCorrect, framework.FinalResultCrash:
		parts = append(parts, resultStyle.With(result.Answer))

	case framework.FinalResultWrong:
		parts = append(parts, resultStyle.ToBackgroundColour().With(result.Answer))

	default:
		parts = append(parts, framework.DefaultDisplayStyle().
			Red().Bold().With("NO RESULT"))
	}

	parts = append(parts, resultStyle.With(result.Result))

	costColour := costColour(result.TimeCost, conf.MethodTimeout)
	cost := makeColourCost(result.TimeCost, costColour, isBest)
	parts = append(parts, cost)

	out.PrintStyleItems(parts...)
}

func printResultTitleWithMultipleResults(out *framework.OutputTable, conf *framework.Configure,
	problem framework.Problem, result *framework.Result, problemResult framework.FinalResult) {

	resultStyle := problemResult.Style()
	args := make([]framework.DisplayStyle, 0, 5)

	args = append(args, resultStyle.With(problem.Id))
	args = append(args, resultStyle.Bold().With(problem.Title))
	args = append(args, resultStyle.With(""))
	args = append(args, resultStyle.With(problemResult))

	costColour := costColour(result.TotalCost(), conf.ProblemTimeout)
	cost := makeColourCost(result.TotalCost(), costColour, false)
	args = append(args, cost)

	out.PrintStyleItems(args...)
}

func printResult(out *framework.OutputTable, conf *framework.Configure, problem framework.Problem, result *framework.Result) {
	if conf.CheckMode {
		result.CheckResult(problem.Answer)
	}

	problemResult, best := result.GetProblemResult()

	if result.Length() == 1 {
		item := result.Results[0]
		printSolutionResult(out, conf, problem, &problem.Id, problem.Title, item, best == 0)

	} else {
		printResultTitleWithMultipleResults(out, conf, problem, result, problemResult)
		for i, item := range result.Results {
			printSolutionResult(out, conf, problem, nil, item.Method, item, best == i)
		}
	}
}
