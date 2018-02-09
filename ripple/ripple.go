// The ripple package implements a material ripple component.
//
// See: https://material.io/components/web/catalog/ripples/
package ripple // import "agamigo.io/material/ripple"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
)

// R is a material ripple component. It should only be created using the New
// function.
type R struct {
	*component.C
	Unbounded bool `js:"unbounded"`
	Disabled  bool `js:"disabled"`
}

// New creates a material ripple component. It is a wrapper around component.New
// which instantiates the component from the MDC library.
func New() (*R, error) {
	newR, err := component.New(component.Ripple)
	if err != nil {
		return nil, err
	}
	return &R{C: newR}, err
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
