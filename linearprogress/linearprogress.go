// The linearprogress package implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/linearprogress"

import (
	"agamigo.io/material/component"
)

// LP is the interface for a material linearprogress component.
type LP interface {
	component.C
}

// linearProgress is the internal implementation of LP made available publicly via
// New().
type linearProgress struct {
	component.C
}

// New creates a material linearprogress component that implement the LP interface.
// It is a wrapper around component.New.
func New() (c LP, err error) {
	newLP, err := component.New(component.LinearProgress)
	if err != nil {
		return nil, err
	}
	return &linearProgress{newLP}, err
}
