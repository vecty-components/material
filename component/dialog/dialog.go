// The dialog package implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/component/dialog"

import (
	"agamigo.io/material/component"
)

// D is the interface for a material dialog component.
type D interface {
	component.C
}

// dialog is the internal implementation of D made available publicly via
// New().
type dialog struct {
	component.C
}

// New creates a material dialog component that implement the D interface.
// It is a wrapper around component.New.
func New() (c D, err error) {
	newD, err := component.New(component.Dialog)
	if err != nil {
		return nil, err
	}
	return &dialog{newD}, err
}
