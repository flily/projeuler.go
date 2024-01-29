package framework

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type WorkerProc struct {
	Proc *os.Process
}

func NewWorkerProc(proc *os.Process) *WorkerProc {
	if proc == nil {
		return nil
	}

	w := &WorkerProc{
		Proc: proc,
	}

	return w
}

func (w *WorkerProc) Pid() int {
	return w.Proc.Pid
}

func (w *WorkerProc) Kill() {
	if w.Proc == nil {
		return
	}

	err := w.Proc.Kill()
	if err != nil {
		fmt.Printf("stop worker (PID=%d) failed: %s", w.Proc.Pid, err)
	} else {
		w.Proc = nil
	}
}

type Configure struct {
	RunnerMode     bool
	TotalTimeout   time.Duration
	ClientMode     bool
	WorkerMode     bool
	RawMode        bool
	DebugMode      bool
	ServePort      int
	RunPort        int
	CheckMode      bool
	ProblemTimeout time.Duration
	MethodTimeout  time.Duration
	Problems       []string
}

func (c *Configure) NewClient(host string) (*Client, error) {
	return NewClient(host, c.RunPort)
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
