// persistentdrawer implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/persistentdrawer"
import (
	"agamigo.io/material/internal/base"
	"github.com/gopherjs/gopherjs/js"
)

// PD is a material persistentdrawer component.
type PD struct {
	mdc  *base.Component
	Open bool `js:"open"`
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *PD) Start(rootElem *js.Object) error {
	return base.Start(c.mdc, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *PD) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *PD) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{}
		c.mdc.Type = base.ComponentType{
			MDCClassName:     "MDCPersistentDrawer",
			MDCCamelCaseName: "drawer",
		}
	}
	return c.mdc
}

// ComponentType implements the ComponentTyper interface.
func (c *PD) ComponentType() base.ComponentType {
	return base.ComponentType{}
}

// TODO: Custom events
// - MDCPersistentDrawer:open
// - MDCPersistentDrawer:close
