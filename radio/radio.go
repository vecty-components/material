// radio implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/radio"

import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// R is a material radio component.
type R struct {
	mdc      *js.Object
	Checked  bool   `js:"checked"`
	Disabled bool   `js:"disabled"`
	Value    string `js:"value"`
}

// ComponentType implements the ComponentTyper interface.
func (c *R) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCRadio",
		MDCCamelCaseName: "radio",
	}
}

// Component implements the material.Componenter interface.
func (c *R) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdc.
func (c *R) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *R) String() string {
	return c.ComponentType().String()
}
