// dialog implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/dialog"

import (
	"agamigo.io/gojs"
	"agamigo.io/material"
)

// D is a material dialog component.
type D struct {
	*material.Component
	IsOpen bool `js:"open"`
}

// ComponentType implements the ComponentTyper interface.
func (c *D) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCDialog",
		MDCCamelCaseName: "dialog",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *D) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *D) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
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
