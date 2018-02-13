// temporarydrawer implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/temporarydrawer"
import "agamigo.io/material"

// TD is a material temporarydrawer component.
type TD struct {
	*material.Component
	Open bool `js:"open"`
}

// ComponentType implements the ComponentTyper interface.
func (c *TD) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCTemporaryDrawer",
		MDCCamelCaseName: "drawer",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *TD) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *TD) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// TODO: Custom events
// - MDCTemporaryDrawer:open
// - MDCTemporaryDrawer:close
