package glog

import (
	"fmt"
	"os"
)

func Example() {

	// Remove any existing backends
	ClearBackends()

	// Add a backend
	SetBackend("default", // Backend name
		NewWriterBackend(
			os.Stderr, // Write to stderr
			"",        // Empty/unspecified module
			Debug,     // Debug-level and above records will be logged
			"",        // Use the default message format
		))

	// Log some messages
	Debugf("Debug message")
	Infof("Info message")
	Noticef("Notice message")
	Warningf("Warning message")
	Errorf("Error message")
	Criticalf("Critical message")

	// Change "default" log level. Note that this simply replaces the previously-defined backend.
	SetBackend("default", // Backend name
		NewWriterBackend(
			os.Stderr, // Write to stderr
			"",        // Empty/unspecified module
			Warning,   // Warning-level and above records will be logged
			"",        // Use the default message format
		))
}

func ExampleFileBackend() {

	// As we will be using a file-based backend,
	// ensure things (eventually) get cleaned up.
	defer Close()

	// Remove any existing backends
	ClearBackends()

	// Add a backend
	SetBackend("default", // Backend name
		NewFileBackend(
			"my_app.log", // Filename
			true,         // Append if file exists otherwise create
			"",           // Empty/unspecified module
			Debug,        // Debug-level and above records will be logged
			"",           // Use the default message format
		))

	// Log some messages
	Debugf("Debug message")
	Infof("Info message")
	Noticef("Notice message")
	Warningf("Warning message")
	Errorf("Error message")
	Criticalf("Critical message")

}

func ExampleListBackend() {

	ClearBackends()
	backendName := "session"
	backend := NewListBackend("", Debug)
	SetBackend(backendName, backend)

	// Produce some log messages

	// Retrieve all Warning-and-above messages since backend creation
	sessionContent := backend.Get(Warning)

	for _, record := range sessionContent {
		fmt.Printf("Level: %s, message: %s", record.Level, record.Message)
	}

	// Clear backend
	backend.Clear()

	// Produce some more log messages

	// Retrieve all Info-and-above messages since backend was last clearec
	sessionContent = backend.Get(Info)
}
