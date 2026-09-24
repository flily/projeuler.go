package main

import (
	"fmt"

	"github.com/flily/projeuler.go/framework"
)

var listTable = []framework.Column{
	{
		Name:  "PID",
		Style: framework.NewIntegerStyle(ColumnWidthPID),
	},
	{
		Name:  "Title",
		Style: framework.NewGenericStyle(ColumnWidthTitle),
	},
	{
		Name:  "Solution",
		Style: framework.NewGenericStyle(ColumnWidthTitle),
	},
}

func doList(args []string, allProblems []*framework.Problem) {
	selectors, err := framework.ParseSelectorCollection(args)
	if err != nil {
		fmt.Printf("invalid problem selectors\n%s\n", err)
		return
	}

	output := framework.NewOutputTableWith(listTable)

	pidStyle := framework.DefaultDisplayStyle().Blue()
	titleStyle := framework.DefaultDisplayStyle().Green()
	solutionStyles := []framework.DisplayStyle{
		framework.DefaultDisplayStyle().Yellow(),
		framework.DefaultDisplayStyle().Green(),
		framework.DefaultDisplayStyle().Magenta(),
		framework.DefaultDisplayStyle().Cyan(),
	}

	output.PrintHeader()
	for _, problem := range allProblems {
		selector, matched := selectors.Match(problem)
		if !matched {
			continue
		}

		methods, count := selector.MatchedSolutions(problem)
		if count <= 0 {
			continue
		}

		if count == 1 {
			method := methods[0]
			output.PrintStyleItems(
				pidStyle.With(problem.Id),
				titleStyle.With(problem.Title),
				solutionStyles[2].With("- "+method.Method),
			)
			continue
		}

		output.PrintStyleItems(
			pidStyle.With(problem.Id),
			titleStyle.With(problem.Title),
			framework.DefaultDisplayStyle().With(fmt.Sprintf("+-- %d solutions", count)),
		)

		solutionStyleIndex := 0
		for _, method := range methods {
			output.PrintStyleItems(
				framework.DefaultDisplayStyle().With(" "),
				framework.DefaultDisplayStyle().With(" "),
				solutionStyles[solutionStyleIndex].With("+ "+method.Method),
			)

			solutionStyleIndex = (solutionStyleIndex + 1) % len(solutionStyles)
		}
	}

	output.PrintSeparator()
}
