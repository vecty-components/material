// linearprogress implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/linearprogress"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/internal/base"
	"github.com/gopherjs/gopherjs/js"
)

// LP is a material libearprogress component.
type LP struct {
	mdc         *base.Component
	Determinate bool    `js:"determinate"`
	Reverse     bool    `js:"reverse"`
	Progress    float64 `js:"progress"`
	Buffer      float64 `js:"buffer"`
	bufferCache float64
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *LP) Start(rootElem *js.Object) error {
	err := base.Start(c.Component(), rootElem)
	if err != nil {
		return err
	}
	err = c.afterStart()
	if err != nil {
		// TODO: handle afterStart + stop error
		_ = c.Stop()
		return err
	}
	return nil
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *LP) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *LP) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{}
		c.mdc.Type = base.ComponentType{
			MDCClassName:     "MDCLinearProgress",
			MDCCamelCaseName: "linearProgress",
		}
	}
	return c.mdc
}

// Open opens the linearProgress component.
func (lp *LP) Open() (err error) {
	gojs.CatchException(&err)
	lp.Component().Call("open")
	return err
}

// Close closes the linearProgress component.
func (lp *LP) Close() (err error) {
	gojs.CatchException(&err)
	lp.Component().Call("close")
	return err
}
