// The persistentdrawer package implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/persistentdrawer"

import (
	"agamigo.io/material/component"
)

// PD is a material persistentdrawer component. It should only be created using
// the New function.
type PD struct {
	*component.C
	Open bool `js:"open"`
}

// New creates a material persistentdrawer component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*PD, error) {
	newPD, err := component.New(component.PersistentDrawer)
	if err != nil {
		return nil, err
	}
	return &PD{C: newPD}, err
}

// TODO: Custom events
// - MDCPersistentDrawer:open
// - MDCPersistentDrawer:close
