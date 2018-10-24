package glog

import (
	"io/ioutil"
	"os"
	"time"

	logging "github.com/shenwei356/go-logging"
)

var globalBackends map[string]logging.Backend

func resetBackends() {

	backends := make([]logging.Backend, 0)

	for _, v := range globalBackends {
		backends = append(backends, v)
	}

	logging.SetBackend(backends...)
}

type FileBackend struct {
	temporary bool
	file      *os.File
	format    string
	logging.Backend
}

// NewFileBackend creates a new file-based backend.
func NewFileBackend(filename string, append bool, module string, level LogLevel, format string) *FileBackend {

	fb := &FileBackend{}

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

		flags := os.O_CREATE | os.O_WRONLY
		if append {
			flags |= os.O_APPEND
		}

		handle, err := os.OpenFile(filename, flags, 0600)

		if err != nil {
			return nil
		}

		fb.file = handle
		fb.temporary = false
	}

	if format == "" {
		format = `%{time:2006-01-02 15:04:05.000} ▶ %{level:-8s} ▶ %{message}`
	}

	fb.format = "plain"
	backendFormatter := logging.MustStringFormatter(format)
	backend := logging.NewLogBackend(fb.file, "", 0)
	formatter := logging.NewBackendFormatter(backend, backendFormatter)
	leveller := logging.AddModuleLevel(formatter)
	vendorLevel, _ := level.toVendorLevel()
	leveller.SetLevel(vendorLevel, module)

	fb.Backend = leveller

	return fb
}

func (fb FileBackend) Close() {

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

// Record encapsulates a single log entry.
type Record struct {
	Time    time.Time
	Level   LogLevel
	Message string
}

type RecordSummary struct {
	Level LogLevel
	Count int
}

// ListBackend provides a simple list-based store of log records.
type ListBackend struct {
	records []Record
}

// NewListBackend creates a new record-list based backend.
func NewListBackend(module string, level LogLevel) *ListBackend {

	return &ListBackend{
		records: make([]Record, 0),
	}
}

// Clear removes all stored records from the backend.
func (backend *ListBackend) Clear() {
	backend.records = make([]Record, 0)
}

// Log implements the Log function required by the Backend interface.
func (backend *ListBackend) Log(vendorLevel logging.Level, calldepth int, record *logging.Record) error {

	//fmt.Printf("<<%v, %v, %v, %v>>\n", level, calldepth, record.Level, record.Message())

	level, _ := fromVendorLevel(vendorLevel)

	backend.records = append(backend.records, Record{
		Time:    record.Time,
		Level:   level,
		Message: record.Message(),
	})

	return nil
}

// Get retrieves all stored log records at or above the specified minimim level.
func (backend *ListBackend) Get(minimumLevel LogLevel) []Record {

	sessionContent := make([]Record, 0)

	for _, r := range backend.records {

		if r.Level >= minimumLevel {
			sessionContent = append(sessionContent, r)
		}
	}

	return sessionContent
}

// Summary retrieves a summary of the counts of all records at each level.
func (backend *ListBackend) Summary() []*RecordSummary {

	summaryMap := make(map[LogLevel]*RecordSummary)
	for _, l := range ListLogLevels() {
		summaryMap[l] = &RecordSummary{
			Level: l,
			Count: 0,
		}
	}

	for _, r := range backend.records {
		summaryMap[r.Level].Count++
	}

	summary := make([]*RecordSummary, 0)
	for _, l := range ListLogLevels() {
		summary = append(summary, summaryMap[l])
	}

	return summary
}
