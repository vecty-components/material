// appbar implements a material appbar component.
//
// See: https://material.io/components/web/catalog/appbar/
package appbar

import (
	"syscall/js"

	"github.com/vecty-material/material/components/base"
)

// A is a material appbar component.
type A struct {
	mdc *base.Component
}

// New returns a new component.
func New() *A {
	c := &A{}
	c.Component()
	return c
}

// Start initializes the component with an existing HAMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *A) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HAMLElement and cleans up
// event listeners, etc.
func (c *A) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *A) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTopAppBar",
				MDCCamelCaseName: "topAppBar",
			},
		}
		fallthrough
	case c.mdc.Value.IsNull() || c.mdc.Value.IsUndefined():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *A) StateMap() base.StateMap {
	return base.StateMap{}
}
