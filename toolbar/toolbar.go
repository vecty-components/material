package toolbar

import (
	"agamigo.io/material/toolbar"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// T is a vecty-material toolbar component.
type T struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild
	SectionStart  vecty.List
	SectionCenter vecty.List
	SectionEnd    vecty.List
	Fixed         bool
}

// Render implements the vecty.Component interface.
func (c *T) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Header(c.Root)
	}

	// Built in root element.
	return elem.Header(
		vecty.Markup(
			c,
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-toolbar__row"),
			),
			vecty.If(c.SectionStart != nil,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-toolbar__section"),
						vecty.Class("mdc-toolbar__section--align-start"),
					),
					c.SectionStart,
				),
			),
			vecty.If(c.SectionCenter != nil,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-toolbar__section"),
					),
					c.SectionCenter,
				),
			),
			vecty.If(c.SectionEnd != nil,
				elem.Section(
					vecty.Markup(
						vecty.Class("mdc-toolbar__section"),
						vecty.Class("mdc-toolbar__section--align-end"),
					),
					c.SectionEnd,
				),
			),
		),
	)
}

func (c *T) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = toolbar.New()
	}
	c.MDC.RootElement = h
	vecty.Markup(
		vecty.Class("mdc-toolbar"),
		vecty.MarkupIf(c.Fixed,
			vecty.Class("mdc-toolbar--fixed"),
		),
	).Apply(h)
}

func Title(title string, mUp []vecty.Applyer) *vecty.HTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("mdc-toolbar__title"),
			vecty.Markup(mUp...),
		),
		vecty.Text(title),
	)
}
