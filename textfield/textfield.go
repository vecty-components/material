// textfield implements a material textfield component.
//
// See: https://material.io/components/web/catalog/input-controls/text-field/
package textfield // import "agamigo.io/material/textfield"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/internal/base"
	"github.com/gopherjs/gopherjs/js"
)

// TF is a material textfield component.
type TF struct {
	mdc *base.Component

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

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *TF) Start(rootElem *js.Object) error {
	return base.Start(c.mdc, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *TF) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *TF) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{}
		c.mdc.Type = base.ComponentType{
			MDCClassName:     "MDCTextField",
			MDCCamelCaseName: "textField",
		}
	}
	return c.mdc
}

// Layout adjusts the dimensions and positions for all sub-elements.
func (tf *TF) Layout() error {
	var err error
	gojs.CatchException(&err)
	tf.mdc.Call("layout")
	return err
}
