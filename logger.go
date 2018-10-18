// +build ignore

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

// Logf logs at the specified level with a format string and set of objects.
func (logger *Logger) Logf(level LogLevel, format string, objects ...interface{}) {

	switch level {
	case Critical:
		logger.Criticalf(format, objects...)
	case Error:
		logger.Errorf(format, objects...)
	case Warning:
		logger.Warningf(format, objects...)
	case Notice:
		logger.Noticef(format, objects...)
	case Info:
		logger.Infof(format, objects...)
	case Debug:
		logger.Debugf(format, objects...)
	default:
		logger.Criticalf(format, objects...)
	}
}

// // ---- ---- ---- ----
// // Panic-level methods

// // Panicf logs at PANIC level with a format string and set of objects.
// func (logger *Logger) Panicf(format string, objects ...interface{}) {
// 	logger.logger.Panicf(format, objects...)
// }

// // ---- ---- ---- ----
// // Fatal-level methods

// // Fatalf logs at FATAL level with a format string and set of objects.
// func (logger *Logger) Fatalf(format string, objects ...interface{}) {
// 	logger.logger.Fatalf(format, objects...)
// }

// ---- ---- ---- ----
// Critical-level methods

// Criticalf logs at CRITICAL level with a format string and set of objects.
func (logger *Logger) Criticalf(format string, objects ...interface{}) {
	logger.logger.Criticalf(format, objects...)
}

// ---- ---- ---- ----
// Error-level methods

// Errorf logs at ERROR level with a format string and set of objects.
func (logger *Logger) Errorf(format string, objects ...interface{}) {
	logger.logger.Errorf(format, objects...)
}

// ---- ---- ---- ----
// Warning-level methods

// Warningf logs at WARNING level with a format string and set of objects.
func (logger *Logger) Warningf(format string, objects ...interface{}) {
	logger.logger.Warningf(format, objects...)
}

// ---- ---- ---- ----
// Notice-level methods

// Noticef logs at NOTICE level with a format string and set of objects.
func (logger *Logger) Noticef(format string, objects ...interface{}) {
	logger.logger.Noticef(format, objects...)
}

// ---- ---- ---- ----
// Info-level methods

// Infof logs at INFO level with a format string and set of objects.
func (logger *Logger) Infof(format string, objects ...interface{}) {
	logger.logger.Infof(format, objects...)
}

// ---- ---- ---- ----
// Debug-level methods

// Debugf logs at DEBUG level with a format string and set of objects.
func (logger *Logger) Debugf(format string, objects ...interface{}) {
	logger.logger.Debugf(format, objects...)
}
