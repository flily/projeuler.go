package framework

import (
	"slices"
	"sort"
)

var problemsList []*Problem
var problemsMap = make(map[int]*Problem)

func SortProblems() {
	sort.Slice(problemsList, func(i, j int) bool {
		return problemsList[i].Id < problemsList[j].Id
	})
}

func RegisterProblem(problem *Problem) {
	problemsList = append(problemsList, problem)
	problemsMap[problem.Id] = problem
}

func GetAllProblems() []*Problem {
	SortProblems()
	return slices.Clone(problemsList)
}

func GetProblem(id int) (*Problem, bool) {
	problem, exists := problemsMap[id]
	return problem, exists
}
