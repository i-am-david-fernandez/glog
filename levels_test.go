package glog

import "testing"

func TestNewLogLevel(t *testing.T) {

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

		if got != c.expected {
			t.Fail()
		}

		if c.ok && (err != nil) {
			t.Fail()
		}

		if !c.ok && (err == nil) {
			t.Fail()
		}
	}
}
