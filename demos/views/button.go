package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/icon"
)

func NewButtonPage() *components.ComponentCatalogPanel {
	return components.NewComponentCatalogPanel(
		"Button",
		"Buttons communicate an action a user can take. They "+
			"are typically placed throughout your UI, in places "+
			"like dialogs, forms, cards, and toolbars.",
		"https://material.io/go/design-buttons",
		"https://material.io/components/web/catalog/buttons/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-button",
		components.NewHeroComponent(
			&ButtonHero{},
		),
		&ButtonDemos{},
	)
}

type ButtonHero struct {
	vecty.Core
}

func (bh *ButtonHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),
		&button.B{
			Root: vecty.Markup(
				vecty.Class("hero-button"),
			),
			Label: vecty.Text("Learn More"),
		},
	)
}

type ButtonDemos struct {
	vecty.Core
}

func (bd *ButtonDemos) renderButtonVariant(title string, variantClass []string) vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading3(
			vecty.Markup(
				vecty.Class("mdc-typography--subtitle1"),
			),
			vecty.Text(title),
		),
		&button.B{
			Root: vecty.Markup(
				vecty.Class("demo-button"),
				vecty.Class(variantClass...),
			),
			Label: vecty.Text("Default"),
		},
		&button.B{
			Root: vecty.Markup(
				vecty.Class("demo-button"),
				vecty.Class(variantClass...),
			),
			Label: vecty.Text("Dense"),
			Dense: true,
		},
		&button.B{
			Root: vecty.Markup(
				vecty.Class("demo-button"),
				vecty.Class(variantClass...),
			),
			Label: vecty.Text("Icon"),
			Icon: &icon.I{
				Name: "favorite",
			},
		},
	)
}

func (bd *ButtonDemos) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/ButtonCatalog.css")

	return elem.Div(
		bd.renderButtonVariant("Text Button", []string{}),
		bd.renderButtonVariant("Raised Button", []string{"mdc-button--raised"}),
		bd.renderButtonVariant("Unelevated Button", []string{"mdc-button--unelevated"}),
		bd.renderButtonVariant("Outlined Button", []string{"mdc-button--outlined"}),
		bd.renderButtonVariant(
			"Shaped Button", []string{"mdc-button--unelevated", "demo-button-shaped"},
		),
	)
}
