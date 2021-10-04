package appbar

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/components/appbar"
)

// A is a vecty-material appbar component.
type A struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild
	SectionStart  vecty.List
	SectionCenter vecty.List
	SectionEnd    vecty.List
	Fixed         bool
}

// Render implements the vecty.Component interface.
func (c *A) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Header(c.Root)
	}

	// Built in root element.
	return elem.Header(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-top-app-bar__row"),
			),
			vecty.If(c.SectionStart != nil,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-top-app-bar__section"),
						vecty.Class("mdc-top-app-bar__section--align-start"),
					),
					c.SectionStart,
				),
			),
			vecty.If(c.SectionCenter != nil,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-top-app-bar__section"),
					),
					c.SectionCenter,
				),
			),
			vecty.If(c.SectionEnd != nil,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-top-app-bar__section"),
						vecty.Class("mdc-top-app-bar__section--align-end"),
					),
					c.SectionEnd,
				),
			),
		),
	)
}

func (c *A) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = appbar.New()
	}
	c.MDC.RootElement = h
	vecty.Markup(
		vecty.Class("mdc-top-app-bar"),
		vecty.MarkupIf(c.Fixed,
			vecty.Class("mdc-top-app-bar--fixed"),
		),
	).Apply(h)
}

func Title(title string, mUp []vecty.Applyer) *vecty.HTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("mdc-top-app-bar__title"),
			vecty.Markup(mUp...),
		),
		vecty.Text(title),
	)
}
