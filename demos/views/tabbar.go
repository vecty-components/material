package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/tabbar"
	"github.com/vecty-material/material/typography"
)

func NewTabbarPage() *components.ComponentPage {
	return components.NewComponentPage(
		"Tabbar",
		"Tabbares allow the user to select multiple options from a set.",
		"https://material.io/go/design-checkboxes",
		"https://material.io/components/web/catalog/checkboxes/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-checkbox",
		components.NewHeroComponent(&TabbarHero{}), &TabbarDemos{},
	)
}

type TabbarHero struct {
	vecty.Core
}

func (bh *TabbarHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),
	)
}

type TabbarDemos struct {
	vecty.Core
}

func (bd *TabbarDemos) Render() vecty.ComponentOrHTML {

	return elem.Div(
		typography.Subtitle1(
			vecty.Text("Tabs with icons next to labels"),
		),

		&tabbar.TB{
			Tabs: []*tabbar.T{
				{
					Label: vecty.Text("First tab"),
				},
				{
					Label: vecty.Text("Second tab"),
				},
			},
		},
	)

}
