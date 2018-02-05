// The tab package implements a material tab component.
//
// See: https://material.io/components/web/catalog/tabs/
package tab // import "agamigo.io/material/component/tab"

import (
	"agamigo.io/material/component"
)

// T is the interface for a material tab component.
type T interface {
	component.C
}

// tab is the internal implementation of T made available publicly via
// New().
type tab struct {
	component.C
}

// New creates a material tab component that implement the T interface.
// It is a wrapper around component.New.
func New() (c T, err error) {
	newT, err := component.New(component.Tab)
	if err != nil {
		return nil, err
	}
	return &tab{newT}, err
}
