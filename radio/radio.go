// radio implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/radio"

import (
	"agamigo.io/material/component"
)

// R is a material radio component.
type R struct {
	*component.C
	Checked  bool   `js:"checked"`
	Disabled bool   `js:"disabled"`
	Value    string `js:"value"`
}

// MDCType implements the MDComponenter interface.
func (c *R) MDCType() component.Type {
	return component.Radio
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *R) MDCClassAttr() string {
	return "mdc-radio"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *R) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *R) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}
