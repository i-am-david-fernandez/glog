// Package glog is a go logging library.
// Primarily, it is a convenience wrapper around shenwei356's fork of go-logging (https://github.com/shenwei356/go-logging)
//
// Its primary intended use is within short-lived, typically-unattended applications, such as those that may be periodically run on a server. The driving force was the desire for logging a set of backup-management tools on a home NAS. These tools are typically run in attended mode initially (for example, to setup repositories or test configuration), such that console logging is important, but thereafter are run unattended (for example, via crond), where file-based logging and retrieval of log content for e-mailing is important. Further, these tools operate on a set of "profiles" (for example, they perform backups for a configured set of users, with each user having an independent configuration or "profile"), such that the notion of session-based logging (i.e., separate logging for each processed profile) is important. To accomodate this, glog includes a session logging backend that may be aribrarily reset/cleared or retrieved for storage or transmission.
package glog

import (
	"fmt"

	logging "github.com/shenwei356/go-logging"
)

var globalLogger *Logger

func init() {
	globalLogger = NewLogger("")
	globalBackends = make(map[string]logging.Backend)

	ResetConsole(Info)
	ResetSession()

	//globalBackends["temp"] = newFileBackend("", Debug, "", "plain")
	resetBackends()
}

// Close shuts down the logging system, performing cleanup such as flushing and closing files.
func Close() {

	for k, v := range globalBackends {
		switch v.(type) {
		case *_FileBackend:
			fmt.Printf("Closing backend for %s\n", k)
			v.(*_FileBackend).Close()
			delete(globalBackends, k)
		}
	}

	resetBackends()
}

// Logf logs at the specified level with a format string and set of objects.
func Logf(level LogLevel, format string, objects ...interface{}) {
	globalLogger.Logf(level, format, objects...)
}

// // ---- ---- ---- ----
// // Panic-level methods

// // Panicf logs at PANIC level with a format string and set of objects.
// func Panicf(format string, objects ...interface{}) {
// 	globalLogger.Panicf(format, objects...)
// }

// // ---- ---- ---- ----
// // Fatal-level methods

// // Fatalf logs at FATAL level with a format string and set of objects.
// func Fatalf(format string, objects ...interface{}) {
// 	globalLogger.Fatalf(format, objects...)
// }

// ---- ---- ---- ----
// Critical-level methods

// Criticalf logs at CRITICAL level with a format string and set of objects.
func Criticalf(format string, objects ...interface{}) {
	globalLogger.Criticalf(format, objects...)
}

// ---- ---- ---- ----
// Error-level methods

// Errorf logs at ERROR level with a format string and set of objects.
func Errorf(format string, objects ...interface{}) {
	globalLogger.Errorf(format, objects...)
}

// ---- ---- ---- ----
// Warning-level methods

// Warningf logs at WARNING level with a format string and set of objects.
func Warningf(format string, objects ...interface{}) {
	globalLogger.Warningf(format, objects...)
}

// ---- ---- ---- ----
// Notice-level methods

// Noticef logs at NOTICE level with a format string and set of objects.
func Noticef(format string, objects ...interface{}) {
	globalLogger.Noticef(format, objects...)
}

// ---- ---- ---- ----
// Info-level methods

// Infof logs at INFO level with a format string and set of objects.
func Infof(format string, objects ...interface{}) {
	globalLogger.Infof(format, objects...)
}

// ---- ---- ---- ----
// Debug-level methods

// Debugf logs at DEBUG level with a format string and set of objects.
func Debugf(format string, objects ...interface{}) {
	globalLogger.Debugf(format, objects...)
}
