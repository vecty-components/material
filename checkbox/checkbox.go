// checkbox implements a material checkbox component.
//
// See: https://material.io/components/web/catalog/input-controls/checkboxes/
package checkbox // import "agamigo.io/material/checkbox"

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// CB is a material checkbox component.
type CB struct {
	mdc           *base.Component
	Checked       bool   `js:"checked"`
	Indeterminate bool   `js:"indeterminate"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *CB) Start(rootElem *js.Object) error {
	return base.Start(c.Component(), rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *CB) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *CB) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{}
		c.mdc.Type = base.ComponentType{
			MDCClassName:     "MDCCheckbox",
			MDCCamelCaseName: "checkbox",
		}
	}
	return c.mdc
}
