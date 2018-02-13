// linearprogress implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/linearprogress"

import (
	"agamigo.io/gojs"
	"agamigo.io/material"
)

// LP is a material libearprogress component.
type LP struct {
	*material.Component
	Determinate bool    `js:"determinate"`
	Reverse     bool    `js:"reverse"`
	Progress    float64 `js:"progress"`
	Buffer      float64 `js:"buffer"`
	bufferCache float64
}

// ComponentType implements the ComponentTyper interface.
func (c *LP) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCLinearProgress",
		MDCCamelCaseName: "linearProgress",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *LP) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *LP) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// Open opens the linearProgress component.
func (lp *LP) Open() (err error) {
	gojs.CatchException(&err)
	lp.GetObject().Call("open")
	return err
}

// Close closes the linearProgress component.
func (lp *LP) Close() (err error) {
	gojs.CatchException(&err)
	lp.GetObject().Call("close")
	return err
}
