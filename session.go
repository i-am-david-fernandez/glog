package glog

import (
	"time"
)

// CloseSession closes the current session log without creating a new session.
func CloseSession() {

	key := "session"

	if elem, ok := globalBackends[key]; ok {
		elem.(*_RetrievableBackend).Close()
		delete(globalBackends, key)
	}
}

//ResetSession closes the current session log and creates a new session backend.
func ResetSession() {

	key := "session"

	CloseSession()
	globalBackends[key] = newRetrievableBackend("", Debug)
	resetBackends()
}

// SessionRecord encapsulates a single session log entry.
type SessionRecord struct {
	Time    time.Time
	Level   LogLevel
	Message string
}

// GetSession retrieves the current session's log records.
func GetSession(minimumLevel LogLevel) []SessionRecord {

	key := "session"

	if elem, ok := globalBackends[key]; ok {
		return elem.(*_RetrievableBackend).Get(minimumLevel)
	}

	return nil
}
