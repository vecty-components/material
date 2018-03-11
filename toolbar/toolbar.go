package toolbar

import (
	"agamigo.io/material/toolbar"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// T is a vecty-material toolbar component.
type T struct {
	*base.Base
	*State
}

type State struct {
	*toolbar.T
	SectionStart  vecty.List
	SectionCenter vecty.List
	SectionEnd    vecty.List
	Fixed         bool
}

func New(p *base.Props, s *State) *T {
	c := &T{}
	if s == nil {
		s = &State{}
	}
	if s.T == nil {
		s.T = toolbar.New()
	}
	c.State = s
	c.Base = base.New(p, nil)
	return c
}

// Render implements the vecty.Component interface.
func (c *T) Render() vecty.ComponentOrHTML {
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
	return elem.Header(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.MarkupIf(c.Props.ID != "", prop.ID(c.Props.ID)),
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
}

func Title(title string, p *base.Props) *vecty.HTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("mdc-toolbar__title"),
		),
		vecty.Text(title),
	)
}
