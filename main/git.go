package main

import (
	"os"
)

func CheckGitRepo(string) {

}

func CloneGitRepo(url string) {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return
	}

	Info("Target tempdir: %s", tempDir)
}
