// checkbox implements a material checkbox component.
//
// See: https://material.io/components/web/catalog/input-controls/checkboxes/
package checkbox // import "agamigo.io/material/checkbox"

import (
	"agamigo.io/material/component"
)

// CB is a material checkbox component.
type CB struct {
	*component.C
	Checked       bool   `js:"checked"`
	Indeterminate bool   `js:"indeterminate"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

// MDCType implements the MDComponenter interface.
func (c *CB) MDCType() component.Type {
	return component.Checkbox
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *CB) MDCClassAttr() string {
	return "mdc-checkbox"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *CB) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *CB) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}
