package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/flily/projeuler.go/framework"
	_ "github.com/flily/projeuler.go/problems"
)

type CommandEntry func(args []string, problems []*framework.Problem)

var supportedCommand = map[string]CommandEntry{
	"run":    doRun,
	"raw":    doRunRaw,
	"client": doClient,
	"worker": doWorker,
	"list":   doList,
}

func initLogger(debugMode bool) {
	if !debugMode {
		log.SetOutput(io.Discard)
	}

	log.SetFlags(log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
}

func usage() {
	fmt.Printf("usage: ./projeuler [COMMAND] [ARGS...]\n")
	fmt.Printf("\n")
	fmt.Printf("supported commands:\n")
	for cmd := range supportedCommand {
		fmt.Printf("  %s\n", cmd)
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 {
		usage()
		return
	}

	command := args[0]
	nextArgs := args[1:]

	allProblems := framework.GetAllProblems()

	entry, ok := supportedCommand[command]
	if !ok {
		fmt.Printf("error: unknown command '%s'\n", command)
		usage()
		return
	}
	entry(nextArgs, allProblems)
}
