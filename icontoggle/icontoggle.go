// icontoggle implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/icontoggle"

import (
	"agamigo.io/material/component"
)

// IT is a material icontoggle component.
type IT struct {
	*component.C
	On       bool `js:"on"`
	Disabled bool `js:"disabled"`
}

// MDCType implements the MDComponenter interface.
func (c *IT) MDCType() component.Type {
	return component.IconToggle
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *IT) MDCClassAttr() string {
	return "mdc-icon-toggle"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *IT) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *IT) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
