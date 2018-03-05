package menu

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// afterStart adds a missing getter to MDCMenu.quickOpen so we can work with
// that variable as expected in Go.
func (c *M) afterStart() error {
	o := c.Component()
	proto := js.Global.Get("Object").Call("getPrototypeOf", c)
	ogSetter := js.Global.Get("Object").Call("getOwnPropertyDescriptor",
		proto, "quickOpen").Get("set").Interface()
	// Adds a getter for M.quickOpen.
	err := base.DefineSetGet(c, "quickOpen",
		ogSetter.(func(v interface{})),
		func() interface{} {
			return o.Get("foundation_").Get("quickOpen_").Bool()
		},
	)
	return err
}
