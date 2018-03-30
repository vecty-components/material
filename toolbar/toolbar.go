package toolbar

import (
	"agamigo.io/material/ripple"
	"agamigo.io/material/toolbar"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// T is a vecty-material toolbar component.
type T struct {
	*toolbar.T
	vecty.Core
	ID            string
	Markup        []vecty.Applyer
	rootElement   *vecty.HTML
	Ripple        bool
	ripple        *ripple.R
	SectionStart  vecty.List
	SectionCenter vecty.List
	SectionEnd    vecty.List
	Fixed         bool
}

// Render implements the vecty.Component interface.
func (c *T) Render() vecty.ComponentOrHTML {
	c.init()
	var start, center, end vecty.List
	if c.SectionStart != nil {
		start = make(vecty.List, len(c.SectionStart))
		for i, v := range c.SectionStart {
			start[i] = elem.Section(
				vecty.Markup(
					vecty.Class("mdc-toolbar__section"),
					vecty.Class("mdc-toolbar__section--align-start")),
				v,
			)
		}
	}
	if c.SectionCenter != nil {
		center = make(vecty.List, len(c.SectionCenter))
		for i, v := range c.SectionCenter {
			center[i] = elem.Section(
				vecty.Markup(
					vecty.Class("mdc-toolbar__section")),
				v,
			)
		}
	}
	if c.SectionEnd != nil {
		end = make(vecty.List, len(c.SectionEnd))
		for i, v := range c.SectionEnd {
			end[i] = elem.Section(
				vecty.Markup(
					vecty.Class("mdc-toolbar__section"),
					vecty.Class("mdc-toolbar__section--align-end")),
				v,
			)
		}
	}
	c.rootElement = elem.Header(
		vecty.Markup(
			vecty.MarkupIf(c.Markup != nil, c.Markup...),
			vecty.MarkupIf(c.ID != "", prop.ID(c.ID)),
			vecty.Class("mdc-toolbar"),
			vecty.MarkupIf(c.Fixed, vecty.Class("mdc-toolbar--fixed"))),
		elem.Div(vecty.Markup(vecty.Class("mdc-toolbar__row")),
			vecty.If(c.SectionStart != nil, elem.Section(
				vecty.Markup(
					vecty.Class("mdc-toolbar__section"),
					vecty.Class("mdc-toolbar__section--align-start")),
				c.SectionStart,
			)),
			center,
			end,
		),
	)
	return c.rootElement
}

func (c *T) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       c,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		RippleC:   c.ripple,
	}
}

func (c *T) Mount() {
	c.MDCRoot().Mount()
}

func (c *T) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *T) init() {
	if c.T == nil {
		c.T = toolbar.New()
	}
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
