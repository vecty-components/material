package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-components/material/demos/components"
	"github.com/vecty-components/material/radio"
	"github.com/vecty-components/material/typography"
)

func NewRadioPage() *components.ComponentPage {
	return components.NewComponentPage(
		"Radio",
		"Radioes allow the user to select multiple options from a set.",
		"https://material.io/go/design-checkboxes",
		"https://material.io/components/web/catalog/checkboxes/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-checkbox",
		components.NewHeroComponent(&RadioHero{}), &RadioDemos{},
	)
}

type RadioHero struct {
	vecty.Core
}

func (bh *RadioHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),

		&radio.R{
			Root: vecty.Markup(
				vecty.Class("demo-checkbox"),
			),
			Name: "radio-hero",
		},

		&radio.R{
			Root: vecty.Markup(
				vecty.Class("demo-checkbox"),
			),
			Name: "radio-hero",
		},
	)
}

type RadioDemos struct {
	vecty.Core
}

func (bd *RadioDemos) Render() vecty.ComponentOrHTML {

	return elem.Div(
		typography.Subtitle1(
			vecty.Text("Enabled"),
		),

		&radio.R{},

		typography.Subtitle1(
			vecty.Text("Disabled"),
		),

		&radio.R{
			Disabled: true,
		},
	)
}
