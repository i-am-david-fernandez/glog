package glog

import (
	"errors"

	logging "github.com/shenwei356/go-logging"
)

// LogLevel represents the available logging levels.
type LogLevel int

// Available log levels.
const (
	Debug LogLevel = iota + 1
	Info
	Notice
	Warning
	Error
	Critical
)

// Mapping from LogLevel to string name/label
var logLevelMap = map[LogLevel]string{
	Debug:    "debug",
	Info:     "info",
	Notice:   "notice",
	Warning:  "warning",
	Error:    "error",
	Critical: "critical",
}

// Reverse mapping from name to LogLevel
var logLevelReverseMap map[string]LogLevel

func init() {

	// Initialise reverse mapping
	logLevelReverseMap = make(map[string]LogLevel)

	for k, v := range logLevelMap {
		logLevelReverseMap[v] = k
	}
}

// NewLogLevel returns a LogLevel for the specified level
func NewLogLevel(name string) (LogLevel, error) {
	if level, ok := logLevelReverseMap[name]; ok {
		return level, nil
	}

	return Debug, errors.New("Invalid LogLevel " + name)
}

func (level LogLevel) String() string {
	return logLevelMap[level]
}

func (level LogLevel) toVendorLevel() (logging.Level, error) {

	switch level {
	case Debug:
		return logging.DEBUG, nil
	case Info:
		return logging.INFO, nil
	case Notice:
		return logging.NOTICE, nil
	case Warning:
		return logging.WARNING, nil
	case Error:
		return logging.ERROR, nil
	case Critical:
		return logging.CRITICAL, nil
	default:
		return logging.INFO, errors.New("invalid level")
	}
}

func fromVendorLevel(level logging.Level) (LogLevel, error) {

	switch level {
	case logging.DEBUG:
		return Debug, nil
	case logging.INFO:
		return Info, nil
	case logging.NOTICE:
		return Notice, nil
	case logging.WARNING:
		return Warning, nil
	case logging.ERROR:
		return Error, nil
	case logging.CRITICAL:
		return Critical, nil
	default:
		return Debug, errors.New("invalid level")
	}
}
