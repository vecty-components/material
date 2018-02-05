// The icontoggle package implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/component/icontoggle"

import (
	"agamigo.io/material/component"
)

// IT is the interface for a material icontoggle component.
type IT interface {
	component.C
}

// icontoggle is the internal implementation of IT made available publicly via
// New().
type iconToggle struct {
	component.C
}

// New creates a material icontoggle component that implement the IT interface.
// It is a wrapper around component.New.
func New() (c IT, err error) {
	newIT, err := component.New(component.IconToggle)
	if err != nil {
		return nil, err
	}
	return &iconToggle{newIT}, err
}
