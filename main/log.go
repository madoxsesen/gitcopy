package main

import (
	"fmt"
	"github.com/fatih/color"
)

var tealColorStringer = color.New(color.FgCyan).SprintfFunc()
var greenColorStringer = color.New(color.FgGreen).SprintfFunc()
var yellowColorStringer = color.New(color.FgYellow).SprintfFunc()
var redColorStringer = color.New(color.FgRed).SprintfFunc()

func Debug(log string, args ...interface{}) {
	fmt.Printf(tealColorStringer("Debug")+":   %s\n", fmt.Sprintf(log, args))
}

func Info(log string, args ...interface{}) {
	fmt.Printf(greenColorStringer("Info")+":    %s\n", fmt.Sprintf(log, args))
}

func Warn(log string, args ...interface{}) {
	fmt.Printf(yellowColorStringer("Warn")+":    %s\n", fmt.Sprintf(log, args))
}

func Error(log string, args ...interface{}) {
	fmt.Printf(redColorStringer("Error")+":   %s\n", fmt.Sprintf(log, args))
}
