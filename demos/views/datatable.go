package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-material/material/datatable"
	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/typography"
)

func NewDatatablePage() *components.ComponentPage {
	return components.NewComponentPage(
		"Datatable",
		"Datatablees allow the user to select multiple options from a set.",
		"https://material.io/go/design-checkboxes",
		"https://material.io/components/web/catalog/checkboxes/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-checkbox",
		components.NewHeroComponent(&DatatableHero{}), &DatatableDemos{},
	)
}

type DatatableHero struct {
	vecty.Core
}

func (bh *DatatableHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),
	)
}

type DatatableDemos struct {
	vecty.Core
}

func (bd *DatatableDemos) Render() vecty.ComponentOrHTML {

	return elem.Div(
		typography.Subtitle1(
			vecty.Text("Tabs with icons next to labels"),
		),

		&datatable.DT{
			Head: &datatable.R{
				Cells: []*datatable.C{
					{Label: vecty.Text("Col A")},
					{Label: vecty.Text("Col B")},
				},
			},
			Rows: []*datatable.R{
				{
					Cells: []*datatable.C{
						{Label: vecty.Text("Row A")},
						{Label: vecty.Text("Row B")},
					},
				},
			},
		},
	)

}
