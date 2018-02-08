// checkbox implements a material checkbox component.
//
// See: https://material.io/components/web/catalog/input-controls/checkboxes/
package checkbox // import "agamigo.io/material/checkbox"

import (
	"agamigo.io/material/component"
)

// CB is a material checkbox component. It should only be created using the New
// function.
type CB struct {
	*component.C
	Checked       bool   `js:"checked"`
	Indeterminate bool   `js:"indeterminate"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

// New creates a material checkbox component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*CB, error) {
	newC, err := component.New(component.Checkbox)
	if err != nil {
		return nil, err
	}
	return &CB{C: newC}, err
}
