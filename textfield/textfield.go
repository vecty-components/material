// The textfield package implements a material textfield component.
//
// See: https://material.io/components/web/catalog/input-controls/text-field/
package textfield // import "agamigo.io/material/textfield"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
)

// TF is a material textfield component. It should only be created using the New
// function.
type TF struct {
	*component.C

	// The current value of the textfield. Changing this will update the
	// textfield’s value.
	Value string `js:"value"`

	// Whether or not the textfield is disabled.
	Disabled bool `js:"disabled"`

	// Valid and Required are updated according to HTML5 validation markup.
	Valid    bool `js:"valid"`
	Required bool `js:"required"`

	// HelperText provides supplemental information and/or validation
	// messages to users. It appears on input field focus and disappears on
	// input field blur by default, or it can be persistent.
	HelperText string `js:"helperText"`
}

// New creates a material textfield component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*TF, error) {
	newT, err := component.New(component.TextField)
	if err != nil {
		return nil, err
	}
	return &TF{C: newT}, err
}

// Layout adjusts the dimensions and positions for all sub-elements.
func (tf *TF) Layout() error {
	var err error
	gojs.CatchException(&err)
	tf.GetObject().Call("layout")
	return err
}
