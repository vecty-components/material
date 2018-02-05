// The formfield package implements a material formfield component.
//
// See: https://material.io/components/web/catalog/input-controls/form-fields/
package formfield // import "agamigo.io/material/formfield"

import (
	"agamigo.io/material/component"
)

// FF is the interface for a material formfield component.
type FF interface {
	component.C
}

// formField is the internal implementation of FF made available publicly via
// New().
type formField struct {
	component.C
}

// New creates a material formfield component that implement the FF interface.
// It is a wrapper around component.New.
func New() (c FF, err error) {
	newFF, err := component.New(component.FormField)
	if err != nil {
		return nil, err
	}
	return &formField{newFF}, err
}
