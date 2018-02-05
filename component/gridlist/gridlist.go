// The gridlist package implements a material gridlist component.
//
// See: https://material.io/components/web/catalog/grid-lists/
package gridlist // import "agamigo.io/material/component/gridlist"

import (
	"agamigo.io/material/component"
)

// GL is the interface for a material gridlist component.
type GL interface {
	component.C
}

// gridlist is the internal implementation of GL made available publicly via
// New().
type gridList struct {
	component.C
}

// New creates a material gridlist component that implement the GL interface.
// It is a wrapper around component.New.
func New() (c GL, err error) {
	newGL, err := component.New(component.GridList)
	if err != nil {
		return nil, err
	}
	return &gridList{newGL}, err
}
