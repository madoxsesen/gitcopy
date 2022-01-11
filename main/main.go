package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage: gitcopy <from_path> [<to_urls>]")
}

func readCommandLineArgs() (string, []string) {
	args := os.Args[1:]

	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	fromPath := args[0]
	gitRepos := args[1:]
	return fromPath, gitRepos
}

func main() {
	fromPath, gitRepos := readCommandLineArgs()

	GitCopy(fromPath, gitRepos)
}
