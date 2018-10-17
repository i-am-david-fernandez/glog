package glog

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/onsi/gomega"
)

func TestSession(t *testing.T) {

	g := gomega.NewGomegaWithT(t)

	// Close console logging; we don't need it for this test
	CloseConsole()

	// Reset the session and ensure it is now empty
	ResetSession()
	g.Expect(len(GetSession(Debug))).To(gomega.Equal(0))

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

		sessionContent := GetSession(Debug)

		// Ensure the session contains the expected number of records
		g.Expect(len(sessionContent)).To(gomega.Equal(index + 1))

		// Verify content of most-recent record
		last := sessionContent[len(sessionContent)-1]
		g.Expect(last.Level).To(gomega.Equal(level))
		g.Expect(last.Message).To(gomega.Equal(message))
	}
}
