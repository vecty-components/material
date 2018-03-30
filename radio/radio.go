package radio

import (
	"agamigo.io/material/radio"
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// R is a vecty-material radio component.
type R struct {
	*radio.R
	vecty.Core
	ID          string
	Markup      []vecty.Applyer
	rootElement *vecty.HTML
	Ripple      bool
	Basic       bool
	ripple      *ripple.R
	OnChange    func(this *R, e *vecty.Event)
	Name        string
	Checked     bool
	Disabled    bool
	Value       string
}

// Render implements the vecty.Component interface.
func (c *R) Render() vecty.ComponentOrHTML {
	c.init()
	c.rootElement = elem.Div(
		vecty.Markup(
			vecty.Markup(c.Markup...),
			vecty.Class("mdc-radio"),
			vecty.MarkupIf(c.Disabled,
				vecty.Class("mdc-radio--disabled"),
			),
		),
		elem.Input(
			vecty.Markup(
				event.Change(c.onChange),
				vecty.Class("mdc-radio__native-control"),
				vecty.MarkupIf(c.ID != "",
					prop.ID(c.ID),
				),
				prop.Type(prop.TypeRadio),
				prop.Checked(c.Checked),
				vecty.MarkupIf(c.Value != "",
					prop.Value(c.Value),
				),
				vecty.MarkupIf(c.Name != "", vecty.Property("name", c.Name)),
				vecty.Property("disabled", c.Disabled),
			),
		),
		elem.Div(
			vecty.Markup(vecty.Class("mdc-radio__background")),
			elem.Div(vecty.Markup(vecty.Class("mdc-radio__outer-circle"))),
			elem.Div(vecty.Markup(vecty.Class("mdc-radio__inner-circle"))),
		),
	)
	return c.rootElement
}

func (c *R) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       c,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		Basic:     c.Basic,
		RippleC:   c.ripple,
	}
}

func (c *R) Mount() {
	c.MDCRoot().Mount()
}

func (c *R) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *R) init() {
	switch {
	case c.R == nil:
		c.R = radio.New()
		fallthrough
	case c.rootElement == nil:
		c.R.Checked = c.Checked
		c.R.Disabled = c.Disabled
		c.R.Value = c.Value
	}
}

func (c *R) onChange(e *vecty.Event) {
	c.Checked = c.R.Checked
	c.Disabled = c.R.Disabled
	c.Value = c.R.Value
	if c.OnChange != nil {
		c.OnChange(c, e)
	}
}
