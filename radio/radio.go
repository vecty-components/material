// radio implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/radio"

import "agamigo.io/material"

// R is a material radio component.
type R struct {
	*material.Component
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

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *R) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *R) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}
