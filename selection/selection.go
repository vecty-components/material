// selection implements a material selection component.
//
// See: https://material.io/components/web/catalog/input-controls/select-menus/
package selection // import "agamigo.io/material/selection"

import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// S is a material selection component.
type S struct {
	mdc           *js.Object
	SelectedIndex int  `js:"selectedIndex"`
	Disabled      bool `js:"disabled"`
}

// ComponentType implements the ComponentTyper interface.
func (c *S) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCSelect",
		MDCCamelCaseName: "select",
	}
}

// Component implements the material.Componenter interface.
func (c *S) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdc.
func (c *S) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *S) String() string {
	return c.ComponentType().String()
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
