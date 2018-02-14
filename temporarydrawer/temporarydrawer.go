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
	mdc  *material.Component
	Open bool `js:"open"`
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *TD) Start(rootElem *js.Object) error {
	return material.Start(c.mdc, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *TD) Stop() error {
	return material.Stop(c.mdc)
}

// Component returns the component's underlying material.Component.
func (c *TD) Component() *material.Component {
	if c.mdc == nil {
		c.mdc = &material.Component{}
		c.mdc.Type = material.ComponentType{
			MDCClassName:     "MDCTemporaryDrawer",
			MDCCamelCaseName: "drawer",
		}
	}
	return c.mdc
}

// TODO: Custom events
// - MDCTemporaryDrawer:open
// - MDCTemporaryDrawer:close
