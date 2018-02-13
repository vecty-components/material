// checkbox implements a material checkbox component.
//
// See: https://material.io/components/web/catalog/input-controls/checkboxes/
package checkbox // import "agamigo.io/material/checkbox"

import (
	"agamigo.io/material"
)

// CB is a material checkbox component.
type CB struct {
	*material.Component
	Checked       bool   `js:"checked"`
	Indeterminate bool   `js:"indeterminate"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

// ComponentType implements the ComponentTyper interface.
func (c *CB) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCCheckbox",
		MDCCamelCaseName: "checkbox",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *CB) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *CB) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}
