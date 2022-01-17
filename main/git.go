package main

import (
	"os"
)

func CheckGitRepo(string) {

}

func CloneGitRepo(url string) {
	tempDir, err := os.MkdirTemp("", "")
	CheckError(err)
	Info("Target tempdir: %s", tempDir)
}
