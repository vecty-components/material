// selection implements a material selection component.
//
// See: https://material.io/components/web/catalog/input-controls/select-menus/
package selection // import "agamigo.io/material/selection"

import (
	"agamigo.io/material/component"
	"github.com/gopherjs/gopherjs/js"
)

// S is a material selection component.
type S struct {
	*component.C
	SelectedIndex int  `js:"selectedIndex"`
	Disabled      bool `js:"disabled"`
}

// MDCType implements the MDComponenter interface.
func (c *S) MDCType() component.Type {
	return component.Select
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *S) MDCClassAttr() string {
	return "mdc-drawer"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *S) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *S) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}

// Selected returns the id of the currently selected option. If no id is present
// on the selected option, its textContent is used. Returns an empty string when
// no option is selected.
func (s *S) SelectedString() string {
	v := s.GetObject().Get("value").String()
	if v == "undefined" {
		return ""
	}
	return v
}

// SelectedElem returns a NodeList of either the currently selected option, or
// an empty js.S if nothing is selected.
func (s *S) SelectedElem() *js.Object {
	return s.GetObject().Get("selectedOptions")
}

// Options returns a slice of menu items comprising the select’s options.
func (s *S) Options() *js.Object {
	return s.GetObject().Get("options")
}
