// The toolbar package implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import (
	"agamigo.io/material/component"
)

// T is the interface for a material toolbar component.
type T interface {
	component.C
}

// toolbar is the internal implementation of T made available publicly via
// New().
type toolbar struct {
	component.C
}

// New creates a material toolbar component that implement the T interface.
// It is a wrapper around component.New.
func New() (c T, err error) {
	newT, err := component.New(component.Toolbar)
	if err != nil {
		return nil, err
	}
	return &toolbar{newT}, err
}
