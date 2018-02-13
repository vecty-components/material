// temporarydrawer implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/temporarydrawer"
import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// TD is a material temporarydrawer component.
type TD struct {
	mdc  *js.Object
	Open bool `js:"open"`
}

// ComponentType implements the ComponentTyper interface.
func (c *TD) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCTemporaryDrawer",
		MDCCamelCaseName: "drawer",
	}
}

// Component implements the material.Componenter interface.
func (c *TD) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdc.
func (c *TD) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *TD) String() string {
	return c.ComponentType().String()
}

// TODO: Custom events
// - MDCTemporaryDrawer:open
// - MDCTemporaryDrawer:close
