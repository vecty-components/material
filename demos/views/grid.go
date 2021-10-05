package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/grid"
	"github.com/vecty-material/material/typography"
)

func NewLayoutGridPage() *components.ComponentPage {
	return components.NewComponentPage(
		"LayoutGrid",
		"LayoutGrides allow the user to select multiple options from a set.",
		"https://material.io/go/design-checkboxes",
		"https://material.io/components/web/catalog/checkboxes/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-checkbox",
		components.NewHeroComponent(&LayoutGridHero{}), &LayoutGridDemos{},
	)
}

type LayoutGridHero struct {
	vecty.Core
}

func (bh *LayoutGridHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),
	)
}

type LayoutGridDemos struct {
	vecty.Core
}

func (bd *LayoutGridDemos) Render() vecty.ComponentOrHTML {

	return elem.Div(
		typography.Subtitle1(
			vecty.Text("Tabs with icons next to labels"),
		),

		&grid.G{
			Cells: []*grid.C{
				{Label: vecty.Text("Cell A")},
				{Label: vecty.Text("Cell B")},
				{
					Cells: []*grid.C{
						{Label: vecty.Text("Cell C")},
						{Label: vecty.Text("Cell D")},
					},
				},
			},
		},
	)

}
