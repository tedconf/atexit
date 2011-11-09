// Simple atexit implementation. Add handlers using Register, you must call
// atexit.Exit to get the handler invoked (and then terminate the program).
package atexit

import (
	"fmt"
	"os"
)

var handlers = []func(){}

func runHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, "error: atexit handler error:", err)
		}
	}()

	handler()
}

// Exit runs all the atexit handlers and them terminates the program using os.Exit(code)
func Exit(code int) {
	for _, handler := range handlers {
		runHandler(handler)
	}

	os.Exit(code)
}

// Add atexit handler, call atexit.Exit to invoke all handlers.
func Register(handler func()) {
	handlers = append(handlers, handler)
}
