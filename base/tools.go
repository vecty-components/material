package base

import (
	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

func DefineSetGet(c Componenter, key string,
	setter func(value interface{}), getter func() interface{}) (err error) {
	gojs.CatchException(&err)
	js.Global.Get("Object").Call("defineProperty",
		c, key,
		js.M{
			"set": setter,
			"get": getter,
		},
	)
	return err
}

// func DefineSetter(c Componenter, key string, set func(v interface{})) error {
// 	var err error
// 	gojs.CatchException(&err)
// 	js.Global.Get("Object").Call("defineProperty",
// 		c, key,
// 		js.M{
// 			"set": setter,
// 		},
// 	)
// 	return err
// }
