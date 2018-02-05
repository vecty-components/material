// The tabbarscroller package implements a material tabbarscroller component.
//
// See: https://material.io/components/web/catalog/tabs/
package tabbarscroller // import "agamigo.io/material/component/tabbarscroller"

import (
	"agamigo.io/material/component"
)

// TBS is the interface for a material tabbarscroller component.
type TBS interface {
	component.C
}

// tabbarscroller is the internal implementation of TBS made available publicly via
// New().
type tabBarScroller struct {
	component.C
}

// New creates a material tabbarscroller component that implement the TBS interface.
// It is a wrapper around component.New.
func New() (c TBS, err error) {
	newTBS, err := component.New(component.TabBarScroller)
	if err != nil {
		return nil, err
	}
	return &tabBarScroller{newTBS}, err
}
