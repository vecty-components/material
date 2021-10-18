package radio

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/base/applyer"
)

// R is a vecty-material radio component.
type R struct {
	*base.MDC
	vecty.Core
	Root     vecty.MarkupOrChild  `vecty:"prop"`
	Input    vecty.MarkupOrChild  `vecty:"prop"`
	OnChange func(e *vecty.Event) `vecty:"prop"`
	Name     string               `vecty:"prop"`
	Checked  bool                 `vecty:"prop"`
	Disabled bool                 `vecty:"prop"`
	Value    string               `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *R) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	input, _ := c.NativeInput()

	// Built-in root element.
	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		input,
		elem.Div(
			vecty.Markup(vecty.Class("mdc-radio__background")),
			elem.Div(vecty.Markup(vecty.Class("mdc-radio__outer-circle"))),
			elem.Div(vecty.Markup(vecty.Class("mdc-radio__inner-circle"))),
		),
	)
}

func (c *R) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCRadio",
				MDCCamelCaseName: "radio",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{
			"checked":  &c.Checked,
			"disabled": &c.Disabled,
		})
	}

	vecty.Markup(
		vecty.Class("mdc-radio"),
		vecty.MarkupIf(c.Disabled,
			vecty.Class("mdc-radio--disabled"),
		),
	).Apply(h)
	c.MDC.RootElement = h
}

func (c *R) onChange(e *vecty.Event) {
	if c.OnChange != nil {
		c.OnChange(e)
	}
}

func (c *R) NativeInput() (element *vecty.HTML, id string) {
	niMarkup := base.MarkupOnly(c.Input)
	if c.Input != nil && niMarkup == nil {
		// User supplied input element.
		element = elem.Input(c.Input)
		id = applyer.FindID(element)
		return
	}

	// Built-in input element.
	element = elem.Input(
		vecty.Markup(
			vecty.MarkupIf(niMarkup != nil, niMarkup),
			event.Change(c.onChange),
			vecty.Class("mdc-radio__native-control"),
			prop.Type(prop.TypeRadio),
			prop.Checked(c.Checked),
			vecty.MarkupIf(c.Value != "", prop.Value(c.Value)),
			vecty.MarkupIf(c.Name != "", vecty.Property("name", c.Name)),
			vecty.Property("disabled", c.Disabled),
		),
	)
	id = applyer.FindID(element)
	return
}
