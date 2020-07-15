package util

import (
	"fmt"
	"os/exec"
	rt "runtime"

	log "github.com/sirupsen/logrus"
)

//LogLevel - specifies at what level to log
type LogLevel string

// Log levels
const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
	Fatal LogLevel = "fatal"
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

//Logf - logs a formated string if IsVerbose is set to true at specified LogLevel
func Logf(level LogLevel, f string, args ...interface{}) {
	switch level {
	case Debug:
		if IsVerbose {
			log.Debugf(f, args)
		}
	case Info:
		if IsVerbose {
			log.Infof(f, args)
		}
	case Warn:
		if IsVerbose {
			log.Warnf(f, args)
		}
	case Error:
		log.Errorf(f, args)
	}
}

//Logln - logs a line if IsVerbose is set to true at specified LogLevel
func Logln(level LogLevel, args ...interface{}) {
	switch level {
	case Debug:
		if IsVerbose {
			log.Debugln(args)
		}
	case Info:
		if IsVerbose {
			log.Infoln(args)
		}
	case Warn:
		if IsVerbose {
			log.Warnln(args)
		}
	case Error:
		log.Errorln(args)
	case Fatal:
		log.Fatalln(args)
	}
}

//LogFatal - log.Fatal interface
func LogFatal(args ...interface{}) {
	log.Fatal(args)
}
