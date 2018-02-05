// The tabbar package implements a material tabbar component.
//
// See: https://material.io/components/web/catalog/tabs/
package tabbar // import "agamigo.io/material/component/tabbar"

import (
	"agamigo.io/material/component"
)

// TB is the interface for a material tabbar component.
type TB interface {
	component.C
}

// tabBar is the internal implementation of TB made available publicly via
// New().
type tabBar struct {
	component.C
}

// New creates a material tabbar component that implement the TB interface.
// It is a wrapper around component.New.
func New() (c TB, err error) {
	newTB, err := component.New(component.TabBar)
	if err != nil {
		return nil, err
	}
	return &tabBar{newTB}, err
}
