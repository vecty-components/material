package menu

import (
	"syscall/js"

	"github.com/vecty-material/material/components/base"
)

// afterStart adds a missing getter to MDCMenu.quickOpen so we can work with
// that variable as expected in Go.
func (c *M) afterStart() error {
	proto := js.Global().Get("Object").Call("getPrototypeOf", c.mdc.Value)
	ogSetter := js.Global().Get("Object").Call("getOwnPropertyDescriptor",
		proto, "quickOpen").Get("set")
	// Adds a getter for M.quickOpen.
	err := base.DefineSetGet(c, "quickOpen",
		ogSetter,
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return c.mdc.Value.Get("foundation_").Get("quickOpen_")
		}),
	)
	return err
}
