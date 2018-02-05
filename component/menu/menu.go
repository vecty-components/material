// The menu package implements a material menu component.
//
// See: https://material.io/components/web/catalog/menus/
package menu // import "agamigo.io/material/component/menu"

import (
	"agamigo.io/material/component"
)

// M is the interface for a material menu component.
type M interface {
	component.C
}

// menu is the internal implementation of M made available publicly via
// New().
type menu struct {
	component.C
}

// New creates a material menu component that implement the M interface.
// It is a wrapper around component.New.
func New() (c M, err error) {
	newM, err := component.New(component.Menu)
	if err != nil {
		return nil, err
	}
	return &menu{newM}, err
}
