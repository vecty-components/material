package appbar

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-components/material/base"
)

// A is a vecty-material appbar component.
type A struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild `vecty:"prop"`
	SectionStart  vecty.List          `vecty:"prop"`
	SectionCenter vecty.List          `vecty:"prop"`
	SectionEnd    vecty.List          `vecty:"prop"`
	Fixed         bool                `vecty:"prop"`
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
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTopAppBar",
				MDCCamelCaseName: "topAppBar",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{})
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
