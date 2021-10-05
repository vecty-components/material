package grid

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type C struct {
	vecty.Core
	Label vecty.ComponentOrHTML
	Cells []*C
}

func (c *C) Render() vecty.ComponentOrHTML {
	if c.Label != nil && c.Cells != nil {
		panic("cell cannot have both label and cells")
	}

	var label vecty.ComponentOrHTML
	if c.Label != nil {
		label = c.Label
	} else {
		cells := make([]vecty.MarkupOrChild, len(c.Cells))
		for i, c := range c.Cells {
			if c.Cells != nil {
				panic("layout grid cannot have depth greater than 2")
			}

			cells[i] = elem.Div(
				vecty.Markup(
					vecty.Class("mdc-layout-grid__cell"),
				),
				elem.Span(c.Label),
			)
		}

		label = elem.Div(
			append([]vecty.MarkupOrChild{
				vecty.Markup(
					vecty.Class("mdc-layout-grid__inner"),
				),
			}, cells...)...,
		)
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-layout-grid__cell"),
		),
		label,
	)
}

type G struct {
	vecty.Core
	Cells []*C
}

func (g *G) Render() vecty.ComponentOrHTML {
	cells := make([]vecty.MarkupOrChild, len(g.Cells))
	for i, c := range g.Cells {
		cells[i] = c
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-layout-grid"),
		),
		elem.Div(
			append([]vecty.MarkupOrChild{
				vecty.Markup(
					vecty.Class("mdc-layout-grid__inner"),
				),
			}, cells...)...,
		),
	)
}
