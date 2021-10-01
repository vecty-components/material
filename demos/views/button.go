package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/demos/components"
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
	return elem.Div()
}

type ButtonDemos struct {
	vecty.Core
}

func (bd *ButtonDemos) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Text("button demos"),
	)
}
