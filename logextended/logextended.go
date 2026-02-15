package logextended

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	*log.Logger
	logLevel LogLevel
}

func NewLogExtended() *LogExtended {
	return &LogExtended{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelInfo,
	}
}

func (l LogLevel) isValid() bool {
	switch l {
	case LogLevelError, LogLevelWarning, LogLevelInfo:
		return true
	default:
		return false
	}
}

func (lg *LogExtended) SetLogLevel(level LogLevel) {
	if !level.isValid() {
		return
	}
	lg.logLevel = level
}

func (lg *LogExtended) println(level LogLevel, prefix, msg string) {
	if lg.logLevel < level {
		return
	}

	lg.Logger.Println(prefix + msg)
}

func (lg *LogExtended) Infoln(msg string) {
	lg.println(LogLevelInfo, "INFO ", msg)
}

func (lg *LogExtended) Warnln(msg string) {
	lg.println(LogLevelWarning, "WARN ", msg)
}

func (lg *LogExtended) Errorln(msg string) {
	lg.println(LogLevelError, "ERR ", msg)
}
