// The toolbar package implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import (
	"agamigo.io/material/component"
)

// T is a material toolbar component. It should only be created using the New
// function.
type T struct {
	*component.C
}

// New creates a material toolbar component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*T, error) {
	newT, err := component.New(component.Toolbar)
	if err != nil {
		return nil, err
	}
	return &T{C: newT}, err
}

// TODO: Handle custom events
// - change
