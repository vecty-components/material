package gojs // import "github.com/vecty-material/gojs"

import "syscall/js"

// catchException recovers any JS exceptions and
// stores the error in the parameter
func CatchException(err *error) {
	e := recover()

	if e == nil {
		return
	}

	if e, ok := e.(*js.Error); ok {
		*err = e
	} else {
		panic(e)
	}
}
