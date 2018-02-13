// toolbar implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// T is a material toolbar component.
type T struct {
	mdc *js.Object
}

// ComponentType implements the ComponentTyper interface.
func (c *T) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCTextField",
		MDCCamelCaseName: "textField",
	}
}

// Component implements the material.Componenter interface.
func (c *T) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdc.
func (c *T) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *T) String() string {
	return c.ComponentType().String()
}

// TODO: Handle events?
// - change
