// persistentdrawer implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/persistentdrawer"
import "agamigo.io/material"

// PD is a material persistentdrawer component.
type PD struct {
	*material.Component
	Open bool `js:"open"`
}

// ComponentType implements the ComponentTyper interface.
func (c *PD) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCPersistentDrawer",
		MDCCamelCaseName: "drawer",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *PD) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *PD) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// TODO: Custom events
// - MDCPersistentDrawer:open
// - MDCPersistentDrawer:close
