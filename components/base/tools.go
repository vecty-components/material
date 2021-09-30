package base

import (
	"syscall/js"

	"github.com/vecty-material/material/gojs"
)

func DefineSetGet(c Componenter, key string,
	setter interface{}, getter interface{}) (err error) {

	for _, f := range []interface{}{setter, getter} {
		switch f.(type) {
		case js.Value: // should precede Wrapper to avoid a loop
		case js.Func:
		default:
			panic("setter/getter must be js.Value or js.Func")
		}
	}

	defer gojs.CatchException(&err)
	js.Global().Get("Object").Call("defineProperty",
		c.Component().Value, key,
		map[string]interface{}{
			"set": setter,
			"get": getter,
		},
	)
	return err
}
