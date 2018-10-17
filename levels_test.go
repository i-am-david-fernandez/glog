package glog

import (
	"fmt"
	"testing"

	"github.com/onsi/gomega"
)

func TestNewLogLevel(t *testing.T) {

	g := gomega.NewGomegaWithT(t)

	cases := []struct {
		input    string
		expected LogLevel
		ok       bool
	}{
		{"debug", Debug, true},
		{"info", Info, true},
		{"notice", Notice, true},
		{"warning", Warning, true},
		{"error", Error, true},
		{"critical", Critical, true},
		{"foobar", Debug, false},
	}

	for _, c := range cases {

		got, err := NewLogLevel(c.input)

		if c.ok {
			g.Expect(err).To(gomega.BeNil())
		} else {
			g.Expect(err).NotTo(gomega.BeNil())
		}

		g.Expect(got).To(gomega.Equal(c.expected))
	}
}

func TestLogLevelString(t *testing.T) {

	g := gomega.NewGomegaWithT(t)

	cases := []struct {
		input    LogLevel
		expected string
	}{
		{Debug, "debug"},
		{Info, "info"},
		{Notice, "notice"},
		{Warning, "warning"},
		{Error, "error"},
		{Critical, "critical"},
	}

	for _, c := range cases {

		got := fmt.Sprint(c.input)

		g.Expect(got).To(gomega.Equal(c.expected))
	}
}
