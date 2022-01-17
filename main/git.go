package main

import (
	"os"
)

func CheckGitRepo(string) {
	//TODO check whether local or remote git repo
}

func CloneGitRepo(url string) {
	tempDir, err := os.MkdirTemp("", "")
	CheckError(err)
	Info("Target tempdir: %s", tempDir)
}
