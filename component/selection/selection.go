// The selection package implements a material selection component.
//
// See: https://material.io/components/web/catalog/input-controls/select-menus/
package selection // import "agamigo.io/material/component/selection"

import (
	"agamigo.io/material/component"
)

// S is the interface for a material selection component.
type S interface {
	component.C
}

// selection is the internal implementation of S made available publicly via
// New().
type selection struct {
	component.C
}

// New creates a material selection component that implement the S interface.
// It is a wrapper around component.New.
func New() (c S, err error) {
	newS, err := component.New(component.Select)
	if err != nil {
		return nil, err
	}
	return &selection{newS}, err
}
