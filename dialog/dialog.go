// dialog implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/dialog"

import (
	"agamigo.io/gojs"
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// D is a material dialog component.
type D struct {
	mdc    *js.Object
	IsOpen bool `js:"open"`
}

// ComponentType implements the ComponentTyper interface.
func (c *D) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCDialog",
		MDCCamelCaseName: "dialog",
	}
}

// Component implements the material.Componenter interface.
func (c *D) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the
// component's base Component with mdc.
func (c *D) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *D) String() string {
	return c.ComponentType().String()
}

// Open shows the dialog. If the dialog is already open then Open is a no-op.
func (c *D) Open() error {
	var err error
	defer gojs.CatchException(&err)
	c.Component().Call("show")
	return err
}

// Close removes the dialog from view. If the dialog is already closed then
// Close is a no-op.
func (c *D) Close() error {
	var err error
	defer gojs.CatchException(&err)
	c.Component().Call("close")
	return err
}
