package management

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	ProblemsDir             = "problems"
	ProblemIndexFilename    = "index.go"
	SolutionIndexFilename   = "index.go"
	ProblemPackagePattern   = "p%04d"
	ProblemTestCaseFilename = "index_test.go"
	EntryNamePrefix         = "Solve"
	ProjectName             = "github.com/flily/projeuler.go"
)

type SolutionInfo struct {
	Name      string
	EntryName string
}

func MakeProblemDir(pid int, dry bool) (string, error) {
	name := MakeProblemDirName(pid)

	var err error
	if !dry {
		err = os.MkdirAll(name, os.ModePerm)
	}

	return name, err
}

func WriteSolutionIndex(pid int, name string, answer *int64, solutions []SolutionInfo, dry bool) (string, error) {
	lines := make([]string, 0, 12)

	packageName := MakePackageName(pid)
	lines = append(lines, "package "+packageName)
	lines = append(lines, "")
	lines = append(lines, "import (")
	lines = append(lines, `	"github.com/flily/projeuler.go/framework"`)
	lines = append(lines, ")")
	lines = append(lines, "")

	title := fmt.Sprintf(`var Problem = framework.InitProblem(%d, "%s").`,
		pid, name)
	lines = append(lines, title)

	if answer != nil {
		ans := fmt.Sprintf("	WithAnswer(%d).", *answer)
		lines = append(lines, ans)
	}

	for _, solution := range solutions {
		sol := fmt.Sprintf(`	Solution("%s", %s).`, solution.Name, solution.EntryName)
		lines = append(lines, sol)
	}

	lines = append(lines, "	WithDescription(")
	lines = append(lines, "		`Description of the problem.`,")
	lines = append(lines, "	)")
	lines = append(lines, "")

	contentText := strings.Join(lines, "\n")

	problemDir := MakeProblemDirName(pid)
	filename := path.Join(problemDir, SolutionIndexFilename)
	var err error
	if !dry {
		err = os.WriteFile(filename, []byte(contentText), os.ModePerm)
	}

	return filename, err
}

func WriteProblemTestCase(pid int, dry bool) (string, error) {
	problemDir := MakeProblemDirName(pid)
	packageName := MakePackageName(pid)
	filename := path.Join(problemDir, ProblemTestCaseFilename)

	lines := make([]string, 0, 10)
	lines = append(lines, "package "+packageName)
	lines = append(lines, "")
	lines = append(lines, "import (")
	lines = append(lines, `	"testing"`)
	lines = append(lines, ")")
	lines = append(lines, "")
	lines = append(lines, "func TestProblem(t *testing.T) {")
	lines = append(lines, "	Problem.Check(t).All()")
	lines = append(lines, "}")
	lines = append(lines, "")

	contentText := strings.Join(lines, "\n")

	var err error
	if !dry {
		err = os.WriteFile(filename, []byte(contentText), os.ModePerm)
	}

	return filename, err
}

func WriteSolution(pid int, solution SolutionInfo, dry bool) (string, error) {
	problemDir := MakeProblemDirName(pid)
	packageName := MakePackageName(pid)
	filenameBase := MakeSolutionFilename(solution.Name)
	filename := path.Join(problemDir, filenameBase)

	lines := make([]string, 0, 12)
	lines = append(lines, "package "+packageName)
	lines = append(lines, "")
	lines = append(lines, "func "+solution.EntryName+"() int64 {")
	lines = append(lines, "	return 0")
	lines = append(lines, "}")
	lines = append(lines, "")

	contentText := strings.Join(lines, "\n")

	var err error
	if !dry {
		err = os.WriteFile(filename, []byte(contentText), os.ModePerm)
	}

	return filename, err
}

func ReadProblemIndex() ([]int, error) {
	filename := path.Join(".", ProblemsDir, ProblemIndexFilename)
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	regex, _ := regexp.Compile("p([0-9]{4})")
	lines := strings.Split(string(content), "\n")
	pids := make([]int, 0, len(lines))
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		pid, _ := strconv.Atoi(matches[1])
		pids = append(pids, pid)
	}

	return pids, nil
}

func WriteProblemIndex(pids []int, dry bool) (string, error) {
	filename := path.Join(".", ProblemsDir, ProblemIndexFilename)

	slices.Sort(pids)

	fd, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = fd.Close()
	}()

	_, _ = fd.WriteString("package problems\n")
	_, _ = fd.WriteString("\n")
	_, _ = fd.WriteString("import (\n")

	for _, pid := range pids {
		line := fmt.Sprintf(`	_ "%s/problems/p%04d"`, ProjectName, pid)
		_, _ = fd.WriteString(line + "\n")
	}
	_, _ = fmt.Fprintf(fd, ")\n")

	return filename, nil
}

func AddToProblemIndex(pids []int, dry bool) (string, error) {
	oldPids, err := ReadProblemIndex()
	if err != nil {
		return "", err
	}

	pidMap := make(map[int]struct{}, len(oldPids))
	for _, pid := range oldPids {
		pidMap[pid] = struct{}{}
	}

	newPids := make([]int, 0, len(oldPids)+len(pids))
	newPids = append(newPids, oldPids...)
	for _, pid := range pids {
		if _, found := pidMap[pid]; !found {
			newPids = append(newPids, pid)
		} else {
			return "", fmt.Errorf("problem ID %d already exists in the index", pid)
		}
	}

	return WriteProblemIndex(newPids, dry)
}

func RemoveFromProblemIndex(pids []int, dry bool) (string, error) {
	oldPids, err := ReadProblemIndex()
	if err != nil {
		return "", err
	}

	oldRemoveMap := make(map[int]struct{}, len(oldPids))
	for _, pid := range oldPids {
		oldRemoveMap[pid] = struct{}{}
	}

	toRemoveMap := make(map[int]struct{}, len(pids))
	for _, pid := range pids {
		if _, found := oldRemoveMap[pid]; !found {
			return "", fmt.Errorf("problem ID %d not found in the index", pid)
		}
		toRemoveMap[pid] = struct{}{}
	}

	newPids := make([]int, 0, len(oldPids))
	for _, pid := range oldPids {
		if _, found := toRemoveMap[pid]; !found {
			newPids = append(newPids, pid)
		}
	}

	return WriteProblemIndex(newPids, dry)
}
