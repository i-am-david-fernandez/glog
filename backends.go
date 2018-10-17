package glog

import (
	"fmt"
	"io/ioutil"
	"os"

	colorable "github.com/mattn/go-colorable"
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

func newConsoleBackend(tag string, level LogLevel) logging.Backend {

	// Create a console logger
	format := logging.MustStringFormatter(
		`%{color} %{level:-8s} ▶ %{message}%{color:reset}`,
	)
	backend := logging.NewLogBackend(colorable.NewColorableStderr(), "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	leveller := logging.AddModuleLevel(formatter)
	vendorLevel, _ := level.toVendorLevel()
	leveller.SetLevel(vendorLevel, tag)

	return leveller
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

type _RetrievableRecord struct {
	calldepth int
	record    *logging.Record
}

type _RetrievableBackend struct {
	records []_RetrievableRecord
}

func newRetrievableBackend(tag string, level LogLevel) *_RetrievableBackend {
	return &_RetrievableBackend{
		records: make([]_RetrievableRecord, 0),
	}
}

func (backend *_RetrievableBackend) Close() {}

func (backend *_RetrievableBackend) Log(level logging.Level, calldepth int, record *logging.Record) error {

	//fmt.Printf("<<%v, %v, %v, %v>>\n", level, calldepth, record.Level, record.Message())

	backend.records = append(backend.records, _RetrievableRecord{
		calldepth: calldepth,
		record:    record,
	})

	return nil
}

func (backend *_RetrievableBackend) Get(minimumLevel LogLevel) []SessionRecord {

	sessionContent := make([]SessionRecord, 0)

	for _, r := range backend.records {

		recordLevel, _ := fromVendorLevel(r.record.Level)

		if recordLevel >= minimumLevel {
			sessionContent = append(sessionContent, SessionRecord{
				Time:    r.record.Time,
				Level:   recordLevel,
				Message: r.record.Message(),
			})
		}
	}

	return sessionContent
}
