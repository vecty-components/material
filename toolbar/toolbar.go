// toolbar implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import "agamigo.io/material"

// T is a material toolbar component.
type T struct {
	*material.Component
}

// ComponentType implements the ComponentTyper interface.
func (c *T) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCTextField",
		MDCCamelCaseName: "textField",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *T) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *T) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// TODO: Handle events?
// - change
