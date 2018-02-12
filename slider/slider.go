// slider implements a material slider component.
//
// See: https://material.io/components/web/catalog/input-controls/sliders/
package slider // import "agamigo.io/material/slider"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
)

// S is a material slider component.
type S struct {
	*component.C

	// The current value of the slider. Changing this will update the slider’s
	// value.
	Value float64 `js:"value"`

	// The minimum value the slider can have. Values set programmatically will
	// be clamped to this minimum value. Changing this property will update the
	// slider’s value if it is lower than the new minimum.
	Min float64 `js:"min"`

	// The maximum value a slider can have. Values set programmatically will be
	// clamped to this maximum value. Changing this property will update the
	// slider’s value if it is greater than the new maximum.
	Max float64 `js:"max"`

	// Specifies the increments at which a slider value can be set. Can be any
	// positive number, or 0 for no step. Changing this property will update the
	// slider’s value to be quantized along the new step increments.
	Step float64 `js:"step"`

	// Whether or not the slider is disabled.
	Disabled bool `js:"disabled"`
}

// MDCType implements the MDComponenter interface.
func (c *S) MDCType() component.Type {
	return component.Slider
}

// MDCClassAttr implements the MDComponenter interface and returns the HTML
// Class Attribute that is expected to be assigned to the component's root
// HTMLElement.
func (c *S) MDCClassAttr() string {
	return "mdc-slider"
}

// SetMDC implements the MDComponenter interface and replaces the component's
// base MDComponent with mdcC.
func (c *S) SetMDC(mdcC *component.C) {
	c.C = mdcC
}

// String returns the component's "MDCType: status" information.
func (c *S) String() string {
	return c.MDCType().String() + ": " + c.C.String()
}

// Layout recomputes the dimensions and re-lays out the component. This should
// be called if the dimensions of the slider itself or any of its parent
// elements change programmatically (it is called automatically on resize).
func (s *S) Layout() error {
	var err error
	gojs.CatchException(&err)
	s.GetObject().Call("layout")
	return err
}

// TODO: Handle custom events
// - MDCSlider:input
// - MDCSlider:change
