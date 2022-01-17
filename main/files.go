package main

import (
	"os"
)

//TODO return a list of LocalFile
// include git hashes?

type LocalFile struct {
	fullPath string
	fileName string // TODO should probably be normalized by folder and not just be the pure fileName
}

func TraverseFilePath(path string) []LocalFile {
	Debug("Found files:")

	rootDirectory := readDirectory(path)
	resultSlice := make([]LocalFile, 0)
	traverseDirectory(rootDirectory, path, &resultSlice)
	return resultSlice
}

func readDirectory(path string) []os.DirEntry {
	directory, err := os.ReadDir(path)

	CheckErrorWithLog("Error while reading directory contents", err)

	return directory
}

func traverseDirectory(directory []os.DirEntry, path string, slice *[]LocalFile) {
	for _, directoryEntry := range directory {
		isDirectory := directoryEntry.IsDir()

		nextFileFullPath := sanitizePath(path) + directoryEntry.Name()

		if isDirectory {
			directoryContents := readDirectory(nextFileFullPath)
			traverseDirectory(directoryContents, nextFileFullPath, slice)
		} else {
			checkReadRightsForFile(nextFileFullPath)
			*slice = append(*slice, LocalFile{
				fullPath: nextFileFullPath,
				fileName: directoryEntry.Name(),
			})
		}
	}

}

func sanitizePath(path string) string {
	if path == "" {
		return path + "/"
	}

	if rune(path[len(path)-1]) != '/' {
		return path + "/"
	}
	return path
}

func checkReadRightsForFile(path string) {
	file, err := os.Open(path)
	CheckErrorWithLog("Error while attempting to open file contents for reading: ", err)

	err = file.Close()
	CheckError(err)

	Debug(path)
}
