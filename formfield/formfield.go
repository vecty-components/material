// formfield implements a material formfield component.
//
// See: https://material.io/components/web/catalog/input-controls/form-fields/
package formfield // import "agamigo.io/material/formfield"

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// FF is a material formfield component.
type FF struct {
	mdc   *base.Component
	input *js.Object `js:"input"`
}

func New() *FF {
	c := &FF{}
	c.newMDC()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *FF) Start(rootElem *js.Object) error {
	// Copy state variables before MDC init() destroys them.
	var input *js.Object
	if c.mdc.Object != nil {
		input = c.input
	}
	err := base.Start(c.Component(), rootElem)
	if err != nil {
		return err
	}
	if input != nil {
		c.input = input
	}
	return err
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *FF) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *FF) Component() *base.Component {
	if c.mdc == nil {
		c.newMDC()
	}
	return c.mdc
}

func (c *FF) newMDC() {
	c.mdc = &base.Component{}
	c.mdc.Type = base.ComponentType{
		MDCClassName:     "MDCFormField",
		MDCCamelCaseName: "formField",
	}
	c.mdc.Object = js.Global.Get("Object").New()
}

func (c *FF) SetInput(i base.Componenter) {
	if c.mdc == nil || c.mdc.Object == nil {
		c.newMDC()
	}
	c.input = i.Component()
}
