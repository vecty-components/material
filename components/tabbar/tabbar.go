// tabbar implements a material tabbar component.
//
// See: https://material.io/components/web/catalog/tabbar/
package tabbar

import (
	"syscall/js"

	"github.com/vecty-material/material/components/base"
)

// TB is a material tabbar component.
type TB struct {
	mdc *base.Component
}

// New returns a new component.
func New() *TB {
	c := &TB{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTBMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *TB) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTBMLElement and cleans up
// event listeners, etc.
func (c *TB) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *TB) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTabBar",
				MDCCamelCaseName: "tabBar",
			},
		}
		fallthrough
	case c.mdc.Value.IsNull() || c.mdc.Value.IsUndefined():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *TB) StateMap() base.StateMap {
	return base.StateMap{}
}
