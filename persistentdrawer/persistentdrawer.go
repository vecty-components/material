// persistentdrawer implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/persistentdrawer"
import (
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// PD is a material persistentdrawer component.
type PD struct {
	mdc  *js.Object
	Open bool `js:"open"`
}

// ComponentType implements the ComponentTyper interface.
func (c *PD) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCPersistentDrawer",
		MDCCamelCaseName: "drawer",
	}
}

// Component implements the material.Componenter interface.
func (c *PD) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the
// component's base Component with mdcC.
func (c *PD) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *PD) String() string {
	return c.ComponentType().String()
}

// TODO: Custom events
// - MDCPersistentDrawer:open
// - MDCPersistentDrawer:close
