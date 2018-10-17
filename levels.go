package glog

import "errors"

// LogLevel represents the available logging levels.
type LogLevel int

// Available log levels.
const (
	Debug = iota + 1
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
