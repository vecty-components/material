package checkbox

import (
	"agamigo.io/material/checkbox"
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// CB is a vecty-material checkbox component.
type CB struct {
	*checkbox.CB
	vecty.Core
	ID            string
	Markup        []vecty.Applyer
	rootElement   *vecty.HTML
	Ripple        bool
	Basic         bool
	ripple        *ripple.R
	OnChange      func(this *CB, e *vecty.Event)
	Checked       bool
	Indeterminate bool
	Disabled      bool
	Value         string
}

// Render implements the vecty.Component interface.
func (c *CB) Render() vecty.ComponentOrHTML {
	c.init()
	c.rootElement = elem.Div(
		vecty.Markup(
			vecty.MarkupIf(c.Markup != nil, vecty.Markup(c.Markup...)),
			vecty.Class("mdc-checkbox"),
			vecty.MarkupIf(c.Disabled,
				vecty.Class("mdc-checkbox--disabled"),
			),
		),
		elem.Input(
			vecty.Markup(
				event.Change(c.onChange),
				vecty.Class("mdc-checkbox__native-control"),
				vecty.MarkupIf(c.ID != "",
					prop.ID(c.ID)),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(c.CB.Checked),
				vecty.MarkupIf(c.Value != "",
					prop.Value(c.Value),
				),
				vecty.Property("disabled", c.Disabled),
				vecty.Property("indeterminate", c.Indeterminate),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-checkbox__background"),
				vecty.UnsafeHTML(
					`<svg class="mdc-checkbox__checkmark"
							viewBox="0 0 24 24">
							<path class="mdc-checkbox__checkmark-path"
								fill="none"
								stroke="white"
								d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
							</svg>`,
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-checkbox__mixedmark"),
				),
			),
		),
	)
	return c.rootElement
}

func (c *CB) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       c,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		Basic:     c.Basic,
		RippleC:   c.ripple,
	}
}

func (c *CB) Mount() {
	c.MDCRoot().Mount()
}

func (c *CB) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *CB) init() {
	switch {
	case c.CB == nil:
		c.CB = checkbox.New()
		fallthrough
	case c.rootElement == nil:
		c.CB.Checked = c.Checked
		c.CB.Indeterminate = c.Indeterminate
		c.CB.Disabled = c.Disabled
		c.CB.Value = c.Value
	}
}

func (c *CB) onChange(e *vecty.Event) {
	c.Checked = c.CB.Checked
	c.Indeterminate = c.CB.Indeterminate
	c.Disabled = c.CB.Disabled
	c.Value = c.CB.Value
	if c.OnChange != nil {
		c.OnChange(c, e)
	}
}
