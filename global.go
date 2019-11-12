// Package glog is a go logging library.
//
// Primarily, it is a convenience wrapper around go-logging (specifically shenwei356's fork at https://github.com/shenwei356/go-logging).
//
// glog provides convenience functions for configuring commonly-used logging backends (console and file) and for submitting log messages via a (package-scoped) global logger, akin to the print-style helper methods in the standard library log package. It also includes additional backends: a convenience file-based backend and an (unlimited-size) in-memory list backend. This list backend is intended for use in relatively short-lived scenarios, such as batch-processing operations where the log output from each batch is to be treated independently (e.g., conditionally stored or transmitted). In such scenarios, one would clear the backend at the beginning of each batch run and decide what to do with the results at the end. A summary (the number of logged messages of each log level) is available to aid in conditional use.
package glog

import (
	"io"
	"reflect"

	logging "github.com/shenwei356/go-logging"
)

var globalLogger *logging.Logger

func init() {
	globalLogger = logging.MustGetLogger("")
	globalBackends = make(map[string]logging.Backend)

	resetBackends()
}

// Close shuts down the logging system, performing cleanup such as flushing and closing files.
func Close() {

	ClearBackends()
	resetBackends()
}

// SetBackend adds or replaces a named backend.
func SetBackend(name string, backend logging.Backend) error {

	if backend == nil || reflect.ValueOf(backend).IsNil() {
		return nil
	}

	if _, ok := globalBackends[name]; ok {
		RemoveBackend(name)
	}

	globalBackends[name] = backend
	resetBackends()

	return nil
}

// RemoveBackend removes a named backend.
func RemoveBackend(name string) error {

	if backend, ok := globalBackends[name]; ok {

		switch backend.(type) {
		case *FileBackend:
			backend.(*FileBackend).Close()
		}

		delete(globalBackends, name)
	}

	resetBackends()

	return nil
}

// ClearBackends removes all backends.
func ClearBackends() error {

	for name := range globalBackends {
		RemoveBackend(name)
	}

	resetBackends()

	return nil
}

// NewWriterBackend creates a new backend for a supplied io.Writer.
func NewWriterBackend(writer io.Writer, module string, level LogLevel, format string) logging.Backend {

	if format == "" {
		format = `%{color} %{level:-8s} ▶ %{message}%{color:reset}`
	}

	formatter := logging.MustStringFormatter(format)
	backend := logging.NewLogBackend(writer, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, formatter)
	leveller := logging.AddModuleLevel(backendFormatter)
	vendorLevel, _ := level.toVendorLevel()
	leveller.SetLevel(vendorLevel, module)

	return leveller
}

// Logf logs at the specified level with a format string and set of objects.
func Logf(level LogLevel, format string, objects ...interface{}) {

	switch level {
	case Critical:
		globalLogger.Criticalf(format, objects...)
	case Error:
		globalLogger.Errorf(format, objects...)
	case Warning:
		globalLogger.Warningf(format, objects...)
	case Notice:
		globalLogger.Noticef(format, objects...)
	case Info:
		globalLogger.Infof(format, objects...)
	case Debug:
		globalLogger.Debugf(format, objects...)
	default:
		globalLogger.Criticalf(format, objects...)
	}
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
