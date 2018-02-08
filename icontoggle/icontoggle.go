// The icontoggle package implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/icontoggle"

import (
	"agamigo.io/material/component"
)

// IT is a material icontoggle component. It should only be created using the
// New function.
type IT struct {
	*component.C
	On       bool `js:"on"`
	Disabled bool `js:"disabled"`
}

// New creates a material icontoggle component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*IT, error) {
	newIT, err := component.New(component.IconToggle)
	if err != nil {
		return nil, err
	}
	return &IT{C: newIT}, err
}

// TODO: Wrap refreshToggleData
// TODO: Handle MDCIconToggle:change events
