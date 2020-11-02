package util

import (
	log "github.com/sirupsen/logrus"
	"os"
)

//LogLevel - specifies at what level to log
type LogLevel string

var logger *log.Logger

// Log levels
const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
	Fatal LogLevel = "fatal"
)

var (
	isVerbose bool = true
)

//Logf - logs a formated string if IsVerbose is set to true at specified LogLevel
func Logf(level LogLevel, f string, args ...interface{}) {
	switch level {
	case Debug:
		if isVerbose {
			logger.Debugf(f, args...)
		}
	case Info:
		if isVerbose {
			logger.Infof(f, args...)
		}
	case Warn:
		if isVerbose {
			logger.Warnf(f, args...)
		}
	case Error:
		logger.Errorf(f, args...)
	}
}

//Logln - logs a line if IsVerbose is set to true at specified LogLevel
func Logln(level LogLevel, args ...interface{}) {
	switch level {
	case Debug:
		if isVerbose {
			logger.Debugln(args...)
		}
	case Info:
		if isVerbose {
			logger.Infoln(args...)
		}
	case Warn:
		if isVerbose {
			logger.Warnln(args...)
		}
	case Error:
		logger.Errorln(args...)
	case Fatal:
		logger.Fatalln(args...)
	}
}

//LogFatal - logger.Fatal interface
func LogFatal(args ...interface{}) {
	logger.Fatal(args...)
}

//LogEnabled decide if informational log should be logged
func LogEnabled(isLogEnabled bool) {
	isVerbose = isLogEnabled
}

func InitLog() {
	logger = &log.Logger{
		Out:       os.Stderr,
		Formatter: new(log.TextFormatter),
		Hooks:     make(log.LevelHooks),
		Level:     log.InfoLevel,
	}
	if isVerbose {
		logger.SetLevel(log.DebugLevel)
	}
}
