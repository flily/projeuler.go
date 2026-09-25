package management

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

func MakePackageName(pid int) string {
	return fmt.Sprintf("p%04d", pid)
}

func MakeProblemDirName(pid int) string {
	name := MakePackageName(pid)
	return path.Join(".", "problems", name)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func MakeEntryNameBase(name string) string {
	regex, _ := regexp.Compile("[a-zA-Z0-9]+")
	matches := regex.FindAllString(name, -1)

	parts := make([]string, 0, len(matches))
	for _, m := range matches {
		lower := strings.ToLower(m)
		cap := capitalize(lower)
		parts = append(parts, cap)
	}

	return strings.Join(parts, "")
}

func MakeEntryName(name string) string {
	return EntryNamePrefix + MakeEntryNameBase(name)
}

func MakeSolutionFilename(name string) string {
	return strings.ToLower(MakeEntryNameBase(name)) + ".go"
}
