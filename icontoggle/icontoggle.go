// icontoggle implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "github.com/vecty-material/material/icontoggle"

import (
	"syscall/js"

	"github.com/vecty-material/material/base"
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
func (c *IT) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *IT) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *IT) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCIconToggle",
				MDCCamelCaseName: "iconToggle",
			},
		}
		fallthrough
	case c.mdc.Value == js.Null():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *IT) StateMap() base.StateMap {
	return base.StateMap{
		"on":       c.On,
		"disabled": c.Disabled,
	}
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
