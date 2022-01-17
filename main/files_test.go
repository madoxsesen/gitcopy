package main

import "testing"

func TestSanitizePath(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{{input: "/var/log", output: "/var/log/"},
		{input: "/var/log/", output: "/var/log/"},
		{input: "", output: "/"},
	}

	for _, elem := range tests {
		result := sanitizePath(elem.input)
		if result != elem.output {
			t.Errorf("Invalid output: Expected %v Got %v", elem.output, result)
		}
	}
}

func TestTraverseFilePath(t *testing.T) {
	//TODO do mock tests on tmpDirs
}
