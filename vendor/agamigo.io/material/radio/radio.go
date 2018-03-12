// radio implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/radio"

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// R is a material radio component.
type R struct {
	mdc      *base.Component
	Checked  bool   `js:"checked"`
	Disabled bool   `js:"disabled"`
	Value    string `js:"value"`
}

// New returns a new component.
func New() *R {
	c := &R{}
	c.Component()
	c.Checked = false
	c.Disabled = false
	c.Value = ""
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *R) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem, js.M{
		"checked":  c.Checked,
		"disabled": c.Disabled,
		"value":    c.Value,
	})
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *R) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *R) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCRadio",
				MDCCamelCaseName: "radio",
			},
		}
	}
	return c.mdc.Component()
}
