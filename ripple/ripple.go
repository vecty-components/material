// ripple implements a material ripple component.
//
// See: https://material.io/components/web/catalog/ripples/
package ripple // import "agamigo.io/material/ripple"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
)

// R is a material ripple component.
type R struct {
	*component.Component
	Unbounded bool `js:"unbounded"`
	Disabled  bool `js:"disabled"`
}

// ComponentType implements the ComponentTyper interface.
func (c *R) ComponentType() component.ComponentType {
	return component.ComponentType{
		MDCClassName:     "MDCRipple",
		MDCCamelCaseName: "ripple",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *R) SetComponent(mdcC *component.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *R) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// Activate triggers an activation of the ripple (the first stage, which happens
// when the ripple surface is engaged via interaction, such as a mousedown or a
// pointerdown event). It expands from the center.
func (r *R) Activate() error {
	var err error
	gojs.CatchException(&err)
	r.GetObject().Call("activate")
	return err
}

// Deactivate triggers a deactivation of the ripple (the second stage, which
// happens when the ripple surface is engaged via interaction, such as a mouseup
// or a pointerup event). It expands from the center.
func (r *R) Deactivate() error {
	var err error
	gojs.CatchException(&err)
	r.GetObject().Call("deactivate")
	return err
}

// Layout recomputes all dimensions and positions for the ripple element. Useful
// if a ripple surface’s position or dimension is changed programmatically.
func (r *R) Layout() error {
	var err error
	gojs.CatchException(&err)
	r.GetObject().Call("layout")
	return err
}
