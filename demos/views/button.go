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
		&ButtonHero{}, &ButtonDemos{},
	)
}

type ButtonHero struct {
	vecty.Core
}

func (bh *ButtonHero) Render() vecty.ComponentOrHTML {
	return &button.B{
		Root: vecty.Markup(
			vecty.Class("hero-button"),
		),
	}
}

type ButtonDemos struct {
	vecty.Core
}

func (bd *ButtonDemos) renderButtonVariant(title, variantClass string) vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading3(
			vecty.Markup(
				vecty.Class("mdc-typography--subtitle1"),
			),
		),
		&button.B{
			Root: vecty.Markup(
				vecty.Class("demo-button", variantClass),
			),
			Label: vecty.Text("Default"),
		},
		&button.B{
			Root: vecty.Markup(
				vecty.Class("demo-button", variantClass),
			),
			Label: vecty.Text("Dense"),
			Dense: true,
		},
		&button.B{
			Root: vecty.Markup(
				vecty.Class("demo-button", variantClass),
			),
			Label: vecty.Text("Icon"),
			Icon: &icon.I{
				Name: "favorite",
			},
		},
	)
}

func (bd *ButtonDemos) Render() vecty.ComponentOrHTML {
	return elem.Div(
		// bd.renderButtonVariant("Text Button", ""),
		bd.renderButtonVariant("Raised Button", "mdc-button--raised"),
		bd.renderButtonVariant("Unelevated Button", "mdc-button--unelevated"),
		bd.renderButtonVariant("Outlined Button", "mdc-button--outlined"),
		// bd.renderButtonVariant("Shaped Button", "mdc-button--unelevated demo-button-shaped"),
	)
}
