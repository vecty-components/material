// The formfield package implements a material formfield component.
//
// See: https://material.io/components/web/catalog/input-controls/form-fields/
package formfield // import "agamigo.io/material/formfield"

import (
	"agamigo.io/material/component"
)

// FF is a material formfield component. It should only be created using the
// New function.
type FF struct {
	*component.C
	// TODO: Automatically create a sub-component if we detect a compatible MDC
	// Input within the form-field element. This should be done in Start*().
	// inputComponent *js.Object `js:"input"`
}

// New creates a material formfield component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*FF, error) {
	newFF, err := component.New(component.FormField)
	if err != nil {
		return nil, err
	}
	return &FF{C: newFF}, err
}
