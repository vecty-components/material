// selection implements a material selection component.
//
// See: https://material.io/components/web/catalog/input-controls/select-menus/
package selection // import "agamigo.io/material/selection"

import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
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
	c.SelectedIndex = 0
	c.Disabled = false
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *S) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem, js.M{
		"selectedIndex": c.SelectedIndex,
		"disabled":      c.Disabled,
	})
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *S) Stop() error {
	return base.Stop(c.Component())
}

// Component returns the component's underlying base.Component.
func (c *S) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCSelect",
				MDCCamelCaseName: "select",
			},
		}
	}
	return c.mdc.Component()
}

// Selected returns the id of the currently selected option. If no id is present
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
func (s *S) SelectedElem() *js.Object {
	return s.mdc.Get("selectedOptions")
}

// Options returns a slice of menu items comprising the select’s options.
func (s *S) Options() *js.Object {
	return s.mdc.Get("options")
}
