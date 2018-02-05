// The slider package implements a material slider component.
//
// See: https://material.io/components/web/catalog/input-controls/sliders/
package slider // import "agamigo.io/material/slider"

import (
	"agamigo.io/material/component"
)

// S is the interface for a material slider component.
type S interface {
	component.C
}

// slider is the internal implementation of S made available publicly via
// New().
type slider struct {
	component.C
}

// New creates a material slider component that implement the S interface.
// It is a wrapper around component.New.
func New() (c S, err error) {
	newS, err := component.New(component.Slider)
	if err != nil {
		return nil, err
	}
	return &slider{newS}, err
}
