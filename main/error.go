package main

import "os"

func CheckError(err error) {
	if err == nil {
		return
	}

	Exception("", err)
	os.Exit(1)
}

func CheckErrorWithLog(log string, err error) {
	if err == nil {
		return
	}

	Exception(log, err)
	os.Exit(1)
}
