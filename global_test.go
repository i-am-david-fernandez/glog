package glog

import (
	"os"
	"testing"

	"github.com/shenwei356/go-logging"
)

func TestPrinters(t *testing.T) {

	//g := gomega.NewGomegaWithT(t)

	cases := []struct {
		method  func(string, ...interface{})
		message string
	}{
		{Criticalf, "Critical message"},
		{Errorf, "Error message"},
		{Warningf, "Warning message"},
		{Noticef, "Notice message"},
		{Infof, "Info message"},
		{Debugf, "Debug message"},
	}

	for _, c := range cases {
		c.method(c.message)

		//g.Expect(err).To(gomega.BeNil())
	}
}

func TestExample(t *testing.T) {

	messages := []struct {
		method  func(string, ...interface{})
		message string
	}{
		{Criticalf, "Critical message"},
		{Errorf, "Error message"},
		{Warningf, "Warning message"},
		{Noticef, "Notice message"},
		{Infof, "Info message"},
		{Debugf, "Debug message"},
	}

	backends := []struct {
		name    string
		backend logging.Backend
	}{
		{"console", NewWriterBackend(os.Stdout, "", Debug, "")},
		{"temporary", NewFileBackend("", "", Debug, "")},
		{"session", NewListBackend("", Debug)},
	}

	ClearBackends()
	//SetBackend("default", NewWriterBackend(os.Stderr, "", Debug, ""))

	for _, b := range backends {

		SetBackend(b.name, b.backend)

		for _, m := range messages {
			m.method(m.message)
		}

		RemoveBackend(b.name)

		//g.Expect(err).To(gomega.BeNil())
	}
}
