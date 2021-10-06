package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-components/material/checkbox"
	"github.com/vecty-components/material/datatable"
	"github.com/vecty-components/material/demos/components"
	"github.com/vecty-components/material/typography"
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
					{Label: &checkbox.CB{}},
					{Label: vecty.Text("Col A")},
					{Label: vecty.Text("Col B")},
				},
			},
			Rows: []*datatable.R{
				{
					Cells: []*datatable.C{
						{Label: &checkbox.CB{}},
						{Label: vecty.Text("Row 1 A")},
						{Label: vecty.Text("Row 1 B")},
					},
				},
				{
					Cells: []*datatable.C{
						{Label: &checkbox.CB{}},
						{Label: vecty.Text("Row 2 A")},
						{Label: vecty.Text("Row 2 B")},
					},
				},
			},
		},
	)

}
