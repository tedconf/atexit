package atexit

import (
	"os"
)

var handlers = []func(){}

func Exit(code int) {
	for _, handler := range handlers {
		handler()
	}

	os.Exit(code)
}

func Register(handler func()) {
	handlers = append(handlers, handler)
}
