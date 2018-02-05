// The persistentdrawer package implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/component/persistentdrawer"

import (
	"agamigo.io/material/component"
)

// PD is the interface for a material persistentdrawer component.
type PD interface {
	component.C
}

// persistentDrawer is the internal implementation of PD made available publicly via
// New().
type persistentDrawer struct {
	component.C
}

// New creates a material persistentdrawer component that implement the PD interface.
// It is a wrapper around component.New.
func New() (c PD, err error) {
	newPD, err := component.New(component.PersistentDrawer)
	if err != nil {
		return nil, err
	}
	return &persistentDrawer{newPD}, err
}
