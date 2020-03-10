// selection implements a material selection component.
//
// See: https://material.io/components/web/catalog/input-controls/select-menus/
package selection // import "github.com/vecty-material/material/material/selection"

import (
	"syscall/js"

	"github.com/vecty-material/material/material/base"
)

// S is a material selection component.
type S struct {
	mdc           *base.Component
	SelectedIndex int  `js:"selectedIndex"`
	Disabled      bool `js:"disabled"`
}

// New returns a new component.
func New() *S {
	c := &S{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *S) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *S) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *S) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCSelect",
				MDCCamelCaseName: "select",
			},
		}
		fallthrough
	case c.mdc.Value.IsNull():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *S) StateMap() base.StateMap {
	return base.StateMap{
		"selectedIndex": c.SelectedIndex,
		"disabled":      c.Disabled,
	}
}

// SelectedString returns the id of the currently selected option. If no id is present
// on the selected option, its textContent is used. Returns an empty string when
// no option is selected.
func (s *S) SelectedString() string {
	v := s.mdc.Get("value").String()
	if v == "undefined" {
		return ""
	}
	return v
}

// SelectedElem returns a NodeList of either the currently selected option, or
// an empty js.S if nothing is selected.
func (s *S) SelectedElem() js.Value {
	return s.mdc.Get("selectedOptions")
}

// Options returns a slice of menu items comprising the select’s options.
func (s *S) Options() js.Value {
	return s.mdc.Get("options")
}
