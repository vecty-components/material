// formfield implements a material formfield component.
//
// See: https://material.io/components/web/catalog/input-controls/form-fields/
package formfield // import "agamigo.io/material/formfield"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// FF is a material formfield component.
type FF struct {
	mdc            *base.Component
	ChildComponent *base.Component
	// hasChildComponent bool
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *FF) Start(rootElem *js.Object) error {
	var err error
	gojs.CatchException(&err)
	err = base.Start(c.Component(), rootElem)
	if err != nil {
		return err
	}
	if c.ChildComponent != nil {
		c.Component().Set("input", c.ChildComponent.Component())
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
		c.mdc = &base.Component{}
		c.mdc.Type = base.ComponentType{
			MDCClassName:     "MDCFormField",
			MDCCamelCaseName: "formField",
		}
	}
	return c.mdc
}

// // SetChildComponent associates a material component with the formfield
// // component.
// func (c *FF) SetChildComponent(childC base.Componenter) {
// 	c.childComponent = childC
// 	// c.hasChildComponent = true
// }
