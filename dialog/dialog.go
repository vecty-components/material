// dialog implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/dialog"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
)

// D is a material dialog component.
type D struct {
	*component.C
	IsOpen bool `js:"open"`
}

// MDCType implements the MDComponenter interface.
func (c *D) MDCType() component.Type {
	return component.Dialog
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *D) MDCClassAttr() string {
	return "mdc-dialog"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *D) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *D) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}

// Open shows the dialog. If the dialog is already open then Open is a no-op.
func (c *D) Open() error {
	var err error
	defer gojs.CatchException(&err)
	c.GetObject().Call("show")
	return err
}

// Close removes the dialog from view. If the dialog is already closed then
// Close is a no-op.
func (c *D) Close() error {
	var err error
	defer gojs.CatchException(&err)
	c.GetObject().Call("close")
	return err
}
