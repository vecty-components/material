package menu

import (
	"syscall/js"

	"github.com/vecty-material/material/base"
)

// afterStart adds a missing getter to MDCMenu.quickOpen so we can work with
// that variable as expected in Go.
func (c *M) afterStart() error {
	// Adds a getter for M.quickOpen.
	err := base.DefineSetGet(c, "quickOpen",
		nil,
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return c.mdc.Value.Get("foundation_").Get("quickOpen_")
		}),
	)
	return err
}
