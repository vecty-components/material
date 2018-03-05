// temporarydrawer implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/temporarydrawer"
import (
	"agamigo.io/material/base"
	"github.com/gopherjs/gopherjs/js"
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
	c.Open = false
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *TD) Start(rootElem *js.Object) error {
	return base.Start(c, rootElem, js.M{
		"open": c.Open,
	})
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *TD) Stop() error {
	return base.Stop(c.Component())
}

// Component returns the component's underlying base.Component.
func (c *TD) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTemporaryDrawer",
				MDCCamelCaseName: "drawer",
			},
		}
	}
	return c.mdc.Component()
}

// TODO: Custom events
// - MDCTemporaryDrawer:open
// - MDCTemporaryDrawer:close
