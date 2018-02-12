// temporarydrawer implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/temporarydrawer"

import (
	"agamigo.io/material/component"
)

// TD is a material temporarydrawer component.
type TD struct {
	*component.C
	Open bool `js:"open"`
}

// MDCType implements the MDComponenter interface.
func (c *TD) MDCType() component.Type {
	return component.TemporaryDrawer
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *TD) MDCClassAttr() string {
	return "mdc-drawer"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *TD) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *TD) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}

// TODO: Custom events
// - MDCTemporaryDrawer:open
// - MDCTemporaryDrawer:close
