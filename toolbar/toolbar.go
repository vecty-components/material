// toolbar implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherwasm/js"
)

// T is a material toolbar component.
type T struct {
	mdc *base.Component
}

// New returns a new component.
func New() *T {
	c := &T{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *T) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *T) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *T) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCToolbar",
				MDCCamelCaseName: "toolbar",
			},
		}
		fallthrough
	case c.mdc.Value == js.Null():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *T) StateMap() base.StateMap {
	return base.StateMap{}
}
