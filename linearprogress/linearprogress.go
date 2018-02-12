// linearprogress implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/linearprogress"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
)

// LP is a material libearprogress component.
type LP struct {
	*component.C
	Determinate bool    `js:"determinate"`
	Reverse     bool    `js:"reverse"`
	Progress    float64 `js:"progress"`
	Buffer      float64 `js:"buffer"`
	bufferCache float64
}

// MDCType implements the MDComponenter interface.
func (c *LP) MDCType() component.Type {
	return component.LinearProgress
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *LP) MDCClassAttr() string {
	return "mdc-linear-progress"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *LP) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *LP) String() string {
	return c.MDCType().String() + ": " + c.C.String()
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
