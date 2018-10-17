package glog

import "testing"

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
