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

//Logf - logs a formated string if IsVerbose is set to true at specified LogLevel
func Logf(level LogLevel, f string, args ...interface{}) {
	switch level {
	case Debug:
		logger.Debugf(f, args...)
	case Info:
		logger.Infof(f, args...)
	case Warn:
		logger.Warnf(f, args...)
	case Error:
		logger.Errorf(f, args...)
	}
}

//Logln - logs a line if IsVerbose is set to true at specified LogLevel
func Logln(level LogLevel, args ...interface{}) {
	switch level {
	case Debug:
		logger.Debugln(args...)
	case Info:
		logger.Infoln(args...)
	case Warn:
		logger.Warnln(args...)
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

func InitLog(isVerbose, isDebug bool) {
	logger = &log.Logger{
		Out:       os.Stderr,
		Formatter: new(log.TextFormatter),
		Hooks:     make(log.LevelHooks),
		Level:     log.ErrorLevel,
	}
	if isVerbose {
		logger.SetLevel(log.InfoLevel)
	}
	if isDebug {
		logger.SetLevel(log.DebugLevel)
	}
}
