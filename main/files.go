package main

import (
	"os"
)

//TODO return a list of structs with the following:
// full file path (for read access)
// cleaned file name (as would be present on a target git repo)
// git hash?

func TraverseFilePath(path string) {
	Debug("Found files:")

	directory := readDirectory(path)
	traverseDirectory(directory, path)
}

func readDirectory(path string) []os.DirEntry {
	directory, err := os.ReadDir(path)

	if err != nil {
		Error("Error while reading directory contents: ", err)
		os.Exit(1)
	}

	return directory
}

func traverseDirectory(directory []os.DirEntry, path string) {
	for _, directoryEntry := range directory {
		isDirectory := directoryEntry.IsDir()

		if isDirectory {
			directoryContents := readDirectory(sanitizePath(path) + directoryEntry.Name())
			traverseDirectory(directoryContents, sanitizePath(path)+directoryEntry.Name())
		} else {
			checkReadRightsForFile(sanitizePath(path) + directoryEntry.Name())
		}
	}

}

func sanitizePath(path string) string {
	if rune(path[len(path)-1]) != '/' {
		return path + "/"
	}
	return path
}

func checkReadRightsForFile(path string) {
	_, err := os.Open(path)

	if err != nil {
		Error("Error while attempting to open file contents for reading: ", err)
		os.Exit(1)
	}

	Debug(path)
}
