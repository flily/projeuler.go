package management

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	IndexFilename           = "index.go"
	ProblemTestCaseFilename = "index_test.go"
	EntryNamePrefix         = "Solve"
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

func WriteProblemIndex(pid int, name string, answer *int64, solutions []SolutionInfo, dry bool) (string, error) {
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
	filename := path.Join(problemDir, IndexFilename)
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
