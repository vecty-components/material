package base

import (
	"agamigo.io/gojs"
	"github.com/gopherjs/gopherwasm/js"
)

func DefineSetGet(c Componenter, key string,
	setter interface{}, getter interface{}) (err error) {
	gojs.CatchException(&err)
	js.Global().Get("Object").Call("defineProperty",
		c, key,
		js.M{
			"set": setter,
			"get": getter,
		},
	)
	return err
}
