// iconbutton implements a material iconbutton component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package iconbutton // import "github.com/vecty-components/material/components/iconbutton"

import (
	"syscall/js"

	"github.com/vecty-components/material/base"
)

// IB is a material iconbutton component.
type IB struct {
	mdc *base.Component
	On  bool `js:"on"`
}

// New returns a new component.
func New() *IB {
	c := &IB{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *IB) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *IB) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *IB) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCIconButtonToggle",
				MDCCamelCaseName: "iconButton",
			},
		}
		fallthrough
	case c.mdc.Value.IsNull() || c.mdc.Value.IsUndefined():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *IB) StateMap() base.StateMap {
	return base.StateMap{
		"on": c.On,
	}
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
