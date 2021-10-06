// snackbar implements a material snackbar component.
//
// See: https://material.io/components/web/catalog/snackbars/
package snackbar // import "github.com/vecty-material/material/components/snackbar"

import (
	"syscall/js"

	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/gojs"
)

// S is a material snackbar component.
type S struct {
	mdc *base.Component

	Timeout int

	CloseOnEscape bool
}

// New returns a new component.
func New() *S {
	c := &S{
		Timeout:       5000,
		CloseOnEscape: true,
	}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *S) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *S) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *S) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCSnackbar",
				MDCCamelCaseName: "snackbar",
			},
		}
		fallthrough
	case c.mdc.Value.IsNull() || c.mdc.Value.IsUndefined():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *S) StateMap() base.StateMap {
	return base.StateMap{
		"timeoutMs":     c.Timeout,
		"closeOnEscape": c.CloseOnEscape,
	}
}

// Open displays the snackbar. If the configuration is invalid an error message
// will be returned and the snackbar will not be shown. For information on
// config requirements look at documentation for S.
func (c *S) Open() error {
	var err error
	gojs.CatchException(&err)

	c.mdc.Call("open")
	return err
}

// TODO: Handle custom events
// - MDCSnackbar:opened
// - MDCSnackbar:closed
