package glog

// CloseConsole closes/stops the console backend.
func CloseConsole() {

	key := "console"

	if _, ok := globalBackends[key]; ok {
		delete(globalBackends, key)
	}
}

//ResetConsole resets the console backend.
func ResetConsole(level LogLevel) {

	key := "console"

	CloseConsole()
	globalBackends[key] = newConsoleBackend("", level)
	resetBackends()
}
