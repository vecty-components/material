// datatable implements a material datatable component.
//
// See: https://material.io/components/web/catalog/datatable/
package datatable

import (
	"syscall/js"

	"github.com/vecty-material/material/base"
)

// DT is a material datatable component.
type DT struct {
	mdc *base.Component
}

// New returns a new component.
func New() *DT {
	c := &DT{}
	c.Component()
	return c
}

// Start initializes the component with an existing HDTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *DT) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HDTMLElement and cleans up
// event listeners, etc.
func (c *DT) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *DT) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCDataTable",
				MDCCamelCaseName: "dataTable",
			},
		}
		fallthrough
	case c.mdc.Value.IsNull() || c.mdc.Value.IsUndefined():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *DT) StateMap() base.StateMap {
	return base.StateMap{}
}
