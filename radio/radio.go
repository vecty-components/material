// The radio package implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/radio"

import (
	"agamigo.io/material/component"
)

// R is the interface for a material radio component.
type R interface {
	component.C
}

// radio is the internal implementation of R made available publicly via
// New().
type radio struct {
	component.C
}

// New creates a material radio component that implement the R interface.
// It is a wrapper around component.New.
func New() (c R, err error) {
	newR, err := component.New(component.Radio)
	if err != nil {
		return nil, err
	}
	return &radio{newR}, err
}
