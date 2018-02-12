// toolbar implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import (
	"agamigo.io/material/component"
)

// T is a material toolbar component.
type T struct {
	*component.C
}

// MDCType implements the MDComponenter interface.
func (c *T) MDCType() component.Type {
	return component.TextField
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *T) MDCClassAttr() string {
	return "mdc-text-field"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *T) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *T) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}

// TODO: Handle events?
// - change
