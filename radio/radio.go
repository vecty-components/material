package radio

import (
	"agamigo.io/material/radio"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/base/applyer"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// R is a vecty-material radio component.
type R struct {
	*base.MDC
	vecty.Core
	Root     vecty.MarkupOrChild
	Input    vecty.MarkupOrChild
	OnChange func(this *R, e *vecty.Event)
	Name     string
	Checked  bool
	Disabled bool
	Value    string
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
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
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
		c.MDC.Component = radio.New()
		if r, ok := c.MDC.Component.(*radio.R); ok {
			r.Checked = c.Checked
			r.Disabled = c.Disabled
			r.Value = c.Value
		}
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
	if r, ok := c.MDC.Component.(*radio.R); ok {
		c.Checked = r.Checked
		c.Disabled = r.Disabled
		c.Value = r.Value
	}
	if c.OnChange != nil {
		c.OnChange(c, e)
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
