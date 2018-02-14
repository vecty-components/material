// icontoggle implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/icontoggle"

import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// IT is a material icontoggle component.
type IT struct {
	mdc      *material.Component
	On       bool `js:"on"`
	Disabled bool `js:"disabled"`
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *IT) Start(rootElem *js.Object) error {
	return material.Start(c.mdc, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *IT) Stop() error {
	return material.Stop(c.mdc)
}

// Component returns the component's underlying material.Component.
func (c *IT) Component() *material.Component {
	if c.mdc == nil {
		c.mdc = &material.Component{}
		c.mdc.Type = material.ComponentType{
			MDCClassName:     "MDCIconToggle",
			MDCCamelCaseName: "iconToggle",
		}
	}
	return c.mdc
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
