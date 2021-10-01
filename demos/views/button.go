package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/demos/components"
)

func NewButtonPage() *components.ComponentCatalogPanel {
	return components.NewComponentCatalogPanel(
		"", "", "", "", "",
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
