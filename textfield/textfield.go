// The textfield package implements a material textfield component.
//
// See: https://material.io/components/web/catalog/input-controls/text-field/
package textfield // import "agamigo.io/material/textfield"

import (
	"agamigo.io/material/component"
)

// T is the interface for a material textfield component.
type T interface {
	component.C
}

// textField is the internal implementation of T made available publicly via
// New().
type textField struct {
	component.C
}

// New creates a material textfield component that implement the T interface.
// It is a wrapper around component.New.
func New() (c T, err error) {
	newT, err := component.New(component.TextField)
	if err != nil {
		return nil, err
	}
	return &textField{newT}, err
}
