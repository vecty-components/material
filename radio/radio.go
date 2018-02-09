// The radio package implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/radio"

import (
	"agamigo.io/material/component"
)

// R is a material radio component. It should only be created using the New
// function.
type R struct {
	*component.C
	Checked  bool   `js:"checked"`
	Disabled bool   `js:"disabled"`
	Value    string `js:"value"`
}

// New creates a material radio component. It is a wrapper around component.New
// which instantiates the component from the MDC library.
func New() (*R, error) {
	newR, err := component.New(component.Radio)
	if err != nil {
		return nil, err
	}
	return &R{C: newR}, err
}
