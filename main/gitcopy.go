package main

import (
	"fmt"
)

func GitCopy(filePath string, gitRepos []string) {
	//TODO check whether git is installed and working

	TraverseFilePath(filePath)

	for _, repo := range gitRepos {
		fmt.Printf("Checking access of git repo: %s\n", repo)
		CheckGitRepo(repo)
	}
}
