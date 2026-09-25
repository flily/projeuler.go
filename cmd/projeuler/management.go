package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/flily/projeuler.go/framework"
	"github.com/flily/projeuler.go/framework/management"
)

type OptionalFlags[T flag.Value] struct {
	Value T
	IsSet bool
}

func NewOptionalFlags[T flag.Value](value T) *OptionalFlags[T] {
	f := &OptionalFlags[T]{
		Value: value,
		IsSet: false,
	}
	return f
}

func (f *OptionalFlags[T]) Set(value string) error {
	err := f.Value.Set(value)
	if err != nil {
		return err
	}

	f.IsSet = true
	return nil
}

func (f *OptionalFlags[T]) String() string {
	if !f.IsSet {
		return "<None>"
	}

	return f.Value.String()
}

type Int64 int64

func NewInt64(value int64) *Int64 {
	i := Int64(value)
	return &i
}

func (i *Int64) Set(value string) error {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	*i = Int64(v)
	return nil
}

func (i *Int64) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func displayActionResult(kind string, path string, err error) bool {
	succStyle := framework.NewGenericStyle(10).Right().
		With(framework.DefaultDisplayStyle().Green().Bold())
	failStyle := framework.NewGenericStyle(10).Right().
		With(framework.DefaultDisplayStyle().Red().Bold())

	style := succStyle
	if err != nil {
		style = failStyle
	}

	fmt.Printf("%s %s\n", style.Apply(kind), path)
	if err != nil {
		fmt.Printf("\n")
		fmt.Printf("Error: %v\n", err)
	}

	return err == nil
}

func addProblemTemplate(pid int, title string, answer *int64, solutions []management.SolutionInfo, dry bool) {
	problemDir, err := management.MakeProblemDir(pid, dry)
	if !displayActionResult("MKDIR", problemDir, err) {
		return
	}

	indexFile, err := management.WriteProblemIndex(pid, title, answer, solutions, dry)
	if !displayActionResult("INDEX", indexFile, err) {
		return
	}

	testCaseFile, err := management.WriteProblemTestCase(pid, dry)
	if !displayActionResult("TESTCASE", testCaseFile, err) {
		return
	}

	for _, solution := range solutions {
		solutionFile, err := management.WriteSolution(pid, solution, dry)
		if !displayActionResult("SOLUTION", solutionFile, err) {
			return
		}
	}
}

func doAdd(args []string, _ []*framework.Problem) {
	set := flag.NewFlagSet("add", flag.ExitOnError)
	title := set.String("title", "", "Title of the problem")
	answer := NewOptionalFlags(NewInt64(0))
	set.Var(answer, "answer", "Answer of the problem")
	yes := set.Bool("yes", false, "Accept the action without confirmation")
	dryrun := set.Bool("dry-run", false, "Perform a dry run without making any changes")
	_ = set.Parse(args)

	pid := 0
	solutions := make([]string, 0, 1)
	if set.NArg() > 0 {
		var err error
		pid, err = strconv.Atoi(set.Arg(0))
		if err != nil {
			fmt.Printf("wrong pid '%s'\n", set.Arg(0))
			return
		}

		if set.NArg() > 1 {
			solutions = set.Args()[1:]
		}
	}

	if len(*title) <= 0 {
		*title = fmt.Sprintf("Problem %d", pid)
	}

	if len(solutions) <= 0 {
		solutions = append(solutions, "naive")
	}

	problemDir := management.MakeProblemDirName(pid)
	if info, err := os.Stat(problemDir); err == nil {
		if info.IsDir() {
			fmt.Printf("problem directory '%s' already exists\n", problemDir)

		} else {
			fmt.Printf("problem path '%s' already exists and is not a directory\n", problemDir)
		}

		return
	}

	style := framework.NewGenericStyle(0)
	field := framework.NewGenericStyle(12).Right()

	fmt.Printf("%s: %s\n", field.Apply("Problem ID"), style.Green().Apply(pid))
	fmt.Printf("%s: %s\n", field.Apply("Title"), style.Yellow().Apply(*title))
	fmt.Printf("%s: %s\n", field.Apply("Answer"), style.Yellow().Apply(answer))
	fmt.Printf("%s: %s\n", field.Apply("Solutions"),
		style.Yellow().Apply(strings.Join(solutions, ", ")),
	)

	if !*yes {
		green := framework.NewGenericStyle(0).
			With(framework.DefaultDisplayStyle().Green().Bold())
		fmt.Printf("Type '%s' to create templates shown above: ", green.Apply("yes"))

		reply := ""
		_, _ = fmt.Scanf("%s", &reply)

		reply = strings.ToLower(reply)
		if reply != "yes" && reply != "y" {
			fmt.Printf("canceled\n")
			return
		}
	}

	answerPtr := (*int64)(nil)
	if answer.IsSet {
		value := int64(*answer.Value)
		answerPtr = &value
	}

	solutionList := make([]management.SolutionInfo, 0, len(solutions))
	for _, s := range solutions {
		entryName := management.MakeEntryName(s)
		info := management.SolutionInfo{
			Name:      s,
			EntryName: entryName,
		}
		solutionList = append(solutionList, info)
	}

	addProblemTemplate(pid, *title, answerPtr, solutionList, *dryrun)
}
