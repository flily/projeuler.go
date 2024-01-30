package framework

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

const (
	DataPathName = "data"
)

var projectDataPath string

func parseFunctionFullName(fullName string) (string, string) {
	dot := strings.LastIndex(fullName, ".")
	packageName := fullName[:dot]
	functionName := fullName[dot+1:]
	return packageName, functionName
}

func Import() ([]byte, error) {
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	fName := f.Name()
	packageName, _ := parseFunctionFullName(fName)

	parts := strings.Split(packageName, "/")
	packageIndex := parts[len(parts)-1]
	dataFilename := fmt.Sprintf("%s/%s.txt", projectDataPath, packageIndex)

	fd, err := os.Open(dataFilename)
	if err != nil {
		wd, _ := os.Getwd()
		return nil, fmt.Errorf("cannot open file '%s' here: %s, %s", dataFilename, wd, runtime.GOROOT())
	}

	defer fd.Close()
	data, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	return data, nil
}
