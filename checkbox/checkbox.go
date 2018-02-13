// checkbox implements a material checkbox component.
//
// See: https://material.io/components/web/catalog/input-controls/checkboxes/
package checkbox // import "agamigo.io/material/checkbox"

import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// CB is a material checkbox component.
type CB struct {
	*component
	Checked       bool   `js:"checked"`
	Indeterminate bool   `js:"indeterminate"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

// component has fields and methods we need to satisfy material.Componenter but
// we do not need to expose to component users.
type component struct {
	*js.Object
}

// ComponentType implements the ComponentTyper interface.
func (c *CB) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCCheckbox",
		MDCCamelCaseName: "checkbox",
	}
}

// Component implements the material.Componenter interface.
func (c *CB) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the material.Componenter interface and replaces the
// component's base Component with mdcC.
func (c *CB) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *CB) String() string {
	return c.ComponentType().String()
}
