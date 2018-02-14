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
	mdc  *material.Component
	Open bool `js:"open"`
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *PD) Start(rootElem *js.Object) error {
	return material.Start(c.mdc, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *PD) Stop() error {
	return material.Stop(c.mdc)
}

// Component returns the component's underlying material.Component.
func (c *PD) Component() *material.Component {
	if c.mdc == nil {
		c.mdc = &material.Component{}
		c.mdc.Type = material.ComponentType{
			MDCClassName:     "MDCPersistentDrawer",
			MDCCamelCaseName: "drawer",
		}
	}
	return c.mdc
}

// ComponentType implements the ComponentTyper interface.
func (c *PD) ComponentType() material.ComponentType {
	return material.ComponentType{}
}

// TODO: Custom events
// - MDCPersistentDrawer:open
// - MDCPersistentDrawer:close
