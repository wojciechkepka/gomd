package util

import (
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
	isVerbose bool = true
)

//Logf - logs a formated string if IsVerbose is set to true at specified LogLevel
func Logf(level LogLevel, f string, args ...interface{}) {
	switch level {
	case Debug:
		if isVerbose {
			log.Debugf(f, args)
		}
	case Info:
		if isVerbose {
			log.Infof(f, args)
		}
	case Warn:
		if isVerbose {
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
		if isVerbose {
			log.Debugln(args)
		}
	case Info:
		if isVerbose {
			log.Infoln(args)
		}
	case Warn:
		if isVerbose {
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

//LogEnabled decide if informational log should be logged
func LogEnabled(isLogEnabled bool) {
	isVerbose = isLogEnabled
}
