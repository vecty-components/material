// radio implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "github.com/vecty-material/material/radio"

import (
	"syscall/js"

	"github.com/vecty-material/material/base"
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
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *R) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *R) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *R) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCRadio",
				MDCCamelCaseName: "radio",
			},
		}
		fallthrough
	case c.mdc.Value == js.Null():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *R) StateMap() base.StateMap {
	if c.Value == "undefined" {
		c.Value = ""
	}
	return base.StateMap{
		"checked":  c.Checked,
		"disabled": c.Disabled,
		"value":    c.Value,
	}
}
