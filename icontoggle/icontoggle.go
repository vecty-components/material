// icontoggle implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/icontoggle"

import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// IT is a material icontoggle component.
type IT struct {
	mdc      *js.Object
	On       bool `js:"on"`
	Disabled bool `js:"disabled"`
}

// ComponentType implements the ComponentTyper interface.
func (c *IT) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCIconToggle",
		MDCCamelCaseName: "iconToggle",
	}
}

// Component implements the material.Componenter interface.
func (c *IT) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the
// component's base Component with mdc.
func (c *IT) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *IT) String() string {
	return c.ComponentType().String()
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
