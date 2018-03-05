// persistentdrawer implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/persistentdrawer"
import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
)

// PD is a material persistentdrawer component.
type PD struct {
	mdc  *base.Component
	Open bool `js:"open"`
}

// New returns a new component.
func New() *PD {
	c := &PD{}
	c.Component()
	c.Open = false
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *PD) Start(rootElem *js.Object) error {
	return base.Start(c.Component(), rootElem, js.M{
		"open": c.Open,
	})
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *PD) Stop() error {
	return base.Stop(c.Component())
}

// Component returns the component's underlying base.Component.
func (c *PD) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCPersistentDrawer",
				MDCCamelCaseName: "drawer",
			},
		}
	}
	return c.mdc.Component()
}

// TODO: Custom events
// - MDCPersistentDrawer:open
// - MDCPersistentDrawer:close
