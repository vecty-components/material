// radio implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/radio"

import (
	"agamigo.io/material/component"
)

// R is a material radio component.
type R struct {
	*component.Component
	Checked  bool   `js:"checked"`
	Disabled bool   `js:"disabled"`
	Value    string `js:"value"`
}

// ComponentType implements the ComponentTyper interface.
func (c *R) ComponentType() component.ComponentType {
	return component.ComponentType{
		MDCClassName:     "MDCRadio",
		MDCCamelCaseName: "radio",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *R) SetComponent(mdcC *component.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *R) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}
