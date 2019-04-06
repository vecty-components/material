package menu

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherwasm/js"
)

// afterStart adds a missing getter to MDCMenu.quickOpen so we can work with
// that variable as expected in Go.
func (c *M) afterStart() error {
	o := c.Component().Value
	proto := js.Global().Get("Object").Call("getPrototypeOf", c)
	ogSetter := js.Global().Get("Object").Call("getOwnPropertyDescriptor",
		proto, "quickOpen").Get("set")
	// Adds a getter for M.quickOpen.
	err := base.DefineSetGet(c, "quickOpen",
		ogSetter,
		func() interface{} {
			return o.Get("foundation_").Get("quickOpen_").Bool()
		},
	)
	return err
}
