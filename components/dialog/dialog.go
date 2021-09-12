// dialog implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "github.com/vecty-material/material/components/dialog"

import (
	"syscall/js"

	"github.com/vecty-material/material/components/base"
	"github.com/vecty-material/material/gojs"
)

// D is a material dialog component.
type D struct {
	mdc *base.Component

	// Open opens and closes the dialog component.
	Open bool `js:"open"`
}

// New returns a new component.
func New() *D {
	c := &D{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *D) Start(rootElem js.Value) error {
	backup := c.stateMap()
	err := base.Start(c, rootElem)
	if err != nil {
		return err
	}
	err = c.afterStart()
	if err != nil {
		return err
	}
	if backup["open"].(bool) == true {
		c.Open = true
	}
	// c.Component().SetState(backup)
	return err
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *D) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *D) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCDialog",
				MDCCamelCaseName: "dialog",
			},
		}
		fallthrough
	case c.mdc.Value.IsNull() || c.mdc.Value.IsUndefined():
		c.mdc.Component().SetState(c.stateMap())
	}
	return c.mdc.Component()
}

// stateMap does not implement the base.StateMapper interface, but it does the
// same thing as StateMap(). This method is private because we cannot restore
// state until after afterStart() is called.
func (c *D) stateMap() base.StateMap {
	return base.StateMap{
		"open": c.Open,
	}
}

// setOpen shows the dialog. If the dialog is already open then setOpen is a
// no-op.
func (c *D) setOpen() error {
	var err error
	defer gojs.CatchException(&err)
	c.Component().Call("show")
	return err
}

// setClose removes the dialog from view. If the dialog is already closed then
// setClose is a no-op.
func (c *D) setClose() error {
	var err error
	defer gojs.CatchException(&err)
	c.Component().Call("close")
	return err
}
