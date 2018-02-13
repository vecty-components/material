// linearprogress implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/linearprogress"

import (
	"agamigo.io/gojs"
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// LP is a material libearprogress component.
type LP struct {
	mdc         *js.Object
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

// Component implements the material.Componenter interface.
func (c *LP) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the
// component's base Component with mdc.
func (c *LP) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *LP) String() string {
	return c.ComponentType().String()
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
