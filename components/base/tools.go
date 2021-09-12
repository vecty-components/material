package base

import (
	"syscall/js"

	"github.com/vecty-material/material/gojs"
	"github.com/vecty-material/material/gojs/jsdom"
)

func DefineSetGet(c Componenter, key string,
	setter interface{}, getter interface{}) (err error) {
	gojs.CatchException(&err)
	js.Global().Get("Object").Call("defineProperty",
		c, key,
		jsdom.M{
			"set": setter,
			"get": getter,
		},
	)
	return err
}
