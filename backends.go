package glog

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	logging "github.com/shenwei356/go-logging"
)

var globalBackends map[string]logging.Backend

func resetBackends() {

	backends := make([]logging.Backend, 0)

	for k, v := range globalBackends {
		fmt.Printf("Adding backend for %s\n", k)
		backends = append(backends, v)
	}

	logging.SetBackend(backends...)
}

type _FileBackend struct {
	temporary bool
	file      *os.File
	format    string
	logging.Backend
}

func newFileBackend(tag string, level LogLevel, filename string, fileFormat string) *_FileBackend {

	fb := &_FileBackend{}

	if filename == "" {
		// Use a temporary file
		handle, err := ioutil.TempFile("", "log")
		if err != nil {
			return nil
		}

		fb.file = handle
		fb.temporary = true
	} else {
		// Create a file logger
		handle, err := os.Create(filename)

		if err != nil {
			return nil
		}

		fb.file = handle
		fb.temporary = false
	}

	var logFormat logging.Formatter

	fmt.Printf("Creating backend with file %s, %s\n", filename, fileFormat)

	if fileFormat == "html" {
		fb.format = "html"
		logFormat = logging.MustStringFormatter(
			`<div class="code %{level}"> %{time:2006-01-02 15:04:05.000} %{level:-8s} %{message}</div>`,
		)

		//fb.file.WriteString(getHTMLHeader())

	} else {
		fb.format = "plain"
		logFormat = logging.MustStringFormatter(
			`%{time:2006-01-02 15:04:05.000} ▶ %{level:-8s} ▶ %{message}`,
		)
	}

	backend := logging.NewLogBackend(fb.file, "", 0)
	formatter := logging.NewBackendFormatter(backend, logFormat)
	leveller := logging.AddModuleLevel(formatter)
	vendorLevel, _ := level.toVendorLevel()
	leveller.SetLevel(vendorLevel, tag)

	fb.Backend = leveller

	return fb
}

func (fb _FileBackend) Close() {

	fb.file.Sync()
	fb.file.Close()

	if fb.temporary {
		if err := os.Remove(fb.file.Name()); err != nil {
			Errorf("Could not remove file backend %v: %v", fb.file.Name(), err)
		} else {
			Noticef("Removed file backend %v", fb.file.Name())
		}
	}
}

// ListRecord encapsulates a single session log entry.
type ListRecord struct {
	Time    time.Time
	Level   LogLevel
	Message string
}

// ListBackend provides a simple list-based store of log records.
type ListBackend struct {
	records []ListRecord
}

// NewListBackend creates a new record-list based backend.
func NewListBackend(module string, level LogLevel) *ListBackend {

	return &ListBackend{
		records: make([]ListRecord, 0),
	}
}

// Clear removes all stored records from the backend.
func (backend *ListBackend) Clear() {
	backend.records = make([]ListRecord, 0)
}

// Log implements the Log function required by the Backend interface.
func (backend *ListBackend) Log(vendorLevel logging.Level, calldepth int, record *logging.Record) error {

	//fmt.Printf("<<%v, %v, %v, %v>>\n", level, calldepth, record.Level, record.Message())

	level, _ := fromVendorLevel(vendorLevel)

	backend.records = append(backend.records, ListRecord{
		Time:    record.Time,
		Level:   level,
		Message: record.Message(),
	})

	return nil
}

// Get retrieves all stored log records at or above the specified minimim level.
func (backend *ListBackend) Get(minimumLevel LogLevel) []ListRecord {

	sessionContent := make([]ListRecord, 0)

	for _, r := range backend.records {

		if r.Level >= minimumLevel {
			sessionContent = append(sessionContent, r)
		}
	}

	return sessionContent
}
