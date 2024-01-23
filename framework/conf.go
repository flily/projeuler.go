package framework

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Configure struct {
	RunnerMode     bool
	TotalTimeout   time.Duration
	ServerMode     bool
	ServerPort     int
	ProblemTimeout time.Duration
	MethodTimeout  time.Duration
	WorkerMode     bool
	Problems       []string
}

type ProblemRunInfo struct {
	ProblemId int
	Method    string
}

func (i ProblemRunInfo) Valid() bool {
	return i.ProblemId > 0
}

func (i ProblemRunInfo) IsAllMethods() bool {
	return i.Method == ""
}

func NewProblemRunInfo(id int, method string) ProblemRunInfo {
	info := ProblemRunInfo{
		ProblemId: id,
		Method:    method,
	}

	return info
}

func ParseProblemId(problemId string) (ProblemRunInfo, error) {
	idString := problemId
	methodString := ""
	if strings.Contains(problemId, ".") {
		parts := strings.SplitN(problemId, ".", 2)
		idString = parts[0]
		methodString = parts[1]
	}

	if id, err := strconv.Atoi(idString); err == nil {
		return NewProblemRunInfo(id, methodString), nil

	} else {
		newErr := fmt.Errorf("invalid problem id: '%s'", problemId)
		return ProblemRunInfo{}, newErr
	}
}

func ParseProblemIdList(problemIds []string) ([]ProblemRunInfo, error) {
	result := make([]ProblemRunInfo, 0, len(problemIds))
	for _, id := range problemIds {
		info, err := ParseProblemId(id)
		if err != nil {
			return nil, err
		}

		result = append(result, info)
	}

	return result, nil
}
