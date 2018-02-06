package component

import "github.com/gopherjs/gopherjs/js"

type Event struct {
	Type      string
	Event     *js.Object
	Component C
}
