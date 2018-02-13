// textfield implements a material textfield component.
//
// See: https://material.io/components/web/catalog/input-controls/text-field/
package textfield // import "agamigo.io/material/textfield"

import (
	"agamigo.io/gojs"
	"agamigo.io/material"
)

// TF is a material textfield component.
type TF struct {
	*material.Component

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

// ComponentType implements the ComponentTyper interface.
func (c *TF) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCTextField",
		MDCCamelCaseName: "textField",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *TF) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *TF) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// Layout adjusts the dimensions and positions for all sub-elements.
func (tf *TF) Layout() error {
	var err error
	gojs.CatchException(&err)
	tf.GetObject().Call("layout")
	return err
}
