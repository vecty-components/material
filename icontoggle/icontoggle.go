// icontoggle implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/icontoggle"

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// IT is a material icontoggle component.
type IT struct {
	mdc      *base.Component
	On       bool `js:"on"`
	Disabled bool `js:"disabled"`
}

// New returns a new component.
func New() *IT {
	c := &IT{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *IT) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem, js.M{
		"on":       c.On,
		"disabled": c.Disabled,
	})
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *IT) Stop() error {
	return base.Stop(c.Component())
}

// Component returns the component's underlying base.Component.
func (c *IT) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCIconToggle",
				MDCCamelCaseName: "iconToggle",
			},
		}
	}
	return c.mdc.Component()
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
