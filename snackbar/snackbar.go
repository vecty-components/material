package snackbar

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/components/snackbar"
)

// S is a vecty-material snackbar component.
type S struct {
	*base.MDC
	vecty.Core
	Root    vecty.MarkupOrChild
	Label   vecty.MarkupOrChild
	Buttons []*button.B
}

// Render implements the vecty.Component interface.
func (c *S) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	buttons := make([]vecty.MarkupOrChild, len(c.Buttons))
	for i, b := range c.Buttons {
		markup := base.ExtractMarkupFromLink(b.Label.(*vecty.HTML))
		buttons[i] = elem.Button(
			vecty.Markup(
				prop.Type("button"),
				vecty.Class("mdc-button", "mdc-snackbar__action"),
				vecty.MarkupIf(
					markup.OnClick != nil && markup.PreventDefault,
					event.Click(markup.OnClick).PreventDefault(),
				),
				vecty.MarkupIf(
					markup.OnClick != nil && !markup.PreventDefault,
					event.Click(markup.OnClick),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-button__ripple"),
				),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("mdc-button__label"),
				),
				markup.Child,
			),
		)
	}

	// Built-in root element
	return elem.Aside(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-snackbar__surface"),
				vecty.Attribute("role", "status"),
				vecty.Attribute("aria-relevant", "additions"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-snackbar__label"),
					vecty.Attribute("aria-atomic", "false"),
				),
				c.Label,
			),
			elem.Div(
				append([]vecty.MarkupOrChild{
					vecty.Markup(
						vecty.Class("mdc-snackbar__actions"),
						vecty.Attribute("aria-atomic", "true"),
					),
				}, buttons...)...,
			),
		),
	)
}

func (c *S) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = snackbar.New()
	}

	vecty.Markup(
		vecty.Class("mdc-snackbar"),
	).Apply(h)
	c.MDC.RootElement = h
}
