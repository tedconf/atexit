// Simple atexit implementation. Add handlers using Register, you must call
// atexit.Exit to get the handler invoked (and then terminate the program).
package atexit

import (
	"os"
)

var handlers = []func(){}

// Exit runs all the atexit handlers and them terminates the program using os.Exit(code)
func Exit(code int) {
	for _, handler := range handlers {
		handler()
	}

	os.Exit(code)
}

// Add atexit handler, call atexit.Exit to invoke all handlers.
func Register(handler func()) {
	handlers = append(handlers, handler)
}
