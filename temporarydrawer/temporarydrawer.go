// The temporarydrawer package implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/temporarydrawer"

import (
	"agamigo.io/material/component"
)

// TD is a material temporarydrawer component. It should only be created using
// the New function.
type TD struct {
	*component.C
	Open bool `js:"open"`
}

// New creates a material temporarydrawer component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*TD, error) {
	newTD, err := component.New(component.TemporaryDrawer)
	if err != nil {
		return nil, err
	}
	return &TD{C: newTD}, err
}

// TODO: Custom events
// - MDCTemporaryDrawer:open
// - MDCTemporaryDrawer:close
