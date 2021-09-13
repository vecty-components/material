package base

import (
	"syscall/js"

	"github.com/vecty-material/material/gojs"
	"github.com/vecty-material/material/gojs/jsdom"
)

func DefineSetGet(c Componenter, key string,
	setter interface{}, getter interface{}) (err error) {
	defer gojs.CatchException(&err)

	js.Global().Get("Object").Call("defineProperty",
		c.Component().Value, key,
		jsdom.M{
			"set": setter,
			"get": getter,
		},
	)
	return err
}
