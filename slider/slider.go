// slider implements a material slider component.
//
// See: https://material.io/components/web/catalog/input-controls/sliders/
package slider // import "github.com/vecty-material/material/slider"

import (
	"github.com/vecty-material/gojs"
	"github.com/vecty-material/material/base"
	"github.com/gopherjs/gopherwasm/js"
)

// S is a material slider component.
type S struct {
	mdc *base.Component

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

// New returns a new component.
func New() *S {
	c := &S{}
	c.Component()
	return c
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *S) Start(rootElem js.Value) error {
	return base.Start(c, rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *S) Stop() error {
	return base.Stop(c)
}

// Component returns the component's underlying base.Component.
func (c *S) Component() *base.Component {
	switch {
	case c.mdc == nil:
		c.mdc = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCSlider",
				MDCCamelCaseName: "slider",
			},
		}
		fallthrough
	case c.mdc.Value == js.Null():
		c.mdc.Component().SetState(c.StateMap())
	}
	return c.mdc.Component()
}

// StateMap implements the base.StateMapper interface.
func (c *S) StateMap() base.StateMap {
	sm := base.StateMap{
		"value":    c.Value,
		"min":      c.Min,
		"max":      c.Max,
		"step":     c.Step,
		"disabled": c.Disabled,
	}
	if c.Component().Value.Get("value").String() == "undefined" {
		sm["value"] = js.ValueOf(c).Get("Value")
	}
	if c.Component().Value.Get("min").String() == "undefined" {
		sm["min"] = js.ValueOf(c).Get("Min")
	}
	if c.Component().Value.Get("max").String() == "undefined" {
		sm["max"] = js.ValueOf(c).Get("Max")
	}
	if c.Component().Value.Get("step").String() == "undefined" {
		sm["step"] = js.ValueOf(c).Get("Step")
	}
	return sm
}

// Layout recomputes the dimensions and re-lays out the component. This should
// be called if the dimensions of the slider itself or any of its parent
// elements change programmatically (it is called automatically on resize).
func (s *S) Layout() error {
	var err error
	gojs.CatchException(&err)
	s.mdc.Call("layout")
	return err
}

// TODO: Handle custom events
// - MDCSlider:input
// - MDCSlider:change
