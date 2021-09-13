package gojs // import "github.com/vecty-material/material/gojs"

import "syscall/js"

// catchException recovers any JS exceptions and
// stores the error in the parameter
func CatchException(err *error) {
	e := recover()

	if e == nil {
		return
	}

	if er, ok := e.(*js.Error); ok {
		*err = er
	} else if er, ok := e.(*js.ValueError); ok {
		*err = er
	} else {
		// panic(e)
	}
}
