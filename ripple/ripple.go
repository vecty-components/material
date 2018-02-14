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
	mdc       *material.Component
	Unbounded bool `js:"unbounded"`
	Disabled  bool `js:"disabled"`
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *R) Start(rootElem *js.Object) error {
	return material.Start(c.mdc, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *R) Stop() error {
	return material.Stop(c.mdc)
}

// Component returns the component's underlying material.Component.
func (c *R) Component() *material.Component {
	if c.mdc == nil {
		c.mdc = &material.Component{}
		c.mdc.Type = material.ComponentType{
			MDCClassName:     "MDCRipple",
			MDCCamelCaseName: "ripple",
		}
	}
	return c.mdc
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
