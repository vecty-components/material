// The snackbar package implements a material snackbar component.
//
// See: https://material.io/components/web/catalog/snackbars/
package snackbar // import "agamigo.io/material/snackbar"

import (
	"agamigo.io/material/component"
)

// S is the interface for a material snackbar component.
type S interface {
	component.C
}

// snackbar is the internal implementation of S made available publicly via
// New().
type snackbar struct {
	component.C
}

// New creates a material snackbar component that implement the S interface.
// It is a wrapper around component.New.
func New() (c S, err error) {
	newS, err := component.New(component.Snackbar)
	if err != nil {
		return nil, err
	}
	return &snackbar{newS}, err
}
