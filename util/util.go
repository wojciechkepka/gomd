package util

import (
	"fmt"
	"os/exec"
	rt "runtime"

	log "github.com/sirupsen/logrus"
)

var (
	//IsVerbose - Should the output be printed
	IsVerbose bool = true
)

//URLOpen - tries to open a url based on OS
func URLOpen(url string) error {
	var cmd *exec.Cmd

	switch os := rt.GOOS; os {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("start", url)
	default:
		return fmt.Errorf("Unsupported os %v", os)
	}

	return cmd.Run()
}

//IsSubDirPath - checks if given path contains any '/'
func IsSubDirPath(basePath, path string) bool {
	return CountChInStr('/', path) > 0
}

//CountChInStr - counts ch character occurances in str string
func CountChInStr(ch rune, str string) int {
	count := 0
	for _, c := range str {
		if c == ch {
			count++
		}
	}

	return count
}

//Logln - logs line if IsVerbose is set to true
func Logln(args ...interface{}) {
	if IsVerbose {
		log.Println(args)
	}
}

//Logf - logs format string with args if IsVerbose is set to true
func Logf(f string, args ...interface{}) {
	if IsVerbose {
		log.Printf(f, args)
	}
}

//LogFatalln - logs a fatal line
func LogFatalln(args ...interface{}) {
	log.Fatalln(args)
}

//LogFatalf - logs a fatal format string with args
func LogFatalf(f string, args ...interface{}) {
	log.Fatalf(f, args)
}

//LogFatal - to unify logging interface
func LogFatal(args ...interface{}) {
	log.Fatal(args)
}
