// The temporarydrawer package implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/component/temporarydrawer"

import (
	"agamigo.io/material/component"
)

// TD is the interface for a material temporarydrawer component.
type TD interface {
	component.C
}

// temporarydrawer is the internal implementation of TD made available publicly via
// New().
type temporaryDrawer struct {
	component.C
}

// New creates a material temporarydrawer component that implement the TD interface.
// It is a wrapper around component.New.
func New() (c TD, err error) {
	newTD, err := component.New(component.TemporaryDrawer)
	if err != nil {
		return nil, err
	}
	return &temporaryDrawer{newTD}, err
}
