// persistentdrawer implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/persistentdrawer"

import (
	"agamigo.io/material/component"
)

// PD is a material persistentdrawer component.
type PD struct {
	*component.C
	Open bool `js:"open"`
}

// MDCType implements the MDComponenter interface.
func (c *PD) MDCType() component.Type {
	return component.PersistentDrawer
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *PD) MDCClassAttr() string {
	return "mdc-drawer"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *PD) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *PD) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}

// TODO: Custom events
// - MDCPersistentDrawer:open
// - MDCPersistentDrawer:close
