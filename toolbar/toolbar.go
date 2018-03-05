// toolbar implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
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
func (c *T) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem, js.M{})
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *T) Stop() error {
	return base.Stop(c.Component())
}

// Component returns the component's underlying base.Component.
func (c *T) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTextField",
				MDCCamelCaseName: "textField",
			},
		}
	}
	return c.mdc.Component()
}

// TODO: Handle events?
// - change
