package glog

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/onsi/gomega"
)

func TestSession(t *testing.T) {

	g := gomega.NewGomegaWithT(t)

	ClearBackends()
	backendName := "session"
	backend := NewListBackend("", Debug)
	SetBackend(backendName, backend)

	// Verify that the list is initially empty
	g.Expect(len(backend.Get(Debug))).To(gomega.Equal(0))

	levels := make([]LogLevel, 0)
	counts := make(map[LogLevel]int)

	for k := range logLevelMap {
		levels = append(levels, k)
		counts[k] = 0
	}

	N := 32
	for index := 0; index < N; index++ {

		// Generate a message for a random level
		r := rand.Intn(len(levels))
		level := levels[r]
		message := fmt.Sprintf("Message at %s level.", level)

		Logf(level, message)

		sessionContent := backend.Get(Debug)

		// Ensure the session contains the expected number of records
		g.Expect(len(sessionContent)).To(gomega.Equal(index + 1))

		// Verify content of most-recent record
		last := sessionContent[len(sessionContent)-1]
		g.Expect(last.Level).To(gomega.Equal(level))
		g.Expect(last.Message).To(gomega.Equal(message))
	}

	backend.Clear()

	// Verify that the list is empty after clearing.
	g.Expect(len(backend.Get(Debug))).To(gomega.Equal(0))

}
