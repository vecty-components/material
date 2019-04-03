// formfield implements a material formfield component.
//
// See: https://material.io/components/web/catalog/input-controls/form-fields/
package formfield // import "agamigo.io/material/formfield"

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherwasm/js"
)

// FF is a material formfield component.
type FF struct {
	mdc   *base.Component
	Input interface{} `js:"input"`
}

func New() *FF {
	c := &FF{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *FF) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *FF) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *FF) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCFormField",
				MDCCamelCaseName: "formField",
			},
		}
		fallthrough
	case c.mdc.Object == nil:
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *FF) StateMap() base.StateMap {
	return base.StateMap{
		"input": js.InternalObject(c).Get("Input"),
	}
}
