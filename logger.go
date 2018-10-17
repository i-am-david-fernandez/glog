package glog

import logging "github.com/shenwei356/go-logging"

// Logger provides management of application and session logging
type Logger struct {
	logger *logging.Logger
}

// NewLogger returns a new Logger
func NewLogger(tag string) *Logger {
	return &Logger{
		logger: logging.MustGetLogger(tag),
	}
}
