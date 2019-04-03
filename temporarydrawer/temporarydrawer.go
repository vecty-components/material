// temporarydrawer implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/temporarydrawer"
import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherwasm/js"
)

// TD is a material temporarydrawer component.
type TD struct {
	mdc  *base.Component
	Open bool `js:"open"`
}

// New returns a new component.
func New() *TD {
	c := &TD{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *TD) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *TD) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *TD) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTemporaryDrawer",
				MDCCamelCaseName: "drawer",
			},
		}
		fallthrough
	case c.mdc.Object == nil:
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *TD) StateMap() base.StateMap {
	sm := base.StateMap{
		"open": c.Open,
	}
	return sm
}
