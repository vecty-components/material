// ripple implements a material ripple component.
//
// See: https://material.io/components/web/catalog/ripples/
package ripple // import "agamigo.io/material/ripple"

import (
	"agamigo.io/gojs"
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// R is a material ripple component.
type R struct {
	mdc       *js.Object
	Unbounded bool `js:"unbounded"`
	Disabled  bool `js:"disabled"`
}

// ComponentType implements the ComponentTyper interface.
func (c *R) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCRipple",
		MDCCamelCaseName: "ripple",
	}
}

// Component implements the material.Componenter interface.
func (c *R) Component() *js.Object {
	return c.mdc
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdc.
func (c *R) SetComponent(mdc *js.Object) {
	c.mdc = mdc
}

// String returns the component's ComponentType MDCClassName.
func (c *R) String() string {
	return c.ComponentType().String()
}

// Activate triggers an activation of the ripple (the first stage, which happens
// when the ripple surface is engaged via interaction, such as a mousedown or a
// pointerdown event). It expands from the center.
func (r *R) Activate() error {
	var err error
	gojs.CatchException(&err)
	r.mdc.Call("activate")
	return err
}

// Deactivate triggers a deactivation of the ripple (the second stage, which
// happens when the ripple surface is engaged via interaction, such as a mouseup
// or a pointerup event). It expands from the center.
func (r *R) Deactivate() error {
	var err error
	gojs.CatchException(&err)
	r.mdc.Call("deactivate")
	return err
}

// Layout recomputes all dimensions and positions for the ripple element. Useful
// if a ripple surface’s position or dimension is changed programmatically.
func (r *R) Layout() error {
	var err error
	gojs.CatchException(&err)
	r.mdc.Call("layout")
	return err
}
